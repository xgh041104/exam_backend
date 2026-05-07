package model

type Admin struct {
	Id            int    `json:"Id" db:"Id"`
	AdminName     string `json:"AdminName" db:"AdminName"`
	AdminAccount  string `json:"AdminAccount" db:"AdminAccount"`
	AdminPassword string `json:"AdminPassword" db:"AdminPassword"`
}
