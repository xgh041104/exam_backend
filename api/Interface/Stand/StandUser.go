package Stand

import (
	model "StudyExamPlatformAPI/Model"
	"StudyExamPlatformAPI/jwt_use"
	"StudyExamPlatformAPI/lib"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginStandUser(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}
	log.Printf("%s", body)

	var standuser model.StandUser
	json.Unmarshal([]byte(body), &standuser)

	newstanduser := new(model.StandUser)

	err = lib.Db.QueryRow("select StandUserId,StandUserName,StandUserAccount,StandId from standuser  where StandUserAccount=? and StandUserPwd=?", standuser.StandUserAccount, standuser.StandUserPwd).
		Scan(&newstanduser.StandUserId, &newstanduser.StandUserName, &newstanduser.StandUserAccount, &newstanduser.StandId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "账号或密码错误",
			"data": "{}",
		})
		return
	}

	if newstanduser.StandUserId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "账号或密码错误",
			"data": "{}",
		})
		return
	} else {
		tokenString, _ := jwt_use.GetToken(newstanduser.StandUserName, 2)
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "登录成功",
			"data":  newstanduser,
			"token": tokenString,
		})
		return
	}
}
