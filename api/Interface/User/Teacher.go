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

	"github.com/gin-gonic/gin"
)

func LoginTeacher(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}
	log.Printf("%s", body)

	var teacher model.Teacher
	json.Unmarshal([]byte(body), &teacher)

	data := []byte(teacher.TeacherPassword)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	newteacher := new(model.Teacher)

	lib.Db.QueryRow("select TeacherId,SchoolId,TeacherAccount,TeacherName,Sex,PhoneNumber,Email,TeacherTitle from teacher where TeacherAccount=? and TeacherPassword=?", teacher.TeacherAccount, md5str1).
		Scan(&newteacher.TeacherId, &newteacher.SchoolId, &newteacher.TeacherAccount, &newteacher.TeacherName, &newteacher.Sex, &newteacher.PhoneNumber, &newteacher.Email, &newteacher.TeacherTitle)

	if newteacher.TeacherId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "账号或密码错误",
			"data": "{}",
		})
	} else {
		tokenString, _ := jwt_use.GetToken(newteacher.TeacherName, 0)
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "登录成功",
			"data":  newteacher,
			"token": tokenString,
		})
	}
}

func EditTeacher(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var teacher model.Teacher
	json.Unmarshal([]byte(body), &teacher)

	sqlStr := "update  teacher set  SchoolId=? , TeacherName=?,Sex=?,PhoneNumber=?,Email=?,TeacherTitle=? where TeacherId=? "

	ret, err := lib.Db.Exec(sqlStr, teacher.SchoolId, teacher.TeacherName, teacher.Sex, teacher.PhoneNumber, teacher.Email, teacher.TeacherTitle, teacher.TeacherId)

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

func AddTeacher(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var teacher model.Teacher
	json.Unmarshal([]byte(body), &teacher)

	TeacherId := 0
	lib.Db.QueryRow("select  TeacherId from teacher where  TeacherAccount=? ", teacher.TeacherAccount).Scan(&TeacherId)

	if TeacherId != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "教师账号重复",
			"data": "{}",
		})
		return
	}

	sqlStr := "insert into  teacher(SchoolId,TeacherAccount,TeacherPassword,TeacherName,Sex,PhoneNumber,Email,TeacherTitle) VALUES (?,?,?,?,?,?,?,?) "

	data := []byte(teacher.TeacherPassword)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	ret, err := lib.Db.Exec(sqlStr, teacher.SchoolId, teacher.TeacherAccount, md5str1, teacher.TeacherName, teacher.Sex, teacher.PhoneNumber, teacher.Email, teacher.TeacherTitle)

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

func DelTeacher(c *gin.Context) {

	tx, err := lib.Db.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var teacher model.Teacher
	json.Unmarshal([]byte(body), &teacher)

	sqlStr := " delete from  teacher   where TeacherId=? "
	ret, err := tx.Exec(sqlStr, teacher.TeacherId)

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	if n < 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "所选的内容可能已删除",
			"data": "{}",
		})
	}

	sqlStr = " update course set TeacherId=0    where TeacherId=? "
	ret, err = tx.Exec(sqlStr, teacher.TeacherId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected() // 操作影响的行数
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	if n < 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "所选的内容可能已删除",
			"data": "{}",
		})
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": "{}",
	})

}

func GetTeacherById(c *gin.Context) {
	TeacherId := c.Query("TeacherId")

	newTeacherview := new(model.TeacherView)

	err := lib.Db.QueryRow(" select a.TeacherId,a.SchoolId,a.TeacherAccount,a.TeacherName,a.Sex,a.PhoneNumber,a.Email,a.TeacherTitle,b.SchoolName,b.SchoolAddress from teacher a   left join school b on a.SchoolId=b.Id where  TeacherId=?", TeacherId).Scan(&newTeacherview.TeacherId, &newTeacherview.SchoolId, &newTeacherview.TeacherAccount, &newTeacherview.TeacherName, &newTeacherview.Sex, &newTeacherview.PhoneNumber, &newTeacherview.Email, &newTeacherview.TeacherTitle, &newTeacherview.SchoolName, &newTeacherview.SchoolAddress)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	if newTeacherview.TeacherId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": newTeacherview,
		})
	}
}

func GetTeacherList(c *gin.Context) {

	rows, err := lib.Db.Query(" select a.TeacherId,a.SchoolId,a.TeacherAccount,a.TeacherName,a.Sex,a.PhoneNumber,a.Email,a.TeacherTitle,COALESCE(b.SchoolName,'') 'SchoolName',COALESCE(b.SchoolAddress,'') 'SchoolAddress' from teacher a   left join school b on a.SchoolId=b.Id  ")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "修改失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	newTeacherviewarr := make([]*model.TeacherView, 0)
	for rows.Next() {
		newTeacherview := new(model.TeacherView)
		err = rows.Scan(&newTeacherview.TeacherId, &newTeacherview.SchoolId, &newTeacherview.TeacherAccount, &newTeacherview.TeacherName, &newTeacherview.Sex, &newTeacherview.PhoneNumber, &newTeacherview.Email, &newTeacherview.TeacherTitle, &newTeacherview.SchoolName, &newTeacherview.SchoolAddress)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "修改失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		newTeacherviewarr = append(newTeacherviewarr, newTeacherview)
	}
	rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": newTeacherviewarr,
	})
}

func EditTeacherPassWordById(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "修改失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var teacher model.Teacher
	json.Unmarshal([]byte(body), &teacher)

	data := []byte(teacher.TeacherPassword)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	resetsql := "update teacher set TeacherPassword=? where TeacherId=?"
	ret, err := lib.Db.Exec(resetsql, md5str1, teacher.TeacherId)

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
