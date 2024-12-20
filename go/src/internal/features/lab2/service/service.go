package service

import (
	"context"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jackc/pgx/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CourseScheduledLecturesWithStudentsAmount struct {
	CourseId int `json:"course_id"`
	CourseName string `json:"course_name"`
	ClassId int `json:"class_id"`
	ClassDate time.Time `json:"class_date"`
	GroupId int `json:"group_id"`
	StudentsAmount   int    `json:"students_amount"`
}

type CourseScheduledLecturesWithStudentsAmountGrouped struct {
	CourseId int `json:"course_id"`
	CourseName string `json:"course_name"`
	ClassId int `json:"class_id"`
	ClassDate time.Time `json:"class_date"`
	GroupesIds []int `json:"group_ids"`
	StudentsAmount   int    `json:"students_amount"`
}

type Service struct {
	elasticCli *elasticsearch.Client
	neoCli     neo4j.DriverWithContext
	pgCli      *pgx.Conn
	redisCli   *redis.Client
	mongoCli   *mongo.Client
}

func (s *Service) getCourseScheduledLecturesAmountOfStudents(ctx context.Context, courseName string, year int, semester int) ([]CourseScheduledLecturesWithStudentsAmountGrouped, error) {
	session := s.neoCli.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func() {
		_ = session.Close(ctx)
	}()

	query := `
		MATCH (cr: Course) -[:HAS_CLASS]-> (cl: Class) -[TAKES_TIME_IN]-> (sc: Schedule) -[:FOR_GROUP]- (gr: Group) -[:HAS_STUDENT]-> (st: Student)
		WHERE cr.name = $courseName
		AND cl.type = 'Лекторная'
		AND date(sc.date).year = $year
		AND (
			($semester = 2 AND date(sc.date).month >= 1 AND date(sc.date).month <= 6) OR
			($semester = 1 AND date(sc.date).month >= 7 AND date(sc.date).month <= 12)
		)
		RETURN cr.id AS courseId, cr.name AS courseName, cl.id AS classId, sc.date AS classDate, gr.id AS groupId, count(st.id) AS studentsAmount;
	`
	params := map[string]any{
		"courseName": courseName,
		"year": year,
		"semester": semester,
	}

	result, err := session.Run(ctx, query, params)

	if err != nil {
		return nil, fmt.Errorf("Error during request: %s", err)
	}

	var schedules []CourseScheduledLecturesWithStudentsAmount

	for result.Next(ctx) {
		record := result.Record()

		courseID, _ := record.Get("courseId")
		courseName, _ := record.Get("courseName")
		groupID, _ := record.Get("groupId")
		classID, _ := record.Get("classId")
		classDate, _ := record.Get("classDate")
		studentsAmount, _ := record.Get("studentsAmount")

		schedule := CourseScheduledLecturesWithStudentsAmount{
			CourseId:   int(courseID.(int64)),
			ClassId:    int(classID.(int64)),
			ClassDate: classDate.(time.Time),
			CourseName: courseName.(string),
			GroupId: int(groupID.(int64)),
			StudentsAmount: int(studentsAmount.(int64)),
		}

		schedules = append(schedules, schedule)
	}

	groupedSchedulesMap := map[string]CourseScheduledLecturesWithStudentsAmountGrouped{}

	for i := 0; i < len(schedules); i++ {
		currentSchedule := schedules[i]

		key := fmt.Sprintf("%d-%d-%s", currentSchedule.CourseId, currentSchedule.ClassId, currentSchedule.ClassDate)

		groupedSchedule, groupedScheduleExists := groupedSchedulesMap[key]

		if groupedScheduleExists {
			exists := false
			for _, groupID := range groupedSchedule.GroupesIds {
				if groupID == currentSchedule.GroupId {
					exists = true
					break
				}
			}
			if !exists {
				groupedSchedule.GroupesIds = append(groupedSchedule.GroupesIds, currentSchedule.GroupId)
				groupedSchedule.StudentsAmount += currentSchedule.StudentsAmount
				groupedSchedulesMap[key] = groupedSchedule
			}

		} else {
			groupedSchedulesMap[key] = CourseScheduledLecturesWithStudentsAmountGrouped{
				CourseId:      currentSchedule.CourseId,
				ClassId:       currentSchedule.ClassId,
				ClassDate:     currentSchedule.ClassDate,
				CourseName:    currentSchedule.CourseName,
				GroupesIds:    []int{currentSchedule.GroupId},
				StudentsAmount: currentSchedule.StudentsAmount,
			}
		}
	}

	groupedSchedules := make([]CourseScheduledLecturesWithStudentsAmountGrouped, 0, len(groupedSchedulesMap))

	for _, value := range groupedSchedulesMap {
		groupedSchedules = append(groupedSchedules, value)
	}

	return groupedSchedules, nil
}

func (s *Service) Execute(ctx context.Context, in *In) ([]CourseScheduledLecturesWithStudentsAmountGrouped, error) {
	if in.Year < 1 {
		return nil, fmt.Errorf("Incorrect year format (year should be greater than 0): %d", in.Year)
	}

	if in.Semester != 1 && in.Semester != 2 {
		return nil, fmt.Errorf("Incorrect year format (semester should be 1 or 2): %d", in.Year)
	}

	courseSchedulesLecturesAmountOfStudents, err := s.getCourseScheduledLecturesAmountOfStudents(ctx, in.CourseName, in.Year, in.Semester)

	if err != nil {
		return nil, err
	}

	return courseSchedulesLecturesAmountOfStudents, nil
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
