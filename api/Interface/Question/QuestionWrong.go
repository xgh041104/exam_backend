package question

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddQuestionWrong(c *gin.Context) {

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
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var questionwrongarr []*model.QuestionWrong
	json.Unmarshal([]byte(body), &questionwrongarr)

	if len(questionwrongarr) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	for i := 0; i < len(questionwrongarr); i++ {
		if questionwrongarr[i].QuestionId == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "参数错误",
				"data": "{}",
			})
			return
		}

		QuestionId := 0
		tx.QueryRow("select QuestionId,Answer from question where QuestionId=?", questionwrongarr[i].QuestionId).Scan(&QuestionId, &questionwrongarr[i].TrueAnswer)
		if QuestionId == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "参数错误",
				"data": "{}",
			})
			return
		}

		isexites := 0
		tx.QueryRow("select Id from questionwrong where QuestionId=? and StudentId=?", questionwrongarr[i].QuestionId, questionwrongarr[i].StudentId).Scan(&isexites)
		if isexites > 0 {
			continue
		}

		sqlStr := " insert into questionwrong(QuestionId,StudentId,CreateTime,AnswerSteps,TrueAnswer) values (?,?,?,?,?) "
		timestr := time.Now().Format(lib.TimeLayoutStr)
		ret, err := tx.Exec(sqlStr, questionwrongarr[i].QuestionId, questionwrongarr[i].StudentId, timestr, questionwrongarr[i].AnswerSteps, questionwrongarr[i].TrueAnswer)

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
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
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
func AddQuestionRecord(c *gin.Context) {

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
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var questionrecordarr []*model.Questionrecord
	json.Unmarshal([]byte(body), &questionrecordarr)

	if len(questionrecordarr) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	for i := 0; i < len(questionrecordarr); i++ {
		if questionrecordarr[i].QuestionId == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "参数错误",
				"data": "{}",
			})
			return
		}

		QuestionId := 0
		tx.QueryRow("select QuestionId,Answer from question where QuestionId=? ", questionrecordarr[i].QuestionId).Scan(&QuestionId, &questionrecordarr[i].TrueAnswer)
		if QuestionId == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "参数错误",
				"data": "{}",
			})
			return
		}

		isexites := 0
		tx.QueryRow("select Id from questionrecord where QuestionId=? and StudentId=? and PlanId=?", questionrecordarr[i].QuestionId, questionrecordarr[i].StudentId, questionrecordarr[i].PlanId).Scan(&isexites)
		if isexites > 0 {

			sqlStr := " update   questionrecord set AnswerSteps=?,TrainScore=? where  QuestionId=? and StudentId=? and PlanId=?"
			ret, err := tx.Exec(sqlStr, questionrecordarr[i].AnswerSteps, questionrecordarr[i].TrainScore, questionrecordarr[i].QuestionId, questionrecordarr[i].StudentId, questionrecordarr[i].PlanId)

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
					"code": 0,
					"msg":  "操作失败",
					"data": "{}",
				})
			}
		} else {

			sqlStr := " insert into questionrecord(QuestionId,StudentId,CreateTime,AnswerSteps,TrueAnswer,TrainScore,PlanId) values (?,?,?,?,?,?,?) "
			timestr := time.Now().Format(lib.TimeLayoutStr)
			ret, err := tx.Exec(sqlStr, questionrecordarr[i].QuestionId, questionrecordarr[i].StudentId, timestr, questionrecordarr[i].AnswerSteps, questionrecordarr[i].TrueAnswer, questionrecordarr[i].TrainScore, questionrecordarr[i].PlanId)

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
					"code": 0,
					"msg":  "操作失败",
					"data": "{}",
				})
			}
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

func DelQuestionWrong(c *gin.Context) {

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
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var idarr []*model.QuestionWrongIdArr
	json.Unmarshal([]byte(body), &idarr)

	if len(idarr) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	for i := 0; i < len(idarr); i++ {
		sqlStr := " delete from  questionwrong where Id=?"
		ret, err := tx.Exec(sqlStr, idarr[i].Id)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
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
				"code": 0,
				"msg":  "操作失败",
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

func GetQuestionWrongByStudentId(c *gin.Context) {

	StudentId := c.Query("StudentId")
	if StudentId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	strsql := "select a.Id,a.QuestionId,c.QuestionName,a.StudentId,a.CreateTime,a.AnswerSteps,a.TrueAnswer from questionwrong a  left join question c on a.QuestionId=c.QuestionId where StudentId=? "
	rows, err := lib.Db.Query(strsql, StudentId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	var questionWrongViewarr []*model.QuestionWrongView
	for rows.Next() {
		questionWrong := new(model.QuestionWrongView)
		rows.Scan(&questionWrong.Id, &questionWrong.QuestionId, &questionWrong.QuestionName, &questionWrong.StudentId, &questionWrong.CreateTime, &questionWrong.AnswerSteps, &questionWrong.TrueAnswer)
		if questionWrong.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		questionWrongViewarr = append(questionWrongViewarr, questionWrong)
	}

	for i := 0; i < len(questionWrongViewarr); i++ {
		questionWrong := questionWrongViewarr[i]

		questionView := new(model.QuestionViewFile)
		lib.Db.QueryRow("SELECT a.QuestionId,a.SchoolId,a.QuestionPoolId,a.QuestionName,a.QuestionType,a.QuestionContent,a.Digree,a.MajorID,a.CollegeId,a.CourseId,a.Answer,b.MajorName,c.CollegeName,d.CourseName FROM  question a  left join  major b  on a.MajorID=b.MajorId   left join  college  c  on a.CollegeId=c.Id 	 left join  course d on a.CourseId=d.Id  where  a.QuestionId=?", questionWrong.QuestionId).
			Scan(&questionView.QuestionId, &questionView.SchoolId, &questionView.QuestionPoolId, &questionView.QuestionName, &questionView.QuestionType, &questionView.QuestionContent, &questionView.Digree, &questionView.MajorID, &questionView.CollegeId, &questionView.CourseId, &questionView.Answer, &questionView.MajorName, &questionView.CollegeName, &questionView.CourseName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		questionView.FileInfo = make([]*model.FileInfo, 0)
		if questionView.QuestionType == 5 {
			fileid := questionView.QuestionContent

			fileinfo := new(model.FileInfo)
			lib.Db.QueryRow("select Id,FileType,FilePath,FileName from fileinfo    where Id=? ", fileid).Scan(&fileinfo.Id, &fileinfo.FileType, &fileinfo.FilePath, &fileinfo.FileName)
			if fileinfo.Id > 0 {
				questionView.FileInfo = append(questionView.FileInfo, fileinfo)
			}
		} else {
			rows, err := lib.Db.Query("select b.Id,b.FileType,b.FilePath,b.FileName from questionrelation  a left join fileinfo b on a.FileId=b.Id where a.QuestionId=? ", questionWrong.QuestionId)
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
				questionView.FileInfo = append(questionView.FileInfo, newfileinfo)
			}
		}
		questionWrong.QuestionViewEntity = questionView
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": questionWrongViewarr,
	})
}
