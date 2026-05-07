package model

type PatrolUser struct {
	PatrolUserId      int    `json:"patrolUserId" db:"patrolUserId"`
	PatrolUserAccount string `json:"patrolUserAccount" db:"patrolUserAccount"`
	PatrolUserPwd     string `json:"patrolUserPwd" db:"patrolUserPwd"`
}

type LoginLogArr struct {
	NowDay    []*LoginLog `json:"nowday" `
	Yesterday []*LoginLog `json:"yesterday" `
	Sevenday  []*LoginLog `json:"sevenday" `
}
type LoginLog struct {
	Hour     string `json:"hour" db:"hour"`
	Day      string `json:"day" db:"day"`
	Total    int    `json:"total" db:"total"`
	DateType int    `json:"dateType" db:"datetype"`
}

type PatroExamPlan struct {
	Id            int64  `json:"id" db:"Id"`
	StartTime     string `json:"startTime" db:"StartTime"`
	EndTime       string `json:"endTime" db:"EndTime"`
	TestPaperName string `json:"testPaperName" db:"TestPaperName"`
	ExamName      string `json:"examName" db:"ExamName"`
	Status        string `json:"status" db:"Status"`
	Personcount   int64  `json:"personCount" db:"Personcount"`
}
type RealTimeExamSituation struct {
	SumCount    int64 `json:"sumCount",db:"sumCount"`
	DoneCount   int64 `json:"doneCount",db:"doneCount"`
	UnDoneCount int64 `json:"unDoneCount",db:"undoneCount"`
}

type StandPostion struct {
	Id          int    `json:"id" db:"Id"`
	StandName   string `json:"standName" db:"standName"`
	SchoolId    int    `json:"schoolId" db:"schoolId"`
	TeacherId   int    `json:"teacherId" db:"TeacherId"`
	Postion     string `json:"postion" db:"postion"`
	PersonCount int64  `json:"personCount",db:"personCount"`
	DoneCount   int64  `json:"doneCount",db:"doneCount"`
	UnDoneCount int64  `json:"unDoneCount",db:"undoneCount"`
}
