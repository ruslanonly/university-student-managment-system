package service

type In struct {
	GroupName           string `json:"group_name"`
	GroupEnrollmentYear string `json:"group_enrollment_year"`
	DisciplineTag       string `json:"discipline_tag"`
}

type ReportItem struct {
	GroupName           string `json:"group_name"`
	GroupEnrollmentYear string `json:"group_enrollment_year"`
	StudentGradeBookID  string `json:"student_grade_book_id"`
	StudentFullName     string `json:"student_full_name"`
	CourseName          string `json:"course_name"`
	PlannedQuantity     string `json:"planned_quantity"`
	VisitedQuantity     string `json:"visited_quantity"`
}

type Out struct {
	Report []ReportItem `json:"report"`
}
