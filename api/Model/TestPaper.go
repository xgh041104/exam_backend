package model

type TestPaper struct {
	Id                int                  `json:"Id" db:"Id"`
	TestPaperName     string               `json:"TestPaperName" db:"TestPaperName"`
	ExamDuration      int                  `json:"ExamDuration" db:"ExamDuration"`
	FullMarks         float32              `json:"FullMarks" db:"FullMarks"`
	PassScore         float32              `json:"PassScore" db:"PassScore"`
	TestPaperType     int                  `json:"TestPaperType" db:"TestPaperType"`
	CollegeId         int                  `json:"CollegeId" db:"CollegeId"`
	MajorId           int                  `json:"MajorId" db:"MajorId"`
	CourseId          int                  `json:"CourseId" db:"CourseId"`
	QuestionScoreJson []*QuestionScoreJson `json:"QuestionScoreJson"  `
	TeacherId         int                  `json:"TeacherId" db:"TeacherId"`
	SchoolId          int                  `json:"SchoolId" db:"SchoolId"`
}

type TestPaperView struct {
	Id            int     `json:"Id" db:"Id"`
	TestPaperName string  `json:"TestPaperName" db:"TestPaperName"`
	ExamDuration  int     `json:"ExamDuration" db:"ExamDuration"`
	FullMarks     float32 `json:"FullMarks" db:"FullMarks"`
	PassScore     float32 `json:"PassScore" db:"PassScore"`
	TestPaperType int     `json:"TestPaperType" db:"TestPaperType"`
	CollegeId     int     `json:"CollegeId" db:"CollegeId"`
	MajorId       int     `json:"MajorId" db:"MajorId"`
	CourseId      int     `json:"CourseId" db:"CourseId"`
	TeacherId     int     `json:"TeacherId" db:"TeacherId"`
	SchoolId      int     `json:"SchoolId" db:"SchoolId"`
	MajorName     string  `json:"MajorName" db:"MajorName"`
	CollegeName   string  `json:"CollegeName" db:"CollegeName"`
	CourseName    string  `json:"CourseName" db:"CourseName"`
}

type TestPaperDetails struct {
	Id                        int                          `json:"Id" db:"Id"`
	TestPaperName             string                       `json:"TestPaperName" db:"TestPaperName"`
	ExamDuration              int                          `json:"ExamDuration" db:"ExamDuration"`
	FullMarks                 float32                      `json:"FullMarks" db:"FullMarks"`
	PassScore                 float32                      `json:"PassScore" db:"PassScore"`
	TestPaperType             int                          `json:"TestPaperType" db:"TestPaperType"`
	CollegeId                 int                          `json:"CollegeId" db:"CollegeId"`
	MajorId                   int                          `json:"MajorId" db:"MajorId"`
	CourseId                  int                          `json:"CourseId" db:"CourseId"`
	TeacherId                 int                          `json:"TeacherId" db:"TeacherId"`
	SchoolId                  int                          `json:"SchoolId" db:"SchoolId"`
	MajorName                 string                       `json:"MajorName" db:"MajorName"`
	CollegeName               string                       `json:"CollegeName" db:"CollegeName"`
	CourseName                string                       `json:"CourseName" db:"CourseName"`
	SchoolName                string                       `json:"SchoolName" db:"SchoolName"`
	TeacherName               string                       `json:"TeacherName" db:"TeacherName"`
	TestPaperQuestionTypeOver []*QuestionTypeOver          `json:"TestPaperQuestionTypeOver" `
	TestPaperQuestionViewFile []*TestPaperQuestionViewFile `json:"TestPaperQuestionViewFile" `
}

type QuestionTypeOver struct {
	QuestionType  int     `json:"QuestionType" db:"QuestionType"`
	QuestionIdNum int     `json:"QuestionIdNum" db:"QuestionIdNum"`
	QuestionScore float32 `json:"QuestionScore" db:"QuestionScore"`
}

type QuestionScoreJson struct {
	QuestionType   int              `json:"QuestionType" db:"QuestionType"`
	QuestionArr    []*QuestionScore `json:"QuestionArr"  `
	FullMarksRatio int              `json:"FullMarksRatio" db:"FullMarksRatio"`
	// QuestionIdNum  int              `json:"QuestionIdNum" db:"QuestionIdNum"`
}

type QuestionScore struct {
	QuestionId   int     `json:"QuestionId" db:"QuestionId"`
	QuestionName string  `json:"QuestionName" db:"QuestionName"`
	Score        float32 `json:"Score" db:"Score"`
}

type TestPaperQuestion struct {
	Id            int     `json:"Id" db:"Id"`
	QuestionId    int     `json:"QuestionId" db:"QuestionId"`
	QuestionType  int     `json:"QuestionType" db:"QuestionType"`
	QuestionName  string  `json:"QuestionName" db:"QuestionName"`
	QuestionScore float32 `json:"QuestionScore" db:"QuestionScore"`
	TestPaperId   int     `json:"TestPaperId" db:"TestPaperId"`
}
