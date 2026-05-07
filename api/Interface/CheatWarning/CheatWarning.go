package cheatwarning

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	goutils "github.com/typa01/go-utils"
)

func AddCheatWarning(c *gin.Context) {

	tx, err := lib.Db.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "系统错误",
			"data": "{}",
		})
		return
	}
	files, err := c.MultipartForm() // 获取文件
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未收到文件",
			"data": "{}",
		})
		return
	}
	if len(files.File["files"]) > 0 {
		file := files.File["files"][0]

		data, _ := c.GetPostForm("data")
		var cheatwarning model.CheatWarning
		json.Unmarshal([]byte(data), &cheatwarning)

		count := 0
		err = tx.QueryRow("select   COUNT(1) from cheatwarning where studentId=?", cheatwarning.StudentId).Scan(&count)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}

		filename := goutils.GUID()
		file_path := "Resources/CheatWarning/" + strconv.FormatInt(cheatwarning.StudentId, 10) + "/" + filename + path.Ext(file.Filename) // 设置保存文件的路径，不要忘了后面的文件名
		err = lib.UploadFile(c, file, file_path)                                                                                          // 保存文件
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}

		if count == 5 { //先删除再新增
			studentImgPath := ""
			err := tx.QueryRow("select  studentImgPath from cheatwarning  where  studentId=? order by createTime asc limit 1 ", cheatwarning.StudentId).Scan(&studentImgPath)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			os.Remove(studentImgPath)
			_, err = tx.Exec(" DELETE FROM cheatwarning where  studentId=? ORDER BY createTime ASC LIMIT 1 ", cheatwarning.StudentId)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			} //  新增人脸监视 如果有就替换
			ret, err := tx.Exec("insert into  cheatwarning(studentId,studentName,studentImgPath,studentImgNum,createTime) values(?,? ,?,?,?)", cheatwarning.StudentId, cheatwarning.StudentName, file_path, 0, time.Now().Unix()) //  新增人脸监视 如果有就替换
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}

			_, err = ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
		} else if count < 5 { // 直接新增
			ret, err := tx.Exec("insert into  cheatwarning(studentId,studentName,studentImgPath,studentImgNum,createTime) values(?,?,?,?,?)", cheatwarning.StudentId, cheatwarning.StudentName, file_path, 0, time.Now().Unix()) //  新增人脸监视 如果有就替换
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
			_, err = ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
		}

		tx.Commit()
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败:未收到图片",
			"data": "{}",
		})
	}
}

func DelCheatWarning(c *gin.Context) {
	tx, err := lib.Db.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "系统错误",
			"data": "{}",
		})
		return
	}
	studentId := c.Query("studentId")
	query := "select  studentImgPath from cheatwarning  where  studentId=? "
	rows, err := tx.Query(query, studentId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "系统错误",
			"data": "{}",
		})
		return
	}
	patharr := make([]string, 0)
	for rows.Next() {
		pathstr := ""
		rows.Scan(&pathstr)
		patharr = append(patharr, pathstr)
	}
	for _, v := range patharr {
		os.Remove(v)
	}

	_, err = tx.Exec("delete from cheatwarning where  studentId=? ", studentId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "系统错误",
			"data": "{}",
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "系统错误",
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

func QueryCheatWarning(c *gin.Context) {

	// studentId := c.Query("studentId")

	rows, err := lib.Db.Query("select * from cheatwarning  ")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "系统错误",
			"data": "{}",
		})
		return
	}
	resp := make([]*model.CheatWarning, 0)
	for rows.Next() {
		temp := new(model.CheatWarning)
		rows.Scan(&temp.CheatWarningId, &temp.StudentId, &temp.StudentName, &temp.StudentImgPath, &temp.StudentImgNum, &temp.CreateTime)
		resp = append(resp, temp)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": resp,
	})
}

func AddFaceMonitor(c *gin.Context) {
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

		data, _ := c.GetPostForm("data")
		var facemonitor model.FaceMonitor
		json.Unmarshal([]byte(data), &facemonitor)
		file_path := "Resources/FaceMonitor/" + strconv.FormatInt(facemonitor.StudentId, 10) + "/" + "1" + path.Ext(file.Filename) // 设置保存文件的路径，不要忘了后面的文件名
		err = lib.UploadFile(c, file, file_path)                                                                                   // 保存文件
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		count := 0
		err = lib.Db.QueryRow("select count(1) from facemonitor where studentId=?", facemonitor.StudentId).Scan(&count)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		if count == 0 {

			ret, err := lib.Db.Exec("insert into  facemonitor(studentId,imgPath) values(?,? )", facemonitor.StudentId, file_path) //  新增人脸监视 如果有就替换
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}

			_, err = ret.RowsAffected()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + err.Error(),
					"data": "{}",
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败:未收到图片",
			"data": "{}",
		})
	}
}
func QueryFaceMonitor(c *gin.Context) {

	imgpath := ""
	studentId := c.Query("studentId")

	sqlStr := "select imgPath from  facemonitor  where studentId=? limit 1 "

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
