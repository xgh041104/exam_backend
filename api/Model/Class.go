package model

type Class struct {
	Id        int    `json:"Id" db:"Id"`
	MajorId   int    `json:"MajorId" db:"MajorId"`
	SchoolId  int    `json:"SchoolId" db:"SchoolId"`
	ClassName string `json:"ClassName" db:"ClassName"`
	TeacherId int    `json:"TeacherId" db:"TeacherId"`
}

type ClassView struct {
	Id          int    `json:"Id" db:"Id"`
	MajorId     int    `json:"MajorId" db:"MajorId"`
	SchoolId    int    `json:"SchoolId" db:"SchoolId"`
	ClassName   string `json:"ClassName" db:"ClassName"`
	TeacherId   int    `json:"TeacherId" db:"TeacherId"`
	SchoolName  string `json:"SchoolName" db:"SchoolName"`
	MajorName   string `json:"MajorName" db:"MajorName"`
	CollegeId   int    `json:"CollegeId" db:"CollegeId"`
	CollegeName string `json:"CollegeName" db:"CollegeName"`
}

type ClassStudent struct {
	Id        int         `json:"Id" db:"Id"`
	ClassName string      `json:"ClassName" db:"ClassName"`
	Students  []*StudentS `json:"Students"  `
}

type StudentS struct {
	Id       int    `json:"Id" db:"Id"`
	TrueName string `json:"TrueName" db:"TrueName"`
}
