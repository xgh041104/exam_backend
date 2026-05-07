package question

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	goutils "github.com/typa01/go-utils"
)

func AddQuestion(c *gin.Context) {

	data, _ := c.GetPostForm("data")
	var question model.Question
	json.Unmarshal([]byte(data), &question)

	if question.QuestionPoolId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误  QuestionName:" + question.QuestionName,
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

	var fileId int64
	files, err := c.MultipartForm()                                 // 获取文件
	if len(files.File["files"]) > 0 && question.QuestionType == 5 { //实操题
		file := files.File["files"]
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "未收到文件",
				"data": "{}",
			})
			return

		}

		i := 0
		filename := goutils.GUID()
		file_path := "Resources/Zip/" + filename + path.Ext(file[i].Filename)
		FileUseTo := "题目文件zip"

		zipExtArr := []string{".7z", ".zip", ".rar"}
		iszip := lib.In(path.Ext(file[i].Filename), zipExtArr)

		xmlExtArr := []string{".xml", ".musicxml"}
		isxml := lib.In(path.Ext(file[i].Filename), xmlExtArr)
		if iszip {

			toolConfig := new(lib.ToolConfig)
			tx.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)
			if toolConfig.StaticResourcesType == 1 {
				errr := lib.UploadFile(c, file[i], file_path) // 保存文件
				if errr != nil {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "操作失败" + err.Error(),
						"data": "{}",
					})
					return
				} // 保存文件
			} else if toolConfig.StaticResourcesType != 1 {
				c.SaveUploadedFile(file[i], file_path)
			}
			err := lib.UnCompressPkg(file_path, "Resources/Zip/"+filename)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "题目压缩包解压失败",
					"data": "{}",
				})
				return
			}
			if toolConfig.StaticResourcesType != 1 {
				os.Remove(file_path)
				cmd := exec.Command("ossutil64", "cp", "-r", "E:/work/StudyExamPlatform/study_exam/api/"+"Resources/Zip/"+filename+"/", "oss://studyexamplatform"+"/"+"Resources/Zip/"+filename+"/")
				err = cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
				if strings.Contains("Resources/Zip/"+filename+"/", "Resources/Zip/") && path.Dir("Resources/Zip/"+filename+"/") != "" && path.Dir("Resources/Zip/"+filename+"/") != "/" {
					err = os.RemoveAll("Resources/Zip/" + filename + "/")
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("file deleted successfully!")
					}
				}
			}
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增题目压缩包解压失败",
					"data": "{}",
				})
				return
			}
		} else if isxml {
			file_path = "Resources/Xml/" + filename + path.Ext(file[i].Filename)
			FileUseTo = "乐谱题目"
			toolConfig := new(lib.ToolConfig)
			tx.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)
			if toolConfig.StaticResourcesType == 1 {
				errr := lib.UploadFile(c, file[i], file_path) // 保存文件
				if errr != nil {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "操作失败" + err.Error(),
						"data": "{}",
					})
					return
				} // 保存文件
			} else if toolConfig.StaticResourcesType != 1 {
				c.SaveUploadedFile(file[i], file_path)
			}

		} else {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "文件格式不对",
				"data": "{}",
			})
			return
		}

		filenameall := path.Base(file[i].Filename)
		filesuffix := path.Ext(filenameall)

		//orginfilename := filenameall[0 : len(filenameall)-len(filesuffix)]
		ret, err := tx.Exec("insert into  fileinfo(FileType,FileName,FileUseTo,SchoolId,FilePath)  values(?,?,?,?,?)", filesuffix, filenameall, FileUseTo, 0, file_path)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增附件失败",
				"data": "{}",
			})
			return
		}
		fileId, err = ret.LastInsertId()

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增题目失败",
				"data": "{}",
			})
			return
		}
		if fileId == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增题目失败",
				"data": "{}",
			})
			return
		}

		sqlStr := "insert into  question(SchoolId,QuestionPoolId,QuestionName,QuestionType,QuestionContent,Digree,MajorID,CollegeId,CourseId,Answer,TeacherId,QuestionCategory) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) "

		ret, err = tx.Exec(sqlStr, question.SchoolId, question.QuestionPoolId, question.QuestionName, question.QuestionType, fileId, question.Digree, question.MajorID, question.CollegeId, question.CourseId, question.Answer, question.TeacherId, question.QuestionCategory)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		//break
		//}

	} else if question.QuestionType != 5 && len(files.File["files"]) > 0 { //其他题  包括附件
		sqlStr := "insert into  question(SchoolId,QuestionPoolId,QuestionName,QuestionType,QuestionContent,Digree,MajorID,CollegeId,CourseId,Answer,TeacherId,QuestionCategory) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) "

		ret, err := tx.Exec(sqlStr, question.SchoolId, question.QuestionPoolId, question.QuestionName, question.QuestionType, question.QuestionContent, question.Digree, question.MajorID, question.CollegeId, question.CourseId, question.Answer, question.TeacherId, question.QuestionCategory)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		questionId, err := ret.LastInsertId() // 操作影响的行数
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		if questionId <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		err = lib.UploadAnnex(c, files, tx, "questionrelation", questionId, "QuestionId", "题目附件")
		if err != nil {
			return
		}
	} else if question.QuestionType != 5 && len(files.File["files"]) == 0 { // 其他题 没有附件
		sqlStr := "insert into  question(SchoolId,QuestionPoolId,QuestionName,QuestionType,QuestionContent,Digree,MajorID,CollegeId,CourseId,Answer,TeacherId,QuestionCategory) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) "

		ret, err := tx.Exec(sqlStr, question.SchoolId, question.QuestionPoolId, question.QuestionName, question.QuestionType, question.QuestionContent, question.Digree, question.MajorID, question.CollegeId, question.CourseId, question.Answer, question.TeacherId, question.QuestionCategory)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "提交事务失败",
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
func EditQuestion(c *gin.Context) {

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	data, _ := c.GetPostForm("data")
	var question model.QuestionEdit
	json.Unmarshal([]byte(data), &question)

	if question.QuestionId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	//删除附件
	if len(question.RemoveFile) > 0 {
		for i := 0; i < len(question.RemoveFile); i++ {
			if question.RemoveFile[i] == 0 {
				continue
			}
			filename := ""
			file_path := ""
			tx.QueryRow("select FileName,FilePath from  fileinfo where  Id=? ", question.RemoveFile[i]).Scan(&filename, &file_path)

			ret, err := tx.Exec("delete from questionrelation where FileId=? and QuestionId=?", question.RemoveFile[i], question.QuestionId)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "删除章节关系表出错",
					"data": "{}",
				})
				return
			}
			n, err := ret.RowsAffected()

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
			ret, err = tx.Exec("delete from fileinfo where Id=?", question.RemoveFile[i])
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
			if n == 0 {
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

				err = lib.DelRemoveAll(path.Dir(file_path)+"/", "Resources/Annex/")
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Folder deleted successfully!")
				}
				// if strings.Contains(path.Dir(file_path), "Resources/Annex/") && path.Dir(file_path) != "" && path.Dir(file_path) != "/" {
				// 	err = os.RemoveAll(path.Dir(file_path) + "/") //删除转换后的文件夹
				// 	if err != nil {
				// 		fmt.Println(err)
				// 	} else {
				// 		fmt.Println("Folder deleted successfully!")
				// 	}
				// }

			}

			err = lib.RemoveFile(file_path)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "文件删除错误",
					"data": "{}",
				})
				return
			}
		}

	}

	var fileId int64
	files, err := c.MultipartForm() // 获取文件

	if len(files.File["files"]) > 0 && question.QuestionType == 5 { // 修改实操课
		file := files.File["files"]
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "未收到文件",
				"data": "{}",
			})
			return

		}

		questionnew := new(model.Question)
		Fileid := 0
		tx.QueryRow("select a.QuestionType,CASE  WHEN a.QuestionType=5 THEN b.FilePath ELSE 	a.QuestionContent END  'QuestionContent' ,b.Id 'Fileid' from question a  left join fileinfo b on a.QuestionContent=b.Id  where a.QuestionId=?", question.QuestionId).Scan(&questionnew.QuestionType, &questionnew.QuestionContent, &Fileid)

		if questionnew.QuestionType == 5 && questionnew.QuestionContent != "" { //删除原来的视频文件

			filesuffix := path.Ext(questionnew.QuestionContent)

			filepath := questionnew.QuestionContent[0 : len(questionnew.QuestionContent)-len(filesuffix)]

			if strings.Contains(questionnew.QuestionContent, "Resources/Zip/") {
				err = lib.DelRemoveAll(filepath+"/", "Resources/Zip/")
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Folder deleted successfully!")
				}
				err = lib.RemoveFile(questionnew.QuestionContent)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("file deleted successfully!")
				}
				if Fileid != 0 {
					ret, err := tx.Exec("delete  from  fileinfo where id=?", Fileid)
					if err != nil {
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "删除文件失败",
							"data": "{}",
						})
						return
					}

					n, err := ret.RowsAffected()
					if err != nil {
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "删除文件失败",
							"data": "{}",
						})
						return
					}
					if n <= 0 {
						if err != nil {
							tx.Rollback()
							c.JSON(http.StatusOK, gin.H{
								"code": 0,
								"msg":  "删除文件失败",
								"data": "{}",
							})
							return
						}
					}
				}
			}
		}

		//for i := 0; i < len(file); i++ {
		i := 0
		//tx.Exec()
		filename := goutils.GUID()
		file_path := "Resources/Zip/" + filename + path.Ext(file[i].Filename)

		zipExtArr := []string{".7z", ".zip", ".rar"}
		iszip := lib.In(path.Ext(file[i].Filename), zipExtArr)

		if iszip {
			toolConfig := new(lib.ToolConfig)
			tx.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)
			if toolConfig.StaticResourcesType == 1 {
				errr := lib.UploadFile(c, file[i], file_path) // 保存文件
				if errr != nil {
					tx.Rollback()
					c.JSON(http.StatusOK, gin.H{
						"code": 0,
						"msg":  "操作失败" + err.Error(),
						"data": "{}",
					})
					return
				} // 保存文件
			} else if toolConfig.StaticResourcesType != 1 {
				c.SaveUploadedFile(file[i], file_path)
			}
			err := lib.UnCompressPkg(file_path, "Resources/Zip/"+filename)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "题目压缩包解压失败",
					"data": "{}",
				})
				return
			}
			if toolConfig.StaticResourcesType != 1 {
				os.Remove(file_path)
				cmd := exec.Command("ossutil64", "cp", "-r", "E:/work/StudyExamPlatform/study_exam/api/"+"Resources/Zip/"+filename+"/", "oss://studyexamplatform"+"/"+"Resources/Zip/"+filename+"/")
				err = cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
				if strings.Contains("Resources/Zip/"+filename+"/", "Resources/Zip/") && path.Dir("Resources/Zip/"+filename+"/") != "" && path.Dir("Resources/Zip/"+filename+"/") != "/" {
					err = os.RemoveAll("Resources/Zip/" + filename + "/")
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("file deleted successfully!")
					}
				}
			}

			// errr := lib.UploadFile(c, file[i], file_path) // 保存文件
			// if errr != nil {
			// 	tx.Rollback()
			// 	c.JSON(http.StatusOK, gin.H{
			// 		"code": 0,
			// 		"msg":  "操作失败" + err.Error(),
			// 		"data": "{}",
			// 	})
			// 	return
			// } // 保存文件

			// _, err := lib.Unzip(file_path, "Resources/Zip/"+filename)

		} else {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "文件格式不对",
				"data": "{}",
			})
			return
		}

		filenameall := path.Base(file[i].Filename)
		filesuffix := path.Ext(filenameall)

		//orginfilename := filenameall[0 : len(filenameall)-len(filesuffix)]
		ret, err := tx.Exec("insert into  fileinfo(FileType,FileName,FileUseTo,SchoolId,FilePath)  values(?,?,?,?,?)", filesuffix, filenameall, "题目文件zip", 0, file_path)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增附件失败",
				"data": "{}",
			})
			return
		}
		fileId, err = ret.LastInsertId()

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增题目失败",
				"data": "{}",
			})
			return
		}
		if fileId == 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "新增题目失败",
				"data": "{}",
			})
			return
		}

		sqlStr := "update  question  set   SchoolId=?,QuestionPoolId=?,QuestionName=?,QuestionType=?,QuestionContent=?,Digree=?,MajorID=?,CollegeId=?,CourseId=?,Answer=?,QuestionCategory=?  where QuestionId=? "

		ret, err = tx.Exec(sqlStr, question.SchoolId, question.QuestionPoolId, question.QuestionName, question.QuestionType, fileId, question.Digree, question.MajorID, question.CollegeId, question.CourseId, question.Answer, question.QuestionCategory, question.QuestionId)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		if n <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
		//break
		//}

	} else if question.QuestionType != 5 { // 修改带附件的非实操课
		sqlStr := "update  question  set   SchoolId=?,QuestionPoolId=?,QuestionName=?,QuestionType=?,QuestionContent=?,Digree=?,MajorID=?,CollegeId=?,CourseId=?,Answer=?  where QuestionId=? "

		ret, err := tx.Exec(sqlStr, question.SchoolId, question.QuestionPoolId, question.QuestionName, question.QuestionType, question.QuestionContent, question.Digree, question.MajorID, question.CollegeId, question.CourseId, question.Answer, question.QuestionId)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		if n < 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		if len(files.File["files"]) > 0 {
			err = lib.UploadAnnex(c, files, tx, "questionrelation", int64(question.QuestionId), "QuestionId", "题目附件")
			if err != nil {
				return
			}

		}

	} else if question.QuestionType == 5 && len(files.File["files"]) == 0 {
		sqlStr := "update  question  set   SchoolId=?,QuestionPoolId=?,QuestionName=?,QuestionType=?, Digree=?,MajorID=?,CollegeId=?,CourseId=?,Answer=?  where QuestionId=? "

		ret, err := tx.Exec(sqlStr, question.SchoolId, question.QuestionPoolId, question.QuestionName, question.QuestionType, question.Digree, question.MajorID, question.CollegeId, question.CourseId, question.Answer, question.QuestionId)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}

		if n < 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败",
				"data": "{}",
			})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "提交事务失败",
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

func DelQuestion(c *gin.Context) {

	tx, err := lib.Db.Begin()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	var question model.Question
	json.Unmarshal([]byte(body), &question)

	if question.QuestionId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	Fileid := 0
	tx.QueryRow("select a.QuestionType,CASE  WHEN a.QuestionType=5 THEN b.FilePath ELSE 	a.QuestionContent END  'QuestionContent' ,b.Id 'Fileid' from question a  left join fileinfo b on a.QuestionContent=b.Id  where a.QuestionId=?", question.QuestionId).Scan(&question.QuestionType, &question.QuestionContent, &Fileid)

	if question.QuestionType == 5 {

		if question.QuestionContent != "" {

			filesuffix := path.Ext(question.QuestionContent)

			filepath := question.QuestionContent[0 : len(question.QuestionContent)-len(filesuffix)]

			if strings.Contains(question.QuestionContent, "Resources/Zip/") {
				err = lib.DelRemoveAll(filepath+"/", "Resources/Zip/")
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Folder deleted successfully!")
				}
				//os.RemoveAll(filepath)
				err = lib.RemoveFile(question.QuestionContent)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("file deleted successfully!")
				}
				if Fileid != 0 {
					ret, err := tx.Exec("delete  from  fileinfo where id=?", Fileid)
					if err != nil {
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "删除文件失败",
							"data": "{}",
						})
						return
					}

					n, err := ret.RowsAffected()
					if err != nil {
						tx.Rollback()
						c.JSON(http.StatusOK, gin.H{
							"code": 0,
							"msg":  "删除文件失败",
							"data": "{}",
						})
						return
					}
					if n <= 0 {
						if err != nil {
							tx.Rollback()
							c.JSON(http.StatusOK, gin.H{
								"code": 0,
								"msg":  "删除文件失败",
								"data": "{}",
							})
							return
						}
					}

				}
			}

		}
	}

	ret, err := tx.Exec("delete from question where QuestionId=? ", question.QuestionId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除文件失败",
			"data": "{}",
		})
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除文件失败",
			"data": "{}",
		})
		return
	}
	if n <= 0 {
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除文件失败",
				"data": "{}",
			})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "提交事务失败",
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

func GetQuestionByQuestionId(c *gin.Context) {

	QuestionId := c.Query("QuestionId")

	questionView := new(model.QuestionViewFile)
	lib.Db.QueryRow("SELECT a.QuestionId,a.SchoolId,a.QuestionPoolId,a.QuestionName,a.QuestionType,a.QuestionContent,a.Digree,a.MajorID,a.CollegeId,a.CourseId,a.Answer,a.QuestionCategory,b.MajorName,c.CollegeName,d.CourseName FROM  question a  left join  college  c  on a.CollegeId=c.Id  left join  major b  on a.MajorID=b.MajorId   	 left join  course d on a.CourseId=d.Id   where  a.QuestionId=?", QuestionId).
		Scan(&questionView.QuestionId, &questionView.SchoolId, &questionView.QuestionPoolId, &questionView.QuestionName, &questionView.QuestionType, &questionView.QuestionContent, &questionView.Digree, &questionView.MajorID, &questionView.CollegeId, &questionView.CourseId, &questionView.Answer, &questionView.QuestionCategory, &questionView.MajorName, &questionView.CollegeName, &questionView.CourseName)
	if questionView.QuestionId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	questionView.FileInfo = make([]*model.FileInfo, 0)
	if questionView.QuestionType == 5 {
		fileid := questionView.QuestionContent

		fileinfo := new(model.FileInfo)
		lib.Db.QueryRow("select Id,FileType,FilePath,FileName from fileinfo    where Id=? ", fileid).Scan(&fileinfo.Id, &fileinfo.FileType, &fileinfo.FilePath, &fileinfo.FileName)
		if fileinfo.Id > 0 {
			questionView.FileInfo = append(questionView.FileInfo, fileinfo)
		}
	} else {
		rows, err := lib.Db.Query("select b.Id,b.FileType,b.FilePath,b.FileName from questionrelation  a left join fileinfo b on a.FileId=b.Id where a.QuestionId=? ", QuestionId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		for rows.Next() {
			newfileinfo := new(model.FileInfo)
			err := rows.Scan(&newfileinfo.Id, &newfileinfo.FileType, &newfileinfo.FilePath, &newfileinfo.FileName)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "失败",
					"data": "{}",
				})
				return
			}
			questionView.FileInfo = append(questionView.FileInfo, newfileinfo)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": questionView,
	})
}

func GetQuestionExeclData(c *gin.Context) {
	files, err := c.MultipartForm() // 获取文件
	//msg := ""
	if len(files.File["files"]) > 0 {
		filesarr := files.File["files"]

		filename := goutils.GUID()
		for i := 0; i < len(filesarr); i++ {
			file := filesarr[i]
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
				return
			}

			filePath := "Resources/Execl/" + filename + path.Ext(file.Filename)
			if err := lib.UploadFile(c, file, filePath); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}

			f, err := excelize.OpenFile(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}

			questionExeclDataarr := make([]*QuestionExeclData, 0)
			if f.GetSheetMap() == nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败0",
					"data": questionExeclDataarr,
				})
				return
			}
			name := f.GetSheetMap()[1]

			rows := f.GetRows(name)
			if len(rows) == 0 {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败1",
					"data": questionExeclDataarr,
				})
				return
			}
			// 遍历每行数据
			for i, row := range rows {
				fmt.Printf("Row %d:\n", i+1)
				if i == 0 {
					continue
				}
				questionExeclData := new(QuestionExeclData)
				questionExeclData.QuestionPoolId = 1
				answerArr := make([]string, 0)
				// 遍历每列数据
				for j, colCell := range row {

					if rows[0][j] == "题目名称" {
						questionExeclData.QuestionName = colCell
					} else if rows[0][j] == "题目类型" {

						if colCell == "单选题" {
							questionExeclData.QuestionType = 1
						} else if colCell == "多选题" {
							questionExeclData.QuestionType = 2
						} else if colCell == "判断题" {
							questionExeclData.QuestionType = 3
						} else if colCell == "填空题" {
							questionExeclData.QuestionType = 4
						} else if colCell == "实操题" {
							questionExeclData.QuestionType = 5
						} else {
							questionExeclData.Status = 2
						}

					} else if rows[0][j] == "难度系数" {
						digree, err := strconv.Atoi(colCell)

						if err == nil {
							questionExeclData.Digree = digree
						} else {
							questionExeclData.Status = 2
						}

					} else if strings.Contains(rows[0][j], "答案") {
						if colCell == "" {
							continue
						}

						if questionExeclData.QuestionType == 2 || questionExeclData.QuestionType == 4 {
							answerArr = append(answerArr, colCell)
						} else {
							questionExeclData.Answer = colCell
						}

					} else if strings.Contains(rows[0][j], "选项") {
						if colCell == "" {
							continue
						}
						questionExeclData.QuestionContent = append(questionExeclData.QuestionContent, colCell)
					}

				}

				if questionExeclData.QuestionType == 2 || questionExeclData.QuestionType == 4 {
					questionExeclData.Answer = lib.ArrayToString(answerArr)
				}

				questionExeclDataarr = append(questionExeclDataarr, questionExeclData)
			}

			// if msg == "" {
			// 	msg = "操作成功"
			// }
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "操作成功",
				"data": questionExeclDataarr,
			})
			os.Remove(filePath) //一直是本地操作
		}
	}
}

func MatchAddQuestion(c *gin.Context) {

	var arrQuestionExeclData []*QuestionExeclData
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	json.Unmarshal([]byte(body), &arrQuestionExeclData)
	SchoolId := 0
	CollegeId := 0
	MajorId := 0
	CourseId := 0

	for i := 0; i < len(arrQuestionExeclData); i++ {

		tempenetity := arrQuestionExeclData[i]

		//if tempenetity.QuestionPoolId == 2 {

		lib.Db.QueryRow("select SchoolId,CollegeId,MajorId,Id from course where Id=?", tempenetity.CourseId).Scan(&SchoolId, &CollegeId, &MajorId, &CourseId)
		if SchoolId != 0 {
			tempenetity.SchoolId = SchoolId
		}
		ret, err := lib.Db.Exec("insert into  question(SchoolId,QuestionPoolId,QuestionName,QuestionType,QuestionContent,Digree,CollegeId,MajorID,CourseId,Answer,TeacherId) values(?,?,?,?,?,?,?,?,?,?,?)", tempenetity.SchoolId, tempenetity.QuestionPoolId, tempenetity.QuestionName, tempenetity.QuestionType, lib.ArrayToString(tempenetity.QuestionContent), tempenetity.Digree, CollegeId, MajorId, CourseId, tempenetity.Answer, tempenetity.TeacherId)
		if err != nil {
			arrQuestionExeclData[i].Status = 0
			continue
		}
		n, err := ret.RowsAffected()
		if err != nil {
			arrQuestionExeclData[i].Status = 0
			continue
		}
		if n <= 0 {
			arrQuestionExeclData[i].Status = 0
			continue
		} else {
			arrQuestionExeclData[i].Status = 1

		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": arrQuestionExeclData,
	})
}
func GetQuestionBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	var questionView []*model.QuestionView
	rows, err := lib.Db.Query("SELECT a.QuestionId,a.SchoolId,a.QuestionPoolId,a.QuestionName,a.QuestionType,a.QuestionContent,a.Digree,a.MajorID,a.CollegeId,a.CourseId,a.Answer, COALESCE(b.MajorName,'') 'MajorName',COALESCE(c.CollegeName,'') 'CollegeName',COALESCE(d.CourseName,'') 'CourseName',a.QuestionCategory  FROM  question a  left join  major b  on a.MajorID=b.MajorId   left join  college  c  on a.CollegeId=c.Id 	 left join  course d on a.CourseId=d.Id  where  a.SchoolId=?", SchoolId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	for rows.Next() {
		newquestionView := new(model.QuestionView)
		err := rows.Scan(&newquestionView.QuestionId, &newquestionView.SchoolId, &newquestionView.QuestionPoolId, &newquestionView.QuestionName, &newquestionView.QuestionType, &newquestionView.QuestionContent, &newquestionView.Digree, &newquestionView.MajorID, &newquestionView.CollegeId, &newquestionView.CourseId, &newquestionView.Answer, &newquestionView.MajorName, &newquestionView.CollegeName, &newquestionView.CourseName, &newquestionView.QuestionCategory)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		questionView = append(questionView, newquestionView)

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": questionView,
	})
}

func GetExamQuestionBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	var questionView []*model.QuestionView
	rows, err := lib.Db.Query("SELECT a.QuestionId,a.SchoolId,a.QuestionPoolId,a.QuestionName,a.QuestionType,a.QuestionContent,a.Digree,a.MajorID,a.CollegeId,a.CourseId,a.Answer, COALESCE(b.MajorName,'') 'MajorName',COALESCE(c.CollegeName,'') 'CollegeName',COALESCE(d.CourseName,'') 'CourseName',a.QuestionCategory FROM  question a  left join  major b  on a.MajorID=b.MajorId   left join  college  c  on a.CollegeId=c.Id 	 left join  course d on a.CourseId=d.Id  where  a.SchoolId=? and a.QuestionCategory!=1", SchoolId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	for rows.Next() {
		newquestionView := new(model.QuestionView)
		err := rows.Scan(&newquestionView.QuestionId, &newquestionView.SchoolId, &newquestionView.QuestionPoolId, &newquestionView.QuestionName, &newquestionView.QuestionType, &newquestionView.QuestionContent, &newquestionView.Digree, &newquestionView.MajorID, &newquestionView.CollegeId, &newquestionView.CourseId, &newquestionView.Answer, &newquestionView.MajorName, &newquestionView.CollegeName, &newquestionView.CourseName, &newquestionView.QuestionCategory)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		questionView = append(questionView, newquestionView)

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": questionView,
	})
}

func GetTrainQuestionBySchoolId(c *gin.Context) {

	SchoolId := c.Query("SchoolId")

	var questionView []*model.QuestionView
	rows, err := lib.Db.Query("SELECT a.QuestionId,a.SchoolId,a.QuestionPoolId,a.QuestionName,a.QuestionType,a.QuestionContent,a.Digree,a.MajorID,a.CollegeId,a.CourseId,a.Answer, COALESCE(b.MajorName,'') 'MajorName',COALESCE(c.CollegeName,'') 'CollegeName',COALESCE(d.CourseName,'') 'CourseName',a.QuestionCategory FROM  question a  left join  major b  on a.MajorID=b.MajorId   left join  college  c  on a.CollegeId=c.Id 	 left join  course d on a.CourseId=d.Id  where  a.SchoolId=? and a.QuestionCategory!=2", SchoolId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	for rows.Next() {
		newquestionView := new(model.QuestionView)
		err := rows.Scan(&newquestionView.QuestionId, &newquestionView.SchoolId, &newquestionView.QuestionPoolId, &newquestionView.QuestionName, &newquestionView.QuestionType, &newquestionView.QuestionContent, &newquestionView.Digree, &newquestionView.MajorID, &newquestionView.CollegeId, &newquestionView.CourseId, &newquestionView.Answer, &newquestionView.MajorName, &newquestionView.CollegeName, &newquestionView.CourseName, &newquestionView.QuestionCategory)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		questionView = append(questionView, newquestionView)

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": questionView,
	})
}
func GetSocietyQuestionByStudentId(c *gin.Context) {

	StudentId := c.Query("StudentId")
	var questionView []*model.QuestionView
	rows, err := lib.Db.Query("SELECT a.QuestionId,a.SchoolId,a.QuestionPoolId,a.QuestionName,a.QuestionType,a.QuestionContent,a.Digree,a.MajorID,a.CollegeId,a.CourseId,a.Answer, COALESCE(b.MajorName,'') 'MajorName',COALESCE(c.CollegeName,'') 'CollegeName',COALESCE(d.CourseName,'') 'CourseName' FROM  question a  left join  major b  on a.MajorID=b.MajorId   left join  college  c  on a.CollegeId=c.Id 	 left join  course d on a.CourseId=d.Id  where  "+
		" a.CourseId in (select d.CourseId  from examstudent a "+
		" 	left join exam b on a.ExamId=b.Id "+
		" LEFT JOIN examsession c on b.Id=c.ExamId "+
		" LEFT JOIN testpaper d on c.TestPaperId=d.Id "+
		" LEFT JOIN  course e on d.CourseId=e.Id "+
		" where StudentId=? and d.CourseId!=0 and e.CourseCode!=''   ) and a.QuestionCategory!=2 ", StudentId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	questionView = make([]*model.QuestionView, 0)
	for rows.Next() {
		newquestionView := new(model.QuestionView)
		err := rows.Scan(&newquestionView.QuestionId, &newquestionView.SchoolId, &newquestionView.QuestionPoolId, &newquestionView.QuestionName, &newquestionView.QuestionType, &newquestionView.QuestionContent, &newquestionView.Digree, &newquestionView.MajorID, &newquestionView.CollegeId, &newquestionView.CourseId, &newquestionView.Answer, &newquestionView.MajorName, &newquestionView.CollegeName, &newquestionView.CourseName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		questionView = append(questionView, newquestionView)

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": questionView,
	})
}

func AddOperPractice(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var operpractice model.OperPractice
	json.Unmarshal([]byte(body), &operpractice)

	if operpractice.QuestionId == 0 && operpractice.StudentId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	ret, err := lib.Db.Exec("insert into operpractice(QuestionId,StudentId,PracticeStep,PracticeAnswer,PracticeScore,CreateTime) values(?,?,?,?,?,now())", operpractice.QuestionId, operpractice.StudentId, operpractice.PracticeStep, operpractice.PracticeAnswer, operpractice.PracticeScore)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作失败",
			"data": "{}",
		})
		return
	}
	if n <= 0 {
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

func GetOperPracticeByQuestionId(c *gin.Context) {

	QuestionId := c.Query("QuestionId")

	if QuestionId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	sql := " select  a.StudentId,b.TrueName,COUNT(StudentId)'StudentNum',MAX(CreateTime) 'MaxTime' from operpractice  a " +
		" LEFT JOIN student b on  a.StudentId=b.Id " +
		"  where   QuestionId=? " +
		"	 GROUP BY StudentId "
	rows, err := lib.Db.Query(sql, QuestionId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "操作错误",
			"data": "{}",
		})
		return
	}
	operpracticeViewarr := make([]*model.OperPracticeView, 0)
	for rows.Next() {
		model := new(model.OperPracticeView)
		err := rows.Scan(&model.StudentId, &model.TrueName, &model.StudentNum, &model.MaxTime)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
			return
		}
		operpracticeViewarr = append(operpracticeViewarr, model)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": operpracticeViewarr,
	})

}

type QuestionExeclData struct {
	QuestionPoolId  int      `json:"QuestionPoolId" db:"QuestionPoolId"`
	QuestionName    string   `json:"QuestionName" db:"QuestionName"`
	QuestionContent []string `json:"QuestionContent" db:"QuestionContent"`
	QuestionType    int      `json:"QuestionType" db:"QuestionType"`
	Digree          int      `json:"Digree" db:"Digree"`
	Answer          string   `json:"Answer" db:"Answer"`
	CourseId        int      `json:"CourseName" db:"CourseName"`
	TeacherId       int      `json:"TeacherId" db:"TeacherId"`
	Status          int      `json:"Status" db:"Status"` //0 标识未导入  1 标识表示导入成功 2 数据错误
	SchoolId        int      `json:"SchoolId" db:"SchoolId"`
}
