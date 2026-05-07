package model

type OperPractice struct {
	Id             int     `json:"Id" db:"Id"`
	QuestionId     int     `json:"QuestionId" db:"QuestionId"`
	StudentId      int     `json:"StudentId" db:"StudentId"`
	PracticeStep   string  `json:"PracticeStep" db:"PracticeStep"`
	PracticeAnswer string  `json:"PracticeAnswer" db:"PracticeAnswer"`
	PracticeScore  float64 `json:"PracticeScore" db:"PracticeScore"`
	CreateTime     string  `json:"CreateTime" db:"CreateTime"`
}

type OperPracticeView struct {
	StudentId  int    `json:"StudentId" db:"StudentId"`
	TrueName   string `json:"TrueName" db:"TrueName"`
	StudentNum int    `json:"StudentNum" db:"StudentNum"`
	MaxTime    string `json:"MaxTime" db:"MaxTime"`
}
