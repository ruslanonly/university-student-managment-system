package service

import "time"

type ReportItem struct {
	StudentFullName string
	Attendance      int
	PeriodStart     time.Time
	PeriodEnd       time.Time
	Phrase          string
}

type In struct {
	Phrase      string
	PeriodStart time.Time
	PeriodEnd   time.Time
}
