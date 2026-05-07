package lib

import (
	model "StudyExamPlatformAPI/Model"
	"database/sql"
	"fmt"
	"mime/multipart"
	"net/http"
	"os/exec"
	"path"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"

	goutils "github.com/typa01/go-utils"
)

func UploadAnnex(c *gin.Context, files *multipart.Form, tx *sql.Tx, tablename string, abilityId int64, abilityField string, abilityname string) error {
	var filearr []*model.FileInfo

	var officearr []*OfficePdf
	if len(files.File["files"]) > 0 {
		file := files.File["files"]

		for i := 0; i < len(file); i++ {
			//tx.Exec()
			filename := goutils.GUID()
			file_path := "Resources/Annex/" + filename + path.Ext(file[i].Filename)
			errr := UploadFile(c, file[i], file_path)

			if errr != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "操作失败" + errr.Error(),
					"data": "{}",
				})
				return errr
			} // 保存文件

			filenameall := path.Base(file[i].Filename)
			filesuffix := path.Ext(filenameall)

			//orginfilename := filenameall[0 : len(filenameall)-len(filesuffix)]
			ret, err := tx.Exec("insert into  fileinfo(FileType,FileName,FileUseTo,SchoolId,FilePath)  values(?,?,?,?,?)", filesuffix, filenameall, abilityname, 0, file_path)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增附件失败",
					"data": "{}",
				})
				return err
			}
			fileId, err := ret.LastInsertId()

			videoExtArr := []string{".avi", ".mp4", ".mov", ".wmv", ".flv", ".mkv", ".mpg", ".rmvb"}
			isvideo := In(path.Ext(file[i].Filename), videoExtArr)

			if isvideo {
				fileinfo := new(model.FileInfo)
				fileinfo.Id = int(fileId)
				fileinfo.FileType = path.Ext(file[i].Filename)
				fileinfo.FileName = file[i].Filename
				fileinfo.FilePath = file_path
				filearr = append(filearr, fileinfo)
			}

			officeExtArr := []string{".xlsx", ".xls", ".docx", ".doc", ".pptx", ".ppt"}
			isoffice := In(path.Ext(file[i].Filename), officeExtArr)
			if isoffice {
				officeinfo := new(OfficePdf)
				officeinfo.InputFile = file_path
				officeinfo.OutputFile = "Resources/Annex"
				officearr = append(officearr, officeinfo)
			}
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增附件失败",
					"data": "{}",
				})
				return err
			}
			if fileId == 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增附件失败",
					"data": "{}",
				})
				return err
			}
			ret, err = tx.Exec("insert into  "+tablename+"("+abilityField+",FileId)  values(?,?)", abilityId, fileId)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增" + abilityname + "文件关系表失败",
					"data": "{}",
				})
				return err
			}
			n, err := ret.RowsAffected()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增" + abilityname + "文件关系表失败",
					"data": "{}",
				})
				return err
			}
			if n == 0 {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "新增" + abilityname + "文件关系表失败",
					"data": "{}",
				})
				return err
			}
		}
	}

	fileinfochanmutex := LoadChan()
	for i := 0; i < len(filearr); i++ {
		fileinfochanmutex.Fileinfochan <- filearr[i]
	}
	go func() {
		if !fileinfochanmutex.Mutex.TryLock() {
			return
		}
		DealVideoChan()

		fileinfochanmutex.Mutex.Unlock()

	}()

	officechanmutex := LoadOfficeChan()
	for i := 0; i < len(officearr); i++ {
		officechanmutex.OfficePdf <- officearr[i]
	}
	go func() {
		if !officechanmutex.Mutex.TryLock() {
			return
		}
		OfficeToPdf()

		officechanmutex.Mutex.Unlock()

	}()

	return nil
}

func DealVideoChan() {

	fileinfochanmutex := LoadChan()

	for {

		select {
		case fileinfo := <-fileinfochanmutex.Fileinfochan:
			fmt.Println(fileinfo)
			filepath, err := VideoToM3u8(fileinfo)
			if err != nil {
				fmt.Println(err)
			}
			ret, err := Db.Exec(" update fileinfo set FilePath=?,FileType=?  where Id=?", filepath, path.Ext(filepath), fileinfo.Id)
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
		case <-time.After(time.Second * 2):
			return
		}

	}
}

func OfficeToPdf() {

	officechanmutex := LoadOfficeChan()

	for {

		//无锁 进入下面
		//上锁

		select {
		case officePdf := <-officechanmutex.OfficePdf:

			con := LoadConfig()

			//dir, _ := os.Getwd()
			cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf:writer_pdf_Export", officePdf.InputFile, "--outdir", officePdf.OutputFile)
			err := cmd.Run()
			if err != nil {
				fmt.Println("转换失败:", err)
				return
			}

			fmt.Println("转换成功！")

			toolConfig := new(ToolConfig)
			Db.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)

			if toolConfig.StaticResourcesType != 1 {
				endpoint := con.ENDPOINT
				// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
				accessKeyId := con.ACCESS_KEY_ID
				accessKeySecret := con.ACCESS_KEY_SECRET
				// yourBucketName填写存储空间名称。
				bucketName := con.BACKET_NAME

				client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
				if err != nil {
					fmt.Println(err.Error())
				}
				// 获取存储空间。
				bucket, err := client.Bucket(bucketName)
				if err != nil {
					fmt.Println(err.Error())
				}

				filenameall := path.Base(officePdf.InputFile)
				filesuffix := path.Ext(filenameall)

				filename := filenameall[0 : len(filenameall)-len(filesuffix)]
				err = bucket.PutObjectFromFile(officePdf.OutputFile+"/"+filename+".pdf", officePdf.OutputFile+"/"+filename+".pdf")
				if err != nil {
					fmt.Println(err.Error())
				}
			}

		case <-time.After(time.Second * 2):
			return
		}

	}

	//soffice --headless --convert-to pdf:writer_pdf_Export "C:\Users\user\Desktop\document.odt" --outdir "C:\Users\user\Desktop\pdf"
	// cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf:writer_pdf_Export", inputFile, "--outdir", outputDir)
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println("转换失败:", err)
	// 	return
	// }

	// fmt.Println("转换成功！")
}

func CountOccurrences(arr []int, value int) int {
	count := 0
	for _, num := range arr {
		if num == value {
			count++
		}
	}
	return count
}
