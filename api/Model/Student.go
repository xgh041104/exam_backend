package model

type Student struct {
	Id             int    `json:"Id" db:"Id"`
	StudentType    int    `json:"StudentType" db:"StudentType"`
	StudentAccount string `json:"StudentAccount" db:"StudentAccount"`
	StudentPwd     string `json:"StudentPwd" db:"StudentPwd"`
	StandId        int    `json:"StandId" db:"StandId"`
	ExamName       string `json:"ExamName" db:"ExamName"`
	SchoolId       int    `json:"SchoolId" db:"SchoolId"`
	CollegeId      int    `json:"CollegeId" db:"CollegeId"`
	MajorId        int    `json:"MajorId" db:"MajorId"`
	ClassId        int    `json:"ClassId" db:"ClassId"`
	TrueName       string `json:"TrueName" db:"TrueName"`
	IDNumber       string `json:"IDNumber" db:"IDNumber"`
	ExamNumber     string `json:"ExamNumber" db:"ExamNumber"`
	Birthday       string `json:"Birthday" db:"Birthday"`
	Phone          string `json:"Phone" db:"Phone"`
	Email          string `json:"Email" db:"Email"`
	IDImage        string `json:"IDImage" db:"IDImage"`
	FaceOpen       int    `json:"FaceOpen" db:"FaceOpen"`
	Sex            int    `json:"Sex" db:"Sex"`
	NativePlace    string `json:"NativePlace" db:"NativePlace"`
}

type StudentView struct {
	Id             int    `json:"Id" db:"Id"`
	StudentType    int    `json:"StudentType" db:"StudentType"`
	StudentAccount string `json:"StudentAccount" db:"StudentAccount"`
	StudentPwd     string `json:"StudentPwd" db:"StudentPwd"`
	StandId        int    `json:"StandId" db:"StandId"`
	ExamName       string `json:"ExamName" db:"ExamName"`
	SchoolId       int    `json:"SchoolId" db:"SchoolId"`
	CollegeId      int    `json:"CollegeId" db:"CollegeId"`
	MajorId        int    `json:"MajorId" db:"MajorId"`
	ClassId        int    `json:"ClassId" db:"ClassId"`
	TrueName       string `json:"TrueName" db:"TrueName"`
	IDNumber       string `json:"IDNumber" db:"IDNumber"`
	ExamNumber     string `json:"ExamNumber" db:"ExamNumber"`
	Birthday       string `json:"Birthday" db:"Birthday"`
	Phone          string `json:"Phone" db:"Phone"`
	Email          string `json:"Email" db:"Email"`
	IDImage        string `json:"IDImage" db:"IDImage"`
	FaceOpen       int    `json:"FaceOpen" db:"FaceOpen"`
	SchoolName     string `json:"SchoolName" db:"SchoolName"`
	CollegeName    string `json:"CollegeName" db:"CollegeName"`
	MajorName      string `json:"MajorName" db:"MajorName"`
	ClassName      string `json:"ClassName" db:"ClassName"`
	StandName      string `json:"StandName" db:"StandName"`
	Sex            int    `json:"Sex" db:"Sex"`
	NativePlace    string `json:"NativePlace" db:"NativePlace"`
}

type ImportStudent struct {
	StudentType    int    `json:"StudentType" db:"StudentType"`
	TrueName       string `json:"TrueName" db:"TrueName"`
	IDNumber       string `json:"IDNumber" db:"IDNumber"`
	SchoolId       int    `json:"SchoolId" db:"SchoolId"`
	CollegeId      int    `json:"CollegeId" db:"CollegeId"`
	MajorId        int    `json:"MajorId" db:"MajorId"`
	ClassId        int    `json:"ClassId" db:"ClassId"`
	ExamNumber     string `json:"ExamNumber" db:"ExamNumber"`
	ExamName       string `json:"ExamName" db:"ExamName"`
	NativePlace    string `json:"NativePlace" db:"NativePlace"`
	Sex            int    `json:"Sex" db:"Sex"`
	Birthday       string `json:"Birthday" db:"Birthday"`
	SchoolName     string `json:"SchoolName" db:"SchoolName"`
	StudentAccount string `json:"StudentAccount" db:"StudentAccount"`
	StudentPwd     string `json:"StudentPwd" db:"StudentPwd"`
	StandName      string `json:"StandName" db:"StandName"`
	StandId        int    `json:"StandId" db:"StandId"`
	Phone          string `json:"Phone" db:"Phone"`
	Email          string `json:"Email" db:"Email"`
	Status         int    `json:"Status" db:"Status"`
}
