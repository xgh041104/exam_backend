package testpaper

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTestPaper(c *gin.Context) {

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
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var testpaper model.TestPaper
	json.Unmarshal([]byte(body), &testpaper)

	if testpaper.TestPaperName == "" || testpaper.ExamDuration == 0 || testpaper.FullMarks == 0 {

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	resustr, err := json.Marshal(testpaper.QuestionScoreJson)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	ret, err := tx.Exec("insert into  testpaper(TestPaperName,ExamDuration,FullMarks,PassScore,TestPaperType,CollegeId,MajorId,CourseId,QuestionScoreJson,TeacherId,SchoolId) values (?,?,?,?,?,?,?,?,?,?,?)",
		testpaper.TestPaperName, testpaper.ExamDuration, testpaper.FullMarks, testpaper.PassScore, testpaper.TestPaperType, testpaper.CollegeId, testpaper.MajorId, testpaper.CourseId, string(resustr), testpaper.TeacherId, testpaper.SchoolId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	testpaperId, err := ret.LastInsertId()

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	if testpaperId <= 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	questscorejson := testpaper.QuestionScoreJson
	//json.Unmarshal([]byte(testpaper.QuestionScoreJson), &questscorejson)

	//sumScore := testpaper.FullMarks

	var testPaperQuestionarr []*model.TestPaperQuestion
	testPaperQuestionarr = make([]*model.TestPaperQuestion, 0)
	for i := 0; i < len(questscorejson); i++ {

		// if questscorejson[i].QuestionIdNum > 0 {
		// 	currentTypescore := (float32(questscorejson[i].FullMarksRatio) * 0.01) * testpaper.FullMarks //这个类型占总分比

		// 	if len(questscorejson[i].QuestionIdArr) == 0 {
		// 		continue
		// 	}
		// 	ereryscore := currentTypescore / float32(len(questscorejson[i].QuestionIdArr)) //每题的分数
		for j := 0; j < len(questscorejson[i].QuestionArr); j++ {
			// sumScore -= ereryscore

			testPaperQuestion := new(model.TestPaperQuestion)
			testPaperQuestion.QuestionId = questscorejson[i].QuestionArr[j].QuestionId
			testPaperQuestion.QuestionType = questscorejson[i].QuestionType
			testPaperQuestion.QuestionScore = questscorejson[i].QuestionArr[j].Score
			testPaperQuestion.TestPaperId = int(testpaperId)
			testPaperQuestionarr = append(testPaperQuestionarr, testPaperQuestion)

		}
		// }
	}

	err = AddTestPaperQuestion(tx, c, testPaperQuestionarr)
	if err != nil {
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
func EditTestPaper(c *gin.Context) {

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
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var testpaper model.TestPaper
	json.Unmarshal([]byte(body), &testpaper)

	if testpaper.Id == 0 {

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	resustr, err := json.Marshal(testpaper.QuestionScoreJson)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	// mark 修改试卷时，判断该试卷有没有在考试，如果有，则新建一个试卷
	ret, err := tx.Exec("update  testpaper set  TestPaperName=?,ExamDuration=?,FullMarks=?,PassScore=?,TestPaperType=?,CollegeId=?,MajorId=?,CourseId=?,QuestionScoreJson=?,TeacherId=?,SchoolId=?  where  Id=? ",
		testpaper.TestPaperName, testpaper.ExamDuration, testpaper.FullMarks, testpaper.PassScore, testpaper.TestPaperType, testpaper.CollegeId, testpaper.MajorId, testpaper.CourseId, string(resustr), testpaper.TeacherId, testpaper.SchoolId, testpaper.Id)

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

	ret, err = tx.Exec(" delete from testpaperquestion where TestPaperId=? ", testpaper.Id)
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

	questscorejson := testpaper.QuestionScoreJson
	//json.Unmarshal([]byte(testpaper.QuestionScoreJson), &questscorejson)

	//sumScore := testpaper.FullMarks

	var testPaperQuestionarr []*model.TestPaperQuestion
	for i := 0; i < len(questscorejson); i++ {

		// if questscorejson[i].QuestionIdNum > 0 {
		// 	currentTypescore := (float32(questscorejson[i].FullMarksRatio) * 0.01) * testpaper.FullMarks //这个类型占总分比

		if len(questscorejson[i].QuestionArr) == 0 {
			continue
		}
		//ereryscore := currentTypescore / float32(len(questscorejson[i].QuestionArr)) //每题的分数
		for j := 0; j < len(questscorejson[i].QuestionArr); j++ {
			//sumScore -= ereryscore

			testPaperQuestion := new(model.TestPaperQuestion)
			testPaperQuestion.QuestionId = questscorejson[i].QuestionArr[j].QuestionId
			testPaperQuestion.QuestionType = questscorejson[i].QuestionType
			testPaperQuestion.QuestionScore = questscorejson[i].QuestionArr[j].Score
			testPaperQuestion.TestPaperId = testpaper.Id
			testPaperQuestionarr = append(testPaperQuestionarr, testPaperQuestion)

		}
		// }

	}

	err = AddTestPaperQuestion(tx, c, testPaperQuestionarr)
	if err != nil {
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

func DelTestPaper(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var testpaper model.TestPaper
	json.Unmarshal([]byte(body), &testpaper)

	ret, err := lib.Db.Exec("update testpaper set IsDel=1 where Id=?", testpaper.Id)
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
			"msg":  "操作失败",
			"data": "{}",
		})
	}
}

func GetTestPaperBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	// 	select a.QuestionId,a.QuestionName,a.QuestionType,a.QuestionPoolId, CASE  WHEN a.QuestionType=5 THEN b.FilePath ELSE 	a.QuestionContent END  'QuestionContent' ,COALESCE(b.FileName,'') 'FileName' ,a.Answer
	//  from question a
	// left join fileinfo b on a.QuestionContent=b.Id

	var testPaperViewarr []*model.TestPaperView
	rows, err := lib.Db.Query(" select   a.Id,a.TestPaperName,a.ExamDuration,a.FullMarks,a.PassScore,a.TestPaperType,a.CollegeId,a.MajorId,a.CourseId,a.TeacherId,a.SchoolId,  COALESCE(b.MajorName,'') 'MajorName',COALESCE(c.CollegeName,'') 'CollegeName',COALESCE(d.CourseName,'') 'CourseName'  FROM  testpaper a  left join  major b  on a.MajorId=b.MajorId   left join  college  c  on a.CollegeId=c.Id 	 left join  course d on a.CourseId=d.Id  where a.IsDel=0  and  a.SchoolId=?", SchoolId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	for rows.Next() {
		newtestPaperView := new(model.TestPaperView)
		err := rows.Scan(&newtestPaperView.Id, &newtestPaperView.TestPaperName, &newtestPaperView.ExamDuration, &newtestPaperView.FullMarks, &newtestPaperView.PassScore, &newtestPaperView.TestPaperType, &newtestPaperView.CollegeId, &newtestPaperView.MajorId, &newtestPaperView.CourseId, &newtestPaperView.TeacherId, &newtestPaperView.SchoolId, &newtestPaperView.MajorName, &newtestPaperView.CollegeName, &newtestPaperView.CourseName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		testPaperViewarr = append(testPaperViewarr, newtestPaperView)

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": testPaperViewarr,
	})
}

func GetDelTestPaperByTeacherId(c *gin.Context) {

	TeacherId := c.Query("TeacherId")
	SchoolId := c.Query("SchoolId")
	tempstr := ""

	strquery := " select   a.Id,a.TestPaperName,a.ExamDuration,a.FullMarks,a.PassScore,a.TestPaperType,a.CollegeId,a.MajorId,a.CourseId,a.TeacherId,a.SchoolId,  COALESCE(b.MajorName,'') 'MajorName',COALESCE(c.CollegeName,'') 'CollegeName',COALESCE(d.CourseName,'') 'CourseName' FROM  testpaper a  left join  major b  on a.MajorId=b.MajorId   left join  college  c  on a.CollegeId=c.Id 	 left join  course d on a.CourseId=d.Id  where a.IsDel=1  "
	if TeacherId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	if TeacherId == "0" {
		strquery += "and  a.SchoolId=?  "

		if SchoolId == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "参数错误",
				"data": "{}",
			})
			return
		}
		tempstr = SchoolId
	} else {
		strquery += "and a.TeacherId=? "
		tempstr = TeacherId
	}
	// 	select a.QuestionId,a.QuestionName,a.QuestionType,a.QuestionPoolId, CASE  WHEN a.QuestionType=5 THEN b.FilePath ELSE 	a.QuestionContent END  'QuestionContent' ,COALESCE(b.FileName,'') 'FileName' ,a.Answer
	//  from question a
	// left join fileinfo b on a.QuestionContent=b.Id

	var testPaperViewarr []*model.TestPaperView
	rows, err := lib.Db.Query(strquery, tempstr)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	for rows.Next() {
		newtestPaperView := new(model.TestPaperView)
		err := rows.Scan(&newtestPaperView.Id, &newtestPaperView.TestPaperName, &newtestPaperView.ExamDuration, &newtestPaperView.FullMarks, &newtestPaperView.PassScore, &newtestPaperView.TestPaperType, &newtestPaperView.CollegeId, &newtestPaperView.MajorId, &newtestPaperView.CourseId, &newtestPaperView.TeacherId, &newtestPaperView.SchoolId, &newtestPaperView.MajorName, &newtestPaperView.CollegeName, &newtestPaperView.CourseName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		testPaperViewarr = append(testPaperViewarr, newtestPaperView)

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": testPaperViewarr,
	})
}

func GetTestPaperByPaperId(c *gin.Context) {

	TestPaperId := c.Query("TestPaperId")
	testPaperDetails := new(model.TestPaper)
	strjson := ""
	lib.Db.QueryRow("select Id,TestPaperName,ExamDuration,FullMarks,PassScore,TestPaperType,CollegeId,MajorId,CourseId,QuestionScoreJson,TeacherId,SchoolId from testpaper where IsDel=0 and Id=?", TestPaperId).
		Scan(&testPaperDetails.Id, &testPaperDetails.TestPaperName, &testPaperDetails.ExamDuration, &testPaperDetails.FullMarks, &testPaperDetails.PassScore, &testPaperDetails.TestPaperType, &testPaperDetails.CollegeId, &testPaperDetails.MajorId, &testPaperDetails.CourseId, &strjson, &testPaperDetails.TeacherId, &testPaperDetails.SchoolId)
	if testPaperDetails.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}
	var questionscorejson []*model.QuestionScoreJson
	json.Unmarshal([]byte(strjson), &questionscorejson)
	testPaperDetails.QuestionScoreJson = questionscorejson
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": testPaperDetails,
	})
}

// mark 需要修改
func GetTestPaperByTestPaperId(c *gin.Context) {

	TestPaperId := c.Query("TestPaperId")
	testPaperDetails := new(model.TestPaperDetails)
	lib.Db.QueryRow("select   a.Id,a.TestPaperName,a.ExamDuration,a.FullMarks,a.PassScore,a.TestPaperType,a.CollegeId,a.MajorId,a.CourseId,a.TeacherId,a.SchoolId,  COALESCE(b.MajorName,'') 'MajorName',COALESCE(c.CollegeName,'') 'CollegeName',COALESCE(d.CourseName,'') 'CourseName',COALESCE(e.SchoolName,'') 'SchoolName',COALESCE(f.TeacherName,'')  'TeacherName' FROM  testpaper a  left join  major b  on a.MajorId=b.MajorId   left join  college  c  on a.CollegeId=c.Id 	 left join  course d on a.CourseId=d.Id 		left join school e on a.SchoolId=e.Id  		LEFT JOIN  teacher f on a.TeacherId=f.TeacherId  where a.IsDel=0  and  a.Id=?", TestPaperId).
		Scan(&testPaperDetails.Id, &testPaperDetails.TestPaperName, &testPaperDetails.ExamDuration, &testPaperDetails.FullMarks, &testPaperDetails.PassScore, &testPaperDetails.TestPaperType, &testPaperDetails.CollegeId, &testPaperDetails.MajorId, &testPaperDetails.CourseId, &testPaperDetails.TeacherId, &testPaperDetails.SchoolId, &testPaperDetails.MajorName, &testPaperDetails.CollegeName, &testPaperDetails.CourseName, &testPaperDetails.SchoolName, &testPaperDetails.TeacherName)
	if testPaperDetails.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}

	strsql := " select	a.QuestionId,COALESCE(b.QuestionName,'') 'QuestionName',COALESCE(b.QuestionPoolId,0) 'QuestionPoolId'  ,COALESCE(b.QuestionType,0) 'QuestionType' " +
		",COALESCE(b.QuestionContent,'')'QuestionContent' ,COALESCE(a.QuestionScore,0) 'QuestionScore',COALESCE(b.Digree,0) 'Digree',COALESCE(b.Answer,0)'Answer' from testpaperquestion a " +
		"    LEFT JOIN  question b on a.QuestionId=b.QuestionId " +
		"		   where a.TestPaperId=? order by b.QuestionType asc "
	rows, err := lib.Db.Query(strsql, testPaperDetails.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	testPaperDetails.TestPaperQuestionTypeOver = make([]*model.QuestionTypeOver, 0)

	for i := 1; i <= 5; i++ {
		temparr := new(model.QuestionTypeOver)
		temparr.QuestionIdNum = 0
		temparr.QuestionScore = 0
		temparr.QuestionType = i
		testPaperDetails.TestPaperQuestionTypeOver = append(testPaperDetails.TestPaperQuestionTypeOver, temparr)
	}
	testPaperDetails.TestPaperQuestionViewFile = make([]*model.TestPaperQuestionViewFile, 0)
	for rows.Next() {
		newTestPaperQuestionViewFile := new(model.TestPaperQuestionViewFile)

		newTestPaperQuestionViewFile.FileInfo = make([]*model.FileInfo, 0)
		err := rows.Scan(&newTestPaperQuestionViewFile.QuestionId, &newTestPaperQuestionViewFile.QuestionName, &newTestPaperQuestionViewFile.QuestionPoolId, &newTestPaperQuestionViewFile.QuestionType, &newTestPaperQuestionViewFile.QuestionContent, &newTestPaperQuestionViewFile.QuestionScore, &newTestPaperQuestionViewFile.Digree, &newTestPaperQuestionViewFile.Answer)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}

		if newTestPaperQuestionViewFile.QuestionType != 0 {
			testPaperDetails.TestPaperQuestionTypeOver[newTestPaperQuestionViewFile.QuestionType-1].QuestionIdNum += 1
			testPaperDetails.TestPaperQuestionTypeOver[newTestPaperQuestionViewFile.QuestionType-1].QuestionScore += newTestPaperQuestionViewFile.QuestionScore
		}

		testPaperDetails.TestPaperQuestionViewFile = append(testPaperDetails.TestPaperQuestionViewFile, newTestPaperQuestionViewFile)
	}

	for i := 0; i < len(testPaperDetails.TestPaperQuestionViewFile); i++ {
		newTestPaperQuestionViewFile := testPaperDetails.TestPaperQuestionViewFile[i]
		if newTestPaperQuestionViewFile.QuestionType == 5 {
			fileid := newTestPaperQuestionViewFile.QuestionContent

			fileinfo := new(model.FileInfo)
			lib.Db.QueryRow("select Id,FileType,FilePath,FileName from fileinfo    where Id=? ", fileid).Scan(&fileinfo.Id, &fileinfo.FileType, &fileinfo.FilePath, &fileinfo.FileName)
			if fileinfo.Id > 0 {
				newTestPaperQuestionViewFile.FileInfo = append(newTestPaperQuestionViewFile.FileInfo, fileinfo)
			}
		} else {
			rowsa, err := lib.Db.Query("select b.Id,b.FileType,b.FilePath,b.FileName from questionrelation  a left join fileinfo b on a.FileId=b.Id where a.QuestionId=? ", newTestPaperQuestionViewFile.QuestionId)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			for rowsa.Next() {
				newfileinfo := new(model.FileInfo)
				err := rowsa.Scan(&newfileinfo.Id, &newfileinfo.FileType, &newfileinfo.FilePath, &newfileinfo.FileName)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "失败",
						"data": "{}",
					})
					return
				}
				newTestPaperQuestionViewFile.FileInfo = append(newTestPaperQuestionViewFile.FileInfo, newfileinfo)
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": testPaperDetails,
	})
}
func AddTestPaperQuestion(tx *sql.Tx, c *gin.Context, testpaperquestionarr []*model.TestPaperQuestion) error {

	for i := 0; i < len(testpaperquestionarr); i++ {

		ret, err := tx.Exec("insert into testpaperquestion(QuestionId,QuestionType,QuestionScore,TestPaperId) values(?,?,?,?) ", testpaperquestionarr[i].QuestionId, testpaperquestionarr[i].QuestionType, testpaperquestionarr[i].QuestionScore, testpaperquestionarr[i].TestPaperId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return err
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return err
		}
		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return errors.New("操作失败")
		}
	}
	return nil

}
