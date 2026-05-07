package model

type Stand struct {
	Id        int    `json:"Id" db:"Id"`
	StandName string `json:"StandName" db:"StandName"`
	SchoolId  int    `json:"SchoolId" db:"SchoolId"`
	TeacherId int    `json:"TeacherId" db:"TeacherId"`
}

type StandId struct {
	StandId  int `json:"StandId" db:"StandId"`
	SchoolId int `json:"SchoolId" db:"SchoolId"`
}
