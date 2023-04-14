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
	token := c.Request.Header.Get("Authorization")
	// token = token[6:]
	auser, _ := util.ConfirmToken(token, mapss)
	getTokenTimer := int64(auser["timer"].(float64))
	if currentTime < getTokenTimer {
		util.Success(c, "登录成功")
	} else {
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
