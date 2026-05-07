package lib

import (
	model "StudyExamPlatformAPI/Model"
	"archive/zip"
	"bytes"
	"crypto/des"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver"
)

var TimeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间

var Pattern = `^d{4}-d{2}-d{2} d{2}:d{2}:d{2}$` //验证是否是日期的正则

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

// 加密DES
func Encrypt(src []byte, key []byte) (string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

// 解密DES
func Decrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func In(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

func VideoToM3u8(fileinfo *model.FileInfo) (string, error) {

	logicalCPU := int(runtime.NumCPU() / 2)

	toolConfig := new(ToolConfig)
	Db.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)

	filenameall := path.Base(fileinfo.FilePath)
	filesuffix := path.Ext(filenameall)

	filename := filenameall[0 : len(filenameall)-len(filesuffix)]

	filedir := path.Dir(fileinfo.FilePath)
	err := os.MkdirAll(filedir+"/"+filename, 0777)
	if err != nil {
		return "", err

	} else {
		fmt.Println("Successfully created directories")
	}

	cmd := exec.Command("ffmpeg",
		"-i", fileinfo.FilePath,
		"-threads", Strval(logicalCPU),
		"-c:v", "copy",
		"-c:a", "copy",
		"-y",
		filedir+"/"+filename+"/"+filename+".mp4")

	err = cmd.Run()
	if err != nil {
		return "", err
	}
	err = os.Remove(fileinfo.FilePath)
	if err != nil {
		return "", err
	}
	cmd = exec.Command("ffmpeg",
		"-i", filedir+"/"+filename+"/"+filename+".mp4",
		//"-profile:v", "baseline",
		"-acodec", "copy",
		"-level", "3.0",
		"-start_number", "0",
		"-hls_time", "20",
		"-hls_list_size", "0",
		"-threads", Strval(logicalCPU),
		"-f", "hls", filedir+"/"+filename+"/"+filename+".m3u8")
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	err = os.Remove(filedir + "/" + filename + "/" + filename + ".mp4")
	if err != nil {
		return "", err
	}

	if toolConfig.StaticResourcesType != 1 {
		cmd = exec.Command("ossutil64", "cp", "-r", "E:/work/StudyExamPlatform/study_exam/api/"+filedir+"/"+filename+"/", "oss://studyexamplatform"+"/"+filedir+"/"+filename+"/")
		err = cmd.Run()
		if err != nil {
			return "", err
		}
		if strings.Contains(filedir+"/"+filename+"/", "Resources") && path.Dir(filedir+"/"+filename) != "" && filedir != "/" {
			err = os.RemoveAll(filedir + "/" + filename + "/") //删除转换后的文件夹
			if err != nil {
				return "", err
			}
		}
	}

	return filedir + "/" + filename + "/" + filename + ".m3u8", err

}

func DelHtmlResources(res string, tx *sql.Tx) bool {

	r := strings.NewReader(res)
	doc, err := goquery.NewDocumentFromReader(r)
	if err == nil {

		flag := true
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			//解析<div>标签
			//h,err := s.Html()
			v, t := s.Attr("src")
			v = strings.Replace(v, "\n", "", -1)
			if t {
				myUrl, _ := url.Parse(v)
				params, _ := url.ParseQuery(myUrl.RawQuery)
				filepathtemp := params.Get("filename")
				if filepathtemp != "" {
					RemoveFile(filepathtemp)

					ret, err := tx.Exec(" delete from fileinfo where FilePath=? ", filepathtemp)

					if err != nil {
						flag = false
					}
					n, err := ret.RowsAffected()
					if err != nil {
						flag = false
					}

					if n < 0 {
						flag = false
					}
				}
			}

		})

		if !flag {
			return flag
		}

		doc.Find("iframe").Each(func(i int, s *goquery.Selection) {
			//解析<div>标签
			//h,err := s.Html()
			v, t := s.Attr("src")
			v = strings.Replace(v, "\n", "", -1)
			if t {
				myUrl, _ := url.Parse(v)
				params, _ := url.ParseQuery(myUrl.RawQuery)
				filepathtemp := params.Get("filename")
				if filepathtemp != "" {
					RemoveFile(filepathtemp)
					ret, err := tx.Exec(" delete from fileinfo where FilePath=? ", filepathtemp)

					if err != nil {
						flag = false
					}
					n, err := ret.RowsAffected()
					if err != nil {
						flag = false
					}

					if n < 0 {
						flag = false
					}
				}
			}

		})
		if !flag {
			return flag
		}

		return true
	} else {
		return false
	}
}

func UpdateHtmlResources(res string, newres string, tx *sql.Tx) bool {

	var resarr []string
	var newresarr []string

	//var removearr []string
	r := strings.NewReader(res)
	doc, err := goquery.NewDocumentFromReader(r)
	if err == nil {
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			//解析<div>标签
			//h,err := s.Html()
			v, t := s.Attr("src")
			v = strings.Replace(v, "\n", "", -1)
			if t {
				myUrl, _ := url.Parse(v)
				params, _ := url.ParseQuery(myUrl.RawQuery)
				filepathtemp := params.Get("filename")
				if filepathtemp != "" {
					resarr = append(resarr, filepathtemp)
				}
			}

		})
		doc.Find("iframe").Each(func(i int, s *goquery.Selection) {
			//解析<div>标签
			//h,err := s.Html()
			v, t := s.Attr("src")
			v = strings.Replace(v, "\n", "", -1)
			if t {
				myUrl, _ := url.Parse(v)
				params, _ := url.ParseQuery(myUrl.RawQuery)
				filepathtemp := params.Get("filename")
				if filepathtemp != "" {
					resarr = append(resarr, filepathtemp)
				}
			}

		})
	}

	newr := strings.NewReader(newres)
	newdoc, err := goquery.NewDocumentFromReader(newr)
	if err == nil {
		newdoc.Find("img").Each(func(i int, s *goquery.Selection) {
			//解析<div>标签
			//h,err := s.Html()
			v, t := s.Attr("src")
			v = strings.Replace(v, "\n", "", -1)
			if t {
				myUrl, _ := url.Parse(v)
				params, _ := url.ParseQuery(myUrl.RawQuery)
				filepathtemp := params.Get("filename")
				if filepathtemp != "" {
					newresarr = append(newresarr, filepathtemp)
				}
			}

		})
		newdoc.Find("iframe").Each(func(i int, s *goquery.Selection) {
			//解析<div>标签
			//h,err := s.Html()
			v, t := s.Attr("src")
			v = strings.Replace(v, "\n", "", -1)
			if t {
				myUrl, _ := url.Parse(v)
				params, _ := url.ParseQuery(myUrl.RawQuery)
				filepathtemp := params.Get("filename")
				if filepathtemp != "" {
					newresarr = append(newresarr, filepathtemp)
				}
			}

		})
	}

	for i := 0; i < len(resarr); i++ {
		flag := true
		for j := 0; j < len(newresarr); j++ {
			if resarr[i] == newresarr[j] {
				continue
			}
			flag = false

		}
		if !flag {
			//removearr = append(removearr, resarr[i])
			RemoveFile(resarr[i])
			ret, err := tx.Exec(" delete from fileinfo where FilePath=? ", resarr[i])

			if err != nil {
				return false
			}
			n, err := ret.RowsAffected()
			if err != nil {
				return false
			}

			if n < 0 {
				return false
			}

		}

	}
	return true
}

// Unzip将解压缩一个zip归档文件,
// 复制所有文件和文件夹
// 在zip文件(参数1)
// 到输出目录(参数2)。
func Unzip(src string, destination string) ([]string, error) {

	// 存储任何文件名的变量
	// 在字符串数组中可用
	var filenames []string

	// OpenReader将打开Zip文件
	// 指定名称并返回一个ReadCloser
	// Readcloser关闭Zip文件,
	// 使其不可用于I / O
	// 它返回两个值:
	// 1.对ReadCloser的指针值
	// 2.错误消息(如果有)

	r, err := zip.OpenReader(src)

	// 如果有任何错误则
	// (err!= nill)变为真
	if err != nil {

		return filenames, err

	}

	defer r.Close()
	// 延迟确保文件关闭
	// 在程序结束时不管怎样。

	for _, f := range r.File {

		// 这个循环将一直运行,
		// 直到源目录中有文件为止,
		// 并将继续存储文件名,
		// 然后进行提取到目标文件夹,直到出现错误为止

		// 存储“路径/文件名”以供稍后返回和使用
		fpath := filepath.Join(destination, f.Name)

		// 检查是否有任何无效的文件路径
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {

			return filenames, fmt.Errorf("％s是非法文件路径", fpath)

		}

		// 现在访问的文件名被附加
		// 在filenames字符串数组中与其路径

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// 创建一个新文件夹
			os.MkdirAll(fpath, os.ModePerm)
			continue

		}

		// 在目标目录中创建文件
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {

			return filenames, err

		}

		// 创建的文件将被存储在
		// 带有写入和/或截断权限的outFile中

		outFile, err := os.OpenFile(fpath,

			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,

			f.Mode())

		// 如果有任何错误,此块将再次执行
		// 过程将返回到main函数
		if err != nil {

			// 用现在收集到的文件名
			// 和err消息

			return filenames, err

		}

		rc, err := f.Open()

		// 再次如果有任何错误,任何错误会在此代码块中执行并返回到主函数
		if err != nil {
			// 将收集到的文件名和错误信息返回给主函数
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// 没有使用defer关闭文件,以便它在循环
		// 进入下一个迭代之前关闭outfile。这样
		// 在最坏情况下节省了迭代的内存和时间
		outFile.Close()
		rc.Close()

		// 再次,如果有任何错误,会在此代码块中执行并返回到主函数
		if err != nil {
			// 将收集到的文件名和错误信息返回给主函数
			return filenames, err
		}
	}

	return filenames, nil
}

// src 压缩包文件路径    destination  输出的目录
func UnCompressPkg(src string, destination string) error {

	//获取文件名带后缀
	filenameWithSuffix := path.Base(src)
	fileSuffix := path.Ext(filenameWithSuffix)

	if fileSuffix == ".rar" {
		r := archiver.NewRar()
		r.OverwriteExisting = true
		r.ContinueOnError = true

		err := r.Unarchive(src, destination)
		if err != nil {
			return err
		}
	} else if fileSuffix == ".zip" {

		_, err := Unzip(src, destination)
		if err != nil {
			return err
		}
	} else if fileSuffix == ".7z" {
		err := extract7z(src, destination)
		if err != nil {
			return err
		}

	} else {
		return errors.New("文件格式错误")
	}
	return nil
}

// 解压7z文件到指定目录
func extract7z(src, dest string) error {
	cmd := exec.Command("7z", "x", src, "-o"+dest) // 确保系统中已安装7z命令行工具，并将其添加到系统路径中
	return cmd.Run()
}
func UploadFile(c *gin.Context, file *multipart.FileHeader, dst string) error {

	con := LoadConfig()

	toolConfig := new(ToolConfig)
	Db.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)

	if toolConfig.StaticResourcesType == 1 {
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "失败" + err.Error(),
				"data": "{}",
			})
			return err
		}
	} else if toolConfig.StaticResourcesType != 1 { // oss操作
		endpoint := con.ENDPOINT
		// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
		accessKeyId := con.ACCESS_KEY_ID
		accessKeySecret := con.ACCESS_KEY_SECRET
		// yourBucketName填写存储空间名称。
		bucketName := con.BACKET_NAME
		// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。

		// err = c.SaveUploadedFile(file, dst)
		// if err != nil {
		// 	return err
		// }
		// // yourLocalFileName填写本地文件的完整路径。
		// //localFileName := toolConfig.LocalUrl + dst
		// 创建OSSClient实例。
		client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
		if err != nil {
			return err
		}
		// 获取存储空间。
		bucket, err := client.Bucket(bucketName)
		if err != nil {
			return err
		}
		// 读取文件内容
		multipartFile, err := file.Open()
		if err != nil {
			return err
		}
		// 上传文件。
		err = bucket.PutObject(dst, multipartFile)
		if err != nil {
			return err
		}

		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			return err
		}
	}

	return nil

}

func RemoveFile(filepath string) error {
	con := LoadConfig()

	toolConfig := new(ToolConfig)
	Db.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)

	if toolConfig.StaticResourcesType == 1 {
		err := os.Remove(filepath)
		if err != nil {
			return err
		}
	} else if toolConfig.StaticResourcesType == 3 {

		endpoint := con.ENDPOINT
		// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
		accessKeyId := con.ACCESS_KEY_ID
		accessKeySecret := con.ACCESS_KEY_SECRET
		// yourBucketName填写存储空间名称。
		bucketName := con.BACKET_NAME

		// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。

		// 创建OSSClient实例。
		client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		// 获取存储空间。
		bucket, err := client.Bucket(bucketName)
		if err != nil {
			return err
		}
		// 删除文件。
		err = bucket.DeleteObject(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}

func DelRemoveAll(filepath string, targetPath string) error {

	con := LoadConfig()

	toolConfig := new(ToolConfig)
	Db.QueryRow("select StaticResourcesType,OSSConfig,CDNConfig,LocalUrl from tool where Id=1").Scan(&toolConfig.StaticResourcesType, &toolConfig.OSSConfig, &toolConfig.CDNConfig, &toolConfig.LocalUrl)

	if toolConfig.StaticResourcesType == 1 {

		if strings.Contains(path.Dir(filepath), targetPath) && path.Dir(filepath) != "" && path.Dir(filepath) != "/" {
			err := os.RemoveAll(path.Dir(filepath) + "/") //删除转换后的文件夹
			if err != nil {
				return err
			} else {
				fmt.Println("Folder deleted successfully!")
				return nil
			}
		}
	} else if toolConfig.StaticResourcesType == 3 {
		endpoint := con.ENDPOINT
		// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
		accessKeyId := con.ACCESS_KEY_ID
		accessKeySecret := con.ACCESS_KEY_SECRET
		// yourBucketName填写存储空间名称。
		bucketName := con.BACKET_NAME

		// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。

		// 创建OSSClient实例。
		client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		// 填写Bucket名称，例如examplebucket。
		bucket, err := client.Bucket(bucketName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		marker := oss.Marker("")
		// 如果您需要删除所有前缀为src的文件，则prefix设置为src。设置为src后，所有前缀为src的非目录文件、src目录以及目录下的所有文件均会被删除。
		prefix := oss.Prefix(path.Dir(filepath) + "/")
		count := 0
		for {
			lor, err := bucket.ListObjects(marker, prefix)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}

			objects := []string{}
			for _, object := range lor.Objects {
				objects = append(objects, object.Key)
			}
			// 将oss.DeleteObjectsQuiet设置为true，表示不返回删除结果。
			delRes, err := bucket.DeleteObjects(objects, oss.DeleteObjectsQuiet(true))
			if err != nil {
				return err
			}

			if len(delRes.DeletedObjects) > 0 {
				fmt.Println("these objects deleted failure,", delRes.DeletedObjects)
			}

			count += len(objects)

			prefix = oss.Prefix(lor.Prefix)
			marker = oss.Marker(lor.NextMarker)
			if !lor.IsTruncated {
				break
			}
		}

	}
	return nil
}

func ArrayToString(arr []string) string {
	var result string
	result += "["
	for index, i := range arr { //遍历数组将元素追加成字符串

		if (len(arr) - 1) == index {
			result += "\"" + i + "\""
		} else {
			result += "\"" + i + "\"" + ","
		}
	}
	result += "]"
	return result
}

type ToolConfig struct {
	StaticResourcesType int    `json:"StaticResourcesType" db:"StaticResourcesType"`
	OSSConfig           string `json:"OSSConfig" db:"OSSConfig"`
	CDNConfig           string `json:"CDNConfig" db:"CDNConfig"`
	LocalUrl            string `json:"LocalUrl" db:"LocalUrl"`
}
