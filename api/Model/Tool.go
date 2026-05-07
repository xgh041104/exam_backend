package model

type ToolConfig struct {
	Id                  int `json:"Id" db:"Id"`
	StaticResourcesType int `json:"StaticResourcesType" db:"StaticResourcesType"`
	// OSSConfig           string `json:"OSSConfig" db:"OSSConfig"`
	// CDNConfig           string `json:"CDNConfig" db:"CDNConfig"`
	// LocalUrl            string `json:"LocalUrl" db:"LocalUrl"`
	FaceVerify         int `json:"FaceVerify" db:"FaceVerify"`
	SeparateFaceVerify int `json:"SeparateFaceVerify" db:"SeparateFaceVerify"`
}

type FaceVerifyModel struct {
	FaceVerify         int `json:"FaceVerify" db:"FaceVerify"`
	SeparateFaceVerify int `json:"SeparateFaceVerify" db:"SeparateFaceVerify"`
	MidiFlag           int `json:"MidiFlag"  `
}
