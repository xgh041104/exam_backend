package model

type CheatWarning struct {
	CheatWarningId int64  `json:"cheatWarningId" db:"cheatWarningId"`
	StudentId      int64  `json:"studentId" db:"studentId"`
	StudentName    string `json:"studentName" db:"studentName"`
	StudentImgPath string `json:"studentImgPath" db:"studentImgPath"`
	StudentImgNum  int64  `json:"studentImgNum" db:"studentImgNum"`
	CreateTime     int64  `json:"createTime" db:"createTime"`
}

type FaceMonitor struct {
	FaceMonitorId int64  `json:"faceMonitorId" db:"faceMonitorId"`
	StudentId     int64  `json:"studentId" db:"studentId"`
	ImgPath       string `json:"imgPath" db:"imgPath"`
}
