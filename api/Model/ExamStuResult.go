package model

type ExamStuResult struct {
	Id            int     `json:"Id" db:"Id"`
	StudentId     int     `json:"StudentId" db:"StudentId"`
	ExamId        int     `json:"ExamId" db:"ExamId"`
	ExamSessionId int     `json:"ExamSessionId" db:"ExamSessionId"`
	StartExamTime string  `json:"StartExamTime" db:"StartExamTime"`
	EndExamTime   string  `json:"EndExamTime" db:"EndExamTime"`
	Score         float32 `json:"Score" db:"Score"`
	ExamStatus    int     `json:"ExamStatus" db:"ExamStatus"`
	ExamType      int     `json:"ExamType" db:"ExamType"`
}

type ExamAnswerSheet struct {
	Id          int     `json:"Id" db:"Id"`
	ExamId      int     `json:"ExamId" db:"ExamId"`
	TestPaperId int     `json:"TestPaperId" db:"TestPaperId"`
	StudentId   int     `json:"StudentId" db:"StudentId"`
	QuestionId  int     `json:"QuestionId" db:"QuestionId"`
	AnswerScore float32 `json:"AnswerScore" db:"AnswerScore"`
	AnswerSteps string  `json:"AnswerSteps" db:"AnswerSteps"`
}

type ExamScoreStuView struct {
	Id            int          `json:"Id" db:"Id"`
	TrueName      string       `json:"TrueName" db:"TrueName"`
	IDNumber      string       `json:"IDNumber" db:"IDNumber"`
	StudentId     int          `json:"StudentId" db:"StudentId"`
	ExamId        int          `json:"ExamId" db:"ExamId"`
	ExamSessionId int          `json:"ExamSessionId" db:"ExamSessionId"`
	StartExamTime string       `json:"StartExamTime" db:"StartExamTime"`
	EndExamTime   string       `json:"EndExamTime" db:"EndExamTime"`
	UseTime       float64      `json:"UseTime"  `
	Score         int          `json:"Score" db:"Score"`
	ExamStatus    int          `json:"ExamStatus"  db:"ExamStatus" `
	ExamImageArr  []*ExamImage `json:"ExamImageArr"   `
	ExamType      int          `json:"ExamType"  db:"ExamType" `
}

type SumitExamEntity struct {
	StudentId      int            `json:"StudentId" db:"StudentId"`
	ExamId         int            `json:"ExamId" db:"ExamId"`
	ExamSessionId  int            `json:"ExamSessionId" db:"ExamSessionId"`
	TestPaperId    int            `json:"TestPaperId" db:"TestPaperId"`
	StartExamTime  string         `json:"StartExamTime" db:"StartExamTime"`
	EndExamTime    string         `json:"EndExamTime" db:"EndExamTime"`
	Score          float32        `json:"Score" db:"Score"`
	IsReTest       int            `json:"IsReTest"  `
	AnswerSheetArr []*AnswerSheet `json:"AnswerSheetArr" `
}

type AnswerSheet struct {
	QuestionId  int     `json:"QuestionId" db:"QuestionId"`
	AnswerScore float32 `json:"AnswerScore" db:"AnswerScore"`
	AnswerSteps string  `json:"AnswerSteps" db:"AnswerSteps"`
	IsTrue      int     `json:"IsTrue" db:"IsTrue"`
}
