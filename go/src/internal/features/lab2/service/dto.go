package service

type ReportItem struct {
	CourseName string `json:"course_name"`
	ClassName  string `json:"class_name"`
	Capacity   int    `json:"capacity"`
}

type In struct {
	CourseName          string `json:"course_name"`
	GroupEnrollmentYear string `json:"group_enrollment_year"`
	GroupSemester       string `json:"group_semester"`
}

type Out struct {
	Report []ReportItem `json:"report"`
}
