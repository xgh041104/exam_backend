package User

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/jwt_use"
	"StudyExamPlatformAPI/lib"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/shakinm/xlsReader/xls"
	goutils "github.com/typa01/go-utils"
)

func LoginStudent(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}
	log.Printf("%s", body)

	var user model.Student
	json.Unmarshal([]byte(body), &user)

	data := []byte(user.StudentPwd)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	newuser := new(model.Student)

	lib.Db.QueryRow("select Id,StudentType,StudentAccount,StudentPwd,StandId,ExamName,SchoolId,CollegeId,MajorId,ClassId,TrueName,IDNumber,ExamNumber,Birthday,Phone,Email,IDImage,FaceOpen,Sex,NativePlace from student   where StudentAccount=? and StudentPwd=?", user.StudentAccount, md5str1).
		Scan(&newuser.Id, &newuser.StudentType, &newuser.StudentAccount, &newuser.StudentPwd, &newuser.StandId, &newuser.ExamName, &newuser.SchoolId, &newuser.CollegeId, &newuser.MajorId, &newuser.ClassId, &newuser.TrueName, &newuser.IDNumber, &newuser.ExamNumber, &newuser.Birthday, &newuser.Phone, &newuser.Email, &newuser.IDImage, &newuser.FaceOpen, &newuser.Sex, &newuser.NativePlace)

	if newuser.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "账号或密码错误",
			"data": "{}",
		})
	} else {
		ret, err := lib.Db.Exec("insert into loginlog(studentId,loginTime) values(?,?)", newuser.Id, time.Now().Unix())
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "系统错误" + err.Error(),
				"data": "{}",
			})
		}
		n, _ := ret.RowsAffected()

		if n > 0 {
			tokenString, _ := jwt_use.GetToken(newuser.TrueName, 1)
			c.JSON(http.StatusOK, gin.H{
				"code":  1,
				"msg":   "登录成功",
				"data":  newuser,
				"token": tokenString,
			})
		}

	}
}

// 获取学生信息列表无筛选条件
func GetStudentViewList(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	var studentviewArr []*model.StudentView

	rows, err := lib.Db.Query("select Id,StudentType,StudentAccount,StudentPwd,StandId,ExamName,SchoolId,CollegeId,MajorId,ClassId,TrueName,IDNumber,ExamNumber,Birthday,Phone,Email,IDImage,FaceOpen,SchoolName,CollegeName,MajorName,ClassName,StandName,Sex,NativePlace   from  studentview where SchoolId=?  ", SchoolId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}

	for rows.Next() {
		// 声明
		newstudentview := new(model.StudentView)
		err := rows.Scan(&newstudentview.Id, &newstudentview.StudentType, &newstudentview.StudentAccount, &newstudentview.StudentPwd, &newstudentview.StandId, &newstudentview.ExamName, &newstudentview.SchoolId, &newstudentview.CollegeId, &newstudentview.MajorId, &newstudentview.ClassId, &newstudentview.TrueName, &newstudentview.IDNumber, &newstudentview.ExamNumber, &newstudentview.Birthday, &newstudentview.Phone, &newstudentview.Email, &newstudentview.IDImage, &newstudentview.FaceOpen, &newstudentview.SchoolName, &newstudentview.CollegeName, &newstudentview.MajorName, &newstudentview.ClassName, &newstudentview.StandName, &newstudentview.Sex, &newstudentview.NativePlace)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		studentviewArr = append(studentviewArr, newstudentview)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": studentviewArr,
	})
}

func GetStudentViewById(c *gin.Context) {

	StudentId := c.Query("StudentId")

	newstudentview := new(model.StudentView)

	lib.Db.QueryRow("select Id,StudentType,StudentAccount,StudentPwd,StandId,ExamName,SchoolId,CollegeId,MajorId,ClassId,TrueName,IDNumber,ExamNumber,Birthday,Phone,Email,IDImage,FaceOpen,SchoolName,CollegeName,MajorName,ClassName,StandName,Sex,NativePlace   from  studentview   where Id=?", StudentId).Scan(&newstudentview.Id, &newstudentview.StudentType, &newstudentview.StudentAccount, &newstudentview.StudentPwd, &newstudentview.StandId, &newstudentview.ExamName, &newstudentview.SchoolId, &newstudentview.CollegeId, &newstudentview.MajorId, &newstudentview.ClassId, &newstudentview.TrueName, &newstudentview.IDNumber, &newstudentview.ExamNumber, &newstudentview.Birthday, &newstudentview.Phone, &newstudentview.Email, &newstudentview.IDImage, &newstudentview.FaceOpen, &newstudentview.SchoolName, &newstudentview.CollegeName, &newstudentview.MajorName, &newstudentview.ClassName, &newstudentview.StandName, &newstudentview.Sex, &newstudentview.NativePlace)

	if newstudentview.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": newstudentview,
		})
	}
}

func AddStudent(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var student model.Student
	json.Unmarshal([]byte(body), &student)

	Id := 0
	lib.Db.QueryRow("select  Id from student where  StudentAccount=? ", student.StudentAccount).Scan(&Id)

	if Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "学生账号重复",
			"data": "{}",
		})
		return
	}

	if student.IDNumber != "" {
		lib.Db.QueryRow("select  Id from student where  IDNumber=? ", student.IDNumber).Scan(&Id)

		if Id != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "身份证号重复",
				"data": "{}",
			})
			return
		}

	}
	if student.ExamNumber != "" {
		lib.Db.QueryRow("select  Id from student where  ExamNumber=? ", student.ExamNumber).Scan(&Id)

		if Id != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "考号重复",
				"data": "{}",
			})
			return
		}

	}

	sqlStr := "insert into  student(StudentType,StudentAccount,StudentPwd,StandId,ExamName,SchoolId,CollegeId,MajorId,ClassId,TrueName,IDNumber,ExamNumber,Birthday,Phone,Email,IDImage,Sex,NativePlace) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) "

	data := []byte(student.StudentPwd)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	ret, err := lib.Db.Exec(sqlStr, student.StudentType, student.StudentAccount, md5str1, student.StandId, student.ExamName, student.SchoolId, student.CollegeId, student.MajorId, student.ClassId, student.TrueName, student.IDNumber, student.ExamNumber, student.Birthday, student.Phone, student.Email, student.IDImage, student.Sex, student.NativePlace)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	if n > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
	}
}

func EditStudent(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var student model.Student
	json.Unmarshal([]byte(body), &student)

	sqlStr := "update  student set  StudentType=? ,StandId=?,ExamName=?,SchoolId=?,CollegeId=?,MajorId=?,ClassId=?,TrueName=?,IDNumber=?,ExamNumber=?,Birthday=?,Phone=?,Email=?,FaceOpen=?,Sex=?,NativePlace=?  where Id=? "

	ret, err := lib.Db.Exec(sqlStr, student.StudentType, student.StandId, student.ExamName, student.SchoolId, student.CollegeId, student.MajorId, student.ClassId, student.TrueName, student.IDNumber, student.ExamNumber, student.Birthday, student.Phone, student.Email, student.FaceOpen, student.Sex, student.NativePlace, student.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	if n > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "没有修改任何信息",
			"data": "{}",
		})
	}
}

func DelStudent(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var student model.Student
	json.Unmarshal([]byte(body), &student)

	sqlStr := " delete from  student   where Id=? "
	ret, err := lib.Db.Exec(sqlStr, student.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	if n > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "所选的内容可能已删除",
			"data": "{}",
		})
	}
}

func EditFaceOpenState(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var student model.Student
	json.Unmarshal([]byte(body), &student)

	sqlStr := "update  student set FaceOpen=? where Id=? "
	ret, err := lib.Db.Exec(sqlStr, student.FaceOpen, student.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	if n > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "所选的内容可能已删除",
			"data": "{}",
		})
	}
}

func EditStudentPassWordById(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "修改失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var student model.Student
	json.Unmarshal([]byte(body), &student)

	data := []byte(student.StudentPwd)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	resetsql := "update student set StudentPwd=? where Id=?"
	ret, err := lib.Db.Exec(resetsql, md5str1, student.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "修改失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	n, err := ret.RowsAffected()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "修改失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	if n < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "修改失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "修改密码成功",
		"data": "{}",
	})
}

func ReSetStudentPassWord(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var student model.Student
	json.Unmarshal([]byte(body), &student)
	studentaccount := ""
	lib.Db.QueryRow("select StudentAccount from student where Id=?", student.Id).Scan(&studentaccount)
	originalPwd := studentaccount + "!q@w#e$r%t"

	data := []byte(originalPwd)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	resetsql := "update student set StudentPwd=? where Id=?"
	ret, err := lib.Db.Exec(resetsql, md5str1, student.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	n, err := ret.RowsAffected()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "重置失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	if n < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "重置失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "重置成功，密码为" + originalPwd,
		"data": "{}",
	})
}

func GetStudentExeclData(c *gin.Context) {
	files, err := c.MultipartForm() // 获取文件

	data := c.PostForm("data")
	var standId model.StandId
	json.Unmarshal([]byte(data), &standId)

	if standId.SchoolId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	if len(files.File["files"]) > 0 {
		filesarr := files.File["files"]

		filename := goutils.GUID()
		for i := 0; i < len(filesarr); i++ {
			file := filesarr[i]
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
				return
			}

			filePath := "Resources/Execl/" + filename + path.Ext(file.Filename)

			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}

			f, err := excelize.OpenFile(filePath)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg":  "操作错误",
					"data": "",
				})
				return
			}

			importStudentarr := make([]*model.ImportStudent, 0)
			for _, name := range f.GetSheetMap() {
				rows := f.GetRows(name)
				if err != nil {
					fmt.Println(err)
					return
				}
				// 遍历每行数据
				for i, row := range rows {
					fmt.Printf("Row %d:\n", i+1)
					if i == 0 {
						continue
					}
					importStudent := new(model.ImportStudent)
					importStudent.StudentType = 0
					importStudent.SchoolId = standId.SchoolId
					// 遍历每列数据
					for j, colCell := range row {

						if rows[0][j] == "姓名" {
							importStudent.TrueName = colCell
						} else if rows[0][j] == "证件号码" {

							if len(colCell) == 18 {

								idnumber := colCell
								importStudent.IDNumber = idnumber[0:6]
								importStudent.Birthday = idnumber[6:10] + "-" + idnumber[10:12] + "-" + idnumber[12:14]

								genderCode := idnumber[16:17]
								sexint, err := strconv.Atoi(genderCode)
								if err != nil {
									importStudent.Sex = 1
								}
								if sexint%2 == 0 {
									importStudent.Sex = 0
								} else {
									importStudent.Sex = 1
								}
							} else {
								importStudent.IDNumber = colCell
							}
						} else if rows[0][j] == "登录账号" {
							importStudent.StudentAccount = colCell
						} else if rows[0][j] == "电话" {
							importStudent.Phone = colCell
						} else if rows[0][j] == "邮箱" {
							importStudent.Email = colCell
						} else if rows[0][j] == "准考证号" {
							importStudent.ExamNumber = colCell
						} else if rows[0][j] == "籍贯" {
							importStudent.NativePlace = colCell
						} else if rows[0][j] == "性别" {
							if colCell == "男" {
								importStudent.Sex = 1
							} else {
								importStudent.Sex = 0
							}
						} else if rows[0][j] == "出生日期" {
							importStudent.Birthday = colCell
						}

					}
					fmt.Println()
					importStudentarr = append(importStudentarr, importStudent)
				}

				break
			}

			os.Remove(filePath) //一直是本地操作
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "操作成功",
				"data": importStudentarr,
			})
		}
	}
}

func MatchAddStudent(c *gin.Context) {

	var arrImportStudent []*model.ImportStudent
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	json.Unmarshal([]byte(body), &arrImportStudent)

	// sqlStr := "insert into  student(StudentType,StudentAccount,StudentPwd,StandId,TrueName,IDNumber,ExamNumber,ExamName,Birthday,Phone,Email,Sex,NativePlace) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) "
	for i := 0; i < len(arrImportStudent); i++ {

		tempenetity := arrImportStudent[i]

		if tempenetity.IDNumber == "" || tempenetity.ExamNumber == "" {
			arrImportStudent[i].Status = 0
			continue
		}

		StudentId := 0
		lib.Db.QueryRow("select  Id from student where StudentAccount=?", tempenetity.StudentAccount).Scan(&StudentId)

		if StudentId > 0 { //账号重复
			arrImportStudent[i].Status = 0
			continue
		}

		if tempenetity.IDNumber != "" {
			lib.Db.QueryRow("select  Id from student where  IDNumber=? ", tempenetity.IDNumber).Scan(&StudentId)

			if StudentId != 0 { //身份证号重复
				arrImportStudent[i].Status = 0
			}

		}
		if tempenetity.ExamNumber != "" {
			lib.Db.QueryRow("select  Id from student where  ExamNumber=? ", tempenetity.ExamNumber).Scan(&StudentId)

			if StudentId != 0 { //准考证号重复
				arrImportStudent[i].Status = 0
			}

		}

		StandId := 0

		lib.Db.QueryRow("select  Id from stand where StandName=?", tempenetity.StandName).Scan(&StandId)

		if StandId > 0 {
			tempenetity.StandId = StandId
			arrImportStudent[i].StandId = StandId

			// lib.Db.Exec("insert ")
		}

		if tempenetity.StudentType == 0 {

			if len(tempenetity.ExamNumber) >= 6 {
				tempenetity.StudentPwd = tempenetity.ExamNumber[len(tempenetity.ExamNumber)-6 : len(tempenetity.ExamNumber)]
			} else {
				tempenetity.StudentPwd = tempenetity.ExamNumber
			}
			//在校生
			sqlStr := "insert into  student(StudentType,TrueName,IDNumber,SchoolId,CollegeId,MajorId,ClassId,ExamNumber,NativePlace,Sex,Birthday,StudentAccount,StudentPwd,Phone,Email) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) "

			data := []byte(tempenetity.StudentPwd)
			has := md5.Sum(data)
			md5str1 := fmt.Sprintf("%x", has)
			ret, err := lib.Db.Exec(sqlStr, 0, tempenetity.TrueName, tempenetity.IDNumber, tempenetity.SchoolId, tempenetity.CollegeId, tempenetity.MajorId, tempenetity.ClassId, tempenetity.ExamNumber, tempenetity.NativePlace, tempenetity.Sex, tempenetity.Birthday, tempenetity.StudentAccount, md5str1, tempenetity.Phone, tempenetity.Email)

			if err != nil {
				arrImportStudent[i].Status = 0
				continue
			}
			n, err := ret.RowsAffected()
			if err != nil {
				arrImportStudent[i].Status = 0
				continue
			}
			if n <= 0 {
				arrImportStudent[i].Status = 0
				continue
			} else {
				arrImportStudent[i].Status = 1
			}
		} else if tempenetity.StudentType == 1 {
			//社会人士
			if len(tempenetity.ExamNumber) >= 6 {
				tempenetity.StudentPwd = tempenetity.ExamNumber[len(tempenetity.ExamNumber)-6 : len(tempenetity.ExamNumber)]
			} else {
				tempenetity.StudentPwd = tempenetity.ExamNumber
			}
			sqlStr := "insert into  student(StudentType,TrueName,IDNumber,SchoolId,ExamNumber,ExamName,NativePlace,Sex,Birthday,StudentAccount,StudentPwd,StandId,Phone,Email) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?) "

			data := []byte(tempenetity.StudentPwd)
			has := md5.Sum(data)
			md5str1 := fmt.Sprintf("%x", has)
			ret, err := lib.Db.Exec(sqlStr, 1, tempenetity.TrueName, tempenetity.IDNumber, tempenetity.SchoolId, tempenetity.ExamNumber, tempenetity.ExamName, tempenetity.NativePlace, tempenetity.Sex, tempenetity.Birthday, tempenetity.StudentAccount, md5str1, tempenetity.StandId, tempenetity.Phone, tempenetity.Email)

			if err != nil {
				arrImportStudent[i].Status = 0
				continue
			}
			n, err := ret.RowsAffected()
			if err != nil {
				arrImportStudent[i].Status = 0
				continue
			}
			if n <= 0 {
				arrImportStudent[i].Status = 0
				continue
			} else {
				arrImportStudent[i].Status = 1

			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": arrImportStudent,
	})
}

func MatchAddStudentIDImg(c *gin.Context) {

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	errarr := make([]string, 0)
	files, err := c.MultipartForm() // 获取文件
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未收到文件",
			"data": "{}",
		})
		return
	}
	if len(files.File["files"]) > 0 {
		filesarr := files.File["files"]

		for i := 0; i < len(filesarr); i++ {
			filetemp := filesarr[i]
			filename := goutils.GUID()

			filenameall := path.Base(filetemp.Filename)
			filesuffix := path.Ext(filenameall)

			Idflag := filenameall[0 : len(filenameall)-len(filesuffix)]
			num := 0
			tx.QueryRow("select count(1) 'num' from student where IDNumber=? or ExamNumber=?", Idflag, Idflag).Scan(&num)

			if num == 0 {
				errarr = append(errarr, filenameall)
				continue
			}
			filepath := "Resources/IDImg/" + filename + filesuffix
			lib.UploadFile(c, filetemp, filepath)

			ret, err := tx.Exec("update student set IDImage=?  where IDNumber=? or ExamNumber=?", filepath, Idflag, Idflag)
			if err != nil {
				errarr = append(errarr, filenameall)
				continue
			}
			n, err := ret.RowsAffected()
			if err != nil {
				errarr = append(errarr, filenameall)
				continue
			}
			if n <= 0 {
				errarr = append(errarr, filenameall)
				continue
			}
		}
		err = tx.Commit()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功,未导入的数据在data",
			"data": errarr,
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未收到文件",
			"data": "{}",
		})
	}

}

func MatchAddSocietyStudentExecl(c *gin.Context) {
	files, err := c.MultipartForm() // 获取文件.
	data := c.PostForm("data")
	var standId model.StandId
	json.Unmarshal([]byte(data), &standId)

	if standId.StandId == 0 && standId.SchoolId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	if len(files.File["files"]) > 0 {
		filesarr := files.File["files"]
		filename := goutils.GUID()
		for i := 0; i < len(filesarr); i++ {
			file := filesarr[i]
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
				return
			}

			filePath := "Resources/Execl/" + filename + path.Ext(file.Filename)

			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
			importStudentarr := make([]*model.ImportStudent, 0)
			myMap := make(map[string]*model.ImportStudent)
			if path.Ext(file.Filename) == ".xls" {
				wb, err := xls.OpenFile(filePath)
				if err != nil {
					fmt.Println(err)
				}
				sheet, err := wb.GetSheet(0)
				if err != nil {
					fmt.Println(err)
				}
				//获得工作表的行数
				rn := sheet.GetNumberRows()
				//循环处理每一行
				for i := 1; i < rn; i++ {
					//获取行
					importStudent := new(model.ImportStudent)
					importStudent.StudentType = 1
					importStudent.NativePlace = "湖北"
					importStudent.SchoolId = standId.SchoolId
					importStudent.StandId = standId.StandId
					if row, e := sheet.GetRow(i); e == nil && row != nil {
						//获取这行的所有列
						cols := row.GetCols()
						if cols == nil || len(cols) < 1 {
							continue
						}
						//获得列数
						colen := len(cols)
						for j := 0; j < colen; j++ {

							if strings.Contains(sheet.GetRows()[0].GetCols()[j].GetString(), "身份证") {

								if len(row.GetCols()[j].GetString()) == 18 {
									idnumber := row.GetCols()[j].GetString()
									importStudent.IDNumber = idnumber[0:6]
									importStudent.Birthday = idnumber[6:10] + "-" + idnumber[10:12] + "-" + idnumber[12:14]

									genderCode := idnumber[16:17]
									sexint, err := strconv.Atoi(genderCode)
									if err != nil {
										importStudent.Sex = 1
									}
									if sexint%2 == 0 {
										importStudent.Sex = 0
									} else {
										importStudent.Sex = 1
									}
								}

							} else if strings.Contains(sheet.GetRows()[0].GetCols()[j].GetString(), "姓名") {
								importStudent.TrueName = row.GetCols()[j].GetString()
							} else if strings.Contains(sheet.GetRows()[0].GetCols()[j].GetString(), "准考证号") {

								if len(row.GetCols()[j].GetString()) >= 6 {
									examnumber := row.GetCols()[j].GetString()
									importStudent.StudentAccount = examnumber
									importStudent.ExamNumber = examnumber

									importStudent.StudentPwd = examnumber[len(examnumber)-6 : len(examnumber)] // 后六位
								}
							}
						}
					}

					_, ok := myMap[importStudent.StudentAccount]

					if !ok {
						myMap[importStudent.StudentAccount] = importStudent
						importStudentarr = append(importStudentarr, importStudent)
					}

				}
			} else {
				f, err := excelize.OpenFile(filePath)
				if err != nil {
					fmt.Println(err)
					return
				}

				// 获取出生年月日

				for _, name := range f.GetSheetMap() {
					rows := f.GetRows(name)
					if err != nil {
						fmt.Println(err)
						return
					}
					// 遍历每行数据
					for i, row := range rows {
						fmt.Printf("Row %d:\n", i+1)
						if i == 0 {
							continue
						}
						importStudent := new(model.ImportStudent)
						importStudent.StudentType = 1
						importStudent.NativePlace = "湖北"
						importStudent.SchoolId = standId.SchoolId
						importStudent.StandId = standId.StandId
						// 遍历每列数据
						for j, colCell := range row {

							if strings.Contains(rows[0][j], "身份证") {
								if len(colCell) == 18 {
									importStudent.IDNumber = colCell[0:6]
									importStudent.Birthday = colCell[6:10] + "-" + colCell[10:12] + "-" + colCell[12:14]

									genderCode := colCell[16:17]
									sexint, err := strconv.Atoi(genderCode)
									if err != nil {
										importStudent.Sex = 1
									}
									if sexint%2 == 0 {
										importStudent.Sex = 0
									} else {
										importStudent.Sex = 1
									}
								}

							} else if strings.Contains(rows[0][j], "姓名") {
								importStudent.TrueName = colCell
							} else if strings.Contains(rows[0][j], "准考证号") {
								if len(colCell) >= 6 {
									importStudent.StudentAccount = colCell

									importStudent.ExamNumber = colCell

									importStudent.StudentPwd = colCell[len(colCell)-6 : len(colCell)] // 后六位
								}
							}
						}
						_, ok := myMap[importStudent.StudentAccount]
						if !ok {
							myMap[importStudent.StudentAccount] = importStudent
							importStudentarr = append(importStudentarr, importStudent)
						}
					}

					break
				}

			}

			myMap = nil
			os.Remove(filePath) //一直是本地操作
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "操作成功",
				"data": importStudentarr,
			})
		}
	}
}
