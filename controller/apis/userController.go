package apis

import (
	"fmt"
	"go_serveadmin/model"
	"go_serveadmin/util"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (UserController) LoginUser(c *gin.Context) {
	currentTime := time.Now().Unix()

	// formatTimeStr := time.Unix(currentTime, 0).Format("2006-01-02 15:04:05")
	// i := currentTime + 1800
	// get := time.Unix(i, 0).Format("2006-01-02 15:04:05")
	// fmt.Println(formatTimeStr, "===================", get)

	auser := make(map[string]interface{})
	user := model.User{}
	name := c.PostForm("name")
	password := c.PostForm("password")
	fmt.Println(user.Password == password, user.Name == name)
	if name != "" && password != "" {
		util.DB.Where("name=?", name).First(&user)
		if user.Password == password {
			auser["aname"] = user.Name
			auser["apass"] = user.Password
			auser["timer"] = currentTime + 20
			token, _ := util.GenerateToken(auser)
			user.Token = token
			util.DB.Model(&user).Update("token", token)
			util.Success(c, user)
		} else {
			util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
		}
	} else {
		util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
	}
}

func (UserController) TokenVerification(c *gin.Context) {
	mapss := make(map[string]interface{})
	currentTime := time.Now().Unix()
	fmt.Println(c.Request.Header.Get("Authorization"))
	token := c.Request.Header.Get("Authorization")
	// token = token[6:]
	fmt.Println(token, "------------------------->")
	fmt.Println(util.ConfirmToken(token, mapss))
	auser, _ := util.ConfirmToken(token, mapss)
	getTokenTimer := int64(auser["timer"].(float64))

	fmt.Println("现在时间=》", currentTime, "====================?>token时间", getTokenTimer) // 输出: 123456
	if currentTime < getTokenTimer {

		fmt.Println("有效")
		util.Success(c, auser)
	} else {
		// fmt.Println("过期")
		util.Error(c, -1, "token过期")
	}
}

func (UserController) SetUser(c *gin.Context) {

	name := c.PostForm("name")
	id := util.Snow()
	password := c.PostForm("password")
	if name != "" {
		user := model.User{Name: name, Id: id, Password: password}
		if err := util.DB.Create(&user).Error; err != nil {
			fmt.Println("插入失败", err)
			util.Error(c, -1, "添加失败")
			return
		} else {
			util.Success(c, err)
		}
	} else {
		util.Error(c, -1, "不能为空。")
	}

}
