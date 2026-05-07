package model

type StandUser struct {
	StandUserId      int    `json:"StandUserId" db:"StandUserId"`
	StandUserName    string `json:"StandUserName" db:"StandUserName"`
	StandUserAccount string `json:"StandUserAccount" db:"StandUserAccount"`
	StandUserPwd     string `json:"StandUserPwd" db:"StandUserPwd"`
	StandId          int    `json:"StandId" db:"StandId"`
}
