package Stand

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditStand(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var stand model.Stand
	json.Unmarshal([]byte(body), &stand)

	sqlStr := "Update  stand set  StandName  =?,SchoolId   =? where Id=? "

	ret, err := lib.Db.Exec(sqlStr, stand.StandName, stand.SchoolId, stand.Id)

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

func AddStand(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var stand model.Stand
	json.Unmarshal([]byte(body), &stand)

	sqlStr := "insert into  stand(StandName,SchoolId,TeacherId) VALUES (?,?,?) "

	ret, err := lib.Db.Exec(sqlStr, stand.StandName, stand.SchoolId, stand.TeacherId)

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

func DelStand(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var stand model.Stand
	json.Unmarshal([]byte(body), &stand)

	sqlStr := "delete from  stand   where Id=? "

	ret, err := lib.Db.Exec(sqlStr, stand.Id)

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
func GetStandListBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	standArr := make([]*model.Stand, 0)

	rows, err := lib.Db.Query("select Id,StandName,SchoolId,TeacherId  from  stand where SchoolId=? ", SchoolId)

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
		newstand := new(model.Stand)
		err := rows.Scan(&newstand.Id, &newstand.StandName, &newstand.SchoolId, &newstand.TeacherId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		standArr = append(standArr, newstand)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": standArr,
	})
}
