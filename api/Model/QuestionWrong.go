package model

type QuestionWrong struct {
	Id          int    `json:"Id" db:"Id"`
	QuestionId  int    `json:"QuestionId" db:"QuestionId"`
	StudentId   int    `json:"StudentId" db:"StudentId"`
	CreateTime  string `json:"CreateTime" db:"CreateTime"`
	AnswerSteps string `json:"AnswerSteps" db:"AnswerSteps"`
	TrueAnswer  string `json:"TrueAnswer" db:"TrueAnswer"`
}

type QuestionWrongView struct {
	Id                 int               `json:"Id" db:"Id"`
	QuestionId         int               `json:"QuestionId" db:"QuestionId"`
	QuestionName       string            `json:"QuestionName" db:"QuestionName"`
	StudentId          int               `json:"StudentId" db:"StudentId"`
	CreateTime         string            `json:"CreateTime" db:"CreateTime"`
	AnswerSteps        string            `json:"AnswerSteps" db:"AnswerSteps"`
	TrueAnswer         string            `json:"TrueAnswer" db:"TrueAnswer"`
	QuestionViewEntity *QuestionViewFile `json:"QuestionViewEntity" `
}

type QuestionWrongIdArr struct {
	Id int `json:"Id" db:"Id"`
}
