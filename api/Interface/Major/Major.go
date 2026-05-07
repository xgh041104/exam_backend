package major

import (
	model "StudyExamPlatformAPI/Model"
	lib "StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditMajor(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var major model.Major
	json.Unmarshal([]byte(body), &major)

	sqlStr := "Update  major set  SchoolId =?,CollegeId =?,MajorName=? where MajorId=? "

	ret, err := lib.Db.Exec(sqlStr, major.SchoolId, major.CollegeId, major.MajorName, major.MajorId)

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

func AddMajor(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var major model.Major
	json.Unmarshal([]byte(body), &major)

	if major.CollegeId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	sqlStr := "insert into  major(SchoolId,CollegeId,MajorName,TeacherId) VALUES (?,?,?,?) "

	ret, err := lib.Db.Exec(sqlStr, major.SchoolId, major.CollegeId, major.MajorName, major.TeacherId)

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

func DelMajor(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var major model.Major
	json.Unmarshal([]byte(body), &major)

	count := 0
	lib.Db.QueryRow("select  count(1) from class where  MajorId=? ", major.MajorId).Scan(&count)

	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该专业绑定的还有班级，请先删除班级！",
			"data": "{}",
		})
		return
	}

	count = 0
	lib.Db.QueryRow("select  count(1) from course where  MajorId=? ", major.MajorId).Scan(&count)

	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该专业绑定的还有课程，请先删除课程！",
			"data": "{}",
		})
		return
	}

	sqlStr := "delete from  major   where MajorId=? "

	ret, err := lib.Db.Exec(sqlStr, major.MajorId)

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

func GetMajorByCollegeId(c *gin.Context) {

	CollegeId := c.Query("CollegeId")

	majorArr := make([]*model.Major, 0)

	rows, err := lib.Db.Query("select MajorId,SchoolId,CollegeId,MajorName,TeacherId  from  major where CollegeId=? ", CollegeId)

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
		newmajor := new(model.Major)
		err := rows.Scan(&newmajor.MajorId, &newmajor.SchoolId, &newmajor.CollegeId, &newmajor.MajorName, &newmajor.TeacherId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		majorArr = append(majorArr, newmajor)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": majorArr,
	})
}

func GetMajorBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	majorviewArr := make([]*model.MajorView, 0)

	rows, err := lib.Db.Query(" SELECT  a.MajorId,a.SchoolId,a.CollegeId,a.MajorName,  COALESCE(b.CollegeName, 0) AS 'CollegeName' ,a.TeacherId FROM major  a   left join college b on a.CollegeId=b.Id where a.SchoolId=? ", SchoolId)

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
		newmajorview := new(model.MajorView)
		err := rows.Scan(&newmajorview.MajorId, &newmajorview.SchoolId, &newmajorview.CollegeId, &newmajorview.MajorName, &newmajorview.CollegeName, &newmajorview.TeacherId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		majorviewArr = append(majorviewArr, newmajorview)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": majorviewArr,
	})
}

func GetMajorListView(c *gin.Context) {

	majorviewArr := make([]*model.MajorView, 0)

	rows, err := lib.Db.Query(" SELECT  a.MajorId,a.SchoolId,a.CollegeId,a.MajorName,COALESCE(b.CollegeName, 0) AS 'CollegeName' ,a.TeacherId FROM major  a   left join college b on a.CollegeId=b.Id  ")

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
		newmajorview := new(model.MajorView)
		err := rows.Scan(&newmajorview.MajorId, &newmajorview.SchoolId, &newmajorview.CollegeId, &newmajorview.MajorName, &newmajorview.CollegeName, &newmajorview.TeacherId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		majorviewArr = append(majorviewArr, newmajorview)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": majorviewArr,
	})
}
