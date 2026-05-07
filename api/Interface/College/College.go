package college

import (
	model "StudyExamPlatformAPI/Model"
	lib "StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditCollege(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var college model.College
	json.Unmarshal([]byte(body), &college)

	sqlStr := "Update  college set  CollegeName =?,SchoolId =? where Id=? "

	ret, err := lib.Db.Exec(sqlStr, college.CollegeName, college.SchoolId, college.Id)

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

func AddCollege(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var college model.College
	json.Unmarshal([]byte(body), &college)

	if college.SchoolId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}
	sqlStr := "insert into  college(CollegeName,SchoolId,TeacherId) VALUES (?,?,?) "

	ret, err := lib.Db.Exec(sqlStr, college.CollegeName, college.SchoolId, college.TeacherId)

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
func DelCollege(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var college model.College
	json.Unmarshal([]byte(body), &college)

	count := 0
	lib.Db.QueryRow("select  count(1) from major where  CollegeId=? ", college.Id).Scan(&count)

	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该学院绑定的还有专业，请先删除专业！",
			"data": "{}",
		})
		return
	}

	count = 0
	lib.Db.QueryRow("select  count(1) from course where  CollegeId=? ", college.Id).Scan(&count)

	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该学院绑定的还有课程，请先删除课程！",
			"data": "{}",
		})
		return
	}

	sqlStr := "delete from  college   where Id=? "

	ret, err := lib.Db.Exec(sqlStr, college.Id)

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

func GetCollegeBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	collegeViewArr := make([]*model.CollegeView, 0)

	rows, err := lib.Db.Query("select a.Id,a.CollegeName,a.SchoolId,a.TeacherId,b.SchoolName  from  college a left join school b on a.SchoolId=b.Id where a.SchoolId=? ", SchoolId)
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
		newcollegeView := new(model.CollegeView)
		err := rows.Scan(&newcollegeView.Id, &newcollegeView.CollegeName, &newcollegeView.SchoolId, &newcollegeView.TeacherId, &newcollegeView.SchoolName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		collegeViewArr = append(collegeViewArr, newcollegeView)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": collegeViewArr,
	})
}

func GetCollegeListAll(c *gin.Context) {

	var collegeViewArr []*model.CollegeView

	rows, err := lib.Db.Query("select a.Id,a.CollegeName,a.SchoolId,a.TeacherId,b.SchoolName  from  college a left join school b on a.SchoolId=b.Id ")

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
		newcollegeView := new(model.CollegeView)
		err := rows.Scan(&newcollegeView.Id, &newcollegeView.CollegeName, &newcollegeView.SchoolId, &newcollegeView.TeacherId, &newcollegeView.SchoolName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		collegeViewArr = append(collegeViewArr, newcollegeView)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": collegeViewArr,
	})
}

func GetCollegeByCollegeId(c *gin.Context) {

	CollegeId := c.Query("CollegeId")

	var newcollege model.College

	lib.Db.QueryRow("select Id,CollegeName,SchoolId,TeacherId  from  college where Id=? ", CollegeId).
		Scan(&newcollege.Id, &newcollege.CollegeName, &newcollege.SchoolId, &newcollege.TeacherId)

	if newcollege.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": newcollege,
	})

}
