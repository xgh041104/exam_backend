package model

type Chapter struct {
	Id             int    `json:"Id" db:"Id"`
	ChapterName    string `json:"ChapterName" db:"ChapterName"`
	CourseId       int    `json:"CourseId" db:"CourseId"`
	ChapterType    string `json:"ChapterType" db:"ChapterType"`
	ChapterOrder   int    `json:"ChapterOrder" db:"ChapterOrder"`
	ChapterContent string `json:"ChapterContent" db:"ChapterContent"`
}

type ChapterEdit struct {
	Id             int    `json:"Id" db:"Id"`
	ChapterName    string `json:"ChapterName" db:"ChapterName"`
	CourseId       int    `json:"CourseId" db:"CourseId"`
	ChapterType    string `json:"ChapterType" db:"ChapterType"`
	ChapterOrder   int    `json:"ChapterOrder" db:"ChapterOrder"`
	ChapterContent string `json:"ChapterContent" db:"ChapterContent"`
	RemoveFile     []int  `json:"RemoveFile" `
}

type ChapterView struct {
	Id             int         `json:"Id" db:"Id"`
	ChapterName    string      `json:"ChapterName" db:"ChapterName"`
	CourseId       int         `json:"CourseId" db:"CourseId"`
	ChapterType    string      `json:"ChapterType" db:"ChapterType"`
	ChapterOrder   int         `json:"ChapterOrder" db:"ChapterOrder"`
	ChapterContent string      `json:"ChapterContent" db:"ChapterContent"`
	FileInfo       []*FileInfo `json:"FileInfo" `
}

type ChapterIdStruct struct {
	ChapterId int `json:"ChapterId" db:"ChapterId"`
}

type ChapterOrderParam struct {
	ChapterOrder []int `json:"ChapterOrder"  `
}

type ChapterVideo struct {
	ChapterId int `json:"ChapterId" db:"ChapterId"`
	SchoolId  int `json:"SchoolId" db:"SchoolId"`
}

type ChapterList struct {
	Id           int    `json:"Id" db:"Id"`
	ChapterName  string `json:"ChapterName" db:"ChapterName"`
	ChapterType  string `json:"ChapterType" db:"ChapterType"`
	ChapterOrder int    `json:"ChapterOrder" db:"ChapterOrder"`
}
