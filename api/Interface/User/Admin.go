package User

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/jwt_use"
	"StudyExamPlatformAPI/lib"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginAdmin(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}
	log.Printf("%s", body)

	var admin model.Admin
	json.Unmarshal([]byte(body), &admin)

	data := []byte(admin.AdminPassword)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	newadmin := new(model.Admin)

	lib.Db.QueryRow("SELECT Id,AdminName,AdminAccount FROM  admin where AdminAccount=? and AdminPassword=?", admin.AdminAccount, md5str1).
		Scan(&newadmin.Id, &newadmin.AdminName, &newadmin.AdminAccount)
	fmt.Println(err)
	if newadmin.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "账号或密码错误",
			"data": "{}",
		})
		return
	} else {
		tokenString, _ := jwt_use.GetToken(newadmin.AdminName, 0)
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "登录成功",
			"data":  newadmin,
			"token": tokenString,
		})
		return
	}
}
