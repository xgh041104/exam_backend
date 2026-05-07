package model

type Plan struct {
	PlanId      int    `json:"PlanId" db:"PlanId"`
	CourseId    int    `json:"CourseId" db:"CourseId"`
	PlanName    string `json:"PlanName" db:"PlanName"`
	CourseRatio int    `json:"CourseRatio" db:"CourseRatio"`
	ExamRatio   int    `json:"ExamRatio" db:"ExamRatio"`
	TrainRatio  int    `json:"TrainRatio" db:"TrainRatio"`
	TeacherId   int    `json:"TeacherId" db:"TeacherId"`
}

type PlanStudentViewarr struct {
	PlanStudentView
}

type PlanStudentCourseResult struct {

	//课程进度
	PlanStudentCourseResult StudyCourse
}

type PlanExamViewResultArr struct {
	//考试进度 【】
	PlanExamViewResultArr StudentExamInfo
}

type PlanQuestionArr struct {
	//训练进度 【】
	PlanQuestionArr []*QuestionrecordArr `json:"PlanQuestionArr" `
}

type PlanArr struct {
	Plan
	CourseName       string `json:"CourseName" db:"CourseName"`
	PlanStudentCount int    `json:"PlanStudentCount" db:"PlanStudentCount"`
	PlanExamCount    int    `json:"PlanExamCount" db:"PlanExamCount"`
	PlanTrainCount   int    `json:"PlanTrainCount" db:"PlanTrainCount"`
}

type PlanExam struct {
	PlanexamId    int `json:"PlanexamId" db:"PlanexamId"`
	PlanId        int `json:"PlanId" db:"PlanId"`
	ExamSessionId int `json:"ExamSessionId" db:"ExamSessionId"`
}

type PlanExamView struct {
	PlanExam
	ExamId int `json:"ExamId" db:"ExamId"`
}

type PlanExamViewName struct {
	PlanExamView
	ExamName string `json:"ExamName" db:"ExamName"`
}

type PlanExamViewResult struct {
	PlanExamView
	ExamName      string  `json:"ExamName" db:"ExamName"`
	TestPaperName string  `json:"TestPaperName" db:"TestPaperName"`
	PassScore     float64 `json:"PassScore" db:"PassScore"`
	StartExamTime string  `json:"StartExamTime" db:"StartExamTime"`
	EndExamTime   string  `json:"EndExamTime" db:"EndExamTime"`
	Score         float32 `json:"Score" db:"Score"`
}

type PlanStudent struct {
	PlanStudentId int `json:"PlanStudentId" db:"PlanStudentId"`
	PlanId        int `json:"PlanId" db:"PlanId"`
	StudentId     int `json:"StudentId" db:"StudentId"`
}

type PlanStudentView struct {
	PlanStudent
	TrueName       string `json:"TrueName" db:"TrueName"`
	StudentAccount string `json:"StudentAccount" db:"StudentAccount"`
}

type PlanTrain struct {
	PlanTrainId  int `json:"PlanTrainId" db:"PlanTrainId"`
	PlanId       int `json:"PlanId" db:"PlanId"`
	QuestionId   int `json:"QuestionId" db:"QuestionId"`
	QuestionType int `json:"QuestionType" db:"QuestionType"`
}
type PlanTrainView struct {
	PlanTrain
	QuestionName string `json:"QuestionName" db:"QuestionName"`
}

type Questionrecord struct {
	QuestionRecordId int    `json:"QuestionRecordId" db:"QuestionRecordId"`
	QuestionId       int    `json:"QuestionId" db:"QuestionId"`
	StudentId        int    `json:"StudentId" db:"StudentId"`
	CreateTime       string `json:"CreateTime" db:"CreateTime"`
	AnswerSteps      string `json:"AnswerSteps" db:"AnswerSteps"`
	TrueAnswer       string `json:"TrueAnswer" db:"TrueAnswer"`
	TrainScore       int    `json:"TrainScore" db:"TrainScore"`
	PlanId           int    `json:"PlanId" db:"PlanId"`
}

type QuestionrecordArr struct {
	PlanTrain
	QuestionName    string  `json:"QuestionName" db:"QuestionName"`
	QuestionContent string  `json:"QuestionContent" db:"QuestionContent"`
	AnswerSteps     string  `json:"AnswerSteps" db:"AnswerSteps"`
	TrainScore      float64 `json:"TrainScore" db:"TrainScore"`
}

type PlanStudentImg struct {
	PlanStudentImgId int64  `json:"PlanStudentImgId" db:"PlanStudentImgId"`
	PlanStudentId    int64  `json:"PlanStudentId" db:"PlanStudentId"`
	ImgPath          string `json:"ImgPath" db:"ImgPath"`
}
