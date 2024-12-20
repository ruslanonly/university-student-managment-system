package service

type Lab2ReportItem struct {
	CourseName string `json:"course_name"`
	ClassName  string `json:"class_name"`
	Capacity   int    `json:"capacity"`
}

type In struct {
	CourseName  string `json:"course_name"`
	Year 		int `json:"year"`
	Semester    int `json:"semester"`
}

type Out struct {
	Report []Lab2ReportItem `json:"report"`
}
