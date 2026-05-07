package model

type Teacher struct {
	TeacherId       int    `json:"TeacherId" db:"TeacherId"`
	SchoolId        int    `json:"SchoolId" db:"SchoolId"`
	TeacherAccount  string `json:"TeacherAccount" db:"TeacherAccount"`
	TeacherPassword string `json:"TeacherPassword" db:"TeacherPassword"`
	TeacherName     string `json:"TeacherName" db:"TeacherName"`
	Sex             int    `json:"Sex" db:"Sex"`
	PhoneNumber     string `json:"PhoneNumber" db:"PhoneNumber"`
	Email           string `json:"Email" db:"Email"`
	TeacherTitle    string `json:"TeacherTitle" db:"TeacherTitle"`
}

type TeacherView struct {
	TeacherId      int    `json:"TeacherId" db:"TeacherId"`
	SchoolId       int    `json:"SchoolId" db:"SchoolId"`
	TeacherAccount string `json:"TeacherAccount" db:"TeacherAccount"`
	TeacherName    string `json:"TeacherName" db:"TeacherName"`
	Sex            int    `json:"Sex" db:"Sex"`
	PhoneNumber    string `json:"PhoneNumber" db:"PhoneNumber"`
	Email          string `json:"Email" db:"Email"`
	TeacherTitle   string `json:"TeacherTitle" db:"TeacherTitle"`
	SchoolName     string `json:"SchoolName" db:"SchoolName"`
	SchoolAddress  string `json:"SchoolAddress" db:"SchoolAddress"`
}
