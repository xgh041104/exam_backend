package class

import (
	model "StudyExamPlatformAPI/Model"
	lib "StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditClass(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var class model.Class
	json.Unmarshal([]byte(body), &class)

	sqlStr := "Update  class set  MajorId =?,ClassName =?,TeacherId=?,SchoolId=? where Id=? "

	ret, err := lib.Db.Exec(sqlStr, class.MajorId, class.ClassName, class.TeacherId, class.SchoolId, class.Id)

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

func AddClass(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var class model.Class
	json.Unmarshal([]byte(body), &class)

	if class.MajorId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	sqlStr := "insert into  class(MajorId,ClassName,TeacherId,SchoolId) VALUES (?,?,?,?) "

	ret, err := lib.Db.Exec(sqlStr, class.MajorId, class.ClassName, class.TeacherId, class.SchoolId)

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

func DelClass(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var class model.Class
	json.Unmarshal([]byte(body), &class)

	count := 0
	lib.Db.QueryRow("select  count(1) from student where  ClassId=? ", class.Id).Scan(&count)

	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该班级绑定的还有学生，不能删除",
			"data": "{}",
		})
		return
	}

	sqlStr := "delete from  class   where Id=? "

	ret, err := lib.Db.Exec(sqlStr, class.Id)

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

func GetClassByMajorId(c *gin.Context) {

	MajorId := c.Query("MajorId")

	classArr := make([]*model.Class, 0)

	rows, err := lib.Db.Query("select  Id,MajorId,ClassName,TeacherId  from  class where MajorId=? ", MajorId)

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
		newcalss := new(model.Class)
		err := rows.Scan(&newcalss.Id, &newcalss.MajorId, &newcalss.ClassName, &newcalss.TeacherId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		classArr = append(classArr, newcalss)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": classArr,
	})
}

func GetClassBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	if SchoolId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	var classviewArr []*model.ClassView

	rows, err := lib.Db.Query("  select a.Id,a.MajorId,a.SchoolId,a.ClassName,a.TeacherId,  COALESCE(b.SchoolName, 0) AS 'SchoolName' , COALESCE(c.MajorName, 0) AS 'MajorName',COALESCE( d.Id, 0) AS 'CollegeId', COALESCE( d.CollegeName, 0) AS 'CollegeName'    from class a   left  join major c on  a.MajorId=c.MajorId   left  join  college d on c.CollegeId=d.Id  left join school b on d.SchoolId=b.Id where a.SchoolId=? ", SchoolId)

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
		newcalssview := new(model.ClassView)
		err := rows.Scan(&newcalssview.Id, &newcalssview.MajorId, &newcalssview.SchoolId, &newcalssview.ClassName, &newcalssview.TeacherId, &newcalssview.SchoolName, &newcalssview.MajorName, &newcalssview.CollegeId, &newcalssview.CollegeName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		classviewArr = append(classviewArr, newcalssview)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": classviewArr,
	})
}

func GetClassView(c *gin.Context) {

	classviewArr := make([]*model.ClassView, 0)
	rows, err := lib.Db.Query("  select a.Id,a.MajorId,a.SchoolId,a.ClassName,a.TeacherId,COALESCE(b.SchoolName, 0) AS 'SchoolName',COALESCE(c.MajorName, 0) AS 'MajorName',COALESCE( d.Id, 0) AS 'CollegeId',COALESCE( d.CollegeName, 0) AS 'CollegeName'  from class a  left  join major c on  a.MajorId=c.MajorId  left  join  college d  on c.CollegeId=d.Id left join school b on a.SchoolId=b.Id   ")

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
		newcalssview := new(model.ClassView)
		err := rows.Scan(&newcalssview.Id, &newcalssview.MajorId, &newcalssview.SchoolId, &newcalssview.ClassName, &newcalssview.TeacherId, &newcalssview.SchoolName, &newcalssview.MajorName, &newcalssview.CollegeId, &newcalssview.CollegeName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		classviewArr = append(classviewArr, newcalssview)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": classviewArr,
	})
}

func GetClassStudentsBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	rows, err := lib.Db.Query("select Id,ClassName from class where SchoolId=? ", SchoolId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}

	var classstudentarr []*model.ClassStudent
	for rows.Next() {
		// 声明
		classstudent := new(model.ClassStudent)
		err := rows.Scan(&classstudent.Id, &classstudent.ClassName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		classstudentarr = append(classstudentarr, classstudent)
	}

	for i := 0; i < len(classstudentarr); i++ {
		classstudent := classstudentarr[i]
		rowstu, err := lib.Db.Query("select Id,TrueName from student where ClassId=? ", classstudent.Id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		classstudent.Students = make([]*model.StudentS, 0)
		for rowstu.Next() {
			studentS := new(model.StudentS)
			err := rowstu.Scan(&studentS.Id, &studentS.TrueName)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			classstudent.Students = append(classstudent.Students, studentS)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": classstudentarr,
	})

}
