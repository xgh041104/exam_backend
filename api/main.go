package main

import (
	chapter "StudyExamPlatformAPI/Interface/Chapter"
	cheatwarning "StudyExamPlatformAPI/Interface/CheatWarning"
	class "StudyExamPlatformAPI/Interface/Class"
	college "StudyExamPlatformAPI/Interface/College"
	"StudyExamPlatformAPI/Interface/Course"
	exam "StudyExamPlatformAPI/Interface/Exam"
	major "StudyExamPlatformAPI/Interface/Major"
	plan "StudyExamPlatformAPI/Interface/Plan"
	question "StudyExamPlatformAPI/Interface/Question"
	school "StudyExamPlatformAPI/Interface/School"
	"StudyExamPlatformAPI/Interface/Stand"
	testpaper "StudyExamPlatformAPI/Interface/TestPaper"
	tool "StudyExamPlatformAPI/Interface/Tool"
	"StudyExamPlatformAPI/Interface/User"
	"StudyExamPlatformAPI/jwt_use"
	"StudyExamPlatformAPI/lib"
	"StudyExamPlatformAPI/lib/ase"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"

	// "github.com/essentialkaos/ek/v12/log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// originalPwd := "admin" + "!q@w#e$r%t"
	// data := []byte(originalPwd)
	// has := md5.Sum(data)
	// md5str1 := fmt.Sprintf("%x", has)
	// fmt.Println(md5str1)
	lib.InitDB() // 调用输出化数据库的函数
	gin.SetMode(gin.DebugMode)

	f, _ := os.Create("gin.log")
	// // gin.DefaultWriter = io.MultiWriter(f)
	// // 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()

	r.Use(Cors()) //开启中间件 允许使用跨域请求
	//学生
	r.POST("/Student/LoginStudent", jwtUserLoginAuthMiddleware(), User.LoginStudent)
	r.POST("/Student/EditStudent", jwtAuthMiddleware(), User.EditStudent)
	//查询学生个人信息
	r.GET("/Student/GetStudentViewById", jwtAuthMiddleware(), User.GetStudentViewById)

	//老师登录
	r.POST("/Manage/LoginTeacher", jwtUserLoginAuthMiddleware(), User.LoginTeacher)
	r.POST("/Manage/EditTeacher", User.EditTeacher)
	r.POST("/Manage/AddTeacher", User.AddTeacher)
	r.POST("/Manage/DelTeacher", User.DelTeacher)
	r.POST("/Manage/EditTeacherPassWordById", User.EditTeacherPassWordById)
	r.GET("/Manage/GetTeacherById", User.GetTeacherById)
	r.GET("/Manage/GetTeacherList", User.GetTeacherList)

	// 管理员接口
	r.POST("/Manage/LoginAdmin", jwtUserLoginAuthMiddleware(), User.LoginAdmin)

	//学校
	r.GET("/Manage/SchoolList", jwtAuthMiddleware(), school.SchoolList)
	r.POST("/Manage/EditSchool", jwtAuthMiddleware(), school.EditSchool)
	r.POST("/Manage/DelSchool", jwtAuthMiddleware(), school.DelSchool)
	r.POST("/Manage/AddSchool", jwtAuthMiddleware(), school.AddSchool)

	//站点
	r.GET("/Manage/GetStandListBySchoolId", jwtAuthMiddleware(), Stand.GetStandListBySchoolId)
	r.POST("/Manage/EditStand", jwtAuthMiddleware(), Stand.EditStand)
	r.POST("/Manage/AddStand", jwtAuthMiddleware(), Stand.AddStand)
	r.POST("/Manage/DelStand", jwtAuthMiddleware(), Stand.DelStand)
	//学院
	r.POST("/Manage/EditCollege", jwtAuthMiddleware(), college.EditCollege)
	r.POST("/Manage/AddCollege", jwtAuthMiddleware(), college.AddCollege)
	r.POST("/Manage/DelCollege", jwtAuthMiddleware(), college.DelCollege)

	r.GET("/Manage/GetCollegeListAll", jwtAuthMiddleware(), college.GetCollegeListAll)
	r.GET("/Manage/GetCollegeByCollegeId", jwtAuthMiddleware(), college.GetCollegeByCollegeId)

	//专业
	r.POST("/Manage/EditMajor", jwtAuthMiddleware(), major.EditMajor)
	r.POST("/Manage/AddMajor", jwtAuthMiddleware(), major.AddMajor)
	r.POST("/Manage/DelMajor", jwtAuthMiddleware(), major.DelMajor)
	r.GET("/Manage/GetMajorBySchoolId", jwtAuthMiddleware(), major.GetMajorBySchoolId)
	r.GET("/Manage/GetMajorListView", jwtAuthMiddleware(), major.GetMajorListView)

	// 班级
	r.POST("/Manage/EditClass", jwtAuthMiddleware(), class.EditClass)
	r.POST("/Manage/AddClass", jwtAuthMiddleware(), class.AddClass)
	r.POST("/Manage/DelClass", jwtAuthMiddleware(), class.DelClass)
	r.GET("/Manage/GetClassBySchoolId", jwtAuthMiddleware(), class.GetClassBySchoolId)
	r.GET("/Manage/GetClassStudentsBySchoolId", jwtAuthMiddleware(), class.GetClassStudentsBySchoolId)
	r.GET("/Manage/GetClassView", jwtAuthMiddleware(), class.GetClassView)

	//学生使用
	r.GET("/Student/GetCurrentStudyCourse", jwtAuthMiddleware(), Course.GetCurrentStudyCourse)
	r.GET("/Student/CourseDetails", jwtAuthMiddleware(), Course.CourseDetails)
	r.GET("/Student/GetChapterById", jwtAuthMiddleware(), chapter.GetChapterById)
	r.POST("/Student/StudyPlanUpload", jwtAuthMiddleware(), chapter.StudyPlanUpload)

	//课程
	r.POST("/Manage/AddCourse", jwtAuthMiddleware(), Course.AddCourse)
	r.POST("/Manage/EditCourse", jwtAuthMiddleware(), Course.EditCourse)
	r.POST("/Manage/EditCourseFile", jwtAuthMiddleware(), Course.EditCourseFile)
	r.POST("/Manage/EditClassRelation", jwtAuthMiddleware(), Course.EditClassRelation)
	r.POST("/Manage/DelCourse", jwtAuthMiddleware(), Course.DelCourse)
	r.GET("/Manage/GetCourseClassRelationById", jwtAuthMiddleware(), Course.GetCourseClassRelationById)
	r.GET("/Manage/GetCourseListByTeacherId", jwtAuthMiddleware(), Course.GetCourseListByTeacherId)
	r.GET("/Manage/GetCourseByCourseId", jwtAuthMiddleware(), Course.GetCourseByCourseId)
	r.GET("/Manage/GetCouresStudyPlanByCourseId", jwtAuthMiddleware(), Course.GetCouresStudyPlanByCourseId)
	r.POST("/Manage/CourseCancel", jwtAuthMiddleware(), Course.CourseCancel)

	//章节
	//TODO jwtAuthMiddleware(),
	r.POST("/Manage/AddChapter", chapter.AddChapter)
	r.POST("/Manage/EditChapter", jwtAuthMiddleware(), chapter.EditChapter)
	r.POST("/Manage/DelChapter", jwtAuthMiddleware(), chapter.DelChapter)
	r.POST("/Manage/UploadChapterFile", jwtAuthMiddleware(), chapter.UploadChapterFile)
	r.GET("/Manage/GetChapterById", jwtAuthMiddleware(), chapter.GetChapterById)
	r.GET("/Manage/GetChapterByCourseId", jwtAuthMiddleware(), chapter.GetChapterByCourseId)
	r.POST("/Manage/UpdateChapterOrder", jwtAuthMiddleware(), chapter.UpdateChapterOrder)

	//TODO jwtAuthMiddleware(),
	r.POST("/Manage/UploadChapterVideoFile", chapter.UploadChapterVideoFile)

	//题目
	r.POST("/Manage/AddQuestion", jwtAuthMiddleware(), question.AddQuestion)
	r.POST("/Manage/EditQuestion", jwtAuthMiddleware(), question.EditQuestion)
	r.POST("/Manage/DelQuestion", jwtAuthMiddleware(), question.DelQuestion)
	r.GET("/Manage/GetQuestionBySchoolId", jwtAuthMiddleware(), question.GetQuestionBySchoolId)
	r.GET("/Manage/GetQuestionByQuestionId", jwtAuthMiddleware(), question.GetQuestionByQuestionId)
	r.GET("/Student/GetQuestionBySchoolId", jwtAuthMiddleware(), question.GetQuestionBySchoolId)
	r.GET("/Manage/GetExamQuestionBySchoolId", jwtAuthMiddleware(), question.GetExamQuestionBySchoolId)
	r.GET("/Student/GetTrainQuestionBySchoolId", jwtAuthMiddleware(), question.GetTrainQuestionBySchoolId)
	r.GET("/Student/GetSocietyQuestionByStudentId", jwtAuthMiddleware(), question.GetSocietyQuestionByStudentId)
	r.GET("/Student/GetQuestionByQuestionId", jwtAuthMiddleware(), question.GetQuestionByQuestionId)
	r.POST("/Manage/GetQuestionExeclData", jwtAuthMiddleware(), question.GetQuestionExeclData)
	r.POST("/Manage/MatchAddQuestion", jwtAuthMiddleware(), question.MatchAddQuestion)
	r.GET("/Manage/GetOperPracticeByQuestionId", jwtAuthMiddleware(), question.GetOperPracticeByQuestionId)
	r.POST("/Student/AddOperPractice", question.AddOperPractice)

	//错题
	r.POST("/Student/AddQuestionWrong", jwtAuthMiddleware(), question.AddQuestionWrong)
	r.POST("/Student/DelQuestionWrong", jwtAuthMiddleware(), question.DelQuestionWrong)
	r.GET("/Student/GetQuestionWrongByStudentId", jwtAuthMiddleware(), question.GetQuestionWrongByStudentId)

	//试卷
	r.POST("/Manage/AddTestPaper", jwtAuthMiddleware(), testpaper.AddTestPaper)
	r.POST("/Manage/EditTestPaper", jwtAuthMiddleware(), testpaper.EditTestPaper)
	r.POST("/Manage/DelTestPaper", jwtAuthMiddleware(), testpaper.DelTestPaper)
	r.GET("/Manage/GetTestPaperBySchoolId", jwtAuthMiddleware(), testpaper.GetTestPaperBySchoolId)
	r.GET("/Manage/GetTestPaperByTestPaperId", jwtAuthMiddleware(), testpaper.GetTestPaperByTestPaperId)
	r.GET("/Manage/GetTestPaperByPaperId", jwtAuthMiddleware(), testpaper.GetTestPaperByPaperId)
	r.GET("/Manage/GetDelTestPaperByTeacherId", jwtAuthMiddleware(), testpaper.GetDelTestPaperByTeacherId)

	//考试与考场
	r.POST("/Manage/AddExam", jwtAuthMiddleware(), exam.AddExam)
	r.POST("/Manage/EditExam", jwtAuthMiddleware(), exam.EditExam)
	r.POST("/Manage/DelExam", jwtAuthMiddleware(), exam.DelExam)

	r.POST("/Manage/AddExamSession", jwtAuthMiddleware(), exam.AddExamSession)
	r.POST("/Manage/EditExamSession", jwtAuthMiddleware(), exam.EditExamSession)
	r.POST("/Manage/DelExamSession", jwtAuthMiddleware(), exam.DelExamSession)
	r.POST("/Manage/AddExamStudent", jwtAuthMiddleware(), exam.AddExamStudent)
	r.POST("/Manage/EditExamStudent", jwtAuthMiddleware(), exam.EditExamStudent)
	r.GET("/Manage/GetExamBySchoolId", jwtAuthMiddleware(), exam.GetExamBySchoolId)
	r.GET("/Manage/GetExamReViewBySchoolId", jwtAuthMiddleware(), exam.GetExamReViewBySchoolId)
	r.POST("/Manage/EditReViewExamByExamId", jwtAuthMiddleware(), exam.EditReViewExamByExamId)
	r.GET("/Manage/GetExamByExamId", jwtAuthMiddleware(), exam.GetExamByExamId)

	r.GET("/Student/GetStudentExamInfo", exam.GetStudentExamInfo)       // jwtAuthMiddleware(),
	r.GET("/Student/GetStudentExamDetails", exam.GetStudentExamDetails) //jwtAuthMiddleware(),
	r.POST("/Student/ExamStudentSumit", jwtAuthMiddleware(), exam.ExamStudentSumit)

	r.GET("/Student/GetStudentExamPaperOver", jwtAuthMiddleware(), exam.GetStudentExamPaperOver)
	r.POST("/Student/UploadExamImage", jwtAuthMiddleware(), exam.UploadExamImage)

	r.POST("/Manage/AddExamBatchSessionStudent", jwtAuthMiddleware(), exam.AddExamBatchSessionStudent)

	r.POST("/Manage/ExamCancel", jwtAuthMiddleware(), exam.ExamCancel)
	//考试标题通知
	r.POST("/Manage/EditExamNotice", jwtAuthMiddleware(), exam.EditExamNotice)
	r.POST("/Manage/AddExamNotice", jwtAuthMiddleware(), exam.AddExamNotice)
	r.POST("/Manage/DelExamNotice", jwtAuthMiddleware(), exam.DelExamNotice)
	r.GET("/Manage/GetExamNoticeByExamId", jwtAuthMiddleware(), exam.GetExamNoticeByExamId)
	r.GET("/Student/GetExamNoticeByExamId", jwtAuthMiddleware(), exam.GetExamNoticeByExamId)
	//成绩查询
	r.GET("/Manage/GetExamSessionResult", jwtAuthMiddleware(), exam.GetExamSessionResult)
	r.GET("/Manage/GetExamOverView", jwtAuthMiddleware(), exam.GetExamOverView)

	r.POST("/StandUser/LoginStandUser", Stand.LoginStandUser)
	r.GET("/StandUser/QueryStudentResultByStandId", jwtAuthMiddleware(), exam.QueryStudentResultByStandId)
	r.GET("/StandUser/QueryStudentResultExeclByStandId", jwtAuthMiddleware(), exam.QueryStudentResultExeclByStandId)

	//成绩查询补考管理
	r.GET("/Manage/GetCurrentExamSessionBKStudent", jwtAuthMiddleware(), exam.GetCurrentExamSessionBKStudent)
	r.POST("/Manage/AddExamReset", jwtAuthMiddleware(), exam.AddExamReset)
	r.GET("/Manage/GetExamResetByExamSessionId", jwtAuthMiddleware(), exam.GetExamResetByExamSessionId)

	//学生
	r.GET("/Manage/GetStudentViewList", jwtAuthMiddleware(), User.GetStudentViewList)
	r.POST("/Manage/AddStudent", jwtAuthMiddleware(), User.AddStudent)
	r.POST("/Manage/EditStudent", jwtAuthMiddleware(), User.EditStudent)
	r.POST("/Manage/DelStudent", jwtAuthMiddleware(), User.DelStudent)
	r.POST("/Manage/EditFaceOpenState", jwtAuthMiddleware(), User.EditFaceOpenState)
	r.POST("/Manage/EditStudentPassWordById", jwtAuthMiddleware(), User.EditStudentPassWordById)
	r.POST("/Manage/ReSetStudentPassWord", jwtAuthMiddleware(), User.ReSetStudentPassWord)
	r.POST("/Manage/GetStudentExeclData", jwtAuthMiddleware(), User.GetStudentExeclData)
	r.POST("/Manage/MatchAddStudentIDImg", jwtAuthMiddleware(), User.MatchAddStudentIDImg)

	r.POST("/Manage/MatchAddStudent", User.MatchAddStudent)
	r.POST("/Manage/MatchAddSocietyStudentExecl", jwtAuthMiddleware(), User.MatchAddSocietyStudentExecl)

	// 公共接口
	r.GET("/Manage/GetCollegeBySchoolId", jwtAuthMiddleware(), college.GetCollegeBySchoolId) //学院查询
	r.GET("/Manage/GetMajorByCollegeId", jwtAuthMiddleware(), major.GetMajorByCollegeId)     //专业查询
	r.GET("/Manage/GetClassByMajorId", jwtAuthMiddleware(), class.GetClassByMajorId)         //班级查询
	r.GET("/Manage/Get1", tool.Get1)
	//公告列表
	r.GET("/Manage/NoticeList", tool.NoticeList)                       // 公告列表 最新的top 5
	r.GET("/Manage/NoticeListAll", tool.NoticeListAll)                 // 公告列表 所有
	r.GET("/Student/NoticeList", tool.NoticeList)                      // 公告列表 最新的top 5
	r.GET("/Student/NoticeListAll", tool.NoticeListAll)                // 公告列表 所有
	r.GET("/Manage/GetNoticeById", tool.GetNoticeById)                 // 公告列表  根据公告id 获取当前公告
	r.GET("/Student/GetNoticeById", tool.GetNoticeById)                // 公告列表  根据公告id 获取当前公告
	r.POST("/Manage/AddNotice", jwtAuthMiddleware(), tool.AddNotice)   // 添加公告
	r.POST("/Manage/EditNotice", jwtAuthMiddleware(), tool.EditNotice) // 修改公告
	r.POST("/Manage/DelNotice", jwtAuthMiddleware(), tool.DelNotice)   // 删除公告公告

	r.GET("/Student/GetTime", jwtAuthMiddleware(), tool.GetTime)
	r.GET("/Manage/GetDiskState", jwtAuthMiddleware(), tool.GetDiskState) // 获取磁盘使用情况

	r.GET("/Manage/GetHttpUrl", tool.GetHttpUrl)
	r.GET("/Student/GetHttpUrl", tool.GetHttpUrl)

	r.GET("/Student/GetFaceVerify", tool.GetFaceVerify)

	r.GET("/Manage/GetTool", jwtAuthMiddleware(), tool.GetTool)    // 获取系统设置
	r.POST("/Manage/EditTool", jwtAuthMiddleware(), tool.EditTool) //  修改系统设置

	//学习计划
	r.POST("/Student/AddQuestionRecord", jwtAuthMiddleware(), question.AddQuestionRecord)
	r.POST("/Manage/AddPlan", jwtAuthMiddleware(), plan.AddPlan)
	r.POST("/Manage/EditPlan", jwtAuthMiddleware(), plan.EditPlan)
	r.POST("/Manage/DelPlan", jwtAuthMiddleware(), plan.DelPlan)
	r.GET("/Manage/QueryPlan", jwtAuthMiddleware(), plan.QueryPlan)
	r.GET("/Manage/QuertStydyPlan", jwtAuthMiddleware(), plan.QuertStydyPlan)
	r.GET("/Manage/QueryPlanCourseProgress", jwtAuthMiddleware(), plan.QueryPlanCourseProgress)
	r.GET("/Manage/QueryPlanExamProgress", jwtAuthMiddleware(), plan.QueryPlanExamProgress)
	r.GET("/Manage/QueryPlanTrainProgress", jwtAuthMiddleware(), plan.QueryPlanTrainProgress)
	r.POST("/Manage/AddExamBatchPlanStudent", jwtAuthMiddleware(), plan.AddExamBatchPlanStudent)

	r.POST("/Manage/AddPlanStudent", jwtAuthMiddleware(), plan.AddPlanStudent)
	r.POST("/Manage/DelAllPlanStudent", jwtAuthMiddleware(), plan.DelAllPlanStudent)
	r.POST("/Manage/DelPlanStudent", jwtAuthMiddleware(), plan.DelPlanStudent)
	r.GET("/Manage/QueryPlanStudent", jwtAuthMiddleware(), plan.QueryPlanStudent)

	r.POST("/Manage/AddPlanExam", jwtAuthMiddleware(), plan.AddPlanExam)
	r.POST("/Manage/DelAllPlanExam", jwtAuthMiddleware(), plan.DelAllPlanExam)
	r.POST("/Manage/DelPlanExam", jwtAuthMiddleware(), plan.DelPlanExam)
	r.GET("/Manage/QueryPlanExam", jwtAuthMiddleware(), plan.QueryPlanExam)

	r.POST("/Manage/AddPlanTrain", jwtAuthMiddleware(), plan.AddPlanTrain)
	r.POST("/Manage/DelPlanTrain", jwtAuthMiddleware(), plan.DelPlanTrain)
	r.POST("/Manage/DelAllPlanTrain", jwtAuthMiddleware(), plan.DelAllPlanTrain)
	r.GET("/Manage/QueryPlanTrain", jwtAuthMiddleware(), plan.QueryPlanTrain)

	r.GET("/Student/QuertStudentPlan", jwtAuthMiddleware(), plan.QuertStudentPlan)
	r.POST("/Student/UploadPlanImage", jwtAuthMiddleware(), plan.UploadPlanImage)

	r.GET("/Student/QueryPlanCourseProgress", jwtAuthMiddleware(), plan.QueryPlanCourseProgress)
	r.GET("/Student/QueryPlanExamProgress", jwtAuthMiddleware(), plan.QueryPlanExamProgress)
	r.GET("/Student/QueryPlanTrainProgress", jwtAuthMiddleware(), plan.QueryPlanTrainProgress)

	r.GET("/Manage/GetStudentExamScoreByExamNumber", exam.GetStudentExamScoreByExamNumber)
	r.GET("/Manage/GetStudentImg", exam.GetStudentImg)

	r.POST("/Patrol/LoginPatrolUser", User.LoginPatrolUser)
	r.GET("/Patrol/GetLoginLog", User.GetLoginLog)
	r.GET("/Patrol/GetPatrolExamPlan", User.GetPatrolExamPlan)
	r.GET("/Patrol/GetPatrolExamOverView", User.GetPatrolExamOverView)
	r.GET("/Patrol/GetPatrolExamSessionResult", User.GetPatrolExamSessionResult)
	r.GET("/Patrol/GetStudentIdImg", User.GetStudentIdImg)
	r.POST("/Patrol/AddFaceMonitor", cheatwarning.AddFaceMonitor)
	r.GET("/Patrol/QueryFaceMonitor", cheatwarning.QueryFaceMonitor)
	r.POST("/Patrol/AddCheatWarning", cheatwarning.AddCheatWarning)
	r.DELETE("/Patrol/DelCheatWarning", cheatwarning.DelCheatWarning)
	r.GET("/Patrol/QueryCheatWarning", cheatwarning.QueryCheatWarning)
	r.GET("/Patrol/GetRealTimeExamSituation", User.GetRealTimeExamSituation)
	r.GET("/Patrol/GetStandGL", User.GetStandGL)

	r.GET("/File/Download", Download)

	r.GET("/SlideImg", SlideImg)

	r.Static("/Resources/Video", "./Resources/Video/")
	r.Static("/Resources/Annex", "./Resources/Annex/")
	r.Static("/Resources/Zip", "./Resources/Zip/")
	r.Static("/Resources/Img", "./Resources/Img/")
	r.Static("/Resources/Xml", "./Resources/Xml/")
	r.Static("/Resources/IDImg", "./Resources/IDImg/")
	r.Static("/Resources/ExamImage", "./Resources/ExamImage/")

	r.Static("/Resources/FaceMonitor", "./Resources/FaceMonitor/")
	r.Static("/Resources/CheatWarning", "./Resources/CheatWarning/")

	// 其他路由设置
	// r.Run(":7566")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002" // 如果本地运行（没有环境变量），依然使用你习惯的 7566
	}

	fmt.Println("服务启动中，监听端口:", port)

	// 启动 Gin
	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("服务启动失败: %v\n", err)
	}
}

// 中间件,认证token合法性
func jwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.Request.Header.Get("authorization")
		if authHandler == "" {
			c.JSON(200, gin.H{"code": 2003, "msg": "请求头部auth为空", "data": "{}"})
			c.Abort()
			return
		}

		// 前两部门可以直接解析出来
		jwt := strings.Split(authHandler, ".")
		cnt := 0
		for _, val := range jwt {
			cnt++
			if cnt == 3 {
				break
			}
			msg, _ := base64.StdEncoding.DecodeString(val)
			fmt.Println("val ->", string(msg))
		}

		// 我们使用之前定义好的解析JWT的函数来解析它,并且在内部解析时判断了token是否过期
		mc, err := jwt_use.ParseToken(authHandler)
		if err != nil {
			fmt.Println("err = ", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
				"data": "{}",
			})

			c.Abort()
			return
		}

		if (strings.Contains(c.FullPath(), "/Manage/") && mc.UserType != 0) ||
			(strings.Contains(c.FullPath(), "/Student/") && mc.UserType != 1) ||
			(strings.Contains(c.FullPath(), "/StandUser/") && mc.UserType != 2) ||
			(strings.Contains(c.FullPath(), "/Patrol/") && mc.UserType != 3) { //3 是巡考
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "Token不匹配请求的接口",
				"data": "{}",
			})
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.UserName)
		c.Set("userType", mc.UserType)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func jwtUserLoginAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		conf := lib.LoadConfig()
		keyStr := "1951EC8DA5"
		key := ase.GenerateKey(keyStr)
		encryptedData, err := ase.Encrypt(key, []byte(conf.YzOnlyMark))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "系统错误" + err.Error(),
				"token": "",
				"data":  "{}",
			})
			c.Abort()
			return
		}
		var reqpost PostReqParam

		reqpost.OnlyMark = encryptedData
		reqpost.HostMAC = conf.YzMAC
		reqpost.SerialNum = conf.YzYLH
		reqpost.Ciphertext = "GIKEuYU1dg361P7ofk960d9FuuFFd136qyAXB5G+N/w="
		// 定义请求的 URL
		url := "http://47.116.207.219/secretverify/verify/user/verifyuser"

		jsonBytes, err := json.Marshal(reqpost)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "系统错误" + err.Error(),
				"token": "",
				"data":  "{}",
			})
			c.Abort()
			return
		}

		// 创建一个 POST 请求
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "系统错误" + err.Error(),
				"token": "",
				"data":  "{}",
			})
			c.Abort()
			return
		}
		// 设置请求头信息，例如设置内容类型为 JSON
		req.Header.Set("Content-Type", "application/json")

		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "系统错误" + err.Error(),
				"token": "",
				"data":  "{}",
			})
			c.Abort()
			return
		}
		defer resp.Body.Close()
		// 读取响应数据
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "系统错误" + err.Error(),
				"token": "",
				"data":  "{}",
			})
			c.Abort()
			return
		}

		var myMap map[string]interface{}

		err = json.Unmarshal(body, &myMap)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "系统错误" + err.Error(),
				"token": "",
				"data":  "{}",
			})
			c.Abort()
			return
		}

		if myMap["code"].(float64) == 0 {
			c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "系统错误",
				"token": "",
				"data":  "{}",
			})
			c.Abort()
			return
		}

	}
}

type PostReqParam struct {
	OnlyMark   string `json:"onlyMark"`
	HostMAC    string `json:"hostMAC"`
	SerialNum  string `json:"serialNum"`
	Ciphertext string `json:"ciphertext"`
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT, ")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,ETag")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}

func Download(c *gin.Context) {

	filename := c.Query("filename")      //截取get请求参数，也就是图片的路径，可是使用绝对路径，也可使用相对路径
	file, _ := ioutil.ReadFile(filename) //把要显示的图片读取到变量中
	c.Writer.WriteString(string(file))   //关键一步，写给前端

}

func SlideImg(c *gin.Context) {

	randomInt := rand.Intn(16) + 1
	file, err := ioutil.ReadFile("Resources/SlideImg/" + lib.Strval(randomInt) + ".png") //把要显示的图片读取到变量中
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "文件获取失败",
			"data": "{}",
		})
		return
	}
	c.Writer.WriteString(string(file)) //关键一步，写给前端
}
