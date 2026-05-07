package model

type Question struct {
	QuestionId       int    `json:"QuestionId" db:"QuestionId"`
	SchoolId         int    `json:"SchoolId" db:"SchoolId"`
	QuestionPoolId   int    `json:"QuestionPoolId" db:"QuestionPoolId"`
	QuestionName     string `json:"QuestionName" db:"QuestionName"`
	QuestionType     int    `json:"QuestionType" db:"QuestionType"`
	QuestionContent  string `json:"QuestionContent" db:"QuestionContent"`
	Digree           int    `json:"Digree" db:"Digree"`
	MajorID          int    `json:"MajorID" db:"MajorID"`
	CollegeId        int    `json:"CollegeId" db:"CollegeId"`
	CourseId         int    `json:"CourseId" db:"CourseId"`
	Answer           string `json:"Answer" db:"Answer"`
	TeacherId        int    `json:"TeacherId" db:"TeacherId"`
	QuestionCategory int    `json:"QuestionCategory" db:"QuestionCategory"`
}

type QuestionView struct {
	QuestionId       int    `json:"QuestionId" db:"QuestionId"`
	SchoolId         int    `json:"SchoolId" db:"SchoolId"`
	QuestionPoolId   int    `json:"QuestionPoolId" db:"QuestionPoolId"`
	QuestionName     string `json:"QuestionName" db:"QuestionName"`
	QuestionType     int    `json:"QuestionType" db:"QuestionType"`
	QuestionContent  string `json:"QuestionContent" db:"QuestionContent"`
	Digree           int    `json:"Digree" db:"Digree"`
	MajorID          int    `json:"MajorID" db:"MajorID"`
	CollegeId        int    `json:"CollegeId" db:"CollegeId"`
	CourseId         int    `json:"CourseId" db:"CourseId"`
	Answer           string `json:"Answer" db:"Answer"`
	MajorName        string `json:"MajorName" db:"MajorName"`
	CollegeName      string `json:"CollegeName" db:"CollegeName"`
	CourseName       string `json:"CourseName" db:"CourseName"`
	QuestionCategory int    `json:"QuestionCategory" db:"QuestionCategory"`
}

type QuestionViewFile struct {
	QuestionId       int         `json:"QuestionId" db:"QuestionId"`
	SchoolId         int         `json:"SchoolId" db:"SchoolId"`
	QuestionPoolId   int         `json:"QuestionPoolId" db:"QuestionPoolId"`
	QuestionName     string      `json:"QuestionName" db:"QuestionName"`
	QuestionType     int         `json:"QuestionType" db:"QuestionType"`
	QuestionContent  string      `json:"QuestionContent" db:"QuestionContent"`
	Digree           int         `json:"Digree" db:"Digree"`
	MajorID          int         `json:"MajorID" db:"MajorID"`
	CollegeId        int         `json:"CollegeId" db:"CollegeId"`
	CourseId         int         `json:"CourseId" db:"CourseId"`
	Answer           string      `json:"Answer" db:"Answer"`
	MajorName        string      `json:"MajorName" db:"MajorName"`
	CollegeName      string      `json:"CollegeName" db:"CollegeName"`
	CourseName       string      `json:"CourseName" db:"CourseName"`
	QuestionCategory int         `json:"QuestionCategory" db:"QuestionCategory"`
	FileInfo         []*FileInfo `json:"FileInfo" `
}

type QuestionEdit struct {
	QuestionId       int    `json:"QuestionId" db:"QuestionId"`
	SchoolId         int    `json:"SchoolId" db:"SchoolId"`
	QuestionPoolId   int    `json:"QuestionPoolId" db:"QuestionPoolId"`
	QuestionName     string `json:"QuestionName" db:"QuestionName"`
	QuestionType     int    `json:"QuestionType" db:"QuestionType"`
	QuestionContent  string `json:"QuestionContent" db:"QuestionContent"`
	Digree           int    `json:"Digree" db:"Digree"`
	MajorID          int    `json:"MajorID" db:"MajorID"`
	CollegeId        int    `json:"CollegeId" db:"CollegeId"`
	CourseId         int    `json:"CourseId" db:"CourseId"`
	Answer           string `json:"Answer" db:"Answer"`
	RemoveFile       []int  `json:"RemoveFile" `
	QuestionCategory int    `json:"QuestionCategory" db:"QuestionCategory"`
}

type TestPaperQuestionViewFile struct {
	QuestionId      int         `json:"QuestionId" db:"QuestionId"`
	QuestionName    string      `json:"QuestionName" db:"QuestionName"`
	QuestionPoolId  int         `json:"QuestionPoolId" db:"QuestionPoolId"`
	QuestionType    int         `json:"QuestionType" db:"QuestionType"`
	QuestionContent string      `json:"QuestionContent" db:"QuestionContent"`
	QuestionScore   float32     `json:"QuestionScore" db:"QuestionScore"`
	Digree          int         `json:"Digree" db:"Digree"`
	Answer          string      `json:"Answer" db:"Answer"`
	FileInfo        []*FileInfo `json:"FileInfo" `
}
