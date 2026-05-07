package school

import (
	model "StudyExamPlatformAPI/Model"
	lib "StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SchoolList(c *gin.Context) {

	var schoolArr []*model.School

	//
	rows, err := lib.Db.Query("select Id,SchoolName,SchoolAddress from  school")

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
		newschool := new(model.School)
		err := rows.Scan(&newschool.Id, &newschool.SchoolName, &newschool.SchoolAddress)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		schoolArr = append(schoolArr, newschool)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": schoolArr,
	})
}

func EditSchool(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var school model.School
	json.Unmarshal([]byte(body), &school)

	Id := 0
	lib.Db.QueryRow("select  Id from school where  SchoolName=?  and  Id!=?", school.SchoolName, school.Id).Scan(&Id)

	if Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "已存在该学校，学校名称重复",
			"data": "{}",
		})
		return
	}

	sqlStr := "Update  school set  SchoolName =?,SchoolAddress =? where Id=? "

	ret, err := lib.Db.Exec(sqlStr, school.SchoolName, school.SchoolAddress, school.Id)

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

func AddSchool(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var school model.School
	json.Unmarshal([]byte(body), &school)

	Id := 0
	lib.Db.QueryRow("select  Id from school where  SchoolName=? ", school.SchoolName).Scan(&Id)

	if Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "已存在该学校",
			"data": "{}",
		})
		return
	}

	sqlStr := "insert into  school(SchoolName,SchoolAddress) VALUES (?,?) "

	ret, err := lib.Db.Exec(sqlStr, school.SchoolName, school.SchoolAddress)

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
func DelSchool(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var school model.School
	json.Unmarshal([]byte(body), &school)

	count := 0
	lib.Db.QueryRow("select  count(1) from college where  SchoolId=? ", school.Id).Scan(&count)

	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该学校绑定的还有学院，请先删除学院！",
			"data": "{}",
		})
		return
	}

	count = 0
	lib.Db.QueryRow("select  count(1) from course where  SchoolId=? ", school.Id).Scan(&count)

	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该学校绑定的还有课程，请先删除课程！",
			"data": "{}",
		})
		return
	}

	sqlStr := "delete from  school   where Id=? "

	ret, err := lib.Db.Exec(sqlStr, school.Id)

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
