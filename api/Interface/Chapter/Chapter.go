package chapter

import (
	model "StudyExamPlatformAPI/Model"
	lib "StudyExamPlatformAPI/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	goutils "github.com/typa01/go-utils"
)

// 附件和章节上传
func AddChapter(c *gin.Context) {

	data, _ := c.GetPostForm("data")
	var chapter model.Chapter
	json.Unmarshal([]byte(data), &chapter)

	if chapter.CourseId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	Count := 0
	lib.Db.QueryRow(" SELECT   COUNT(1)  'Count' FROM `course`   where  (NOW()  BETWEEN CourseStartTime AND CourseEndTime  or  NOW()>CourseEndTime) and Status=1  and id=? ", chapter.CourseId).Scan(&Count)
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

	tx.QueryRow("select  COALESCE( MAX(ChapterOrder)+1,0) 'ChapterOrder' from  chapter where CourseId=?", chapter.CourseId).Scan(&chapter.ChapterOrder)
	ret, err := tx.Exec("insert into  chapter(ChapterName,CourseId,ChapterType,ChapterOrder,ChapterContent)  values(?,?,?,?,?) ", chapter.ChapterName, chapter.CourseId, chapter.ChapterType, chapter.ChapterOrder, chapter.ChapterContent)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "新增章节失败",
			"data": "{}",
		})
		return
	}
	ChapterId, err := ret.LastInsertId()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "新增章节失败",
			"data": "{}",
		})
		return
	}
	if ChapterId == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "新增章节失败",
			"data": "{}",
		})
		return
	}
	files, err := c.MultipartForm() // 获取文件
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未收到文件",
			"data": "{}",
		})
		return

	}
	err = lib.UploadAnnex(c, files, tx, "chapterrelation", ChapterId, "ChapterId", "章节附件")

	if err != nil {
		return
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
		"data": "{\"ChapterId\":" + lib.Strval(ChapterId) + "}",
	})

}

func EditChapter(c *gin.Context) {

	data, _ := c.GetPostForm("data")
	var chapter model.ChapterEdit
	json.Unmarshal([]byte(data), &chapter)

	if chapter.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	Count := 0
	lib.Db.QueryRow(" SELECT   COUNT(1)  'Count' FROM `course`   where  (NOW()  BETWEEN CourseStartTime AND CourseEndTime  or  NOW()>CourseEndTime) and Status=1   and id=? ", chapter.CourseId).Scan(&Count)
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
	newchaptertemp := new(model.Chapter)
	err = tx.QueryRow("select  Id,ChapterType,ChapterContent from  chapter where Id=? ", chapter.Id).Scan(&newchaptertemp.Id, &newchaptertemp.ChapterType, &newchaptertemp.ChapterContent)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	fileinfmodel := new(model.FileInfo)
	if chapter.ChapterContent == "-1" {
		tx.QueryRow("select Id,FileName,FilePath from fileinfo where Id=?", newchaptertemp.ChapterContent).Scan(&fileinfmodel.Id, &fileinfmodel.FileName, &fileinfmodel.FilePath)
		ret, err := tx.Exec(" delete from fileinfo where Id=? ", fileinfmodel.Id)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
		}

		if n < 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
		}

		filenameall := path.Base(fileinfmodel.FilePath)
		filesuffix := path.Ext(filenameall)

		filename := filenameall[0 : len(filenameall)-len(filesuffix)]

		//删视频
		err = lib.RemoveFile("Resources/Video/" + path.Base(filename) + path.Ext(fileinfmodel.FileName)) // 删除原mp4或其他视频格式
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file deleted successfully!")
		}

		err = lib.DelRemoveAll(fileinfmodel.FilePath+"/", "Resources/Video/")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file deleted successfully!")
		}

	}

	if newchaptertemp.ChapterType == "1" && chapter.ChapterType == "0" { //视频课转图文课
		tx.QueryRow("select Id,FileName,FilePath from fileinfo where Id=?", newchaptertemp.ChapterContent).Scan(&fileinfmodel.Id, &fileinfmodel.FileName, &fileinfmodel.FilePath)

		ret, err := tx.Exec(" delete from fileinfo where Id=? ", fileinfmodel.Id)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
		}
		n, err := ret.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
		}

		if n < 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
		}
		filenameall := path.Base(fileinfmodel.FilePath)
		filesuffix := path.Ext(filenameall)

		filename := filenameall[0 : len(filenameall)-len(filesuffix)]

		//删原视频
		err = lib.RemoveFile("Resources/Video/" + path.Base(filename) + path.Ext(fileinfmodel.FileName)) // 删除原mp4或其他视频格式
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file deleted successfully!")
		}

		// err = lib.DelRemoveAll("./Resources/Video/"+path.Base(filename)+path.Ext(fileinfmodel.FileName), "Resources/Video/")
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println("file deleted successfully!")
		// }

		err = lib.DelRemoveAll(fileinfmodel.FilePath+"/", "Resources/Video/")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file deleted successfully!")
		}
	}

	if newchaptertemp.ChapterType == "0" && chapter.ChapterType == "1" { // 图文课转视频课
		//删富文本中的视频和图片

		flag := lib.DelHtmlResources(newchaptertemp.ChapterContent, tx)
		if !flag {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
		}
	}

	if newchaptertemp.ChapterType == "0" && chapter.ChapterType == "0" { // 图文课 正常编辑

		flag := lib.UpdateHtmlResources(newchaptertemp.ChapterContent, chapter.ChapterContent, tx)
		if !flag {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作错误",
				"data": "{}",
			})
		}
	}

	ret, err := tx.Exec("Update chapter set ChapterName=?,CourseId=?,ChapterType=?,ChapterContent=? where Id=?", chapter.ChapterName, chapter.CourseId, chapter.ChapterType, chapter.ChapterContent, chapter.Id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "修改章节失败",
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "修改章节失败",
			"data": "{}",
		})
		return
	}
	if n < 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "新增章节失败",
			"data": "{}",
		})
		return
	}

	if len(chapter.RemoveFile) > 0 {
		for i := 0; i < len(chapter.RemoveFile); i++ {
			if chapter.RemoveFile[i] == 0 {
				continue
			}
			filename := ""
			file_path := ""
			tx.QueryRow("select FileName,FilePath from  fileinfo where  Id=? ", chapter.RemoveFile[i]).Scan(&filename, &file_path)

			ret, err := tx.Exec("delete from chapterrelation where FileId=? and ChapterId=?", chapter.RemoveFile[i], chapter.Id)
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
			ret, err = tx.Exec("delete from fileinfo where Id=?", chapter.RemoveFile[i])
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

				err = lib.DelRemoveAll(file_path+"/", "Resources/Annex/")
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("file deleted successfully!")
				}

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
	var filearr []*model.FileInfo
	files, err := c.MultipartForm() // 获取文件
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未收到文件",
			"data": "{}",
		})
		return

	}

	if len(files.File["files"]) > 0 {
		err = lib.UploadAnnex(c, files, tx, "chapterrelation", int64(chapter.Id), "ChapterId", "章节附件")

		if err != nil {
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
	fileinfochanmutex := lib.LoadChan()
	for i := 0; i < len(filearr); i++ {
		fileinfochanmutex.Fileinfochan <- filearr[i]
	}
	go func() {
		if !fileinfochanmutex.Mutex.TryLock() {
			return
		}
		lib.DealVideoChan()

		fileinfochanmutex.Mutex.Unlock()

	}()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": "{}",
	})
}

func DelChapter(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var chapter model.ChapterIdStruct
	json.Unmarshal([]byte(body), &chapter)

	if chapter.ChapterId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
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

	var htmtcontent string
	tx.QueryRow("select ChapterContent from  chapter   where Id=?", chapter.ChapterId).Scan(&htmtcontent)

	flag := lib.DelHtmlResources(htmtcontent, tx)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "富文本文件错误",
			"data": "{}",
		})
		return
	}
	ret, err := tx.Exec("Delete from  chapter   where Id=?", chapter.ChapterId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除失败",
			"data": "{}",
		})
		return
	}
	n, err := ret.RowsAffected()
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

	rows, err := tx.Query("select FileId from  chapterrelation where  ChapterId=? ", chapter.ChapterId)
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
		ret, err = tx.Exec(delstr+" and ChapterId=?", chapter.ChapterId)
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

func GetChapterById(c *gin.Context) {

	ChapterId := c.Query("ChapterId")

	ChapterView := new(model.ChapterView)
	lib.Db.QueryRow("select Id,ChapterName,CourseId,ChapterType,ChapterOrder,CASE   WHEN ChapterType='0' THEN 	ChapterContent ELSE (select  FilePath from  fileinfo where id=ChapterContent) END  'ChapterContent' from  chapter where Id=?", ChapterId).Scan(&ChapterView.Id, &ChapterView.ChapterName, &ChapterView.CourseId, &ChapterView.ChapterType, &ChapterView.ChapterOrder, &ChapterView.ChapterContent)

	rows, err := lib.Db.Query("select b.Id,b.FileType,b.FilePath,b.FileName from chapterrelation  a left join fileinfo b on a.FileId=b.Id where a.ChapterId=?", ChapterId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	ChapterView.FileInfo = make([]*model.FileInfo, 0)
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
		ChapterView.FileInfo = append(ChapterView.FileInfo, newfileinfo)

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": ChapterView,
	})
}

func GetChapterByCourseId(c *gin.Context) {

	CourseId := c.Query("CourseId")
	var newchapterarr []*model.Chapter
	rows, err := lib.Db.Query("select Id,ChapterName,CourseId,ChapterType,ChapterOrder from  chapter where CourseId=?   order by ChapterOrder asc", CourseId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	for rows.Next() {
		newchapter := new(model.Chapter)
		err := rows.Scan(&newchapter.Id, &newchapter.ChapterName, &newchapter.CourseId, &newchapter.ChapterType, &newchapter.ChapterOrder)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "{}",
			})
			return
		}
		newchapterarr = append(newchapterarr, newchapter)

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": newchapterarr,
	})
}

// 富文本编辑器上传文件
func UploadChapterFile(c *gin.Context) {

	files, err := c.MultipartForm() // 获取文件

	SchoolId := 0
	data, _ := c.GetPostForm("data")
	json.Unmarshal([]byte(data), &SchoolId)

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

		filetype := path.Ext(file.Filename)
		videoExtArr := []string{".avi", ".mp4", ".mov", ".wmv", ".flv", ".mkv", ".mpg", ".rmvb"}
		isvideo := lib.In(filetype, videoExtArr)

		file_path := "" // 设置保存文件的路径，不要忘了后面的文件名

		FileUseTo := ""
		filename := goutils.GUID()
		if isvideo {
			//是视频
			file_path = "Resources/Video/" + filename + filetype
			FileUseTo = "章节视频课视频"

		} else {
			//是图片
			file_path = "Resources/Img/" + filename + filetype
			FileUseTo = "章节图片"
		}

		err = lib.UploadFile(c, file, file_path) // 保存文件
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "上传失败",
				"data": "{}",
			})
			return
		}

		ret, err := lib.Db.Exec("insert into  fileinfo(FileType,FileName,FileUseTo,SchoolId,FilePath) values(?,?,?,?,?)", filetype, filename, FileUseTo, SchoolId, file_path) // 新增文件表
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		if n <= 0 {
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
			"data": file_path,
		})
	}
}

// 视频课视频文件
func UploadChapterVideoFile(c *gin.Context) {

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
	ChapterVideo := new(model.ChapterVideo)
	data, _ := c.GetPostForm("data")
	json.Unmarshal([]byte(data), &ChapterVideo)

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

		filetype := path.Ext(file.Filename)
		videoExtArr := []string{".avi", ".mp4", ".mov", ".wmv", ".flv", ".mkv", ".mpg", ".rmvb"}
		isvideo := lib.In(filetype, videoExtArr)

		file_path := "" // 设置保存文件的路径，不要忘了后面的文件名

		FileUseTo := ""
		filename := goutils.GUID()
		if isvideo {
			//是视频
			file_path = "Resources/Video/" + filename + filetype
			FileUseTo = "章节视频课视频"

		}
		err = lib.UploadFile(c, file, file_path) // 保存文件
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "上传失败",
				"data": "{}",
			})
			return
		}

		filenameall := path.Base(file.Filename)
		filesuffix := path.Ext(filenameall)

		orginfilename := filenameall[0 : len(filenameall)-len(filesuffix)]

		ret, err := tx.Exec("insert into  fileinfo(FileType,FileName,FileUseTo,SchoolId,FilePath) values(?,?,?,?,?)", filetype, filenameall, FileUseTo, ChapterVideo.SchoolId, file_path) // 新增文件表
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		fileid, err := ret.LastInsertId()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		if fileid <= 0 {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}

		ret, err = tx.Exec(" update chapter set ChapterContent=? where Id=?", fileid, ChapterVideo.ChapterId)
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
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "操作失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		if n < 0 {
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
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "操作成功",
			"data": "{}",
		})
		fileinfochanmutex := lib.LoadChan()
		fileinfo := new(model.FileInfo)
		fileinfo.Id = int(fileid)
		fileinfo.FileType = filetype
		fileinfo.FileName = orginfilename
		fileinfo.FilePath = file_path
		fileinfochanmutex.Fileinfochan <- fileinfo

		go func() {
			if !fileinfochanmutex.Mutex.TryLock() {
				return
			}
			lib.DealVideoChan()

			fileinfochanmutex.Mutex.Unlock()

		}()

		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未收到文件",
			"data": "{}",
		})
		return
	}
}

func DealVideo() {

	queue := lib.LoadQueue()

	for {
		if len(queue.QueueArray) == 0 {
			break //如果count>=10则退出
		}

		//无锁 进入下面
		//上锁
		fileinfo, err := queue.DeleteQueue()
		if err != nil {
			fmt.Println(err)
		}
		filepath, err := lib.VideoToM3u8(fileinfo)
		if err != nil {
			fmt.Println(err)
		}

		ret, err := lib.Db.Exec(" update fileinfo set FilePath=?,FileType=?  where Id=?", filepath, path.Ext(filepath), fileinfo.Id)
		if err != nil {
			fmt.Println(err)
		}

		n, err := ret.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		if n > 0 {
			fmt.Println("操作成功")
		} else {
			fmt.Println("操作失败")
		}
	}
}

func UpdateChapterOrder(c *gin.Context) {

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
	//var chapter model.ChapterIdStruct
	var chapterOrderParam []*model.ChapterOrderParam
	json.Unmarshal([]byte(body), &chapterOrderParam)
	for i := 0; i < len(chapterOrderParam); i++ {
		ChapterId := chapterOrderParam[i].ChapterOrder[0] //0 是Id  1 是order
		ChapterOrder := chapterOrderParam[i].ChapterOrder[1]
		ret, err := tx.Exec("update chapter set ChapterOrder=? where Id=?", ChapterOrder, ChapterId)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return
		}
		n, err := ret.RowsAffected()
		if n < 0 && err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
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
			"msg":  "失败" + err.Error(),
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

func StudyPlanUpload(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}
	var studyplan model.StudyPlan
	json.Unmarshal([]byte(body), &studyplan)

	if studyplan.CourseId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
			"data": "{}",
		})
		return
	}

	ret, err := lib.Db.Exec(" update studyplan set ChapterId=?,LearningRate=?,ChapterOrder=?,IsComplete=? where  CourseId=? and StudentId=?", studyplan.ChapterId, studyplan.LearningRate, studyplan.ChapterOrder, studyplan.IsComplete, studyplan.CourseId, studyplan.StudentId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
			"data": "{}",
		})
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "失败" + err.Error(),
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
			"msg":  "操作失败，数据未更新",
			"data": "{}",
		})
	}
}
