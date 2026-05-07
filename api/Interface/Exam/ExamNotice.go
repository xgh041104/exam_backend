package exam

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditExamNotice(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var examnotice model.ExamNotice
	json.Unmarshal([]byte(body), &examnotice)

	sqlStr := "Update  examnotice set  Title=?,SchoolName=?,CourseName=?,CourseCode=?,Context=? where ExamId=? "

	ret, err := lib.Db.Exec(sqlStr, examnotice.Title, examnotice.SchoolName, examnotice.CourseName, examnotice.CourseCode, examnotice.Context, examnotice.ExamId)

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

	if n >= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
	}
}

func AddExamNotice(c *gin.Context) {

	tx, err := lib.Db.Begin()

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var examnotice model.ExamNotice
	json.Unmarshal([]byte(body), &examnotice)

	if examnotice.ExamId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	Count := 0

	tx.QueryRow("select  count(1) 'Count' from examnotice  where  ExamId=? ", examnotice.ExamId).Scan(&Count)

	if Count > 0 {
		sqlStr := "delete from  examnotice   where ExamId=? "

		ret, err := tx.Exec(sqlStr, examnotice.ExamId)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		if n <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
	}

	sqlStr := "insert into  examnotice(ExamId,Title,SchoolName,CourseName,CourseCode,Context) VALUES (?,?,?,?,?,?) "

	ret, err := tx.Exec(sqlStr, examnotice.ExamId, examnotice.Title, examnotice.SchoolName, examnotice.CourseName, examnotice.CourseCode, examnotice.Context)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		tx.Rollback()
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		tx.Rollback()
		return
	}

	if n > 0 {

		err = tx.Commit()

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "提交事务失败",
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
}
func DelExamNotice(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var examNotice model.ExamNotice
	json.Unmarshal([]byte(body), &examNotice)

	sqlStr := "delete from  examnotice   where ExamId=? "

	ret, err := lib.Db.Exec(sqlStr, examNotice.ExamId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
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

func GetExamNoticeByExamId(c *gin.Context) {

	ExamId := c.Query("ExamId")

	if ExamId == "0" || ExamId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	examnotice := new(model.ExamNotice)

	lib.Db.QueryRow("select Id,ExamId,Title,SchoolName,CourseName,CourseCode,Context from  examnotice where ExamId=?", ExamId).Scan(&examnotice.Id, &examnotice.ExamId, &examnotice.Title, &examnotice.SchoolName, &examnotice.CourseName, &examnotice.CourseCode, &examnotice.Context)

	if examnotice.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "成功",
			"data": examnotice,
		})
	}

}
