package model

type College struct {
	Id          int    `json:"Id" db:"Id"`
	CollegeName string `json:"CollegeName" db:"CollegeName"`
	SchoolId    int    `json:"SchoolId" db:"SchoolId"`
	TeacherId   int    `json:"TeacherId" db:"TeacherId"`
}

type CollegeView struct {
	Id          int    `json:"Id" db:"Id"`
	CollegeName string `json:"CollegeName" db:"CollegeName"`
	SchoolId    int    `json:"SchoolId" db:"SchoolId"`
	TeacherId   int    `json:"TeacherId" db:"TeacherId"`
	SchoolName  string `json:"SchoolName" db:"SchoolName"`
}
