package tool

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func EditNotice(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var notice model.Notice
	json.Unmarshal([]byte(body), &notice)

	sqlStr := "Update  notice set  Time=?,NoticeTitle=?,NoticeContent=?,SendUser=?,NoticeLevel=? ,NoticeType=? where Id=? "

	timestr := time.Now().Format("2006-01-02 15:04:05")
	ret, err := lib.Db.Exec(sqlStr, timestr, notice.NoticeTitle, notice.NoticeContent, notice.SendUser, notice.NoticeLevel, notice.NoticeType, notice.Id)

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
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
	}
}

func AddNotice(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var notice model.Notice
	json.Unmarshal([]byte(body), &notice)

	sqlStr := "insert into  notice(Time,NoticeTitle,NoticeContent,SendUser,NoticeLevel,NoticeType) VALUES (?,?,?,?,?,?) "

	timestr := time.Now().Format("2006-01-02 15:04:05")
	ret, err := lib.Db.Exec(sqlStr, timestr, notice.NoticeTitle, notice.NoticeContent, notice.SendUser, notice.NoticeLevel, notice.NoticeType)

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

func DelNotice(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var notice model.Notice
	json.Unmarshal([]byte(body), &notice)

	sqlStr := "delete from  notice   where Id=? "

	ret, err := lib.Db.Exec(sqlStr, notice.Id)

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

func GetNoticeById(c *gin.Context) {

	NoticeId := c.Query("NoticeId")

	newnotice := new(model.Notice)
	//
	lib.Db.QueryRow("select Id,Time,NoticeTitle,NoticeContent,SendUser,NoticeLevel,NoticeType from  notice  where Id=? ", NoticeId).
		Scan(&newnotice.Id, &newnotice.Time, &newnotice.NoticeTitle, &newnotice.NoticeContent, &newnotice.SendUser, &newnotice.NoticeLevel, &newnotice.NoticeType)

	if newnotice.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}

	if lib.Strval(newnotice.Id) == NoticeId {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "成功",
			"data": newnotice,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}
}

func NoticeList(c *gin.Context) {

	noticelArr := make([]*model.Notice, 0)
	//
	rows, err := lib.Db.Query("select Id,Time,NoticeTitle,NoticeContent,SendUser,NoticeLevel,NoticeType from  notice order by  Time desc LIMIT 5;")

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
		newnotice := new(model.Notice)
		err := rows.Scan(&newnotice.Id, &newnotice.Time, &newnotice.NoticeTitle, &newnotice.NoticeContent, &newnotice.SendUser, &newnotice.NoticeLevel, &newnotice.NoticeType)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		noticelArr = append(noticelArr, newnotice)

	}
	rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": noticelArr,
	})
}

func NoticeListAll(c *gin.Context) {

	var noticelArr []*model.Notice
	noticelArr = make([]*model.Notice, 0)
	//
	rows, err := lib.Db.Query("select Id,Time,NoticeTitle,NoticeContent,SendUser,NoticeLevel,NoticeType from  notice order by  Time  ")

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
		newnotice := new(model.Notice)
		err := rows.Scan(&newnotice.Id, &newnotice.Time, &newnotice.NoticeTitle, &newnotice.NoticeContent, &newnotice.SendUser, &newnotice.NoticeLevel, &newnotice.NoticeType)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		noticelArr = append(noticelArr, newnotice)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": noticelArr,
	})
}

func Get1(c *gin.Context) {

	c.String(http.StatusOK, "1")
	c.Done()
}
