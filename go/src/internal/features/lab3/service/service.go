package service

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	mongoDBName               = "my_database"
	mongoGroupsCollectionName = "groups"
	lectureMultiplier         = 2
)

type Service struct {
	elasticCli *elasticsearch.Client
	neoCli     neo4j.DriverWithContext
	pgCli      *pgx.Conn
	redisCli   *redis.Client
	mongoCli   *mongo.Client
}

type Student struct {
	StudentID int    `bson:"student_id" json:"student_id"`
	FullName  string `bson:"full_name" json:"full_name"`
}

type Group struct {
	GroupID   int       `bson:"group_id" json:"group_id"`
	GroupName string    `bson:"group_name" json:"group_name"`
	Students  []Student `bson:"students" json:"students"`
}

type Course struct {
	CourseID    int64   `bson:"id" json:"id"`
	CourseName  string  `bson:"name" json:"name"`
	ScheduleIDs []int64 `bson:"scheduleIds" json:"scheduleIds"`
}

func (s *Service) getGroupByName(ctx context.Context, groupName string) (Group, error) {
	var group Group

	result := s.mongoCli.Database(mongoDBName).Collection(mongoGroupsCollectionName).FindOne(ctx, bson.M{"group_name": groupName})

	err := result.Decode(&group)
	if err != nil {
		return Group{}, err
	}

	return group, nil
}

func (s *Service) getGroupCourses(ctx context.Context, groupId int) ([]Course, error) {
	session := s.neoCli.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func() {
		_ = session.Close(ctx)
	}()

	query := `
        MATCH (:Group {id: $groupId})-[:FOR_GROUP]- (sched:Schedule)-[:SCHEDULED_AT]- (cls:Class)-[:HAS_CLASS]- (c:Course {special: true})
        WITH c.id AS courseId, c.name AS courseName, collect(sched.id) AS scheduleIds
        RETURN collect({
            id: courseId,
            name: courseName,
            scheduleIds: scheduleIds
        }) AS courses
    `

	params := map[string]interface{}{
		"groupId": groupId,
	}

	result, err := session.Run(ctx, query, params)

	if err != nil {
		return nil, fmt.Errorf("error getting response: %s", err)
	}

	courses := make([]Course, 0)

	for result.Next(ctx) {
		record := result.Record()

		coursesData, ok := record.Get("courses")
		if !ok {
			return nil, fmt.Errorf("'courses' field not found in the record")
		}

		coursesList, ok := coursesData.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid type for 'courses' field")
		}

		for _, c := range coursesList {
			courseMap, ok := c.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid course data format")
			}

			id, ok := courseMap["id"].(int64)
			if !ok {
				switch v := courseMap["id"].(type) {
				case int:
					id = int64(v)
				case int32:
					id = int64(v)
				case float64:
					id = int64(v)
				default:
					return nil, fmt.Errorf("invalid type for course id")
				}
			}

			name, ok := courseMap["name"].(string)
			if !ok {
				return nil, fmt.Errorf("invalid type for course name")
			}

			schedIDsInterface, ok := courseMap["scheduleIds"].([]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid type for scheduleIds")
			}

			scheduleIDs := make([]int64, 0, len(schedIDsInterface))
			for _, sid := range schedIDsInterface {
				var sidInt int64
				switch v := sid.(type) {
				case int:
					sidInt = int64(v)
				case int32:
					sidInt = int64(v)
				case int64:
					sidInt = v
				case float64:
					sidInt = int64(v)
				default:
					return nil, fmt.Errorf("invalid type for schedule id")
				}
				scheduleIDs = append(scheduleIDs, sidInt)
			}

			courses = append(courses, Course{
				CourseID:    id,
				CourseName:  name,
				ScheduleIDs: scheduleIDs,
			})
		}
	}

	return courses, nil
}

type StudentAttendanceItem struct {
	StudentID            int `json:"student_id"`
	AttendancePercentage int `json:"attendance_percentage"`
}

type Attendance map[int]int

func (s *Service) GetStudentsWithAttendances(ctx context.Context, scheduleIDs []int64) (Attendance, error) {
	query := `
        SELECT
            student_id,
            COUNT(*) FILTER (WHERE status = '+') AS attendance
        FROM
            attendances
        WHERE
            schedule_id = ANY($1)
        GROUP BY
            student_id;
    `

	rows, err := s.pgCli.Query(ctx, query, pq.Array(scheduleIDs))
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	attendances := make(Attendance)

	for rows.Next() {
		var item StudentAttendanceItem
		if err := rows.Scan(&item.StudentID, &item.AttendancePercentage); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		attendances[item.StudentID] = item.AttendancePercentage
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return attendances, nil
}

type Lab3ReportItem struct {
	GroupName          string              `json:"group_name"`
	CourseName         string              `json:"course_name"`
	Lab3ReportStudents []Lab3ReportStudent `json:"students_attendance"`
}

type Lab3ReportStudent struct {
	StudentGradeBookID int    `json:"student_grade_book_id"`
	StudentFullName    string `json:"student_full_name"`
	VisitedQuantity    int    `json:"visited_quantity"`
	PlannedQuantity    int    `json:"planned_quantity"`
}

func (s *Service) Execute(ctx context.Context, in *In) ([]Lab3ReportItem, error) {
	group, err := s.getGroupByName(ctx, in.GroupName)
	if err != nil {
		return nil, err
	}
	courses, err := s.getGroupCourses(ctx, group.GroupID)
	if err != nil {
		return nil, err
	}

	report := make([]Lab3ReportItem, 0)
	for _, course := range courses {
		attendance, err := s.GetStudentsWithAttendances(ctx, course.ScheduleIDs)
		if err != nil {
			return nil, err
		}
		reportItem := Lab3ReportItem{}
		reportItem.GroupName = group.GroupName
		reportItem.CourseName = course.CourseName
		reportItem.Lab3ReportStudents = make([]Lab3ReportStudent, 0)
		for _, student := range group.Students {
			studentItem := Lab3ReportStudent{}
			studentItem.StudentGradeBookID = student.StudentID
			studentItem.StudentFullName = student.FullName
			studentItem.PlannedQuantity = len(course.ScheduleIDs) * lectureMultiplier
			studentItem.VisitedQuantity = attendance[student.StudentID] * lectureMultiplier
			reportItem.Lab3ReportStudents = append(reportItem.Lab3ReportStudents, studentItem)
		}
		report = append(report, reportItem)
	}

	return report, nil
}

func New(elasticCli *elasticsearch.Client, neoCli neo4j.DriverWithContext, pgCli *pgx.Conn, redisCli *redis.Client, mongoCli *mongo.Client) *Service {
	return &Service{
		elasticCli: elasticCli,
		neoCli:     neoCli,
		pgCli:      pgCli,
		redisCli:   redisCli,
		mongoCli:   mongoCli,
	}
}
