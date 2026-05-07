package model

type School struct {
	Id            int    `json:"Id" db:"Id"`
	SchoolName    string `json:"SchoolName" db:"SchoolName"`
	SchoolAddress string `json:"SchoolAddress" db:"SchoolAddress"`
}
