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
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginPatrolUser(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}
	log.Printf("%s", body)

	var patrol model.PatrolUser
	json.Unmarshal([]byte(body), &patrol)

	data := []byte(patrol.PatrolUserPwd)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	newpatrol := new(model.PatrolUser)

	err = lib.Db.QueryRow("SELECT patrolUserId,patrolUserAccount,patrolUserPwd FROM  patroluser where patrolUserAccount=? and patrolUserPwd=?", patrol.PatrolUserAccount, md5str1).
		Scan(&newpatrol.PatrolUserId, &newpatrol.PatrolUserAccount, &newpatrol.PatrolUserPwd)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
			"data": "{}",
		})
		return
	}

	if newpatrol.PatrolUserId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "账号或密码错误",
			"data": "{}",
		})
		return
	} else {
		tokenString, _ := jwt_use.GetToken(newpatrol.PatrolUserAccount, 3)
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "登录成功",
			"data":  newpatrol,
			"token": tokenString,
		})
		return
	}
}

func GetLoginLog(c *gin.Context) {
	query1 := `SELECT
	DATE_FORMAT( t.time_slot , '%H') as  hour,
		''  as day,
	COUNT(e.loginLogId)    as 'total',
		'1' as  'datetype' 
	FROM
	time_intervals t
	left  JOIN
	loginlog e
	ON
	 DATE_FORMAT(FROM_UNIXTIME(e.loginTime) , '%H')  = HOUR(t.time_slot)  and DATE(FROM_UNIXTIME(e.loginTime)) = CURDATE()
	
	GROUP BY   hour
	order by   hour`

	query2 := `SELECT
    DATE_FORMAT( t.time_slot , '%H') as  'hour',
		'' as day,
    COUNT(e.loginLogId)    as 'total',
		'2' as  'datetype' 
FROM
    time_intervals t
left  JOIN
    loginlog e
ON
     DATE_FORMAT(FROM_UNIXTIME(e.loginTime) , '%H')  = HOUR(t.time_slot)  and TO_DAYS(NOW()) - TO_DAYS(FROM_UNIXTIME(e.loginTime)) = 1
GROUP BY   hour
order by   hour`

	query3 := `WITH  RECURSIVE date_range AS (
		SELECT CURDATE() AS day
		UNION ALL
		SELECT day - INTERVAL 1 DAY
		FROM date_range
		WHERE day > CURDATE() - INTERVAL 7 DAY
	  )
	  SELECT
		'' hour,
	   t.day,
	   COALESCE(count(a.loginTime), 0) AS 'total',
		   '3' as  'datetype'
	  FROM date_range t
	  LEFT JOIN loginlog a   ON t.day = DATE_FORMAT(FROM_UNIXTIME(a.loginTime) , '%Y-%m-%d')
	  
	  GROUP BY day
	  ORDER BY  day;`

	rows, err := lib.Db.Query(query1)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
	}

	loginlogarrresp := new(model.LoginLogArr)
	nowday := make([]*model.LoginLog, 0)
	for rows.Next() {

		tempmodel := new(model.LoginLog)
		rows.Scan(&tempmodel.Hour, &tempmodel.Day, &tempmodel.Total, &tempmodel.DateType)
		nowday = append(nowday, tempmodel)
	}
	loginlogarrresp.NowDay = nowday

	rows, err = lib.Db.Query(query2)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
	}
	yesterday := make([]*model.LoginLog, 0)
	for rows.Next() {
		tempmodel := new(model.LoginLog)
		rows.Scan(&tempmodel.Hour, &tempmodel.Day, &tempmodel.Total, &tempmodel.DateType)
		yesterday = append(yesterday, tempmodel)
	}
	loginlogarrresp.Yesterday = yesterday

	rows, err = lib.Db.Query(query3)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
	}
	sevenday := make([]*model.LoginLog, 0)
	for rows.Next() {
		tempmodel := new(model.LoginLog)
		rows.Scan(&tempmodel.Hour, &tempmodel.Day, &tempmodel.Total, &tempmodel.DateType)
		sevenday = append(sevenday, tempmodel)
	}
	loginlogarrresp.Sevenday = sevenday
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": loginlogarrresp,
	})
}

func GetPatrolExamPlan(c *gin.Context) {
	strsqlQuery := ` select  a.Id,a.StartTime,a.EndTime,b.TestPaperName,c.ExamName,
	CASE
		  WHEN NOW() BETWEEN a.StartTime AND a.EndTime THEN '正在进行'
		  WHEN NOW() < a.StartTime THEN '未开始'
		  WHEN NOW() >  a.EndTime THEN '已过期'
	  END AS 'Status',
	  (SELECT COUNT(1) FROM examstudent WHERE ExamId=a.ExamId) 'Personcount'
   from examsession a 
  left join testpaper b on a.TestPaperId=b.Id
  left join exam c on a.ExamId=c.Id
  where a.IsDel=0 and  c.ReviewFlag=1 `
	rows, err := lib.Db.Query(strsqlQuery)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	var patroExamPlanarr []*model.PatroExamPlan
	for rows.Next() {
		patroExamPlan := new(model.PatroExamPlan)
		rows.Scan(&patroExamPlan.Id, &patroExamPlan.StartTime, &patroExamPlan.EndTime, &patroExamPlan.TestPaperName, &patroExamPlan.ExamName, &patroExamPlan.Status, &patroExamPlan.Personcount)
		patroExamPlanarr = append(patroExamPlanarr, patroExamPlan)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": patroExamPlanarr,
	})
}

func GetPatrolExamOverView(c *gin.Context) {

	strsqlQuery := "select a.Id 'ExamSessionId',a.ExamId, CONCAT(d.ExamName,'/',b.TestPaperName) as 'ExamSession',e.MajorName,CONCAT(b.PassScore,'/',b.FullMarks) 'Score',a.StartTime" +
		" ,(select COALESCE(avg(Score),0)  from  examsturesult   where    ExamId=a.ExamId   and   ExamSessionId=a.Id and Score<>-1	) 'AvgScore', " +
		"	(SELECT COUNT(1) 'wks'  FROM `examsturesult`  where ExamId=a.ExamId and  ExamSessionId=a.Id and Score=-1) 'UnexaminedNum'," +
		"(SELECT COUNT(1) 'yks'  FROM `examsturesult`  where ExamId=a.ExamId and  ExamSessionId=a.Id and Score<>-1 and ExamStatus=1) 'ExaminedNum',b.ExamDuration" +
		"  from examsession a " +
		" left join testpaper b on a.TestPaperId=b.Id " +
		" left JOIN exam d on a.ExamId=d.Id " +
		" LEFT JOIN major e on b.MajorId=e.MajorId WHERE  d.ReviewFlag=1"

	rows, err := lib.Db.Query(strsqlQuery)
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

func GetPatrolExamSessionResult(c *gin.Context) {

	ExamSessionId := c.Query("ExamSessionId")
	strsqlQuery := "  select a.Id,b.TrueName,b.IDNumber,a.StudentId,a.ExamId, a.ExamSessionId,COALESCE( a.StartExamTime,'') 'StartExamTime' " +
		",COALESCE( a.EndExamTime,'') 'EndExamTime',a.Score,COALESCE(a.ExamStatus,0)  'ExamStatus', COALESCE(a.ExamType,0) 'ExamType'" +
		"	from examsturesult a" +
		" LEFT JOIN student b on a.StudentId=b.Id" +
		"	where   ExamSessionId=?    ORDER BY a.score asc"
	rows, err := lib.Db.Query(strsqlQuery, ExamSessionId)
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

func GetStudentIdImg(c *gin.Context) {
	imgpath := ""
	studentId := c.Query("studentId")

	sqlStr := "select IDImage from  student  where Id=? limit 1 "

	err := lib.Db.QueryRow(sqlStr, studentId).Scan(&imgpath)
	if err != nil {
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
		"data": imgpath,
	})
}

func GetRealTimeExamSituation(c *gin.Context) {

	resp := new(model.RealTimeExamSituation)

	sqlStr := `SELECT count(1) 'sumCount',(SELECT  count(1)   FROM examsturesult  where Score >=0) 'doneCount',(SELECT  count(1)   FROM examsturesult  where Score <0)  'undoneCount' FROM examsturesult`

	err := lib.Db.QueryRow(sqlStr).Scan(&resp.SumCount, &resp.DoneCount, &resp.UnDoneCount)

	if err != nil {
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
		"data": resp,
	})
}

func GetStandGL(c *gin.Context) {

	query := `SELECT a.*,COALESCE(b.postion,'')  'postion',
	(select count(1) from student where StandId=a.Id) 'personCount',
	  (SELECT  count(1)   FROM examsturesult c left join  student d on c.StudentId=d.Id     where d.standId=a.Id and  c.Score >=0) 'doneCount', 
	(SELECT  count(1)   FROM examsturesult c left join  student d on c.StudentId=d.Id     where d.standId=a.Id and     c.Score < 0) 'undoneCount' 
	 FROM  stand a
	left join standpostion b on a.Id=b.standId`

	rows, err := lib.Db.Query(query)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var standpostion []*model.StandPostion
	for rows.Next() {
		temp := new(model.StandPostion)
		err = rows.Scan(&temp.Id, &temp.StandName, &temp.SchoolId, &temp.TeacherId, &temp.Postion, &temp.PersonCount, &temp.DoneCount, &temp.UnDoneCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		standpostion = append(standpostion, temp)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": standpostion,
	})
}
