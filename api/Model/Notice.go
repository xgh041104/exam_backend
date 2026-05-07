package model

type Notice struct {
	Id            int    `json:"Id" db:"Id"`
	Time          string `json:"Time" db:"Time"`
	NoticeTitle   string `json:"NoticeTitle" db:"NoticeTitle"`
	NoticeContent string `json:"NoticeContent" db:"NoticeContent"`
	SendUser      string `json:"SendUser" db:"SendUser"`
	NoticeLevel   int    `json:"NoticeLevel" db:"NoticeLevel"`
	NoticeType    string `json:"NoticeType" db:"NoticeType"`
}
