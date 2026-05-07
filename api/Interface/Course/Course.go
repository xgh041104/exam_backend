package Course

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	goutils "github.com/typa01/go-utils"
)

func EditCourseFile(c *gin.Context) {

	tx, err := lib.Db.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	files, err := c.MultipartForm() // 获取文件

	if len(files.File["files"]) > 0 {
		file := files.File["files"][0]
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "未收到文件",
				"data": "{}",
			})
			return

		}

		println(file)

		data, _ := c.GetPostForm("data")
		var courseparam model.CourseIdModel
		json.Unmarshal([]byte(data), &courseparam)

		CourseId := courseparam.CourseId

		if CourseId == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "参数错误",
				"data": "{}",
			})
			return
		}
		fileId := int64(0)
		filename := goutils.GUID()
		file_path := "Resources/Img/" + filename + path.Ext(file.Filename) // 设置保存文件的路径，不要忘了后面的文件名

		course := new(model.Course)
		tx.QueryRow("select Id,FileId,SchoolId from  course where  Id=? ", CourseId).Scan(&course.Id, &course.FileId, &course.SchoolId)

		if course.FileId == 0 {
			//新增
			errr := lib.UploadFile(c, file, file_path) // 保存文件
			if errr != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			} // 保存文件

			ret, err := tx.Exec("insert into  fileinfo(FileType,FileName,FileUseTo,SchoolId,FilePath) values(?,?,?,?,?)", path.Ext(file.Filename), filename, "课程图片", course.SchoolId, file_path) // 新增文件表
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			fileId, err = ret.LastInsertId()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			ret, err = tx.Exec("update course  set  FileId=? where Id=?", fileId, CourseId) // 新增文件表
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}

			n, err := ret.RowsAffected()
			if n == 0 {
				println("操作无效")
			}
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			err = tx.Commit()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			} // 保存文件
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "操作成功",
				"data": "{}",
			})

		} else {
			//修改.

			tx.QueryRow("select FileName,FilePath from  fileinfo where  Id=? ", course.FileId).Scan(&filename, &file_path)
			errr := lib.UploadFile(c, file, file_path) // 保存文件
			if errr != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			} // 保存文件
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "操作成功",
				"data": "{}",
			})
		}
	}

}

func EditClassRelation(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var courseclassstudentJson model.CourseClassStudentJson
	json.Unmarshal([]byte(body), &courseclassstudentJson)

	Count := 0
	lib.Db.QueryRow(" SELECT   COUNT(1)  'Count' FROM `course`   where  (NOW()  BETWEEN CourseStartTime AND CourseEndTime  or  NOW()>CourseEndTime) and Status=1  and id=? ", courseclassstudentJson.Id).Scan(&Count)
	if Count == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该课程已开始或已完成，不能编辑",
			"data": "{}",
		})
		return
	}

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	//删除课程的班级
	removeclassstr := ""
	if len(courseclassstudentJson.RemoveClass) > 0 {

		for i := 0; i < len(courseclassstudentJson.RemoveClass); i++ {
			if i == 0 {
				s := fmt.Sprintf("'%d'", courseclassstudentJson.RemoveClass[i])
				removeclassstr += s
			} else {
				s := fmt.Sprintf(",'%d'", courseclassstudentJson.RemoveClass[i])
				removeclassstr += s
			}
		}

		//删除班级里面的学生进度表
		querystr := fmt.Sprintf(" select Id from  student where ClassId in (%s)   ", removeclassstr)
		rows, err := tx.Query(querystr) //把所选班级的学生id 全部查询到
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		//添加班级里面的学生进度表
		strstudent := ""
		for rows.Next() {
			// 声明
			newstudent := new(model.Student)
			err := rows.Scan(&newstudent.Id)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			strstudent += "'" + lib.Strval(newstudent.Id) + "',"
		}

		if strstudent != "" {

			deletestr := fmt.Sprintf(" delete  from studyplan where   StudentId in  (%s)  ", strings.Trim(strstudent, ","))
			ret, err := tx.Exec(deletestr+" and  CourseId=?  ", courseclassstudentJson.Id)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if n == 0 {
				println("数据一样 没有修改")
			}
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
		}
	}

	//删除课程的学生
	if len(courseclassstudentJson.RemoveStudent) > 0 {
		var ss string
		for i := 0; i < len(courseclassstudentJson.RemoveStudent); i++ {
			if i == 0 {
				s := fmt.Sprintf("'%d'", courseclassstudentJson.RemoveStudent[i])
				ss += s
			} else {
				s := fmt.Sprintf(",'%d'", courseclassstudentJson.RemoveStudent[i])
				ss += s
			}
		}

		deletestr := fmt.Sprintf(" delete  from studyplan where StudentId in (%s)   ", ss)

		ret, err := tx.Exec(deletestr+" and CourseId=?", courseclassstudentJson.Id)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if n == 0 {
			println("数据一样 没有修改")
			// tx.Rollback()
			// c.JSON(http.StatusOK, gin.H{
			// 	"code": 0,
			// 	"msg":  "失败",
			// 	"data": "{}",
			// })
			// return
		}
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
	}

	//删除班级课程关系表
	if len(courseclassstudentJson.RemoveClass) > 0 {
		deletestr := fmt.Sprintf("delete from   classcourserelation  where  ClassId  in (%s)   ", removeclassstr)
		ret, err := tx.Exec(deletestr+" and  CourseId=?", courseclassstudentJson.Id)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if n == 0 {
			println("数据一样 没有修改")
			// tx.Rollback()
			// c.JSON(http.StatusOK, gin.H{
			// 	"code": 0,
			// 	"msg":  "失败",
			// 	"data": "{}",
			// })
			// return
		}
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}

	}
	// 查询班级所有学生

	if len(courseclassstudentJson.AddClass) > 0 {
		addclassstr := ""

		for i := 0; i < len(courseclassstudentJson.AddClass); i++ {
			if i == 0 {
				s := fmt.Sprintf("'%d'", courseclassstudentJson.AddClass[i])
				addclassstr += s
			} else {
				s := fmt.Sprintf(",'%d'", courseclassstudentJson.AddClass[i])
				addclassstr += s
			}
		}
		query := fmt.Sprintf("select Id from  student where ClassId in (%s) ", addclassstr)

		rows, err := tx.Query(query) //把所选班级的学生id 全部查询到
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}

		studentarr := make([]*model.Student, 0)
		//添加班级里面的学生进度表
		for rows.Next() {
			// 声明
			newstudent := new(model.Student)
			err := rows.Scan(&newstudent.Id)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			studentarr = append(studentarr, newstudent)
		}

		sqlStr := "insert into  studyplan(StudentId,CourseId,ChapterId,LearningRate,ChapterOrder ) VALUES (?,?,?,?,?) "

		for i := 0; i < len(studentarr); i++ {
			newstudent := studentarr[i]
			ret, err := tx.Exec(sqlStr, newstudent.Id, courseclassstudentJson.Id, 0, 0, 0)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if n == 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
		}

	}

	// 添加数组里面的学生进度表

	if len(courseclassstudentJson.AddStudent) > 0 {
		for i := 0; i < len(courseclassstudentJson.AddStudent); i++ {
			sqlStr := "insert into  studyplan(StudentId,CourseId,ChapterId,LearningRate,ChapterOrder ) VALUES (?,?,?,?,?) "

			ret, err := tx.Exec(sqlStr, courseclassstudentJson.AddStudent[i], courseclassstudentJson.Id, 0, 0, 0)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			if n == 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
		}
	}

	//添加班级课程关系表

	if len(courseclassstudentJson.AddClass) > 0 {
		for i := 0; i < len(courseclassstudentJson.AddClass); i++ {

			sqlStr := "insert into  classcourserelation(CourseId,ClassId ) VALUES (?,?) "

			ret, err := tx.Exec(sqlStr, courseclassstudentJson.Id, courseclassstudentJson.AddClass[i])
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()
			if n == 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
		}
	}
	err = tx.Commit()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
		return
	} else {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}

}

func AddCourse(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var course model.Course
	json.Unmarshal([]byte(body), &course)

	//添加课程
	sqlStr := "insert into  course(CourseName,Digest,SchoolId,CollegeId,MajorId,TeacherId,CourseStartTime,CourseEndTime,Status,CourseCode) VALUES (?,?,?,?,?,?,?,?,?,?) "

	ret, err := lib.Db.Exec(sqlStr, course.CourseName, course.Digest, course.SchoolId, course.CollegeId, course.MajorId, course.TeacherId, course.CourseStartTime, course.CourseEndTime, course.Status, course.CourseCode)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected() //查看受影响行数

	if n == 0 {
		// 数据一样 没有修改
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
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

func EditCourse(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var course model.Course
	json.Unmarshal([]byte(body), &course)

	Count := 0
	lib.Db.QueryRow(" SELECT   COUNT(1)  'Count' FROM `course`   where  (NOW()  BETWEEN CourseStartTime AND CourseEndTime  or  NOW()>CourseEndTime) and Status=1   and id=? ", course.Id).Scan(&Count)
	if Count == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "该课程已开始或已完成，不能编辑",
			"data": "{}",
		})
		return
	}

	//file_path := "Score/" + file.Filename // 设置保存文件的路径，不要忘了后面的文件名

	//添加课程
	sqlStr := "update  course set  CourseName=?,Digest=?,SchoolId=?,CollegeId=?,MajorId=?,TeacherId=?,CourseStartTime=?,CourseEndTime=?,Status=?,CourseCode=? where Id=?"

	ret, err := lib.Db.Exec(sqlStr, course.CourseName, course.Digest, course.SchoolId, course.CollegeId, course.MajorId, course.TeacherId, course.CourseStartTime, course.CourseEndTime, course.Status, course.CourseCode, course.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected() // 获取课程id

	if n == 0 {
		// 数据一样 没有修改
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "数据一样 没有修改",
			"data": "{}",
		})
		return
	}
	if err != nil {
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

func DelCourse(c *gin.Context) { // 删除所有关系 图片 学生进度

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var courseIdModel model.CourseIdModel
	json.Unmarshal([]byte(body), &courseIdModel)

	tx, err := lib.Db.Begin()

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	course := new(model.Course)
	tx.QueryRow("select Id,FileId from  course where  Id=? ", courseIdModel.CourseId).Scan(&course.Id, &course.FileId)

	ret, err := tx.Exec("delete from  course where Id=?", courseIdModel.CourseId)

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected() // 获取课程id

	if n == 0 {
		tx.Rollback()
		// 数据一样 没有修改
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if course.FileId > 0 {
		fileinfo := new(model.FileInfo)

		tx.QueryRow("select Id,FilePath  from  fileinfo where  Id=? ", course.FileId).Scan(&fileinfo.Id, &fileinfo.FilePath)

		err = lib.RemoveFile(fileinfo.FilePath)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "文件删除失败 :" + err.Error(),
				"data": "{}",
			})
			return
		}
		tx.Exec(" delete from fileinfo where id=?", course.FileId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err = ret.RowsAffected() // 获取课程id

		if n == 0 {
			tx.Rollback()
			// 数据一样 没有修改
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
	}

	ret, err = tx.Exec("delete from studyplan where CourseId=?", courseIdModel.CourseId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected() // 获取课程id

	if n < 0 {
		tx.Rollback()
		// 数据一样 没有修改
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	ret, err = tx.Exec("delete from classcourserelation where CourseId=?", courseIdModel.CourseId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected()

	if n < 0 {
		tx.Rollback()
		// 数据一样 没有修改
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	//mark 需要删除所有章节
	ChapterIdarr := make([]int, 0)

	rows, err := tx.Query("select Id from chapter where CourseId=?", courseIdModel.CourseId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	for rows.Next() {
		Id := 0
		err = rows.Scan(&Id)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		ChapterIdarr = append(ChapterIdarr, Id)
	}
	for i := 0; i < len(ChapterIdarr); i++ {
		var htmtcontent string
		tx.QueryRow("select ChapterContent from  chapter   where Id=?", ChapterIdarr[i]).Scan(&htmtcontent)

		flag := lib.DelHtmlResources(htmtcontent, tx)
		if !flag {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "富文本文件错误",
				"data": "{}",
			})
			return
		}
		ret, err = tx.Exec("Delete from  chapter   where Id=?", ChapterIdarr[i])
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除失败",
				"data": "{}",
			})
			return
		}
		n, err = ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除失败",
				"data": "{}",
			})
			return
		}
		if n < 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除失败",
				"data": "{}",
			})
			return
		}

		rows, err := tx.Query("select FileId from  chapterrelation where  ChapterId=? ", ChapterIdarr[i])
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "查询章节关系表出错",
				"data": "{}",
			})
			return
		}
		var fileidarr []string
		fileidstr := ""
		for rows.Next() {
			fileid := ""
			rows.Scan(&fileid)
			fileidarr = append(fileidarr, fileid)
			fileidstr += "'" + fileid + "',"
		}

		if fileidstr != "" {

			delstr := fmt.Sprintf("delete from chapterrelation where FileId in (%s) ", strings.Trim(fileidstr, ","))
			ret, err = tx.Exec(delstr+" and ChapterId=?", ChapterIdarr[i])
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "删除章节关系表出错",
					"data": "{}",
				})
				return
			}

			n, err = ret.RowsAffected()

			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "删除章节关系表出错",
					"data": "{}",
				})
				return
			}
			if n == 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "删除章节关系表出错",
					"data": "{}",
				})
				return
			}
		}
		for i := 0; i < len(fileidarr); i++ {
			file_path := ""
			filename := ""
			tx.QueryRow("select FilePath,FileName  from fileinfo  where Id=?", fileidarr[i]).Scan(&file_path, &filename)

			ret, err = tx.Exec("delete from fileinfo where Id=?", fileidarr[i])
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "删除文件表出错",
					"data": "{}",
				})
				return
			}
			n, err = ret.RowsAffected()

			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "删除文件表出错",
					"data": "{}",
				})
				return
			}
			if n < 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "删除文件表出错",
					"data": "{}",
				})
				return
			}
			videoExtArr := []string{".avi", ".mp4", ".mov", ".wmv", ".flv", ".mkv", ".mpg", ".rmvb"}
			isvideo := lib.In(path.Ext(filename), videoExtArr)
			if isvideo {

				err = lib.DelRemoveAll(file_path+"/", "Resources/Annex/") //删除m3u8文件夹
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("file deleted successfully!")
				}

				err = lib.RemoveFile("Resources/Annex/" + filename) //删除视频的原视频
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("file deleted successfully!")
				}
			} else {
				err = lib.RemoveFile(file_path)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("file deleted successfully!")
				}
			}

		}

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
func GetCourseClassRelationById(c *gin.Context) { // 获取课程关联的学生 和班级

	CourseId := c.Query("CourseId")

	sqlstr := "SELECT  0 'StudentId',ClassId FROM `classcourserelation` where CourseId=?  UNION select  a.StudentId,b.ClassId from studyplan a LEFT JOIN  student b on a.StudentId= b.id where   b.ClassId  not in  (SELECT  ClassId FROM classcourserelation where CourseId=?) and  a.CourseId=? "

	rows, err := lib.Db.Query(sqlstr, CourseId, CourseId, CourseId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	newCourseRelation := new(model.CourseRelation)
	for rows.Next() {
		// 声明
		newCourseRelationdb := new(model.CourseRelationDb)

		err := rows.Scan(&newCourseRelationdb.StudentId, &newCourseRelationdb.ClassId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}

		if newCourseRelationdb.StudentId != 0 {
			var newstudentarr []int
			newstudentarr = append(newstudentarr, newCourseRelationdb.ClassId)
			newstudentarr = append(newstudentarr, newCourseRelationdb.StudentId)
			newCourseRelation.StudentArr = append(newCourseRelation.StudentArr, newstudentarr)
		}
		if newCourseRelationdb.ClassId != 0 && newCourseRelationdb.StudentId == 0 {
			newCourseRelation.ClassArr = append(newCourseRelation.ClassArr, newCourseRelationdb.ClassId)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": newCourseRelation,
	})

}

func GetCourseListByTeacherId(c *gin.Context) { // 带图片查看所有课程 带老师Id

	TeacherId := c.Query("TeacherId")

	strsql := "	select a.Id,a.CourseName,a.Digest,a.SchoolId,a.CollegeId,a.MajorId,a.FileId,a.TeacherId,a.CourseStartTime,a.CourseEndTime,a.`Status`,COALESCE(f.FilePath, ''),COALESCE(b.SchoolName,''),COALESCE(c.CollegeName,''),COALESCE(d.MajorName,''),COALESCE(e.TeacherName,''),CourseCode   from course a 	left join fileinfo f on a.FileId=f.Id  	left join school b on a.SchoolId=b.Id  	  left join college c on a.CollegeId=c.Id   left join major d on a.MajorId=d.MajorId left join teacher e on a.TeacherId=e.TeacherId  "
	var rows *sql.Rows
	var err error
	if TeacherId != "0" {
		strsql += " where a.TeacherId=?"
		rows, err = lib.Db.Query(strsql, TeacherId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
	} else {
		rows, err = lib.Db.Query(strsql)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
	}
	var courseViewArr []*model.CourseView
	for rows.Next() {
		newCourseView := new(model.CourseView)
		err := rows.Scan(&newCourseView.Id, &newCourseView.CourseName, &newCourseView.Digest, &newCourseView.SchoolId, &newCourseView.CollegeId, &newCourseView.MajorId, &newCourseView.FileId, &newCourseView.TeacherId, &newCourseView.CourseStartTime, &newCourseView.CourseEndTime, &newCourseView.Status, &newCourseView.FilePath, &newCourseView.SchoolName, &newCourseView.CollegeName, &newCourseView.MajorName, &newCourseView.TeacherName, &newCourseView.CourseCode)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		courseViewArr = append(courseViewArr, newCourseView)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": courseViewArr,
	})
}
func GetCourseByCourseId(c *gin.Context) { //带图片 查看当前课程内容

	CurseId := c.Query("CourseId")

	strsql := "	select a.Id,a.CourseName,a.Digest,a.SchoolId,a.CollegeId,a.MajorId,a.FileId,a.TeacherId,a.CourseStartTime,a.CourseEndTime,a.`Status`,COALESCE(f.FilePath, ''),COALESCE(b.SchoolName,''),COALESCE(c.CollegeName,''),COALESCE(d.MajorName,'') ,COALESCE(e.TeacherName,''),CourseCode  from course a 	left join fileinfo f on a.FileId=f.Id  	left join school b on a.SchoolId=b.Id  	  left join college c on a.CollegeId=c.Id   left join major d on a.MajorId=d.MajorId left join teacher e on a.TeacherId=e.TeacherId    where a.Id=?"
	newCourseView := new(model.CourseView)
	lib.Db.QueryRow(strsql, CurseId).Scan(&newCourseView.Id, &newCourseView.CourseName, &newCourseView.Digest, &newCourseView.SchoolId, &newCourseView.CollegeId, &newCourseView.MajorId, &newCourseView.FileId, &newCourseView.TeacherId, &newCourseView.CourseStartTime, &newCourseView.CourseEndTime, &newCourseView.Status, &newCourseView.FilePath, &newCourseView.SchoolName, &newCourseView.CollegeName, &newCourseView.MajorName, &newCourseView.TeacherName, &newCourseView.CourseCode)
	if newCourseView.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作错误",
			"data": "{}",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": newCourseView,
		})
	}
}

func GetCurrentStudyCourse(c *gin.Context) {

	StudentId := c.Query("StudentId")

	querysql := "	select  b.CourseName,  " +
		" IFNULL(b.Digest,'') 'Digest', " +
		" a.CourseId, " +
		" b.TeacherId,	IFNULL(d.TeacherName,'')'TeacherName', " +
		"	IFNULL(e.SchoolName,'')  'SchoolName',	IFNULL(f.CollegeName,'')  'CollegeName'  , b. MajorId,	IFNULL(g.MajorName,'') 'MajorName' , " +
		"			(select COUNT(1) from chapter   WHERE CourseId= a.CourseId ) ' ChapterSum',  " +
		"			(select COUNT(1) from studyplan   WHERE CourseId= a.CourseId ) ' StudentSum',  " +
		"	  b.CourseStartTime, " +
		"		b.CourseEndTime, " +
		"a.ChapterOrder  , " +
		"		 a.LearningRate  , " +
		"		 IFNULL(c.FilePath,'') 'FilePath', " +
		" (CASE    " +
		"	WHEN b.CourseStartTime < now() AND b.CourseEndTime >= now() and  a.IsComplete=0  THEN  1   " +
		"	WHEN b.CourseStartTime > now()  and  a.IsComplete=0  THEN   0  " +
		"	WHEN a.IsComplete=1 THEN   2   " +
		"	WHEN b.CourseEndTime < now()  and  a.IsComplete=0 THEN   3  " +
		" ELSE 0 END   ) as 'IsCurrentStudy' " +
		" from studyplan a  " +
		"left join course b  on a.CourseId=b.Id " +
		"LEFT JOIN fileinfo c on b.FileId=c.Id " +
		"LEFT JOIN teacher d on b.TeacherId=d.TeacherId " +
		"LEFT JOIN school e on b.SchoolId=e.Id " +
		"LEFT JOIN college f on b.CollegeId=f.Id " +
		"LEFT JOIN major g  on b.MajorId=g.MajorId " +
		" WHERE a.StudentId= ? and b.Status=1  "

	rows, err := lib.Db.Query(querysql, StudentId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var StudyCourseViewArr []*model.StudyCourse
	for rows.Next() {
		newStudyCourse := new(model.StudyCourse)
		err := rows.Scan(&newStudyCourse.CourseName, &newStudyCourse.Digest, &newStudyCourse.CourseId, &newStudyCourse.TeacherId, &newStudyCourse.TeacherName, &newStudyCourse.SchoolName, &newStudyCourse.CollegeName, &newStudyCourse.MajorId, &newStudyCourse.MajorName, &newStudyCourse.ChapterSum, &newStudyCourse.StudentSum, &newStudyCourse.CourseStartTime, &newStudyCourse.CourseEndTime, &newStudyCourse.ChapterOrder, &newStudyCourse.LearningRate, &newStudyCourse.FilePath, &newStudyCourse.IsCurrentStudy)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		StudyCourseViewArr = append(StudyCourseViewArr, newStudyCourse)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": StudyCourseViewArr,
	})
}

func CourseDetails(c *gin.Context) { //课程详情  章节列表（显示自己学到当前课程的某个章节的进度）

	StudentId := c.Query("StudentId")
	CourseId := c.Query("CourseId")
	querysql := "	select  b.CourseName,  " +
		" IFNULL(b.Digest,'') 'Digest', " +
		" a.CourseId, " +
		" b.TeacherId,	IFNULL(d.TeacherName,'')'TeacherName', " +
		"	IFNULL(e.SchoolName,'')  'SchoolName',	IFNULL(f.CollegeName,'')  'CollegeName'  , b. MajorId,	IFNULL(g.MajorName,'') 'MajorName' , " +
		"			(select COUNT(1) from chapter WHERE CourseId= a.CourseId ) ' ChapterSum',  " +
		"			(select COUNT(1) from studyplan WHERE CourseId= a.CourseId ) ' StudentSum',  " +
		"	  b.CourseStartTime, " +
		"		b.CourseEndTime, " +
		"a.ChapterOrder  , " +
		"		 a.LearningRate  , " +
		"		 IFNULL(c.FilePath,'') 'FilePath', " +
		" (CASE    " +
		"	WHEN b.CourseStartTime < now() AND b.CourseEndTime > now() THEN  1   " +
		"	WHEN b.CourseStartTime > now()  THEN   0  " +
		"	WHEN b.CourseEndTime < now()  and a.IsComplete=1 THEN   2   " +
		"	WHEN b.CourseEndTime < now()  and  a.IsComplete=0 THEN   3  " +
		" ELSE 0 END   ) as 'IsCurrentStudy' " +
		" from studyplan a  " +
		"left join course b  on a.CourseId=b.Id " +
		"LEFT JOIN fileinfo c on b.FileId=c.Id " +
		"LEFT JOIN teacher d on b.TeacherId=d.TeacherId " +
		"LEFT JOIN school e on b.SchoolId=e.Id " +
		"LEFT JOIN college f on b.CollegeId=f.Id " +
		"LEFT JOIN major g  on b.MajorId=g.MajorId " +
		" WHERE a.StudentId= ? and b.Status=1  and a.CourseId=? "
	newStudyCourse := new(model.StudyCourseChapter)
	lib.Db.QueryRow(querysql, StudentId, CourseId).Scan(&newStudyCourse.CourseName, &newStudyCourse.Digest, &newStudyCourse.CourseId, &newStudyCourse.TeacherId, &newStudyCourse.TeacherName, &newStudyCourse.SchoolName, &newStudyCourse.CollegeName, &newStudyCourse.MajorId, &newStudyCourse.MajorName, &newStudyCourse.ChapterSum, &newStudyCourse.StudentSum, &newStudyCourse.CourseStartTime, &newStudyCourse.CourseEndTime, &newStudyCourse.ChapterOrder, &newStudyCourse.LearningRate, &newStudyCourse.FilePath, &newStudyCourse.IsCurrentStudy)
	if newStudyCourse.CourseId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	// ChapterList
	// append(newStudyCourse.ChapterList)

	rows, err := lib.Db.Query("select Id,ChapterName,ChapterType,ChapterOrder from chapter where CourseId=?", newStudyCourse.CourseId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	newStudyCourse.ChapterList = make([]*model.ChapterList, 0)

	for rows.Next() {
		chapterList := new(model.ChapterList)
		err := rows.Scan(&chapterList.Id, &chapterList.ChapterName, &chapterList.ChapterType, &chapterList.ChapterOrder)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		newStudyCourse.ChapterList = append(newStudyCourse.ChapterList, chapterList)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": newStudyCourse,
	})
}

func GetCouresStudyPlanByCourseId(c *gin.Context) {

	CourseId := c.Query("CourseId")

	sql := " select a.Id,a.StudentId,c.TrueName,a.CourseId,b.CourseName,a.ChapterId,a.LearningRate,a.ChapterOrder, " +
		"(select COUNT(1) from chapter   WHERE CourseId= b.Id ) ' ChapterSum',  a.IsComplete " +
		" from  studyplan a " +
		" left join course b on a.CourseId=b.Id " +
		" left JOIN student c on c.Id=a.StudentId " +
		" where a.CourseId=? "

	rows, err := lib.Db.Query(sql, CourseId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	courseStudyPlanViewarr := make([]*model.CourseStudyPlanView, 0)

	for rows.Next() {
		model := new(model.CourseStudyPlanView)
		err := rows.Scan(&model.Id, &model.StudentId, &model.TrueName, &model.CourseId, &model.CourseName, &model.ChapterId, &model.LearningRate, &model.ChapterOrder, &model.ChapterNum, &model.IsComplete)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		courseStudyPlanViewarr = append(courseStudyPlanViewarr, model)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": courseStudyPlanViewarr,
	})

}

func CourseCancel(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var courseIdModel model.CourseIdModel
	json.Unmarshal([]byte(body), &courseIdModel)

	tx, err := lib.Db.Begin()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "事务打开失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	//新增课程
	// 先查询之前的课程
	coursemodel := new(model.Course)
	tx.QueryRow("SELECT Id,CourseName,Digest,SchoolId,CollegeId,MajorId,FileId,TeacherId,CourseStartTime,CourseEndTime,`Status`,CourseCode FROM course where Id=?", courseIdModel.CourseId).
		Scan(&coursemodel.Id, &coursemodel.CourseName, &coursemodel.Digest, &coursemodel.SchoolId, &coursemodel.CollegeId, &coursemodel.MajorId, &coursemodel.FileId, &coursemodel.TeacherId, &coursemodel.CourseStartTime, &coursemodel.CourseEndTime, &coursemodel.Status, &coursemodel.CourseCode)

	if coursemodel.Id == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "事务打开失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	//添加课程
	sqlStr := "insert into  course(CourseName,Digest,SchoolId,CollegeId,MajorId,FileId,TeacherId,CourseStartTime,CourseEndTime,Status,CourseCode) VALUES (?,?,?,?,?,?,?,?,?,?,?) "

	ret, err := tx.Exec(sqlStr, coursemodel.CourseName, coursemodel.Digest, coursemodel.SchoolId, coursemodel.CollegeId, coursemodel.MajorId, coursemodel.FileId, coursemodel.TeacherId, coursemodel.CourseStartTime, coursemodel.CourseEndTime, 0, coursemodel.CourseCode)

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	newCourseId, err := ret.LastInsertId() //查看受影响行数

	if newCourseId == 0 {
		// 数据一样 没有修改
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	//先查询之前的关系表
	oldClassCourseRelationarr := make([]*model.ClassCourseRelation, 0)
	rows, err := tx.Query("select CourseId,ClassId  from  classcourserelation where  CourseId=? ", coursemodel.Id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	for rows.Next() {
		temp := new(model.ClassCourseRelation)
		err = rows.Scan(&temp.CourseId, &temp.ClassId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		oldClassCourseRelationarr = append(oldClassCourseRelationarr, temp)
	}
	//新增班级关系表
	for i := 0; i < len(oldClassCourseRelationarr); i++ {
		sqlStrinsertclass := "insert into  classcourserelation(CourseId,ClassId ) VALUES (?,?) "

		ret, err = tx.Exec(sqlStrinsertclass, newCourseId, oldClassCourseRelationarr[i].ClassId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if n == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
	}

	//查询之前的学生进度
	oldStudyPlanarr := make([]*model.StudyPlan, 0)
	rows, err = tx.Query("select Id,StudentId,CourseId from  studyplan where  CourseId=? ", coursemodel.Id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	for rows.Next() {
		temp := new(model.StudyPlan)
		err = rows.Scan(&temp.Id, &temp.StudentId, &temp.CourseId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		oldStudyPlanarr = append(oldStudyPlanarr, temp)
	}
	//新增学生学习进度
	sqlStraddStudyPlan := "insert into  studyplan(StudentId,CourseId,ChapterId,LearningRate,ChapterOrder ) VALUES (?,?,?,?,?) "

	for i := 0; i < len(oldStudyPlanarr); i++ {
		ret, err = tx.Exec(sqlStraddStudyPlan, oldStudyPlanarr[i].StudentId, newCourseId, 0, 0, 0)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if n == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
	}
	//修改之前的章结 改为当前的课程id 	//章节关系表不用动
	ret, err = tx.Exec("update chapter set CourseId=? where CourseId=?", newCourseId, coursemodel.Id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败",
			"data": "{}",
		})
		return
	}

	//删除课程的所有数据 除章节外 因为章节在只变更了课程id所以不用删
	course := new(model.Course)
	tx.QueryRow("select Id,FileId from  course where  Id=? ", courseIdModel.CourseId).Scan(&course.Id, &course.FileId)

	ret, err = tx.Exec("delete from  course where Id=?", courseIdModel.CourseId)

	//错误处理
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	n, err = ret.RowsAffected() // 获取课程id

	if n == 0 {
		tx.Rollback()
		// 数据一样 没有修改
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}

	//不用删除原课程照片
	// if course.FileId > 0 {
	// 	fileinfo := new(model.FileInfo)
	// 	tx.QueryRow("select Id,FilePath  from  fileinfo where  Id=? ", course.FileId).Scan(&fileinfo.Id, &fileinfo.FilePath)
	// 	err = lib.RemoveFile(fileinfo.FilePath)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"code": 0,
	// 			"msg":  "文件删除失败 :" + err.Error(),
	// 			"data": "{}",
	// 		})
	// 		return
	// 	}
	// 	tx.Exec(" delete from fileinfo where id=?", course.FileId)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"code": 0,
	// 			"msg":  "操作失败" + err.Error(),
	// 			"data": "{}",
	// 		})
	// 		return
	// 	}
	// 	n, err = ret.RowsAffected() // 获取课程id
	// 	if n == 0 {
	// 		tx.Rollback()
	// 		// 数据一样 没有修改
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"code": 0,
	// 			"msg":  "操作失败",
	// 			"data": "{}",
	// 		})
	// 		return
	// 	}
	// 	if err != nil {
	// 		tx.Rollback()
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"code": 0,
	// 			"msg":  "操作失败",
	// 			"data": "{}",
	// 		})
	// 		return
	// 	}
	// }

	ret, err = tx.Exec("delete from studyplan where CourseId=?", courseIdModel.CourseId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected() // 获取课程id

	if n < 0 {
		tx.Rollback()
		// 数据一样 没有修改
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	ret, err = tx.Exec("delete from classcourserelation where CourseId=?", courseIdModel.CourseId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	n, err = ret.RowsAffected()

	if n < 0 {
		tx.Rollback()
		// 数据一样 没有修改
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
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
