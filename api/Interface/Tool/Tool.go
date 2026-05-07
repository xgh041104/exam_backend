package tool

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
)

func GetTool(c *gin.Context) {

	sqlStr := "SELECT Id,StaticResourcesType,FaceVerify from tool   where Id=1 "

	toolconfig := new(model.ToolConfig)
	lib.Db.QueryRow(sqlStr).Scan(&toolconfig.Id, &toolconfig.StaticResourcesType, &toolconfig.FaceVerify)

	if toolconfig.Id == 0 {
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
		"data": toolconfig,
	})

}

func EditTool(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}
	log.Printf("%s", body)

	var toolconfig model.ToolConfig
	json.Unmarshal([]byte(body), &toolconfig)

	sqlStr := "Update  tool set  StaticResourcesType=?,FaceVerify=?  where Id=1 "

	ret, err := lib.Db.Exec(sqlStr, toolconfig.StaticResourcesType, toolconfig.FaceVerify)

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

func GetFaceVerify(c *gin.Context) {

	faceVerify := new(model.FaceVerifyModel)
	lib.Db.QueryRow("select FaceVerify,SeparateFaceVerify from   tool where id=1  ").Scan(&faceVerify.FaceVerify, &faceVerify.SeparateFaceVerify)

	if faceVerify.FaceVerify < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "配置加载失败",
			"data": "{}",
		})
		return
	}
	faceVerify.MidiFlag = 0

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": faceVerify,
	})

}

func GetTime(c *gin.Context) {
	timestr := time.Now().Format(lib.TimeLayoutStr)
	c.String(http.StatusOK, timestr)
}

func GetDiskState(c *gin.Context) {
	dir, _ := os.Getwd()

	println(dir)
	d, _ := disk.Usage(dir)
	fmt.Println(d.Total, d.Used, d.UsedPercent)

	totalGB := float64(d.Total) / (1024 * 1024 * 1024)
	freeGB := float64(d.Free) / (1024 * 1024 * 1024)
	bfbdisk := (freeGB / totalGB) * 100
	formattedNum := fmt.Sprintf("%.2f", bfbdisk)

	c.String(http.StatusOK, formattedNum)

}
func GetHttpUrl(c *gin.Context) {
	con := lib.LoadConfig()

	toolConfig := new(lib.ToolConfig)
	lib.Db.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)

	if toolConfig.StaticResourcesType == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "",
		})
	}
	urlEntity := new(UrlEntity)
	if toolConfig.StaticResourcesType == 1 {
		urlEntity.StaticResourcesType = toolConfig.StaticResourcesType
		urlEntity.OSShttp = ""
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": urlEntity,
		})
	} else {
		urlEntity.StaticResourcesType = toolConfig.StaticResourcesType
		urlEntity.OSShttp = con.OSShttp
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": urlEntity,
		})
	}
}

type UrlEntity struct {
	OSShttp             string `json:"OSShttp" `
	StaticResourcesType int    `json:"StaticResourcesType" `
}
