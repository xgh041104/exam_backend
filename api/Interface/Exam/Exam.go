package exam

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	goutils "github.com/typa01/go-utils"
)

func AddExam(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var exam model.Exam
	json.Unmarshal([]byte(body), &exam)

	sqlStr := " insert into  exam(SchoolId,ExamName,ExamDescribe,ExamStatus,FaceVerify,TeacherId) values(?,?,?,?,?,?)  "

	ret, err := lib.Db.Exec(sqlStr, exam.SchoolId, exam.ExamName, exam.ExamDescribe, exam.ExamStatus, exam.FaceVerify, exam.TeacherId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	examid, err := ret.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if examid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": "{\"ExamId\":" + lib.Strval(examid) + "}",
	})
}

func EditExam(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var exam model.Exam
	json.Unmarshal([]byte(body), &exam)
	if exam.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	Count := 0

	sessionsql := " select Count(1) 'Count' from examsession a " +
		" left join exam b on a.ExamId=b.Id " +
		" where " +
		" (  NOW()  BETWEEN a.StartTime AND a.EndTime  or  NOW()>a.EndTime) and a.ExamId=?  and b.examStatus=1 "

	lib.Db.QueryRow(sessionsql, exam.Id).Scan(&Count)

	if Count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该考试已开始或已完成，不能编辑",
			"data": "{}",
		})
		return
	}

	sqlStr := " update    exam set SchoolId=?,ExamName=?,ExamDescribe=?,ExamStatus=?,FaceVerify=?  where  Id=?   "

	ret, err := lib.Db.Exec(sqlStr, exam.SchoolId, exam.ExamName, exam.ExamDescribe, exam.ExamStatus, exam.FaceVerify, exam.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if n < 0 {
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

func DelExam(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var exam model.DelExam
	json.Unmarshal([]byte(body), &exam)
	if exam.ExamId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	// Count := 0
	// lib.Db.QueryRow(" select Count(1) 'Count' from examsession where (  NOW()  BETWEEN StartTime AND EndTime  or  NOW()>EndTime) and ExamId=? ", exam.ExamId).Scan(&Count)
	// if Count > 0 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 0,
	// 		"msg":  "该考试已开始或已完成，不能删除",
	// 		"data": "{}",
	// 	})
	// 	return
	// }
	tx, err := lib.Db.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	//mark 删除之前判断是否有学生已考试的判断不要了
	// sqlquery := " select  COUNT(Score) 'isdel' from examsturesult where  ExamId=? and Score>=0"

	// isdel := 0
	// tx.QueryRow(sqlquery, exam.ExamId).Scan(&isdel)

	// if isdel == 0 {

	sqlexamdel := "update exam set IsDel=1 where Id=?"

	ret, err := tx.Exec(sqlexamdel, exam.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()

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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	sqlexamdel = "update examsession set IsDel=1 where ExamId=?"

	ret, err = tx.Exec(sqlexamdel, exam.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected()

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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	sqlexamdel = "update examsturesult set IsDel=1 where ExamId=?"

	ret, err = tx.Exec(sqlexamdel, exam.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected()

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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
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
	return
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 0,
	// 		"msg":  "考试已归档不能删除",
	// 		"data": "{}",
	// 	})
	// 	return
	// }
}

func AddExamSession(c *gin.Context) {

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var examsession model.ExamSession
	json.Unmarshal([]byte(body), &examsession)

	sqlStr := " insert into  examsession(ExamId,StartTime,EndTime,TestPaperId) values(?,?,?,?)  "

	ret, err := tx.Exec(sqlStr, examsession.ExamId, examsession.StartTime, examsession.EndTime, examsession.TestPaperId)

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	examsessionid, err := ret.LastInsertId()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if examsessionid == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
	}

	rows, err := tx.Query("select StudentId from examstudent where ExamId=?", examsession.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	//var StudentIdArr []int
	StudentIdArr := make([]int, 0)
	for rows.Next() {
		StudentId := 0
		err = rows.Scan(&StudentId)
		StudentIdArr = append(StudentIdArr, StudentId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

	}
	rows.Close()
	for i := 0; i < len(StudentIdArr); i++ {
		sqlStr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  "
		ret, err := tx.Exec(sqlStr, StudentIdArr[i], examsession.ExamId, examsessionid, -1)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
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
func EditExamSession(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var examsession model.ExamSession
	json.Unmarshal([]byte(body), &examsession)

	if examsession.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	Count := 0

	sessionsql := " select Count(1) 'Count' from examsession a " +
		" left join exam b on a.ExamId=b.Id " +
		" where " +
		" (  NOW()  BETWEEN a.StartTime AND a.EndTime  or  NOW()>a.EndTime) and a.Id=?  and b.examStatus=1 "
	lib.Db.QueryRow(sessionsql, examsession.Id).Scan(&Count)
	if Count == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该场次已开考或已完成，不能编辑",
			"data": "{}",
		})
		return
	}

	// mark 修改考场时，判断该试卷有没有在考试，如果有，则新建一个考场
	sqlStr := " update  examsession  set ExamId=?,StartTime=?,EndTime=?,TestPaperId=?   where  Id=? "

	ret, err := lib.Db.Exec(sqlStr, examsession.ExamId, examsession.StartTime, examsession.EndTime, examsession.TestPaperId, examsession.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if n < 0 {
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

func DelExamSession(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var examsession model.ExamSession
	json.Unmarshal([]byte(body), &examsession)

	if examsession.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	Count := 0

	sessionsql := " select Count(1) 'Count' from examsession a " +
		" left join exam b on a.ExamId=b.Id " +
		" where " +
		" (  NOW()  BETWEEN a.StartTime AND a.EndTime  or  NOW()>a.EndTime) and a.Id=?  and b.examStatus=1 "
	lib.Db.QueryRow(sessionsql, examsession.Id).Scan(&Count)
	if Count == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该场次已开考或已完成，不能编辑",
			"data": "{}",
		})
		return
	}

	tx, err := lib.Db.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	sqlStr := " delete from   examsession     where  Id=? "

	ret, err := tx.Exec(sqlStr, examsession.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if n == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
	}

	sqlStr = " delete from   examsturesult     where  ExamId=? and ExamSessionId=? "

	ret, err = tx.Exec(sqlStr, examsession.ExamId, examsession.Id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if n < 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	sqlStr = " delete from   examanswersheet     where  ExamId=? and ExamSessionId=? "

	ret, err = tx.Exec(sqlStr, examsession.ExamId, examsession.Id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if n < 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
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
func AddExamStudent(c *gin.Context) {

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var examstudent model.ExamStudent
	json.Unmarshal([]byte(body), &examstudent)

	var sessionidarr []int
	rowsnew, err := tx.Query("select Id from examsession where ExamId=?  ", examstudent.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	for rowsnew.Next() {
		examsessionid := 0
		rowsnew.Scan(&examsessionid)
		sessionidarr = append(sessionidarr, examsessionid)
	}

	sqladdStr := " insert into  examstudent(ExamId,StudentId ) values(?,?)  "
	for i := 0; i < len(examstudent.AddStudentIdArr); i++ {

		ret, err := tx.Exec(sqladdStr, examstudent.ExamId, examstudent.AddStudentIdArr[i])
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
		}

		sqlStr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  "

		for j := 0; j < len(sessionidarr); j++ {

			ret, err := tx.Exec(sqlStr, examstudent.AddStudentIdArr[i], examstudent.ExamId, sessionidarr[j], -1)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}

			if n <= 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
		}

	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
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

func EditExamStudent(c *gin.Context) {

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var examstudent model.ExamStudent
	json.Unmarshal([]byte(body), &examstudent)

	Count := 0

	sessionsql := " select Count(1) 'Count' from examsession a " +
		" left join exam b on a.ExamId=b.Id " +
		" where " +
		" (  NOW()  BETWEEN a.StartTime AND a.EndTime  or  NOW()>a.EndTime) and a.ExamId=?  and b.examStatus=1 "

	tx.QueryRow(sessionsql, examstudent.ExamId).Scan(&Count)

	if Count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该考试已开始或已完成，不能编辑",
			"data": "{}",
		})
		return
	}

	var sessionidarr []int
	rowsnew, err := tx.Query("select Id from examsession where ExamId=?  ", examstudent.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	for rowsnew.Next() {
		examsessionid := 0
		rowsnew.Scan(&examsessionid)
		sessionidarr = append(sessionidarr, examsessionid)
	}

	sqldelStr := " delete  from   examstudent where ExamId=? and StudentId=? "
	for i := 0; i < len(examstudent.RemoveStudentIdArr); i++ {
		ret, err := tx.Exec(sqldelStr, examstudent.ExamId, examstudent.RemoveStudentIdArr[i])
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		if n < 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
		}
		sqlStr := "  delete from   examsturesult where StudentId=? and ExamId=? and ExamSessionId=?  "

		for j := 0; j < len(sessionidarr); j++ {

			ret, err := tx.Exec(sqlStr, examstudent.RemoveStudentIdArr[i], examstudent.ExamId, sessionidarr[j])
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}

			if n < 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
		}
	}

	sqladdStr := " insert into  examstudent(ExamId,StudentId ) values(?,?)  "
	for i := 0; i < len(examstudent.AddStudentIdArr); i++ {

		ret, err := tx.Exec(sqladdStr, examstudent.ExamId, examstudent.AddStudentIdArr[i])
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
		}

		sqlStr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  "

		for j := 0; j < len(sessionidarr); j++ {

			ret, err := tx.Exec(sqlStr, examstudent.AddStudentIdArr[i], examstudent.ExamId, sessionidarr[j], -1)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}

			if n <= 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
		}

	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
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

func GetExamBySchoolId(c *gin.Context) {

	schoolid := c.Query("SchoolId")

	rows, err := lib.Db.Query("select Id,SchoolId,ExamName,ExamDescribe,ExamStatus,FaceVerify,ReviewFlag from exam where SchoolId=?  and isdel=0 ", schoolid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	layout := "2006-01-02 15:04:05"
	var examViewarr []*model.ExamView
	for rows.Next() {
		ExamView := new(model.ExamView)

		rows.Scan(&ExamView.Id, &ExamView.SchoolId, &ExamView.ExamName, &ExamView.ExamDescribe, &ExamView.ExamStatus, &ExamView.FaceVerify, &ExamView.ReviewFlag)
		examViewarr = append(examViewarr, ExamView)

		//examViewarr = append(examViewarr, ExamView)
	}

	for i := 0; i < len(examViewarr); i++ {
		ExamView := examViewarr[i]
		//找到学生集合
		rownewstu, err := lib.Db.Query("select  a.StudentId,b.TrueName from  examstudent a left join student b on a.StudentId=b.Id  where  a.ExamId=?", ExamView.Id)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		ExamView.ExamStudentArr = make([]*model.ExamStudentView, 0)
		for rownewstu.Next() {
			examstudentView := new(model.ExamStudentView)

			rownewstu.Scan(&examstudentView.StudentId, &examstudentView.TrueName)
			if examstudentView.StudentId == 0 {
				continue
			}
			ExamView.ExamStudentArr = append(ExamView.ExamStudentArr, examstudentView)
		}

		//找到考试场次集合.

		rownew, err := lib.Db.Query("select a.Id,a.ExamId,a.StartTime,a.EndTime,a.TestPaperId,b.TestPaperName,b.ExamDuration,b.FullMarks,b.PassScore from  examsession a left join testpaper b on a.TestPaperId=b.Id where  a.ExamId=?", ExamView.Id)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		ExamView.ExamSessionArr = make([]*model.ExamSessionView, 0)
		flagarr := make([]int, 0)
		for rownew.Next() {
			examSessionView := new(model.ExamSessionView)

			rownew.Scan(&examSessionView.Id, &examSessionView.ExamId, &examSessionView.StartTime, &examSessionView.EndTime, &examSessionView.TestPaperId, &examSessionView.TestPaperName, &examSessionView.ExamDuration, &examSessionView.FullMarks, &examSessionView.PassScore)
			if examSessionView.Id == 0 {
				continue
			}

			starttime, _ := time.Parse(examSessionView.StartTime, layout)
			endtime, _ := time.Parse(examSessionView.EndTime, layout)
			nowtime := time.Now()
			if nowtime.After(starttime) && nowtime.Before(endtime) {
				examSessionView.State = 1
				flagarr = append(flagarr, 1)
			} else if nowtime.After(endtime) {
				examSessionView.State = 2
				flagarr = append(flagarr, 2)
			} else if nowtime.Before(starttime) {
				examSessionView.State = 0
				flagarr = append(flagarr, 0)
			}
			ExamView.ExamSessionArr = append(ExamView.ExamSessionArr, examSessionView)
		}

		if lib.CountOccurrences(flagarr, 1) > 0 {
			ExamView.State = 1
		} else if len(flagarr) == 0 {
			ExamView.State = 0
		} else if lib.CountOccurrences(flagarr, 2) == len(flagarr) {
			ExamView.State = 2
		} else if lib.CountOccurrences(flagarr, 0) == len(flagarr) {
			ExamView.State = 0
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": examViewarr,
	})
}

func GetExamReViewBySchoolId(c *gin.Context) {

	schoolid := c.Query("SchoolId")

	rows, err := lib.Db.Query("select Id,SchoolId,ExamName,ExamDescribe,ExamStatus,FaceVerify,ReviewFlag from exam where SchoolId=?  and isdel=0 and ReviewFlag=0", schoolid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	layout := "2006-01-02 15:04:05"
	var examViewarr []*model.ExamView
	examViewarr = make([]*model.ExamView, 0)
	for rows.Next() {
		ExamView := new(model.ExamView)

		rows.Scan(&ExamView.Id, &ExamView.SchoolId, &ExamView.ExamName, &ExamView.ExamDescribe, &ExamView.ExamStatus, &ExamView.FaceVerify, &ExamView.ReviewFlag)
		examViewarr = append(examViewarr, ExamView)

	}

	for i := 0; i < len(examViewarr); i++ {
		ExamView := examViewarr[i]
		//找到学生集合
		rownewstu, err := lib.Db.Query("select  a.StudentId,b.TrueName from  examstudent a left join student b on a.StudentId=b.Id  where  a.ExamId=?", ExamView.Id)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		ExamView.ExamStudentArr = make([]*model.ExamStudentView, 0)
		for rownewstu.Next() {
			examstudentView := new(model.ExamStudentView)

			rownewstu.Scan(&examstudentView.StudentId, &examstudentView.TrueName)
			if examstudentView.StudentId == 0 {
				continue
			}
			ExamView.ExamStudentArr = append(ExamView.ExamStudentArr, examstudentView)
		}

		//找到考试场次集合.

		rownew, err := lib.Db.Query("select a.Id,a.ExamId,a.StartTime,a.EndTime,a.TestPaperId,b.TestPaperName,b.ExamDuration,b.FullMarks,b.PassScore from  examsession a left join testpaper b on a.TestPaperId=b.Id where  a.ExamId=?", ExamView.Id)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		ExamView.ExamSessionArr = make([]*model.ExamSessionView, 0)
		flagarr := make([]int, 0)
		for rownew.Next() {
			examSessionView := new(model.ExamSessionView)

			rownew.Scan(&examSessionView.Id, &examSessionView.ExamId, &examSessionView.StartTime, &examSessionView.EndTime, &examSessionView.TestPaperId, &examSessionView.TestPaperName, &examSessionView.ExamDuration, &examSessionView.FullMarks, &examSessionView.PassScore)
			if examSessionView.Id == 0 {
				continue
			}

			starttime, _ := time.Parse(examSessionView.StartTime, layout)
			endtime, _ := time.Parse(examSessionView.EndTime, layout)
			nowtime := time.Now()
			if nowtime.After(starttime) && nowtime.Before(endtime) {
				examSessionView.State = 1
				flagarr = append(flagarr, 1)
			} else if nowtime.After(endtime) {
				examSessionView.State = 2
				flagarr = append(flagarr, 2)
			} else if nowtime.Before(starttime) {
				examSessionView.State = 0
				flagarr = append(flagarr, 0)
			}
			ExamView.ExamSessionArr = append(ExamView.ExamSessionArr, examSessionView)
		}
		if lib.CountOccurrences(flagarr, 1) > 0 {
			ExamView.State = 1
		} else if len(flagarr) == 0 {
			ExamView.State = 0
		} else if lib.CountOccurrences(flagarr, 2) == len(flagarr) {
			ExamView.State = 2
		} else if lib.CountOccurrences(flagarr, 0) == len(flagarr) {
			ExamView.State = 0
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": examViewarr,
	})
}

func EditReViewExamByExamId(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var exam model.DelExam
	json.Unmarshal([]byte(body), &exam)
	if exam.ExamId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	ret, err := lib.Db.Exec("update exam set ReviewFlag=1 where Id=?", exam.ExamId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if n < 0 {
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

func GetExamByExamId(c *gin.Context) {

	examid := c.Query("ExamId")

	if examid == "" || examid == "0" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	ExamView := new(model.ExamView)
	err := lib.Db.QueryRow("select Id,SchoolId,ExamName,ExamDescribe,ExamStatus,FaceVerify from exam where Id=?", examid).Scan(&ExamView.Id, &ExamView.SchoolId, &ExamView.ExamName, &ExamView.ExamDescribe, &ExamView.ExamStatus, &ExamView.FaceVerify)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	//找到学生集合
	rownewstu, err := lib.Db.Query("select  a.StudentId,b.TrueName from  examstudent a left join student b on a.StudentId=b.Id  where  a.ExamId=?", ExamView.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	ExamView.ExamStudentArr = make([]*model.ExamStudentView, 0)
	for rownewstu.Next() {
		examstudentView := new(model.ExamStudentView)

		rownewstu.Scan(&examstudentView.StudentId, &examstudentView.TrueName)
		if examstudentView.StudentId == 0 {
			continue
		}
		ExamView.ExamStudentArr = append(ExamView.ExamStudentArr, examstudentView)
	}

	//找到考试场次集合.

	rownew, err := lib.Db.Query("select a.Id,a.ExamId,a.StartTime,a.EndTime,a.TestPaperId,b.TestPaperName,b.ExamDuration,b.FullMarks,b.PassScore from  examsession a left join testpaper b on a.TestPaperId=b.Id where  a.ExamId=?", ExamView.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	ExamView.ExamSessionArr = make([]*model.ExamSessionView, 0)
	for rownew.Next() {
		examSessionView := new(model.ExamSessionView)

		rownew.Scan(&examSessionView.Id, &examSessionView.ExamId, &examSessionView.StartTime, &examSessionView.EndTime, &examSessionView.TestPaperId, &examSessionView.TestPaperName, &examSessionView.ExamDuration, &examSessionView.FullMarks, &examSessionView.PassScore)
		if examSessionView.Id == 0 {
			continue
		}
		ExamView.ExamSessionArr = append(ExamView.ExamSessionArr, examSessionView)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": ExamView,
	})
}

func GetExamOverView(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	strsqlQuery := "select a.Id 'ExamSessionId',a.ExamId, CONCAT(d.ExamName,'/',b.TestPaperName) as 'ExamSession',e.MajorName,CONCAT(b.PassScore,'/',b.FullMarks) 'Score',a.StartTime" +
		" ,(select COALESCE(avg(Score),0)  from  examsturesult   where    ExamId=a.ExamId   and   ExamSessionId=a.Id and Score<>-1	) 'AvgScore', " +
		"	(SELECT COUNT(1) 'wks'  FROM `examsturesult`  where ExamId=a.ExamId and  ExamSessionId=a.Id and Score=-1) 'UnexaminedNum'," +
		"(SELECT COUNT(1) 'yks'  FROM `examsturesult`  where ExamId=a.ExamId and  ExamSessionId=a.Id and Score<>-1 and ExamStatus=1) 'ExaminedNum',b.ExamDuration" +
		"  from examsession a " +
		" left join testpaper b on a.TestPaperId=b.Id " +
		" left JOIN exam d on a.ExamId=d.Id " +
		" LEFT JOIN major e on b.MajorId=e.MajorId WHERE d.SchoolId=?  and d.ReviewFlag=1"

	rows, err := lib.Db.Query(strsqlQuery, SchoolId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	var examScoreViewarr []*model.ExamScoreView
	for rows.Next() {
		newexamScoreView := new(model.ExamScoreView)

		rows.Scan(&newexamScoreView.ExamSessionId, &newexamScoreView.ExamId, &newexamScoreView.ExamSession, &newexamScoreView.MajorName, &newexamScoreView.Score, &newexamScoreView.StartTime, &newexamScoreView.AvgScore, &newexamScoreView.UnexaminedNum, &newexamScoreView.ExaminedNum, &newexamScoreView.ExamDuration)
		newexamScoreView.ExamPeopleSumNum = newexamScoreView.UnexaminedNum + newexamScoreView.ExaminedNum
		examScoreViewarr = append(examScoreViewarr, newexamScoreView)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": examScoreViewarr,
	})
}

func GetExamSessionResult(c *gin.Context) {

	ExamId := c.Query("ExamId")
	ExamSessionId := c.Query("ExamSessionId")
	strsqlQuery := "  select a.Id,b.TrueName,b.IDNumber,a.StudentId,a.ExamId, a.ExamSessionId,COALESCE( a.StartExamTime,'') 'StartExamTime' " +
		",COALESCE( a.EndExamTime,'') 'EndExamTime',a.Score,COALESCE(a.ExamStatus,0)  'ExamStatus', COALESCE(a.ExamType,0) 'ExamType'" +
		"	from examsturesult a" +
		" LEFT JOIN student b on a.StudentId=b.Id" +
		"	where ExamId=? and  ExamSessionId=?    ORDER BY a.score asc"
	rows, err := lib.Db.Query(strsqlQuery, ExamId, ExamSessionId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	var ExamScoreStuViewarr []*model.ExamScoreStuView
	for rows.Next() {
		ExamScoreStuView := new(model.ExamScoreStuView)
		rows.Scan(&ExamScoreStuView.Id, &ExamScoreStuView.TrueName, &ExamScoreStuView.IDNumber, &ExamScoreStuView.StudentId, &ExamScoreStuView.ExamId, &ExamScoreStuView.ExamSessionId, &ExamScoreStuView.StartExamTime, &ExamScoreStuView.EndExamTime, &ExamScoreStuView.Score, &ExamScoreStuView.ExamStatus, &ExamScoreStuView.ExamType)
		st, err := time.Parse(lib.TimeLayoutStr, ExamScoreStuView.StartExamTime)
		if err == nil {
			et, err := time.Parse(lib.TimeLayoutStr, ExamScoreStuView.EndExamTime)
			if err == nil {
				minutes := st.Sub(et).Minutes()
				intminutes := int32(minutes)
				ExamScoreStuView.UseTime = math.Abs(float64(intminutes))
			}
		}

		ExamScoreStuViewarr = append(ExamScoreStuViewarr, ExamScoreStuView)

	}

	for i := 0; i < len(ExamScoreStuViewarr); i++ {
		ExamScoreStuView := ExamScoreStuViewarr[i]
		examimgarr := make([]*model.ExamImage, 0)
		temprows, err := lib.Db.Query("SELECT Id,ImagePath,CreateTime FROM examimage where ExamId=? and ExamSessionId=? and  StudentId=?", ExamScoreStuView.ExamId, ExamScoreStuView.ExamSessionId, ExamScoreStuView.StudentId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		for temprows.Next() {
			examimg := new(model.ExamImage)
			err := temprows.Scan(&examimg.Id, &examimg.ImagePath, &examimg.CreateTime)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败",
					"data": "{}",
				})
				return
			}
			examimg.ExamId = ExamScoreStuView.ExamId
			examimg.ExamSessionId = ExamScoreStuView.ExamSessionId
			examimg.StudentId = ExamScoreStuView.StudentId
			examimgarr = append(examimgarr, examimg)
		}
		ExamScoreStuView.ExamImageArr = examimgarr
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": ExamScoreStuViewarr,
	})
}

func GetCurrentExamSessionBKStudent(c *gin.Context) {

	ExamSessionId := c.Query("ExamSessionId")

	strsqlQuery := " 	SELECT * from ( 	SELECT a.StudentId,b.TrueName FROM examsturesult a " +
		"	left  JOIN student b on a.StudentId=b.Id " +
		"	left join  examsession c on a.ExamSessionId=c.Id " +
		"	LEFT JOIN testpaper d on c.TestPaperId=d.Id " +
		"	 where a.ExamStatus=1 and a.Score<d.PassScore   and   a.ExamSessionId=? " +
		"	UNION   " +
		"	 SELECT   a.StudentId,b.TrueName FROM examsturesult a " +
		"	 left  JOIN student b on a.StudentId=b.Id " +
		"	 left join  examsession c on a.ExamSessionId=c.Id " +
		"	where ExamStatus=0  and  NOW() >= c.EndTime and NOW() >c.StartTime   and   a.ExamSessionId=?  			) newtable " +
		"	where newtable.StudentId not in(		" +
		"  SELECT StudentId from examretest WHERE OldExamSessionId=? and Status=0 and   EndExamTime> NOW()  )"

	rows, err := lib.Db.Query(strsqlQuery, ExamSessionId, ExamSessionId, ExamSessionId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	var ExamStudentViewarr []*model.ExamStudentView
	for rows.Next() {
		examStudentView := new(model.ExamStudentView)
		rows.Scan(&examStudentView.StudentId, &examStudentView.TrueName)
		ExamStudentViewarr = append(ExamStudentViewarr, examStudentView)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": ExamStudentViewarr,
	})

}
func AddExamReset(c *gin.Context) {

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	var examRetest model.ExamRetest
	json.Unmarshal([]byte(body), &examRetest)

	if len(examRetest.StudentIdArr) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	for i := 0; i < len(examRetest.StudentIdArr); i++ {
		ret, err := tx.Exec("insert into examretest(StudentId,OldExamSessionId,OldExamId,StartExamTime,EndExamTime)  values(?,?,?,?,?) ", examRetest.StudentIdArr[i], examRetest.OldExamSessionId, examRetest.OldExamId, examRetest.StartExamTime, examRetest.EndExamTime)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		n, err := ret.RowsAffected()
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
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}

		retnew, err := tx.Exec("update examsturesult set  ExamType=1 where StudentId=? and ExamId=? and ExamSessionId=?", examRetest.StudentIdArr[i], examRetest.OldExamId, examRetest.OldExamSessionId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		nnew, err := retnew.RowsAffected()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		if nnew <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		tx.Rollback()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": "{}",
	})
}

func GetExamResetByExamSessionId(c *gin.Context) {

	ExamSessionId := c.Query("ExamSessionId")

	strsqlQuery := " 	SELECT   a.Id,a.StudentId,b.TrueName,a.OldExamId,a.OldExamSessionId,a.StartExamTime,a.EndExamTime,Score, a.`Status`,d.TestPaperName " +
		" from examretest a  " +
		" left  JOIN student b on a.StudentId=b.Id  " +
		"		 left join  examsession c on a.OldExamSessionId=c.Id " +
		" 		LEFT JOIN testpaper d on c.TestPaperId=d.Id  	" +
		" WHERE a.OldExamSessionId=?  "

	rows, err := lib.Db.Query(strsqlQuery, ExamSessionId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	var ExamRetestViewarr []*model.ExamRetestView
	for rows.Next() {
		examRetestView := new(model.ExamRetestView)
		rows.Scan(&examRetestView.Id, &examRetestView.StudentId, &examRetestView.TrueName, &examRetestView.OldExamId, &examRetestView.OldExamSessionId, &examRetestView.StartExamTime, &examRetestView.EndExamTime, &examRetestView.Score, &examRetestView.Status, &examRetestView.TestPaperName)
		if examRetestView.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		ExamRetestViewarr = append(ExamRetestViewarr, examRetestView)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": ExamRetestViewarr,
	})
}

func GetStudentExamInfo(c *gin.Context) {

	StudentId := c.Query("StudentId")

	strsqlQuery := " SELECT a.Id,a.StudentId,a.ExamId,b.ExamName,a.ExamSessionId,COALESCE(a.StartExamTime, '') 'StartExamTime',COALESCE(a.EndExamTime,'')  'EndExamTime',a.Score,a.ExamStatus,a.ExamType,   " +
		" (  CASE " +
		" 	WHEN (a.ExamStatus=0 and a.ExamType=0 and NOW()>c.EndTime) ||  (a.ExamType=1 and NOW()>e.EndExamTime ) THEN 0 " +
		" 	  	WHEN ( a.ExamType=2 || (a.ExamStatus=1  and  a.ExamType=0  )  )  THEN 1    " +
		" 		 		 	WHEN (a.ExamStatus=0 || a.ExamType=0)   and NOW()<=c.EndTime   and NOW()>=c.StartTime THEN 2  " +
		" 			 		 	WHEN (a.ExamStatus=0 || a.ExamType=0)    and NOW()<c.StartTime THEN 3 " +
		" 						WHEN a.ExamType=1 and NOW()<=e.EndExamTime   and NOW()>=e.StartExamTime     THEN 4 " +
		" 	ELSE 0 END  ) as 'ExamZT', " +
		" 	d.MajorId,d.CourseId,COALESCE(f.MajorName,'') 'MajorName' ,COALESCE(g.CourseName,'') 'CourseName', " +
		" 	d.FullMarks,d.PassScore, " +
		" 	d.ExamDuration,	 (select SessionNum from ( SELECT ROW_NUMBER() OVER(order BY StartTime asc) as 'SessionNum',Id FROM examsession where ExamId=a.ExamId) newtable where newtable.id=c.Id) 'SessionNum' , " +
		" (select COUNT(1) from testpaperquestion where TestPaperId=d.Id ) 'QuestionNum', " +
		" 		COALESCE(e.StartExamTime,'') 'ResetStartExamTime',COALESCE(e.EndExamTime,'') 'ResetEndExamTime', " +
		" c.StartTime 'SessionStartExamTime', c.EndTime  'SessionEndExamTime',b.FaceVerify " +
		"  FROM examsturesult a " +
		" left join exam b on a.ExamId=b.Id " +
		" left join examsession c  on a.ExamSessionId=c.Id " +
		" LEFT JOIN testpaper d on c.TestPaperId=d.Id " +
		" LEFT JOIN  examretest e  on a.ExamSessionId=e.OldExamSessionId and a.StudentId=e.StudentId " +
		" LEFT JOIN major f on d.MajorId=f.MajorId " +
		"  LEFT JOIN course g on d.CourseId=g.Id " +
		" where a.StudentId=?  and b.ExamStatus=1 and b.ReviewFlag=1 and a.IsDel=0  ORDER BY ExamZT  asc"

	rows, err := lib.Db.Query(strsqlQuery, StudentId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	var studentExamInfoArr []*model.StudentExamInfo
	for rows.Next() {
		studentExamInfo := new(model.StudentExamInfo)
		err := rows.Scan(&studentExamInfo.Id, &studentExamInfo.StudentId, &studentExamInfo.ExamId, &studentExamInfo.ExamName, &studentExamInfo.ExamSessionId,
			&studentExamInfo.StartExamTime, &studentExamInfo.EndExamTime, &studentExamInfo.Score, &studentExamInfo.ExamStatus, &studentExamInfo.ExamType,
			&studentExamInfo.ExamZT, &studentExamInfo.MajorId, &studentExamInfo.CourseId, &studentExamInfo.MajorName, &studentExamInfo.CourseName, &studentExamInfo.FullMarks, &studentExamInfo.PassScore,
			&studentExamInfo.ExamDuration, &studentExamInfo.SessionNum, &studentExamInfo.QuestionNum, &studentExamInfo.ResetStartExamTime, &studentExamInfo.ResetEndExamTime, &studentExamInfo.SessionStartExamTime, &studentExamInfo.SessionEndExamTime, &studentExamInfo.FaceVerify)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		if studentExamInfo.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		studentExamInfoArr = append(studentExamInfoArr, studentExamInfo)
	}
	rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": studentExamInfoArr,
	})
}

func GetStudentExamDetails(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic", err)
		}
	}()

	StudentId := c.Query("StudentId")
	ExamSessionId := c.Query("ExamSessionId")
	ExamId := c.Query("ExamId")

	strsqlQuery := "	SELECT a.Id,a.StudentId,a.ExamId,b.ExamName,a.ExamSessionId,  c.TestPaperId,d.TestPaperName,d.TestPaperType, " +
		"  COALESCE(f.MajorName,'') 'MajorName' ,COALESCE(g.CourseName,'') 'CourseName',  " +
		"		d.FullMarks,d.PassScore, 	d.ExamDuration, " +
		"		(select SessionNum from ( SELECT ROW_NUMBER() OVER(order BY StartTime asc) as 'SessionNum',Id FROM examsession  where ExamId=a.ExamId) newtable where newtable.id=c.Id) 'SessionNum' " +
		"				,COALESCE(h.SchoolName,'') 'SchoolName',COALESCE(i.TeacherName,'') 'TeacherName',b.FaceVerify " +
		"			 FROM examsturesult a  " +
		"			left join exam b on a.ExamId=b.Id " +
		"			left join examsession c  on a.ExamSessionId=c.Id " +
		"			LEFT JOIN testpaper d on c.TestPaperId=d.Id " +
		"			LEFT JOIN major f on d.MajorId=f.MajorId " +
		"			 LEFT JOIN course g on d.CourseId=g.CollegeId" +
		"				left join school h on d.SchoolId=h.Id  	 LEFT JOIN  teacher i on d.TeacherId=i.TeacherId " +
		"				where  a.StudentId=? and a.ExamSessionId=? and a.ExamId=? "

	returnEntity := new(model.StudentExamDetails)

	err := lib.Db.QueryRow(strsqlQuery, StudentId, ExamSessionId, ExamId).Scan(&returnEntity.Id, &returnEntity.StudentId, &returnEntity.ExamId, &returnEntity.ExamName, &returnEntity.ExamSessionId,
		&returnEntity.TestPaperId, &returnEntity.TestPaperName, &returnEntity.TestPaperType, &returnEntity.MajorName, &returnEntity.CourseName, &returnEntity.FullMarks, &returnEntity.PassScore,
		&returnEntity.ExamDuration, &returnEntity.SessionNum, &returnEntity.SchoolName, &returnEntity.TeacherName, &returnEntity.FaceVerify)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		log.Error(666, err)
		return
	}
	if returnEntity.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		log.Error(7, err)
		return
	}

	returnEntity.AnswerSheetarr = make([]*model.AnswerSheet, 0)

	returnEntity.TestPaperQuestionTypeOver = make([]*model.QuestionTypeOver, 0)

	for i := 1; i <= 5; i++ {
		temparr := new(model.QuestionTypeOver)
		temparr.QuestionIdNum = 0
		temparr.QuestionScore = 0
		temparr.QuestionType = i
		returnEntity.TestPaperQuestionTypeOver = append(returnEntity.TestPaperQuestionTypeOver, temparr)
	}

	returnEntity.TestPaperQuestionViewFile = make([]*model.TestPaperQuestionViewFile, 0)

	strquerstion := " select	a.QuestionId,COALESCE(b.QuestionName,'') 'QuestionName',COALESCE(b.QuestionPoolId,0) 'QuestionPoolId'  ,COALESCE(b.QuestionType,0) 'QuestionType' " +
		",COALESCE(b.QuestionContent,'')'QuestionContent' ,COALESCE(a.QuestionScore,0) 'QuestionScore',COALESCE(b.Digree,0) 'Digree',COALESCE(b.Answer,0)'Answer' from testpaperquestion a " +
		"    LEFT JOIN  question b on a.QuestionId=b.QuestionId " +
		"		   where a.TestPaperId=? order by b.QuestionType asc "
	rows, err := lib.Db.Query(strquerstion, returnEntity.TestPaperId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		log.Error(2, err)
		return
	}
	for rows.Next() {
		testpaperEntity := new(model.TestPaperQuestionViewFile)
		err := rows.Scan(&testpaperEntity.QuestionId, &testpaperEntity.QuestionName, &testpaperEntity.QuestionPoolId, &testpaperEntity.QuestionType,
			&testpaperEntity.QuestionContent, &testpaperEntity.QuestionScore, &testpaperEntity.Digree, &testpaperEntity.Answer)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			log.Error(3, err)
			return
		}
		if testpaperEntity.QuestionType != 0 {
			returnEntity.TestPaperQuestionTypeOver[testpaperEntity.QuestionType-1].QuestionIdNum += 1
			returnEntity.TestPaperQuestionTypeOver[testpaperEntity.QuestionType-1].QuestionScore += testpaperEntity.QuestionScore
		}

		returnEntity.TestPaperQuestionViewFile = append(returnEntity.TestPaperQuestionViewFile, testpaperEntity)
	}
	rows.Close()
	for i := 0; i < len(returnEntity.TestPaperQuestionViewFile); i++ {

		testpaperEntity := returnEntity.TestPaperQuestionViewFile[i]
		testpaperEntity.FileInfo = make([]*model.FileInfo, 0)
		if testpaperEntity.QuestionType == 5 {
			fileid := testpaperEntity.QuestionContent

			fileinfo := new(model.FileInfo)
			lib.Db.QueryRow("select Id,FileType,FilePath,FileName from fileinfo    where Id=? ", fileid).Scan(&fileinfo.Id, &fileinfo.FileType, &fileinfo.FilePath, &fileinfo.FileName)
			if fileinfo.Id > 0 {
				testpaperEntity.FileInfo = append(testpaperEntity.FileInfo, fileinfo)
			}
		} else {
			rows1, err := lib.Db.Query("select b.Id,b.FileType,b.FilePath,b.FileName from questionrelation  a left join fileinfo b on a.FileId=b.Id where a.QuestionId=? ", testpaperEntity.QuestionId)

			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败" + err.Error(),
					"data": "{}",
				})
				log.Error(4444, err)
				return
			}
			for rows1.Next() {
				newfileinfo := new(model.FileInfo)
				err := rows1.Scan(&newfileinfo.Id, &newfileinfo.FileType, &newfileinfo.FilePath, &newfileinfo.FileName)
				if err != nil {

					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "失败",
						"data": "{}",
					})
					log.Error(5, err)
					return
				}
				testpaperEntity.FileInfo = append(testpaperEntity.FileInfo, newfileinfo)
			}
			rows1.Close()
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": returnEntity,
	})

}

func ExamStudentSumit(c *gin.Context) {

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败2" + err.Error(),
			"data": "{}",
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败3" + err.Error(),
			"data": "{}",
		})
		return
	}
	var sumitExamEntity model.SumitExamEntity
	json.Unmarshal([]byte(body), &sumitExamEntity)
	examsql := ""

	if sumitExamEntity.IsReTest == 1 {
		examsql = "update examsturesult set Score=?,ExamStatus=1,StartExamTime=?,EndExamTime=now(),ExamType=2 where   ExamId=? and ExamSessionId=? and StudentId=?"

	} else {
		oldscore := 0.0
		tx.QueryRow("select  Score from examsturesult where   ExamId=? and ExamSessionId=? and StudentId=? LIMIT 1", sumitExamEntity.ExamId, sumitExamEntity.ExamSessionId, sumitExamEntity.StudentId).
			Scan(&oldscore)

		if oldscore > float64(sumitExamEntity.Score) {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 2,
				"msg":  "提交成功",
				"data": "{}",
			})
			return
		}
		// "当前分数没有之前的大，考试数据不提交",
		examsql = "update examsturesult set Score=?,ExamStatus=1,StartExamTime=?,EndExamTime=now() where   ExamId=? and ExamSessionId=? and StudentId=?"
	}

	ret, err := tx.Exec(examsql, sumitExamEntity.Score, sumitExamEntity.StartExamTime, sumitExamEntity.ExamId, sumitExamEntity.ExamSessionId, sumitExamEntity.StudentId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败4" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败5" + err.Error(),
			"data": "{}",
		})
		return
	}
	if n < 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败6" + err.Error(),
			"data": "{}",
		})
		return
	}

	retdel, err := tx.Exec("delete from examanswersheet where ExamId=? and  ExamSessionId=? and  StudentId=? and TestPaperId=? ", sumitExamEntity.ExamId, sumitExamEntity.ExamSessionId, sumitExamEntity.TestPaperId, sumitExamEntity.StudentId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败7" + err.Error(),
			"data": "{}",
		})
		return
	}
	deln, err := retdel.RowsAffected()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败8" + err.Error(),
			"data": "{}",
		})
		return
	}
	if deln < 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败9" + err.Error(),
			"data": "{}",
		})
		return
	}

	for i := 0; i < len(sumitExamEntity.AnswerSheetArr); i++ {
		rettemp, err := tx.Exec("insert into  examanswersheet(ExamId,ExamSessionId,TestPaperId,StudentId,QuestionId,AnswerScore,AnswerSteps,IsTrue)  values(?,?,?,?,?,?,?,?)",
			sumitExamEntity.ExamId, sumitExamEntity.ExamSessionId, sumitExamEntity.TestPaperId, sumitExamEntity.StudentId, sumitExamEntity.AnswerSheetArr[i].QuestionId,
			sumitExamEntity.AnswerSheetArr[i].AnswerScore, sumitExamEntity.AnswerSheetArr[i].AnswerSteps, sumitExamEntity.AnswerSheetArr[i].IsTrue)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败7" + err.Error(),
				"data": "{}",
			})
			return
		}

		n, err := rettemp.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败8" + err.Error(),
				"data": "{}",
			})
			return
		}
		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败9" + err.Error(),
				"data": "{}",
			})
			return
		}

	}

	if sumitExamEntity.IsReTest == 1 {
		ret, err := tx.Exec(" delete from examretest where StudentId=? and OldExamSessionId=? and OldExamId=? ", sumitExamEntity.StudentId, sumitExamEntity.ExamSessionId, sumitExamEntity.ExamId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败4" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败5" + err.Error(),
				"data": "{}",
			})
			return
		}
		if n < 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败6" + err.Error(),
				"data": "{}",
			})
			return
		}
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

func GetStudentExamPaperOver(c *gin.Context) {

	StudentId := c.Query("StudentId")
	ExamSessionId := c.Query("ExamSessionId")
	ExamId := c.Query("ExamId")

	strsqlQuery := "	SELECT a.Id,a.StudentId,a.ExamId,b.ExamName,a.ExamSessionId,  c.TestPaperId,d.TestPaperName,d.TestPaperType, " +
		"  COALESCE(f.MajorName,'') 'MajorName' ,COALESCE(g.CourseName,'') 'CourseName',  " +
		"		d.FullMarks,d.PassScore, 	d.ExamDuration, " +
		"		(select SessionNum from ( SELECT ROW_NUMBER() OVER(order BY StartTime asc) as 'SessionNum',Id FROM examsession  where ExamId=a.ExamId) newtable where newtable.id=c.Id) 'SessionNum' " +
		"				,COALESCE(h.SchoolName,'') 'SchoolName',COALESCE(i.TeacherName,'') 'TeacherName',a.Score,a.StartExamTime,a.EndExamTime" +
		"			 FROM examsturesult a  " +
		"			left join exam b on a.ExamId=b.Id " +
		"			left join examsession c  on a.ExamSessionId=c.Id " +
		"			LEFT JOIN testpaper d on c.TestPaperId=d.Id " +
		"			LEFT JOIN major f on d.MajorId=f.MajorId " +
		"			 LEFT JOIN course g on d.CourseId=g.CollegeId" +
		"				left join school h on d.SchoolId=h.Id  	 LEFT JOIN  teacher i on d.TeacherId=i.TeacherId " +
		"				where  a.StudentId=? and a.ExamSessionId=? and a.ExamId=? "

	returnEntity := new(model.StudentExamAnswerDetails)

	lib.Db.QueryRow(strsqlQuery, StudentId, ExamSessionId, ExamId).Scan(&returnEntity.Id, &returnEntity.StudentId, &returnEntity.ExamId, &returnEntity.ExamName, &returnEntity.ExamSessionId,
		&returnEntity.TestPaperId, &returnEntity.TestPaperName, &returnEntity.TestPaperType, &returnEntity.MajorName, &returnEntity.CourseName, &returnEntity.FullMarks, &returnEntity.PassScore,
		&returnEntity.ExamDuration, &returnEntity.SessionNum, &returnEntity.SchoolName, &returnEntity.TeacherName, &returnEntity.Score, &returnEntity.StartExamTime, &returnEntity.EndExamTime)

	if returnEntity.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	returnEntity.TestPaperQuestionTypeOver = make([]*model.QuestionTypeOver, 0)

	for i := 1; i <= 5; i++ {
		temparr := new(model.QuestionTypeOver)
		temparr.QuestionIdNum = 0
		temparr.QuestionScore = 0
		temparr.QuestionType = i
		returnEntity.TestPaperQuestionTypeOver = append(returnEntity.TestPaperQuestionTypeOver, temparr)
	}

	returnEntity.AnswerSheetarr = make([]*model.AnswerSheet, 0)

	strsqlQueryAnswer := "SELECT QuestionId,AnswerScore,AnswerSteps,IsTrue FROM examanswersheet where StudentId=? and ExamSessionId=? and ExamId=? "
	rowanswerrow, err := lib.Db.Query(strsqlQueryAnswer, StudentId, ExamSessionId, ExamId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	for rowanswerrow.Next() {
		answerSheet := new(model.AnswerSheet)
		rowanswerrow.Scan(&answerSheet.QuestionId, &answerSheet.AnswerScore, &answerSheet.AnswerSteps, &answerSheet.IsTrue)
		returnEntity.AnswerSheetarr = append(returnEntity.AnswerSheetarr, answerSheet)
	}

	strquerstion := " select	a.QuestionId,COALESCE(b.QuestionName,'') 'QuestionName',COALESCE(b.QuestionPoolId,0) 'QuestionPoolId'  ,COALESCE(b.QuestionType,0) 'QuestionType' " +
		",COALESCE(b.QuestionContent,'')'QuestionContent' ,COALESCE(a.QuestionScore,0) 'QuestionScore',COALESCE(b.Digree,0) 'Digree',COALESCE(b.Answer,0)'Answer' from testpaperquestion a " +
		"    LEFT JOIN  question b on a.QuestionId=b.QuestionId " +
		"		   where a.TestPaperId=? order by b.QuestionType asc "
	rows, err := lib.Db.Query(strquerstion, returnEntity.TestPaperId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	returnEntity.TestPaperQuestionViewFile = make([]*model.TestPaperQuestionViewFile, 0)
	for rows.Next() {
		testpaperEntity := new(model.TestPaperQuestionViewFile)
		err := rows.Scan(&testpaperEntity.QuestionId, &testpaperEntity.QuestionName, &testpaperEntity.QuestionPoolId, &testpaperEntity.QuestionType,
			&testpaperEntity.QuestionContent, &testpaperEntity.QuestionScore, &testpaperEntity.Digree, &testpaperEntity.Answer)

		if testpaperEntity.QuestionType != 0 {
			returnEntity.TestPaperQuestionTypeOver[testpaperEntity.QuestionType-1].QuestionIdNum += 1
			returnEntity.TestPaperQuestionTypeOver[testpaperEntity.QuestionType-1].QuestionScore += testpaperEntity.QuestionScore
		}
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		returnEntity.TestPaperQuestionViewFile = append(returnEntity.TestPaperQuestionViewFile, testpaperEntity)
	}

	for i := 0; i < len(returnEntity.TestPaperQuestionViewFile); i++ {
		testpaperEntity := returnEntity.TestPaperQuestionViewFile[i]

		testpaperEntity.FileInfo = make([]*model.FileInfo, 0)
		if testpaperEntity.QuestionType == 5 {
			fileid := testpaperEntity.QuestionContent

			fileinfo := new(model.FileInfo)
			lib.Db.QueryRow("select Id,FileType,FilePath,FileName from fileinfo    where Id=? ", fileid).Scan(&fileinfo.Id, &fileinfo.FileType, &fileinfo.FilePath, &fileinfo.FileName)
			if fileinfo.Id > 0 {
				testpaperEntity.FileInfo = append(testpaperEntity.FileInfo, fileinfo)
			}
		} else {
			rows, err := lib.Db.Query("select b.Id,b.FileType,b.FilePath,b.FileName from questionrelation  a left join fileinfo b on a.FileId=b.Id where a.QuestionId=? ", testpaperEntity.QuestionId)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			for rows.Next() {
				newfileinfo := new(model.FileInfo)
				err := rows.Scan(&newfileinfo.Id, &newfileinfo.FileType, &newfileinfo.FilePath, &newfileinfo.FileName)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "失败",
						"data": "{}",
					})
					return
				}
				testpaperEntity.FileInfo = append(testpaperEntity.FileInfo, newfileinfo)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": returnEntity,
	})
}

func UploadExamImage(c *gin.Context) {

	data, _ := c.GetPostForm("data")
	var examimage model.ExamImage
	json.Unmarshal([]byte(data), &examimage)

	if examimage.StudentId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	files, err := c.MultipartForm()

	if len(files.File["files"]) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	file := files.File["files"]
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未收到文件",
			"data": "{}",
		})
		return

	}

	i := 0
	filename := goutils.GUID()
	file_path := "Resources/ExamImage/" + lib.Strval(examimage.ExamSessionId) + "/" + lib.Strval(examimage.StudentId) + "/" + filename + ".png"

	sqlStr := "insert into  examimage(ExamId,ExamSessionId,StudentId,ImagePath,CreateTime) VALUES (?,?,?,?,? ) "
	timestr := time.Now().Format(lib.TimeLayoutStr)
	ret, err := lib.Db.Exec(sqlStr, examimage.ExamId, examimage.ExamSessionId, examimage.StudentId, file_path, timestr)

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

	if n <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	lib.UploadFile(c, file[i], file_path)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": "{}",
	})

}

func GetStudentExamScoreByExamNumber(c *gin.Context) {

	ExamNumber := c.Query("ExamNumber")
	CourseCode := c.Query("CourseCode")

	sql := "SELECT a.StudentId,a.ExamId,a.ExamSessionId,a.Score,d.CourseCode  FROM examsturesult   a  " +
		" left join examsession b on a.ExamSessionId=b.Id" +
		" left join testpaper c on c.Id=b.TestPaperId " +
		" LEFT JOIN course d on c.CourseId =d.Id " +
		" LEFT JOIN  student e on e.Id=a.StudentId " +
		" where   e.ExamNumber=? and d.CourseCode=? " +
		" and   !ISNULL(a.StartExamTime) and !ISNULL(a.EndExamTime) ORDER BY Score desc LIMIT 1 "

	export := new(model.ExportExam)
	lib.Db.QueryRow(sql, ExamNumber, CourseCode).Scan(&export.StudentId, &export.ExamId, &export.ExamSessionId, &export.Score, &export.CourseCode)

	c.JSON(http.StatusOK, export)

}

func GetStudentImg(c *gin.Context) {

	examNumber := c.Query("ExamNumber")
	imagePath := new(ImagePath)
	lib.Db.QueryRow("SELECT ImagePath FROM examimage a left join student b on a.StudentId=b.Id  where b.ExamNumber=?  LIMIT 1;", examNumber).Scan(&imagePath.ImagePath)
	if imagePath.ImagePath == "" {
		c.String(http.StatusBadRequest, "")
		return
	}
	c.String(http.StatusOK, imagePath.ImagePath)
}

func AddExamBatchSessionStudent(c *gin.Context) {

	files, err := c.MultipartForm()
	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未找到文件",
			"data": "{}",
		})
		return
	}
	if len(files.File["files"]) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未找到文件",
			"data": "{}",
		})
		return
	}

	tx, err := lib.Db.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	temoBatchExamSessionStudentarr := make([]*BatchExamSessionStudent, 0)
	if len(files.File["files"]) > 0 {
		filesarr := files.File["files"]
		filename := goutils.GUID()

		for i := 0; i < len(filesarr); i++ {
			file := filesarr[i]
			filePath := "Resources/Execl/" + filename + path.Ext(file.Filename)
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}

			f, err := excelize.OpenFile(filePath)
			if err != nil {
				os.Remove(filePath)
				fmt.Println(err)
				return
			}

			// 遍历工作表
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
					examStudent := new(BatchExamSessionStudent)
					examStudent.RowNum = i + 1
					// 遍历每列数据
					for j, colCell := range row {

						if strings.Contains(rows[0][j], "准考证号") {
							examStudent.ExamNumber = colCell
							tx.QueryRow("select Id as 'StudentId' from student where ExamNumber=?", examStudent.ExamNumber).Scan(&examStudent.StudentId)
							if examStudent.StudentId == 0 {
								tx.Rollback()
								c.JSON(http.StatusOK, gin.H{
									"code": 0,
									"msg":  "操作失败,该execl 第" + lib.Strval(examStudent.RowNum) + " 行，准考证号在系统找不到该学生",
									"data": "{}",
								})
								return
							}
						} else if strings.Contains(rows[0][j], "课程代码") {

							examStudent.CourseCode = colCell
							if examStudent.CourseCode == "" {
								tx.Rollback()
								c.JSON(http.StatusOK, gin.H{
									"code": 0,
									"msg":  "操作失败,该execl 第" + lib.Strval(examStudent.RowNum) + " 行，课程代码为空",
									"data": "{}",
								})
								return
							}
						}
					}

					if examStudent.StudentId != 0 {
						temoBatchExamSessionStudentarr = append(temoBatchExamSessionStudentarr, examStudent)
					}
				}

				break
			}
			os.Remove(filePath)
		}

		for k := 0; k < len(temoBatchExamSessionStudentarr); k++ {
			sql := "select a.Id,a.ExamId from examsession a  LEFT JOIN testpaper b  on a.TestPaperId=b.Id LEFT JOIN course c on b.CourseId=c.Id where   CourseCode=?"
			rows, err := tx.Query(sql, temoBatchExamSessionStudentarr[k].CourseCode)
			if err != nil {
				temoBatchExamSessionStudentarr[k].Status = 0
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行 ，课程代码找不到考试场次",
					"data": "{}",
				})
				return
			}

			hasData := false
			batchExamTemparr := make([]*BatchExamTemp, 0)
			for rows.Next() {
				hasData = true
				batchExamTemp := new(BatchExamTemp)

				err := rows.Scan(&batchExamTemp.ExamSessionId, &batchExamTemp.ExamId)
				if err != nil {
					temoBatchExamSessionStudentarr[k].Status = 0
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 系统找不到该课程的考场",
						"data": "{}",
					})
					return
				}

				batchExamTemparr = append(batchExamTemparr, batchExamTemp)

			}
			if !hasData {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 系统找不到该课程的考场",
					"data": "{}",
				})
				return
			}

			for i := 0; i < len(batchExamTemparr); i++ {

				numtemp := 0

				tx.QueryRow("select count(1) from examstudent where  ExamId=? and StudentId=? ", batchExamTemparr[i].ExamId, temoBatchExamSessionStudentarr[k].StudentId).
					Scan(&numtemp)

				if numtemp == 0 {

					sqladdStr := " insert into  examstudent(ExamId,StudentId ) values(?,?)  "
					ret, err := tx.Exec(sqladdStr, batchExamTemparr[i].ExamId, temoBatchExamSessionStudentarr[k].StudentId)

					if err != nil {
						temoBatchExamSessionStudentarr[k].Status = 0
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 学生没有插入成功",
							"data": "{}",
						})
						return
					}
					n, err := ret.RowsAffected()
					if err != nil || n == 0 {
						temoBatchExamSessionStudentarr[k].Status = 0
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 学生没有插入成功",
							"data": "{}",
						})
						return
					}

					sqladdsturesultstr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  "
					ret1, err := tx.Exec(sqladdsturesultstr, temoBatchExamSessionStudentarr[k].StudentId, batchExamTemparr[i].ExamId, batchExamTemparr[i].ExamSessionId, -1)
					if err != nil {
						temoBatchExamSessionStudentarr[k].Status = 0
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 成绩表没有插入成功",
							"data": "{}",
						})
						return
					}
					n1, err := ret1.RowsAffected()
					if err != nil || n1 == 0 {
						temoBatchExamSessionStudentarr[k].Status = 0
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行，成绩表没有插入成功",
							"data": "{}",
						})
						return
					}
				}

			}

		}

	}
	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "系统内部错误",
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

func ExamCancel(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var exam model.DelExam
	json.Unmarshal([]byte(body), &exam)
	if exam.ExamId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	tx, err := lib.Db.Begin()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "事务打开失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	newexam := new(model.Exam) //新的考试数据
	// 先查询考试  在新增 新增考试  得到新增的考试id
	err = tx.QueryRow("select SchoolId,ExamName,ExamDescribe,ExamStatus,FaceVerify,TeacherId  from exam where id=?", exam.ExamId).
		Scan(&newexam.SchoolId, &newexam.ExamName, &newexam.ExamDescribe, &newexam.ExamStatus, &newexam.FaceVerify, &newexam.TeacherId)

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	sqlStr := " insert into  exam(SchoolId,ExamName,ExamDescribe,ExamStatus,FaceVerify,TeacherId) values(?,?,?,?,?,?)  "

	ret, err := tx.Exec(sqlStr, newexam.SchoolId, newexam.ExamName, newexam.ExamDescribe, 0, newexam.FaceVerify, newexam.TeacherId)

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	newexamid, err := ret.LastInsertId()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if newexamid == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
	}

	//新增场次

	// 查询以前老场次
	rows, err := tx.Query("select StartTime,EndTime,TestPaperId from examsession where ExamId=?", exam.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	newsessionidarr := make([]*model.ExamSession, 0)
	for rows.Next() {
		examSession := new(model.ExamSession)
		err := rows.Scan(&examSession.StartTime, &examSession.EndTime, &examSession.TestPaperId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		newsessionidarr = append(newsessionidarr, examSession)
	}
	rows.Close()

	for i := 0; i < len(newsessionidarr); i++ {
		sqlStr := " insert into  examsession(ExamId,StartTime,EndTime,TestPaperId) values(?,?,?,?)  "

		ret1, err := tx.Exec(sqlStr, newexamid, newsessionidarr[i].StartTime, newsessionidarr[i].EndTime, newsessionidarr[i].TestPaperId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		newsessionid, err := ret1.LastInsertId()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		if newsessionid <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}

		newsessionidarr[i].Id = int(newsessionid)
	}

	//新增学生
	// 先查询考试  在新增 新增考试  得到新增的考试id
	rowsExamStudent, err := tx.Query("	SELECT  ExamId,StudentId  FROM examstudent where ExamId=?", exam.ExamId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		tx.Rollback()
		return
	}
	newExamStudentModelarr := make([]*model.ExamStudentModel, 0)
	for rowsExamStudent.Next() {
		newExamStudentModel := new(model.ExamStudentModel) //新的考试数据

		err := rowsExamStudent.Scan(&newExamStudentModel.ExamId, &newExamStudentModel.StudentId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		newExamStudentModelarr = append(newExamStudentModelarr, newExamStudentModel)
	}
	rowsExamStudent.Close()

	for i := 0; i < len(newExamStudentModelarr); i++ {
		sqlStr := " insert into  examstudent(ExamId,StudentId  ) values(?,?)  "

		ret, err := tx.Exec(sqlStr, newexamid, newExamStudentModelarr[i].StudentId)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		n, err := ret.RowsAffected()
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
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			tx.Rollback()
			return
		}
		for i := 0; i < len(newsessionidarr); i++ {
			sqlStr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  "

			ret, err := tx.Exec(sqlStr, newExamStudentModelarr[i].StudentId, newExamStudentModelarr[i].ExamId, newsessionidarr[i].Id, -1)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败",
					"data": "{}",
				})
				tx.Rollback()
				return
			}
			n, err := ret.RowsAffected()
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
					"code": 0,
					"msg":  "操作失败",
					"data": "{}",
				})
				tx.Rollback()
				return
			}
		}
	}

	//新增考试公告  先查询公告

	examnotice := new(model.ExamNotice)

	tx.QueryRow("select Id,ExamId,Title,SchoolName,CourseName,CourseCode,Context from  examnotice where ExamId=?", exam.ExamId).Scan(&examnotice.Id, &examnotice.ExamId, &examnotice.Title, &examnotice.SchoolName, &examnotice.CourseName, &examnotice.CourseCode, &examnotice.Context)
	if examnotice.Id != 0 {
		sqlStrnotice := "insert into  examnotice(ExamId,Title,SchoolName,CourseName,CourseCode,Context) VALUES (?,?,?,?,?,?) "

		ret1, err := tx.Exec(sqlStrnotice, newexamid, examnotice.Title, examnotice.SchoolName, examnotice.CourseName, examnotice.CourseCode, examnotice.Context)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret1.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}

		if n == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
		}
	}

	//mark 删除之前判断是否有学生已考试的判断不要了
	// sqlquery := " select  COUNT(Score) 'isdel' from examsturesult where  ExamId=? and Score>=0"

	// isdel := 0
	// tx.QueryRow(sqlquery, exam.ExamId).Scan(&isdel)

	// if isdel == 0 {

	// 往后都是删除考试
	sqlexamdel := "update exam set IsDel=1 where Id=?"

	ret1, err := tx.Exec(sqlexamdel, exam.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	n, err := ret1.RowsAffected()

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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	sqlexamdel = "update examsession set IsDel=1 where ExamId=?"

	ret1, err = tx.Exec(sqlexamdel, exam.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	n, err = ret1.RowsAffected()

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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	sqlexamdel = "update examsturesult set IsDel=1 where ExamId=?"

	ret1, err = tx.Exec(sqlexamdel, exam.ExamId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	n, err = ret1.RowsAffected()

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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
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

func QueryStudentResultByStandId(c *gin.Context) {

	err, studentResultarr := StudentResultRow2arr(c)

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": studentResultarr,
	})
}

func QueryStudentResultExeclByStandId(c *gin.Context) {

	err, studentResultarr := StudentResultRow2arr(c)

	if err != nil {
		return
	}
	f := excelize.NewFile()

	// 设置表头
	header := []string{"序号", "姓名", "课程名称", "课程代码", "准考证号", "学校名称", "成绩"}

	f.SetCellValue("Sheet1", "A1", header[0])
	f.SetCellValue("Sheet1", "B1", header[1])
	f.SetCellValue("Sheet1", "C1", header[2])
	f.SetCellValue("Sheet1", "D1", header[3])
	f.SetCellValue("Sheet1", "E1", header[4])
	f.SetCellValue("Sheet1", "F1", header[5])
	f.SetCellValue("Sheet1", "G1", header[6])

	for i := 0; i < len(studentResultarr); i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), studentResultarr[i].RowNumber)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), studentResultarr[i].TrueName)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), studentResultarr[i].CourseName)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), studentResultarr[i].CourseCode)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), studentResultarr[i].ExamNumber)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), studentResultarr[i].SchoolName)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), studentResultarr[i].Score)
	}
	filename := goutils.GUID()
	// 保存Excel文件
	err = f.SaveAs("Resources/Execl/temp/" + filename + ".xlsx")
	if err != nil {
		log.Fatal(err)
	}
	c.File("Resources/Execl/temp/" + filename + ".xlsx")
	os.Remove("Resources/Execl/temp/" + filename + ".xlsx")
}
func StudentResultRow2arr(c *gin.Context) (error, []*StudentResult) {
	StandId := c.Query("StandId")

	if StandId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误,请传入站点",
			"data": "{}",
		})
		return fmt.Errorf("参数错误,请传入站点"), nil
	}

	CourseId := c.Query("CourseId")       //可为空  不为空的时候只能为数字
	IsPassScore := c.Query("IsPassScore") //0或者空值 都查  1 查及格 2 查不及格

	queStudentResult := " select  ROW_NUMBER() over(ORDER BY f.TrueName) as 'RowNumber',f.TrueName,e.CourseName,e.CourseCode,f.ExamNumber,g.SchoolName, scoreview.Score" +
		" from ( select a.ExamId,a.ExamSessionId,a.StudentId 'StudentId',MAX(a.Score) 'Score'" +
		"	   from examsturesult a " +
		"		 where a.Score>0 " +
		"	GROUP BY a.ExamId,a.ExamSessionId,a.StudentId ) scoreview  " +
		"	LEFT JOIN examsession c on scoreview.examsessionId=c.id " +
		"	LEFT JOIN  testpaper d on  c.testpaperId=d.id " +
		"	LEFT JOIN course e on d.CourseId=e.id  " +
		"	LEFT JOIN  student f on scoreview.StudentId=f.Id " +
		"	LEFT JOIN school g on f.SchoolId=g.Id " +
		"	where   e.CourseCode is not null   and    f.StandId=?  "

	if IsPassScore == "1" {
		queStudentResult += " and  (scoreview.Score>=d.PassScore)  " //查及格
	} else if IsPassScore == "2" {
		queStudentResult += " and  (scoreview.Score<d.PassScore)  " //查不及格
	}
	queStudentResult += "order by  f.TrueName"

	if CourseId != "" {

		_, err := strconv.Atoi(CourseId)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "参数错误,科目选择错误",
				"data": "{}",
			})
			return fmt.Errorf("参数错误,请传入站点"), nil
		}

		queStudentResult += " and d.CourseId=? "
	}

	var rows *sql.Rows
	var err error

	if CourseId != "" {
		rows, err = lib.Db.Query(queStudentResult, StandId, CourseId)
	} else {
		rows, err = lib.Db.Query(queStudentResult, StandId)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作错误",
			"data": "{}",
		})
		return fmt.Errorf("参数错误,请传入站点"), nil
	}
	studentResultarr := make([]*StudentResult, 0)
	for rows.Next() {
		temp := new(StudentResult)
		err = rows.Scan(&temp.RowNumber, &temp.TrueName, &temp.CourseName, &temp.CourseCode, &temp.ExamNumber, &temp.SchoolName, &temp.Score)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
			return fmt.Errorf("参数错误,请传入站点"), nil
		}
		studentResultarr = append(studentResultarr, temp)
	}
	return nil, studentResultarr
}

type ImagePath struct {
	ImagePath string `json:"ImagePath" db:"ImagePath"`
}

type BatchExamSessionStudent struct {
	StudentId     int    `json:"StudentId" db:"StudentId"`
	ExamNumber    string `json:"ExamNumber" db:"ExamNumber"`
	CourseCode    string `json:"CourseCode" db:"CourseCode"`
	ExamId        int    `json:"ExamId" db:"ExamId"`
	ExamSessionId int    `json:"ExamSessionId" db:"ExamSessionId"`
	Status        int    `json:"Status" db:"Status"`
	RowNum        int    `json:"RowNum" db:"RowNum"`
}

type BatchExamTemp struct {
	ExamSessionId int `json:"ExamSessionId" db:"ExamSessionId"`
	ExamId        int `json:"ExamId" db:"ExamId"`
}

type StudentResult struct {
	RowNumber  int     `json:"RowNumber" db:"RowNumber"`
	TrueName   string  `json:"TrueName" db:"TrueName"`
	CourseName string  `json:"CourseName" db:"CourseName"`
	CourseCode string  `json:"CourseCode" db:"CourseCode"`
	ExamNumber string  `json:"ExamNumber" db:"ExamNumber"`
	SchoolName string  `json:"SchoolName" db:"SchoolName"`
	Score      float64 `json:"Score" db:"Score"`
}
