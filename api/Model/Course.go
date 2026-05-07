package model

type Course struct {
	Id                int                    `json:"Id" db:"Id"`
	CourseName        string                 `json:"CourseName" db:"CourseName"`
	Digest            string                 `json:"Digest" db:"Digest"`
	SchoolId          int                    `json:"SchoolId" db:"SchoolId"`
	CollegeId         int                    `json:"CollegeId" db:"CollegeId"`
	MajorId           int                    `json:"MajorId" db:"MajorId"`
	FileId            int                    `json:"FileId" db:"FileId"`
	TeacherId         int                    `json:"TeacherId" db:"TeacherId"`
	CourseStudentJson CourseClassStudentJson `json:"CourseStudentJson" db:"CourseStudentJson"`
	CourseStartTime   string                 `json:"CourseStartTime" db:"CourseStartTime"`
	CourseEndTime     string                 `json:"CourseEndTime" db:"CourseEndTime"`
	Status            int                    `json:"Status" db:"Status"`
	CourseCode        string                 `json:"CourseCode" db:"CourseCode"`
}

type CourseView struct {
	Id                int    `json:"Id" db:"Id"`
	CourseName        string `json:"CourseName" db:"CourseName"`
	Digest            string `json:"Digest" db:"Digest"`
	SchoolId          int    `json:"SchoolId" db:"SchoolId"`
	CollegeId         int    `json:"CollegeId" db:"CollegeId"`
	MajorId           int    `json:"MajorId" db:"MajorId"`
	FileId            int    `json:"FileId" db:"FileId"`
	TeacherId         int    `json:"TeacherId" db:"TeacherId"`
	CourseStudentJson string `json:"CourseStudentJson" db:"CourseStudentJson"`
	CourseStartTime   string `json:"CourseStartTime" db:"CourseStartTime"`
	CourseEndTime     string `json:"CourseEndTime" db:"CourseEndTime"`
	Status            int    `json:"Status" db:"Status"`
	SchoolName        string `json:"SchoolName" db:"SchoolName"`
	CollegeName       string `json:"CollegeName" db:"CollegeName"`
	MajorName         string `json:"MajorName" db:"MajorName"`
	FilePath          string `json:"FilePath" db:"FilePath"`
	TeacherName       string `json:"TeacherName" db:"TeacherName"`
	CourseCode        string `json:"CourseCode" db:"CourseCode"`
}

type ClassCourseRelation struct {
	Id       int `json:"Id" db:"Id"`
	CourseId int `json:"CourseId" db:"CourseId"`
	ClassId  int `json:"ClassId" db:"ClassId"`
}

type StudyPlan struct {
	Id           int `json:"Id" db:"Id"`
	StudentId    int `json:"StudentId" db:"StudentId"`
	CourseId     int `json:"CourseId" db:"CourseId"`
	ChapterId    int `json:"ChapterId" db:"ChapterId"`
	LearningRate int `json:"LearningRate" db:"LearningRate"`
	ChapterOrder int `json:"ChapterOrder" db:"ChapterOrder"`
	IsComplete   int `json:"IsComplete" db:"IsComplete"`
}

type CourseStudyPlanView struct {
	Id           int    `json:"Id" db:"Id"`
	StudentId    int    `json:"StudentId" db:"StudentId"`
	TrueName     string `json:"TrueName" db:"TrueName"`
	CourseId     int    `json:"CourseId" db:"CourseId"`
	CourseName   string `json:"CourseName" db:"CourseName"`
	ChapterId    int    `json:"ChapterId" db:"ChapterId"`
	LearningRate int    `json:"LearningRate" db:"LearningRate"`
	ChapterOrder int    `json:"ChapterOrder" db:"ChapterOrder"`
	ChapterNum   int    `json:"ChapterNum" db:"ChapterNum"`
	IsComplete   int    `json:"IsComplete" db:"IsComplete"`
}

type CourseClassStudentJson struct {
	Id            int   `json:"Id" db:"Id"`
	AddClass      []int `json:"AddClass" `
	AddStudent    []int `json:"AddStudent" `
	RemoveClass   []int `json:"RemoveClass"  `
	RemoveStudent []int `json:"RemoveStudent"  `
}

type CourseIdModel struct {
	CourseId int `json:"CourseId"  `
}

type CourseRelation struct {
	ClassArr   []int   `json:"ClassArr"  `
	StudentArr [][]int `json:"StudentArr"  `
}

type CourseRelationDb struct {
	StudentId int `json:"StudentId"   db:"StudentId"`
	ClassId   int `json:"ClassId"    db:"ClassId"`
}

type StudyCourse struct {
	CourseName      string `json:"CourseName"   db:"CourseName"`
	Digest          string `json:"Digest"    db:"Digest"`
	CourseId        int    `json:"CourseId"   db:"CourseId"`
	TeacherId       int    `json:"TeacherId"    db:"TeacherId"`
	TeacherName     string `json:"TeacherName"   db:"TeacherName"`
	SchoolName      string `json:"SchoolName"    db:"SchoolName"`
	CollegeName     string `json:"CollegeName"   db:"CollegeName"`
	MajorId         int    `json:"MajorId"    db:"MajorId"`
	MajorName       string `json:"MajorName"   db:"MajorName"`
	ChapterSum      int    `json:"ChapterSum"    db:"ChapterSum"`
	StudentSum      int    `json:"StudentSum"   db:"StudentSum"`
	CourseStartTime string `json:"CourseStartTime"    db:"CourseStartTime"`
	CourseEndTime   string `json:"CourseEndTime"   db:"CourseEndTime"`
	ChapterOrder    int    `json:"ChapterOrder"    db:"ChapterOrder"`
	LearningRate    int    `json:"LearningRate"   db:"LearningRate"`
	FilePath        string `json:"FilePath"    db:"FilePath"`
	IsCurrentStudy  int    `json:"IsCurrentStudy"    db:"IsCurrentStudy"`
}

type StudyCourseChapter struct {
	CourseName      string         `json:"CourseName"   db:"CourseName"`
	Digest          string         `json:"Digest"    db:"Digest"`
	CourseId        int            `json:"CourseId"   db:"CourseId"`
	TeacherId       int            `json:"TeacherId"    db:"TeacherId"`
	TeacherName     string         `json:"TeacherName"   db:"TeacherName"`
	SchoolName      string         `json:"SchoolName"    db:"SchoolName"`
	CollegeName     string         `json:"CollegeName"   db:"CollegeName"`
	MajorId         int            `json:"MajorId"    db:"MajorId"`
	MajorName       string         `json:"MajorName"   db:"MajorName"`
	ChapterSum      int            `json:"ChapterSum"    db:"ChapterSum"`
	StudentSum      int            `json:"StudentSum"   db:"StudentSum"`
	CourseStartTime string         `json:"CourseStartTime"    db:"CourseStartTime"`
	CourseEndTime   string         `json:"CourseEndTime"   db:"CourseEndTime"`
	ChapterOrder    int            `json:"ChapterOrder"    db:"ChapterOrder"`
	LearningRate    int            `json:"LearningRate"   db:"LearningRate"`
	FilePath        string         `json:"FilePath"    db:"FilePath"`
	IsCurrentStudy  int            `json:"IsCurrentStudy"    db:"IsCurrentStudy"`
	ChapterList     []*ChapterList `json:"ChapterList"   `
}
