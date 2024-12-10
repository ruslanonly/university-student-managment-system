package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/mpvl/unique"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"strings"
	"time"
)

const (
	materialIndexName = "material"
)

type StudentAttendanceItem struct {
	StudentID  int
	Attendance int
}

type Service struct {
	elasticCli *elasticsearch.Client
	neoCli     neo4j.DriverWithContext
	pgCli      *pgx.Conn
}

func (s *Service) getClassIDsByPhrase(ctx context.Context, phrase string) ([]int, error) {
	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"content": "%s"
			}
		}
	}`, phrase)

	req := esapi.SearchRequest{
		Index: []string{materialIndexName},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(ctx, s.elasticCli)

	if err != nil {
		return nil, fmt.Errorf("error getting response: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.IsError() {
		return nil, fmt.Errorf("[%s] error in search request", res.Status())
	}

	var results map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	hits := results["hits"].(map[string]interface{})["hits"].([]interface{})

	classIDs := make([]int, 0)

	for _, hit := range hits {
		hitSource := hit.(map[string]interface{})["_source"]
		classID := int(hitSource.(map[string]interface{})["class_id"].(float64))
		classIDs = append(classIDs, classID)
	}

	return classIDs, nil
}

func (s *Service) getStudentIDsByAttendance(ctx context.Context, courseIDs []int) ([]int, error) {
	session := s.neoCli.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func() {
		_ = session.Close(ctx)
	}()

	query := `
			MATCH (c: Class) -[:SCHEDULED_AT]-> (sch: Schedule) -[:FOR_GROUP]-> (g: Group) -[:HAS_STUDENT]-> (st: Student)
			WHERE c.id IN $courseIDs
			RETURN st.id;
		`

	params := map[string]any{
		"courseIDs": courseIDs,
	}

	result, err := session.Run(ctx, query, params)

	if err != nil {
		return nil, fmt.Errorf("error getting response: %s", err)
	}

	ids := make([]int, 0)

	for result.Next(ctx) {
		record := result.Record()
		id, ok := record.Values[0].(int64)
		if !ok {
			continue
		}
		ids = append(ids, int(id))
	}

	unique.Ints(&ids)

	return ids, nil
}

func (s *Service) getStudentWithBaddestAttendance(ctx context.Context, studentIDs []int, periodStart, periodEnd time.Time) ([]StudentAttendanceItem, error) {
	query := `
		SELECT
			a.student_id,
			ROUND(SUM(CASE WHEN a.status = '+' THEN 1 ELSE 0 END) * 100.0 / COUNT(*), 2) AS attendance_percentage
		FROM (
				SELECT *
				FROM attendances
				WHERE student_id = ANY($1)
				AND date >= $2
				AND date <= $3
			 ) AS a
		GROUP BY a.student_id
		ORDER BY attendance_percentage
		LIMIT 10;
	`

	args := make([]interface{}, 0)
	args = append(args, pq.Array(studentIDs), periodStart, periodEnd)

	rows, err := s.pgCli.Query(ctx, query, args...)

	if err != nil {
		return nil, fmt.Errorf("error getting response: %s", err)
	}

	defer rows.Close()
	result := make([]StudentAttendanceItem, 0)

	for rows.Next() {
		var item StudentAttendanceItem
		err := rows.Scan(&item.StudentID, &item.Attendance)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %s", err)
		}
		result = append(result, item)
	}

	return result, nil
}

func (s *Service) getStudentFullName(ctx context.Context, studentID int) (string, error) {
	return "", nil
}

func (s *Service) Execute(ctx context.Context, in In) ([]ReportItem, error) {
	courseIDs, err := s.getClassIDsByPhrase(ctx, in.Phrase)

	if err != nil {
		return nil, err
	}

	studentIDs, err := s.getStudentIDsByAttendance(ctx, courseIDs)

	if err != nil {
		return nil, err
	}

	studentAttendance, err := s.getStudentWithBaddestAttendance(ctx, studentIDs, in.PeriodStart, in.PeriodEnd)

	if err != nil {
		return nil, err
	}

	result := make([]ReportItem, 0)

	for _, sa := range studentAttendance {
		studentFullName, err := s.getStudentFullName(ctx, sa.StudentID)

		if err != nil {
			return nil, err
		}

		result = append(result, ReportItem{
			StudentFullName: studentFullName,
			Attendance:      sa.Attendance,
			PeriodStart:     in.PeriodStart,
			PeriodEnd:       in.PeriodEnd,
			Phrase:          in.Phrase,
		})
	}

	return result, nil
}

func New(elasticCli *elasticsearch.Client, neoCli neo4j.DriverWithContext, pgCli *pgx.Conn) *Service {
	return &Service{
		elasticCli: elasticCli,
		neoCli:     neoCli,
		pgCli:      pgCli,
	}
}
