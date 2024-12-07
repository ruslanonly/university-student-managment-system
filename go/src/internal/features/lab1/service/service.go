package service

import (
	"context"
	"time"
)

type StudentAttendanceItem struct {
	StudentID  int
	Attendance int
}

type Service struct{}

func (s *Service) getCourseIDsByPhrase(ctx context.Context, phrase string) ([]int, error) {
	return make([]int, 0), nil
}

func (s *Service) getStudentIDsByAttendance(ctx context.Context, courseIDs []int, periodStart, periodEnd time.Time) ([]int, error) {
	return make([]int, 0), nil
}

func (s *Service) getStudentWithBaddestAttendance(ctx context.Context, studentIDs []int, periodStart, periodEnd time.Time) ([]StudentAttendanceItem, error) {
	return []StudentAttendanceItem{}, nil
}

func (s *Service) getStudentFullName(ctx context.Context, studentID int) (string, error) {
	return "", nil
}

func (s *Service) Execute(ctx context.Context, in In) ([]ReportItem, error) {
	courseIDs, err := s.getCourseIDsByPhrase(ctx, in.Phrase)

	if err != nil {
		return nil, err
	}

	studentIDs, err := s.getStudentIDsByAttendance(ctx, courseIDs, in.PeriodStart, in.PeriodEnd)

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

func New() *Service {
	return &Service{}
}
