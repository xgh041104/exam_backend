package model

type ExamNotice struct {
	Id         int    `json:"Id" db:"Id"`
	ExamId     int    `json:"ExamId" db:"ExamId"`
	Title      string `json:"Title" db:"Title"`
	SchoolName string `json:"SchoolName" db:"SchoolName"`
	CourseName string `json:"CourseName" db:"CourseName"`
	CourseCode string `json:"CourseCode" db:"CourseCode"`
	Context    string `json:"Context" db:"Context"`
}
