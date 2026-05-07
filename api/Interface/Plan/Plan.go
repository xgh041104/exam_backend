package plan

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	goutils "github.com/typa01/go-utils"
)

func AddPlan(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var plan model.Plan
	json.Unmarshal([]byte(body), &plan)

	if (plan.CourseRatio + plan.ExamRatio + plan.TrainRatio) != 100 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "占比有误，请重新设置比例",
			"data": "{}",
		})
		return
	}

	sqlStr := "insert into  plan(CourseId,PlanName,CourseRatio,ExamRatio,TrainRatio,TeacherId) VALUES (?,?,?,?,?,?) "

	ret, err := lib.Db.Exec(sqlStr, plan.CourseId, plan.PlanName, plan.CourseRatio, plan.ExamRatio, plan.TrainRatio, plan.TeacherId)

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

func EditPlan(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var plan model.Plan
	json.Unmarshal([]byte(body), &plan)

	if (plan.CourseRatio + plan.ExamRatio + plan.TrainRatio) != 100 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "占比有误，请重新设置比例",
			"data": "{}",
		})
		return
	}

	sqlStr := "update  plan set  PlanName=?,CourseRatio=?,ExamRatio=?,TrainRatio=?   where  PlanId=?  "

	ret, err := lib.Db.Exec(sqlStr, plan.PlanName, plan.CourseRatio, plan.ExamRatio, plan.TrainRatio, plan.PlanId)

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

func DelPlan(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var plan model.Plan
	json.Unmarshal([]byte(body), &plan)

	tx, err := lib.Db.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	//删除该计划的计划记录
	sqlStr := "delete from   plan  where  PlanId=?  "

	ret, err := tx.Exec(sqlStr, plan.PlanId)

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
		return
	}
	//删除该计划的考试表记录
	sqlStr = "delete from   planexam  where  PlanId=?  "

	ret, err = tx.Exec(sqlStr, plan.PlanId)
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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	//删除该计划的学生关联记录
	sqlStr = "delete from   planstudent  where  PlanId=?  "

	ret, err = tx.Exec(sqlStr, plan.PlanId)
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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	//删除该计划的训练表记录
	sqlStr = "delete from   plantrain  where  PlanId=?  "

	ret, err = tx.Exec(sqlStr, plan.PlanId)
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
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	//删除该计划的训练记录
	sqlStr = "delete from   questionrecord  where  PlanId=?  "

	ret, err = tx.Exec(sqlStr, plan.PlanId)
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
}

func QueryPlan(c *gin.Context) {
	query := " select a.PlanId,a.CourseId,a.PlanName,a.CourseRatio,a.ExamRatio,a.TrainRatio,a.TeacherId, " +
		" ( select COUNT(1) from planstudent  where planId=a.planId  ) 'PlanStudentCount', " +
		" ( select COUNT(1) from planexam  where planId=a.planId  ) 'PlanExamCount', " +
		" ( select COUNT(1) from plantrain  where planId=a.planId  ) 'PlanTrainCount',COALESCE(b.CourseName,'') 'CourseName'  " +
		" from plan  a   left join  course b on a.CourseId=b.Id "

	rows, err := lib.Db.Query(query)

	planarr := make([]*model.PlanArr, 0)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	for rows.Next() {
		temp := new(model.PlanArr)
		rows.Scan(&temp.PlanId, &temp.CourseId, &temp.PlanName, &temp.CourseRatio, &temp.ExamRatio, &temp.TrainRatio, &temp.TeacherId, &temp.PlanStudentCount, &temp.PlanExamCount, &temp.PlanTrainCount, &temp.CourseName)
		planarr = append(planarr, temp)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": planarr,
	})

}

func QuertStydyPlan(c *gin.Context) {

	// planId := c.Query("PlanId")

	// if planId == "0" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 0,
	// 		"msg":  "参数错误",
	// 		"data": "{}",
	// 	})
	// 	return
	// }

	// var resp model.Plan
	// lib.Db.QueryRow("select PlanId,CourseId,PlanName,CourseRatio,ExamRatio,TrainRatio,TeacherId	 from plan where PlanId=? ", planId).
	// 	Scan(&resp.PlanId, &resp.CourseId, &resp.PlanName, &resp.CourseRatio, &resp.ExamRatio, &resp.TrainRatio, &resp.TeacherId)

	// if resp.PlanId == 0 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 0,
	// 		"msg":  "查询失败",
	// 		"data": "{}",
	// 	})
	// 	return
	// }

	// rows, err := lib.Db.Query("select a.PlanStudentId,a.PlanId,a.StudentId,b.TrueName,b.StudentAccount from planstudent a 	left JOIN student  b on a.StudentId=b.Id where a.PlanId=? ", planId)
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 0,
	// 		"msg":  "查询学生失败",
	// 		"data": "{}",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 1,
	// 	"msg":  "操作成功",
	// 	"data": resp,
	// })
}

// 赋值当前学生的课程内容
func QueryPlanCourseProgress(c *gin.Context) {
	studentId := c.Query("StudentId")
	planId := c.Query("PlanId")

	var resp model.Plan
	lib.Db.QueryRow("select PlanId,CourseId,PlanName,CourseRatio,ExamRatio,TrainRatio,TeacherId	 from plan where PlanId=? ", planId).
		Scan(&resp.PlanId, &resp.CourseId, &resp.PlanName, &resp.CourseRatio, &resp.ExamRatio, &resp.TrainRatio, &resp.TeacherId)

	if resp.PlanId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "查询失败",
			"data": "{}",
		})
		return
	}

	planStudentCourseResult := new(model.StudyCourse)

	querysql := "	select  b.CourseName,  " +
		" IFNULL(b.Digest,'') 'Digest', " +
		" a.CourseId, " +
		" b.TeacherId,	IFNULL(d.TeacherName,'')'TeacherName', " +
		"	IFNULL(e.SchoolName,'')  'SchoolName',	IFNULL(f.CollegeName,'')  'CollegeName'  , b. MajorId,	IFNULL(g.MajorName,'') 'MajorName' , " +
		"			(select COUNT(1) from chapter   WHERE CourseId= a.CourseId ) ' ChapterSum',  " +
		"			(select COUNT(1) from studyplan   WHERE CourseId= a.CourseId ) ' StudentSum',  " +
		"	  b.CourseStartTime, " +
		"		b.CourseEndTime, " +
		"a.ChapterOrder  , " +
		"		 a.LearningRate  , " +
		"		 IFNULL(c.FilePath,'') 'FilePath', " +
		" (CASE    " +
		"	WHEN b.CourseStartTime < now() AND b.CourseEndTime >= now() and  a.IsComplete=0  THEN  1   " +
		"	WHEN b.CourseStartTime > now()  and  a.IsComplete=0  THEN   0  " +
		"	WHEN a.IsComplete=1 THEN   2   " +
		"	WHEN b.CourseEndTime < now()  and  a.IsComplete=0 THEN   3  " +
		" ELSE 0 END   ) as 'IsCurrentStudy' " +
		" from studyplan a  " +
		"left join course b  on a.CourseId=b.Id " +
		"LEFT JOIN fileinfo c on b.FileId=c.Id " +
		"LEFT JOIN teacher d on b.TeacherId=d.TeacherId " +
		"LEFT JOIN school e on b.SchoolId=e.Id " +
		"LEFT JOIN college f on b.CollegeId=f.Id " +
		"LEFT JOIN major g  on b.MajorId=g.MajorId " +
		" WHERE a.StudentId= ?  and b.Id=? and b.Status=1  "

	err := lib.Db.QueryRow(querysql, studentId, resp.CourseId).
		Scan(&planStudentCourseResult.CourseName, &planStudentCourseResult.Digest,
			&planStudentCourseResult.CourseId, &planStudentCourseResult.TeacherId,
			&planStudentCourseResult.TeacherName, &planStudentCourseResult.SchoolName,
			&planStudentCourseResult.CollegeName, &planStudentCourseResult.MajorId,
			&planStudentCourseResult.MajorName, &planStudentCourseResult.ChapterSum,
			&planStudentCourseResult.StudentSum, &planStudentCourseResult.CourseStartTime,
			&planStudentCourseResult.CourseEndTime, &planStudentCourseResult.ChapterOrder,
			&planStudentCourseResult.LearningRate, &planStudentCourseResult.FilePath,
			&planStudentCourseResult.IsCurrentStudy)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "没有课程数据",
			"data": "{}",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": planStudentCourseResult,
	})

}

//赋值当前学生的考试内容

func QueryPlanExamProgress(c *gin.Context) {

	studentId := c.Query("StudentId")
	planId := c.Query("PlanId")
	querysql := " SELECT a.Id,a.StudentId,a.ExamId,b.ExamName,a.ExamSessionId,COALESCE(a.StartExamTime, '') 'StartExamTime',COALESCE(a.EndExamTime,'')  'EndExamTime',a.Score,a.ExamStatus,a.ExamType,   " +
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
		" where a.StudentId=?  and a.ExamSessionId in ( select examsessionId from planexam where PlanId=?) and a.IsDel=0  ORDER BY ExamZT  asc" // b.ExamStatus=1 and b.ReviewFlag=1 and
	rows, err := lib.Db.Query(querysql, studentId, planId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	planExamViewResultArr := make([]*model.StudentExamInfo, 0)
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

		planExamViewResultArr = append(planExamViewResultArr, studentExamInfo)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": planExamViewResultArr,
	})
}

//赋值当前学生的训练内容

func QueryPlanTrainProgress(c *gin.Context) {
	studentId := c.Query("StudentId")
	planId := c.Query("PlanId")
	querysql := "	SELECT   a.PlanTrainId,a.PlanId,a.QuestionId,a.QuestionType,b.QuestionName,b.QuestionContent, " +
		" COALESCE(c.AnswerSteps,'')  'AnswerSteps', " +
		" COALESCE(c.TrainScore,-1)  'TrainScore'  FROM  plantrain  a  " +
		" left join question  b on a.QuestionId=b.QuestionId " +
		" left join  questionrecord c on c.QuestionId=a.QuestionId and c.StudentId=? and   " +
		"  c.QuestionRecordId = (select QuestionRecordId from questionrecord where QuestionId=a.QuestionId and StudentId=?  ORDER BY TrainScore desc LIMIT 1 )  " +
		" where  a.PlanId=? "
	rows, err := lib.Db.Query(querysql, studentId, studentId, planId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	planQuestionArr := make([]*model.QuestionrecordArr, 0)
	for rows.Next() {
		temoenetity := new(model.QuestionrecordArr)
		err := rows.Scan(&temoenetity.PlanTrainId, &temoenetity.PlanId, &temoenetity.QuestionId, &temoenetity.QuestionType, &temoenetity.QuestionName,
			&temoenetity.QuestionContent, &temoenetity.AnswerSteps, &temoenetity.TrainScore)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		if temoenetity.PlanTrainId == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		planQuestionArr = append(planQuestionArr, temoenetity)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": planQuestionArr,
	})

}

// /添加学生
func AddPlanStudent(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var planstudentarr []*model.PlanStudent
	json.Unmarshal([]byte(body), &planstudentarr)

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if len(planstudentarr) == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "数据为空",
			"data": "{}",
		})
		return
	}

	courseid := 0

	querycourse := "select a.CourseId from plan  a " +
		" where  a.PlanId=? "

	tx.QueryRow(querycourse, planstudentarr[0].PlanId).Scan(&courseid)

	if courseid == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "当前计划课程id为空",
			"data": "{}",
		})
		return
	}

	query := "select a.PlanexamId,a.PlanId,a.ExamSessionId,b.ExamId from planexam a " +
		" left JOIN examsession  b on a.ExamSessionId=b.Id " +
		" where  a.PlanId=? "

	rows, err := tx.Query(query, planstudentarr[0].PlanId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var planexamarr []*model.PlanExamView
	for rows.Next() {
		temp := new(model.PlanExamView)
		rows.Scan(&temp.PlanexamId, &temp.PlanId, &temp.ExamSessionId, &temp.ExamId)
		planexamarr = append(planexamarr, temp)
	}
	insertstr := "insert into planstudent(PlanId,StudentId) values(?,?) "

	sqladdStr := " insert into  examstudent(ExamId,StudentId ) values(?,?)  "                        //考试添加sql
	sqlStr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  " //考试成绩sql

	sqlstudyStr := "insert into  studyplan(StudentId,CourseId,ChapterId,LearningRate,ChapterOrder ) VALUES (?,?,?,?,?) "

	for i := 0; i < len(planstudentarr); i++ {
		entitytemp := planstudentarr[i]

		coursenumtemp := 0 //判断该学生是否在计划中 标记
		tx.QueryRow("select count(1) from planstudent where  PlanId=? and StudentId=?", entitytemp.PlanId, entitytemp.StudentId).Scan(&coursenumtemp)
		if coursenumtemp == 0 {
			ret, err := tx.Exec(insertstr, entitytemp.PlanId, entitytemp.StudentId)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增学生失败",
					"data": "{}",
				})
				return
			}
			if n <= 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增学生失败",
					"data": "{}",
				})
				return
			}
		}

		coursenumtemp = 0 //判断该学生是否在课程中  标记
		tx.QueryRow("select count(1) from studyplan where  CourseId=? and StudentId=?", courseid, entitytemp.StudentId).Scan(&coursenumtemp)
		if coursenumtemp == 0 {
			ret, err := tx.Exec(sqlstudyStr, entitytemp.StudentId, courseid, 0, 0, 0)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if n == 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
		}

		for j := 0; j < len(planexamarr); j++ {
			numtemp := 0

			tempexam := planexamarr[j]
			err = tx.QueryRow("select count(1) from examstudent where  ExamId=? and StudentId=? ", tempexam.ExamId, entitytemp.StudentId).
				Scan(&numtemp)

			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增学生失败",
					"data": "{}",
				})
				return
			}
			if numtemp == 0 {
				//给考试添加该计划的该学生

				ret, err := tx.Exec(sqladdStr, tempexam.ExamId, entitytemp.StudentId)
				if err != nil {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "失败" + err.Error(),
						"data": "{}",
					})
					return
				}
				n, err := ret.RowsAffected()
				if err != nil {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "新增学生失败",
						"data": "{}",
					})
					return
				}
				if n <= 0 {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "新增学生失败",
						"data": "{}",
					})
					return
				}
				ret, err = tx.Exec(sqlStr, entitytemp.StudentId, tempexam.ExamId, tempexam.ExamSessionId, -1)
				if err != nil {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "失败" + err.Error(),
						"data": "{}",
					})
					return
				}
				n, err = ret.RowsAffected()
				if err != nil {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "新增考试成绩失败",
						"data": "{}",
					})
					return
				}
				if n <= 0 {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "新增考试成绩失败",
						"data": "{}",
					})
					return
				}
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
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
func DelPlanStudent(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var planstudent model.PlanStudent
	json.Unmarshal([]byte(body), &planstudent)

	ret, err := lib.Db.Exec("delete from planstudent where PlanStudentId=? ", planstudent.PlanStudentId)
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
			"msg":  "删除学生失败",
			"data": "{}",
		})
		return
	}
	if n <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除学生失败",
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

func DelAllPlanStudent(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var planstudent model.PlanStudent
	json.Unmarshal([]byte(body), &planstudent)

	ret, err := lib.Db.Exec("delete from planstudent where PlanId=? ", planstudent.PlanId)
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
			"msg":  "删除学生失败",
			"data": "{}",
		})
		return
	}
	if n <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除学生失败",
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

func QueryPlanStudent(c *gin.Context) {

	planId := c.Query("PlanId")

	resparr := make([]*model.PlanStudentView, 0)
	rows, err := lib.Db.Query("select a.PlanStudentId,a.PlanId,a.StudentId,b.TrueName,b.StudentAccount from planstudent a 	left JOIN student  b on a.StudentId=b.Id where a.PlanId=? ", planId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "查询失败",
			"data": "{}",
		})
		return
	}

	for rows.Next() {
		temp := new(model.PlanStudentView)
		rows.Scan(&temp.PlanStudentId, &temp.PlanId, &temp.StudentId, &temp.TrueName, &temp.StudentAccount)
		resparr = append(resparr, temp)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": resparr,
	})
}

// 添加考试计划
func AddPlanExam(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var planexamarr []*model.PlanExamView
	json.Unmarshal([]byte(body), &planexamarr)

	tx, err := lib.Db.Begin()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	if len(planexamarr) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "数据为空",
			"data": "{}",
		})
		return
	}

	planstudentarr := make([]*model.PlanStudent, 0)

	rows, err := tx.Query("select PlanStudentId,PlanId,StudentId from planstudent where PlanId=? ", planexamarr[0].PlanId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	for rows.Next() {
		temp := new(model.PlanStudent)
		rows.Scan(&temp.PlanStudentId, &temp.PlanId, &temp.StudentId)
		planstudentarr = append(planstudentarr, temp)
	}

	insertexamsql := "insert into planexam(PlanId,ExamSessionId) values(?,?)"
	for i := 0; i < len(planexamarr); i++ {

		tempentity := planexamarr[i]
		ret, err := tx.Exec(insertexamsql, tempentity.PlanId, tempentity.ExamSessionId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增学生失败",
				"data": "{}",
			})
			return
		}
		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增学生失败",
				"data": "{}",
			})
			return
		}

	}

	sqladdStr := " insert into  examstudent(ExamId,StudentId ) values(?,?)  "                        //考试添加sql
	sqlStr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  " //考试成绩sql
	for i := 0; i < len(planstudentarr); i++ {
		entitytemp := planstudentarr[i]

		for j := 0; j < len(planexamarr); j++ {
			examentity := planexamarr[j]
			numtemp := 0
			err = tx.QueryRow("select count(1) from examstudent where  ExamId=?  and StudentId=? ", examentity.ExamId, entitytemp.StudentId).
				Scan(&numtemp)

			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增学生失败",
					"data": "{}",
				})
				return
			}
			if numtemp > 0 { // 表示当前学生已经在考试表里面不需要重复添加
				continue
			}
			//给考试添加该计划的该学生

			ret, err := tx.Exec(sqladdStr, examentity.ExamId, entitytemp.StudentId)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增学生失败",
					"data": "{}",
				})
				return
			}
			if n <= 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增学生失败",
					"data": "{}",
				})
				return
			}
			ret, err = tx.Exec(sqlStr, entitytemp.StudentId, examentity.ExamId, examentity.ExamSessionId, -1)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			n, err = ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增考试成绩失败",
					"data": "{}",
				})
				return
			}
			if n <= 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增考试成绩失败",
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
			"msg":  "失败" + err.Error(),
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

func DelPlanExam(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var planexam model.PlanExam
	json.Unmarshal([]byte(body), &planexam)

	ret, err := lib.Db.Exec("delete from planexam where PlanexamId=? ", planexam.PlanexamId)
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
			"msg":  "删除考试条目失败",
			"data": "{}",
		})
		return
	}
	if n <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除考试条目失败",
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

func DelAllPlanExam(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var planexam model.PlanExam
	json.Unmarshal([]byte(body), &planexam)

	ret, err := lib.Db.Exec("delete from planexam where PlanId=? ", planexam.PlanId)
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
			"msg":  "批量删除考试条目失败",
			"data": "{}",
		})
		return
	}
	if n <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "批量删除考试条目失败",
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
func QueryPlanExam(c *gin.Context) {

	planId := c.Query("PlanId")
	resparr := make([]*model.PlanExamViewName, 0)

	query := " select a.PlanexamId,a.PlanId,a.ExamSessionId,b.ExamId,c.ExamName from planexam a " +
		" left JOIN examsession  b on a.ExamSessionId=b.Id " +
		" 	left JOIN exam  c on b.ExamId=c.Id  " +
		" where a.PlanId=? "
	rows, err := lib.Db.Query(query, planId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "查询失败",
			"data": "{}",
		})
		return
	}

	for rows.Next() {
		temp := new(model.PlanExamViewName)
		rows.Scan(&temp.PlanexamId, &temp.PlanId, &temp.ExamSessionId, &temp.ExamId, &temp.ExamName)
		resparr = append(resparr, temp)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": resparr,
	})
}

// 添加训练

func AddPlanTrain(c *gin.Context) {

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
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	planTrainarr := make([]*model.PlanTrain, 0)
	json.Unmarshal([]byte(body), &planTrainarr)

	if len(planTrainarr) == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "没接收到数据",
			"data": "{}",
		})
		return
	}
	insertsql := "insert into plantrain(PlanId,QuestionId,QuestionType) values(?,?,?)"

	for i := 0; i < len(planTrainarr); i++ {
		tempentity := planTrainarr[i]
		ret, err := tx.Exec(insertsql, tempentity.PlanId, tempentity.QuestionId, tempentity.QuestionType)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
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
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
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
func DelPlanTrain(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var plantrain model.PlanTrain
	json.Unmarshal([]byte(body), &plantrain)

	ret, err := lib.Db.Exec("delete from plantrain where PlanTrainId=? ", plantrain.PlanTrainId)
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
			"msg":  "删除计划内题目失败",
			"data": "{}",
		})
		return
	}
	if n <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除计划内题目失败",
			"data": "{}",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除计划内题目失败",
		"data": "{}",
	})
}
func DelAllPlanTrain(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var plantrain model.PlanTrain
	json.Unmarshal([]byte(body), &plantrain)

	ret, err := lib.Db.Exec("delete from plantrain where PlanId=? ", plantrain.PlanId)
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
			"msg":  "删除计划题目失败",
			"data": "{}",
		})
		return
	}
	if n <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除计划题目失败",
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

func QueryPlanTrain(c *gin.Context) {
	planId := c.Query("PlanId")

	resparr := make([]*model.PlanTrainView, 0)
	rows, err := lib.Db.Query("select a.PlanTrainId,a.PlanId,a.QuestionId, b.QuestionName,a.QuestionType from plantrain a  left JOIN question  b on a.QuestionId=b.QuestionId where a.PlanId=? ", planId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "查询失败",
			"data": "{}",
		})
		return
	}

	for rows.Next() {
		temp := new(model.PlanTrainView)
		rows.Scan(&temp.PlanTrainId, &temp.PlanId, &temp.QuestionId, &temp.QuestionName, &temp.QuestionType)
		resparr = append(resparr, temp)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": resparr,
	})
}

func QuertStudentPlan(c *gin.Context) {

	studentid := c.Query("StudentId")

	if studentid == "0" || studentid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	query := "  select a.PlanId,a.CourseId,a.PlanName,a.CourseRatio,a.ExamRatio,a.TrainRatio,a.TeacherId,  " +
		" ( select COUNT(1) from planstudent  where planId=a.planId  ) 'PlanStudentCount', " +
		" ( select COUNT(1) from planexam  where planId=a.planId  ) 'PlanExamCount', " +
		" ( select COUNT(1) from plantrain  where planId=a.planId  ) 'PlanTrainCount',COALESCE(b.CourseName,'') 'CourseName'  " +
		" FROM plan a   left join  course b on a.CourseId=b.Id where   " +
		" a.PlanId in (select PlanId from planstudent  where StudentId=?) "

	rows, err := lib.Db.Query(query, studentid)

	planarr := make([]*model.PlanArr, 0)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	for rows.Next() {
		temp := new(model.PlanArr)
		rows.Scan(&temp.PlanId, &temp.CourseId, &temp.PlanName, &temp.CourseRatio, &temp.ExamRatio, &temp.TrainRatio, &temp.TeacherId, &temp.PlanStudentCount, &temp.PlanExamCount, &temp.PlanTrainCount, &temp.CourseName)
		planarr = append(planarr, temp)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": planarr,
	})

}

func UploadPlanImage(c *gin.Context) {

	data, _ := c.GetPostForm("data")

	data = strings.ReplaceAll(data, "\\", "")

	var planimg model.PlanStudentImg
	json.Unmarshal([]byte(data), &planimg)
	logrus.Println(data)
	if planimg.PlanStudentId == 0 {
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
	file_path := "Resources/studyplan/" + lib.Strval(planimg.PlanStudentId) + "/" + filename + ".png"

	sqlStr := "insert into  planstudentimg(PlanStudentId,ImgPath ) VALUES (?,?  ) "
	ret, err := lib.Db.Exec(sqlStr, planimg.PlanStudentId, file_path)

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

type BatchExamSessionStudent struct {
	StudentId     int    `json:"StudentId" db:"StudentId"`
	ExamNumber    string `json:"ExamNumber" db:"ExamNumber"`
	CourseCode    string `json:"CourseCode" db:"CourseCode"`
	ExamId        int    `json:"ExamId" db:"ExamId"`
	ExamSessionId int    `json:"ExamSessionId" db:"ExamSessionId"`
	Status        int    `json:"Status" db:"Status"`
	RowNum        int    `json:"RowNum" db:"RowNum"`
}
type BatchPlanTemp struct {
	PlanId   int `json:"PlanId" db:"PlanId"`
	CourseId int `json:"CourseId" db:"CourseId"`
}

func AddExamBatchPlanStudent(c *gin.Context) {

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
			sql := "select a.PlanId,a.CourseId from plan a  LEFT JOIN course c on a.CourseId=c.Id where   CourseCode=?"
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
			batchExamTemparr := make([]*BatchPlanTemp, 0)
			for rows.Next() {
				hasData = true
				batchExamTemp := new(BatchPlanTemp)

				err := rows.Scan(&batchExamTemp.PlanId, &batchExamTemp.CourseId)
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

			sqlstudyStr := "insert into  studyplan(StudentId,CourseId,ChapterId,LearningRate,ChapterOrder ) VALUES (?,?,?,?,?) " //添加课程学生
			sqladdStr := " insert into  planstudent(PlanId,StudentId ) values(?,?)  "                                            //添加计划学生

			sqladdexamStr := " insert into  examstudent(ExamId,StudentId ) values(?,?)  "                    //考试添加学生sql
			sqlStr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  " //考试成绩学生sql
			for i := 0; i < len(batchExamTemparr); i++ {

				query := "select a.PlanexamId,a.PlanId,a.ExamSessionId,b.ExamId from planexam a " +
					" left JOIN examsession  b on a.ExamSessionId=b.Id " +
					" where  a.PlanId=? "

				rows, err = tx.Query(query, batchExamTemparr[i].PlanId)
				if err != nil {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 系统找不到该课程的考场",
						"data": "{}",
					})
					return
				}
				var planexamarr []*model.PlanExamView
				for rows.Next() {
					temp := new(model.PlanExamView)
					rows.Scan(&temp.PlanexamId, &temp.PlanId, &temp.ExamSessionId, &temp.ExamId)
					planexamarr = append(planexamarr, temp)
				}

				for _, v := range planexamarr {
					numtemp := 0

					tx.QueryRow("select count(1) from examstudent where  ExamId=? and StudentId=? ", v.ExamId, temoBatchExamSessionStudentarr[k].StudentId).
						Scan(&numtemp)

					if numtemp == 0 {
						// sqladdexamStr := " insert into  examstudent(ExamId,StudentId ) values(?,?)  "                    //考试添加学生sql
						// sqlStr := " insert into  examsturesult(StudentId,ExamId,ExamSessionId,Score ) values(?,?,?,?)  " //考试成绩学生sql
						ret, err := tx.Exec(sqladdexamStr, v.ExamId, temoBatchExamSessionStudentarr[k].StudentId)
						if err != nil {
							tx.Rollback()
							c.JSON(http.StatusOK, gin.H{
								"code": 0,
								"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 在新增考场学生时报错",
								"data": "{}",
							})
							return
						}
						n, err := ret.RowsAffected()
						if n == 0 || err != nil {
							tx.Rollback()
							c.JSON(http.StatusOK, gin.H{
								"code": 0,
								"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 在新增考场学生时报错",
								"data": "{}",
							})
							return
						}

						ret, err = tx.Exec(sqlStr, temoBatchExamSessionStudentarr[k].StudentId, v.ExamId, v.ExamSessionId, -1)
						if err != nil {
							tx.Rollback()
							c.JSON(http.StatusOK, gin.H{
								"code": 0,
								"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 在新增考场学生时报错",
								"data": "{}",
							})
							return
						}
						n, err = ret.RowsAffected()
						if n == 0 || err != nil {
							tx.Rollback()
							c.JSON(http.StatusOK, gin.H{
								"code": 0,
								"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 在新增考场学生时报错",
								"data": "{}",
							})
							return
						}

					}

				}
				coursenumtemp := 0
				tx.QueryRow("select count(1) from planstudent where  PlanId=? and StudentId=?", batchExamTemparr[i].PlanId, temoBatchExamSessionStudentarr[k].StudentId).Scan(&coursenumtemp)
				if coursenumtemp == 0 {
					ret, err := tx.Exec(sqladdStr, batchExamTemparr[i].PlanId, temoBatchExamSessionStudentarr[k].StudentId)

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
				}
				coursenumtemp = 0
				tx.QueryRow("select count(1) from studyplan where  CourseId=? and StudentId=?", batchExamTemparr[i].CourseId, temoBatchExamSessionStudentarr[k].StudentId).Scan(&coursenumtemp)
				if coursenumtemp == 0 {
					ret, err := tx.Exec(sqlstudyStr, temoBatchExamSessionStudentarr[k].StudentId, batchExamTemparr[i].CourseId, 0, 0, 0)
					if err != nil {
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 在新增课程学生时报错",
							"data": "{}",
						})
						return
					}
					n, err := ret.RowsAffected()
					if n == 0 || err != nil {
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "操作失败,该execl 第" + lib.Strval(temoBatchExamSessionStudentarr[k].RowNum) + " 行， 在新增课程学生时报错",
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
