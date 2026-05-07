package model

type Major struct {
	MajorId   int    `json:"MajorId" db:"MajorId"`
	SchoolId  int    `json:"SchoolId" db:"SchoolId"`
	CollegeId int    `json:"CollegeId" db:"CollegeId"`
	MajorName string `json:"MajorName" db:"MajorName"`
	TeacherId int    `json:"TeacherId" db:"TeacherId"`
}

type MajorView struct {
	MajorId     int    `json:"MajorId" db:"MajorId"`
	SchoolId    int    `json:"SchoolId" db:"SchoolId"`
	CollegeId   int    `json:"CollegeId" db:"CollegeId"`
	MajorName   string `json:"MajorName" db:"MajorName"`
	TeacherId   int    `json:"TeacherId" db:"TeacherId"`
	CollegeName string `json:"CollegeName" db:"CollegeName"`
}
