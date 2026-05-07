package model

type FileInfo struct {
	Id        int    `json:"Id" db:"Id"`
	FileType  string `json:"FileType" db:"FileType"`
	FileName  string `json:"FileName" db:"FileName"`
	FileUseTo string `json:"FileUseTo" db:"FileUseTo"`
	SchoolId  int    `json:"SchoolId" db:"SchoolId"`
	FilePath  string `json:"FilePath" db:"FilePath"`
}
