package model

type Exam struct {
	Id           int    `json:"Id" db:"Id"`
	SchoolId     int    `json:"SchoolId" db:"SchoolId"`
	ExamName     string `json:"ExamName" db:"ExamName"`
	ExamDescribe string `json:"ExamDescribe" db:"ExamDescribe"`
	ExamStatus   int    `json:"ExamStatus" db:"ExamStatus"`
	FaceVerify   int    `json:"FaceVerify" db:"FaceVerify"`
	TeacherId    int    `json:"TeacherId" db:"TeacherId"`
	IsReset      int    `json:"IsReset" db:"IsReset"`
}

type ExamView struct {
	Id             int                `json:"Id" db:"Id"`
	SchoolId       int                `json:"SchoolId" db:"SchoolId"`
	ExamName       string             `json:"ExamName" db:"ExamName"`
	ExamDescribe   string             `json:"ExamDescribe" db:"ExamDescribe"`
	ExamStatus     int                `json:"ExamStatus" db:"ExamStatus"`
	FaceVerify     int                `json:"FaceVerify" db:"FaceVerify"`
	State          int                `json:"State"`
	ReviewFlag     int                `json:"ReviewFlag"  db:"ReviewFlag"`
	ExamSessionArr []*ExamSessionView `json:"ExamSessionArr"  `
	ExamStudentArr []*ExamStudentView `json:"ExamStudentArr"  `
}

type ExamSession struct {
	Id          int    `json:"Id" db:"Id"`
	ExamId      int    `json:"ExamId" db:"ExamId"`
	StartTime   string `json:"StartTime" db:"StartTime"`
	EndTime     string `json:"EndTime" db:"EndTime"`
	TestPaperId int    `json:"TestPaperId" db:"TestPaperId"`
}

type ExamSessionView struct {
	Id            int     `json:"Id" db:"Id"`
	ExamId        int     `json:"ExamId" db:"ExamId"`
	StartTime     string  `json:"StartTime" db:"StartTime"`
	EndTime       string  `json:"EndTime" db:"EndTime"`
	TestPaperId   int     `json:"TestPaperId" db:"TestPaperId"`
	TestPaperName string  `json:"TestPaperName" db:"TestPaperName"`
	ExamDuration  int     `json:"ExamDuration" db:"ExamDuration"`
	FullMarks     float32 `json:"FullMarks" db:"FullMarks"`
	PassScore     float32 `json:"PassScore" db:"PassScore"`
	State         int     `json:"State"`
}

type ExamStudentView struct {
	StudentId int    `json:"StudentId" db:"StudentId"`
	TrueName  string `json:"TrueName" db:"TrueName"`
}
type ExamStudent struct {
	Id                 int   `json:"Id" db:"Id"`
	ExamId             int   `json:"ExamId" db:"ExamId"`
	AddStudentIdArr    []int `json:"AddStudentIdArr"  `
	RemoveStudentIdArr []int `json:"RemoveStudentIdArr"  `
}

type ExamStudentModel struct {
	ExamId    int `json:"ExamId" db:"ExamId"`
	StudentId int `json:"StudentId" db:"StudentId"`
}

type ExamScoreView struct {
	ExamSessionId    int    `json:"ExamSessionId" db:"ExamSessionId"`
	ExamId           int    `json:"ExamId" db:"ExamId"`
	ExamSession      string `json:"ExamSession" db:"ExamSession"`
	MajorName        string `json:"MajorName" db:"MajorName"`
	Score            string `json:"Score" db:"Score"`
	StartTime        string `json:"StartTime" db:"StartTime"`
	AvgScore         int    `json:"AvgScore" db:"AvgScore"`
	UnexaminedNum    int    `json:"UnexaminedNum" db:"UnexaminedNum"`
	ExaminedNum      int    `json:"ExaminedNum" db:"ExaminedNum"`
	ExamDuration     int    `json:"ExamDuration" db:"ExamDuration"`
	ExamPeopleSumNum int    `json:"ExamPeopleSumNum"  `
}

type ExamRetest struct {
	Id               int     `json:"Id" db:"Id"`
	StudentIdArr     []int   `json:"StudentIdArr"  `
	OldExamSessionId int     `json:"OldExamSessionId" db:"OldExamSessionId"`
	OldExamId        int     `json:"OldExamId" db:"OldExamId"`
	StartExamTime    string  `json:"StartExamTime" db:"StartExamTime"`
	EndExamTime      string  `json:"EndExamTime" db:"EndExamTime"`
	Score            float32 `json:"Score" db:"Score"`
	Status           int     `json:"Status" db:"Status"`
}

type ExamRetestView struct {
	Id               int     `json:"Id" db:"Id"`
	StudentId        int     `json:"StudentId"  db:"StudentId" `
	TrueName         string  `json:"TrueName" db:"TrueName"`
	TestPaperName    string  `json:"TestPaperName" db:"TestPaperName"`
	OldExamSessionId int     `json:"OldExamSessionId" db:"OldExamSessionId"`
	OldExamId        int     `json:"OldExamId" db:"OldExamId"`
	StartExamTime    string  `json:"StartExamTime" db:"StartExamTime"`
	EndExamTime      string  `json:"EndExamTime" db:"EndExamTime"`
	Score            float32 `json:"Score" db:"Score"`
	Status           int     `json:"Status" db:"Status"`
}

type StudentExamInfo struct {
	// 	-- 0 已过期
	// -- 1 已完成
	// -- 2 未参加考试
	// -- 3 考试未开始
	// -- 4 未补考
	Id                   int     `json:"Id" db:"Id"`
	StudentId            int     `json:"StudentId"  db:"StudentId" `
	ExamId               int     `json:"ExamId" db:"ExamId"`
	ExamName             string  `json:"ExamName" db:"ExamName"`
	ExamSessionId        int     `json:"ExamSessionId" db:"ExamSessionId"`
	StartExamTime        string  `json:"StartExamTime" db:"StartExamTime"`
	EndExamTime          string  `json:"EndExamTime" db:"EndExamTime"`
	Score                float32 `json:"Score" db:"Score"`
	ExamStatus           int     `json:"ExamStatus" db:"ExamStatus"`
	ExamType             int     `json:"ExamType" db:"ExamType"`
	ExamZT               int     `json:"ExamZT" db:"ExamZT"`
	MajorId              int     `json:"MajorId" db:"MajorId"`
	CourseId             int     `json:"CourseId" db:"CourseId"`
	MajorName            string  `json:"MajorName" db:"MajorName"`
	CourseName           string  `json:"CourseName" db:"CourseName"`
	FullMarks            int     `json:"FullMarks" db:"FullMarks"`
	PassScore            int     `json:"PassScore" db:"PassScore"`
	SessionNum           int     `json:"SessionNum" db:"SessionNum"`
	ExamDuration         int     `json:"ExamDuration" db:"ExamDuration"`
	QuestionNum          int     `json:"QuestionNum" db:"QuestionNum"`
	ResetStartExamTime   string  `json:"ResetStartExamTime" db:"ResetStartExamTime"`
	ResetEndExamTime     string  `json:"ResetEndExamTime" db:"ResetEndExamTime"`
	SessionStartExamTime string  `json:"SessionStartExamTime" db:"SessionStartExamTime"`
	SessionEndExamTime   string  `json:"SessionEndExamTime" db:"SessionEndExamTime"`
	FaceVerify           int     `json:"FaceVerify" db:"FaceVerify"`
}

type StudentExamDetails struct {
	Id                        int                          `json:"Id" db:"Id"`
	StudentId                 int                          `json:"StudentId"  db:"StudentId" `
	ExamId                    int                          `json:"ExamId" db:"ExamId"`
	ExamName                  string                       `json:"ExamName" db:"ExamName"`
	ExamSessionId             int                          `json:"ExamSessionId" db:"ExamSessionId"`
	TestPaperId               int                          `json:"TestPaperId" db:"TestPaperId"`
	TestPaperName             string                       `json:"TestPaperName" db:"TestPaperName"`
	TestPaperType             int                          `json:"TestPaperType" db:"TestPaperType"`
	ExamType                  int                          `json:"ExamType" db:"ExamType"`
	MajorName                 string                       `json:"MajorName" db:"MajorName"`
	CourseName                string                       `json:"CourseName" db:"CourseName"`
	FullMarks                 int                          `json:"FullMarks" db:"FullMarks"`
	PassScore                 int                          `json:"PassScore" db:"PassScore"`
	ExamDuration              int                          `json:"ExamDuration" db:"ExamDuration"`
	SessionNum                int                          `json:"SessionNum" db:"SessionNum"`
	SchoolName                string                       `json:"SchoolName" db:"SchoolName"`
	TeacherName               string                       `json:"TeacherName" db:"TeacherName"`
	FaceVerify                int                          `json:"FaceVerify" db:"FaceVerify"`
	AnswerSheetarr            []*AnswerSheet               `json:"AnswerSheetarr" `
	TestPaperQuestionTypeOver []*QuestionTypeOver          `json:"TestPaperQuestionTypeOver" `
	TestPaperQuestionViewFile []*TestPaperQuestionViewFile `json:"TestPaperQuestionViewFile" `
}

type StudentExamAnswerDetails struct {
	Id                        int                          `json:"Id" db:"Id"`
	StudentId                 int                          `json:"StudentId"  db:"StudentId" `
	ExamId                    int                          `json:"ExamId" db:"ExamId"`
	ExamName                  string                       `json:"ExamName" db:"ExamName"`
	ExamSessionId             int                          `json:"ExamSessionId" db:"ExamSessionId"`
	TestPaperId               int                          `json:"TestPaperId" db:"TestPaperId"`
	TestPaperName             string                       `json:"TestPaperName" db:"TestPaperName"`
	TestPaperType             int                          `json:"TestPaperType" db:"TestPaperType"`
	ExamType                  int                          `json:"ExamType" db:"ExamType"`
	MajorName                 string                       `json:"MajorName" db:"MajorName"`
	CourseName                string                       `json:"CourseName" db:"CourseName"`
	FullMarks                 int                          `json:"FullMarks" db:"FullMarks"`
	PassScore                 int                          `json:"PassScore" db:"PassScore"`
	ExamDuration              int                          `json:"ExamDuration" db:"ExamDuration"`
	SessionNum                int                          `json:"SessionNum" db:"SessionNum"`
	SchoolName                string                       `json:"SchoolName" db:"SchoolName"`
	TeacherName               string                       `json:"TeacherName" db:"TeacherName"`
	Score                     float32                      `json:"Score" db:"Score"`
	AnswerSheetarr            []*AnswerSheet               `json:"AnswerSheetarr" `
	StartExamTime             string                       `json:"StartExamTime" db:"StartExamTime"`
	EndExamTime               string                       `json:"EndExamTime" db:"EndExamTime"`
	TestPaperQuestionTypeOver []*QuestionTypeOver          `json:"TestPaperQuestionTypeOver" `
	TestPaperQuestionViewFile []*TestPaperQuestionViewFile `json:"TestPaperQuestionViewFile" `
}

type DelExam struct {
	ExamId int `json:"ExamId" db:"ExamId"`
}

type ExamImage struct {
	Id            int    `json:"Id" db:"Id"`
	ExamId        int    `json:"ExamId" db:"ExamId"`
	ExamSessionId int    `json:"ExamSessionId" db:"ExamSessionId"`
	StudentId     int    `json:"StudentId" db:"StudentId"`
	ImagePath     string `json:"ImagePath" db:"ImagePath"`
	CreateTime    string `json:"CreateTime" db:"CreateTime"`
}

type ExportExam struct {
	StudentId     int     `json:"StudentId" db:"StudentId"`
	ExamId        int     `json:"ExamId" db:"ExamId"`
	ExamSessionId int     `json:"ExamSessionId" db:"ExamSessionId"`
	Score         float32 `json:"Score" db:"Score"`
	CourseCode    string  `json:"CourseCode" db:"CourseCode"`
}
