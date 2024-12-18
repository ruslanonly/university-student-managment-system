package service

import "time"

type ReportItem struct {
	StudentFullName string
	Attendance      float64
	PeriodStart     time.Time
	PeriodEnd       time.Time
	Phrase          string
}

type In struct {
	Phrase      string
	PeriodStart time.Time
	PeriodEnd   time.Time
}
