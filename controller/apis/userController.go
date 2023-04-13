package apis

import (
	"fmt"
	"go_serveadmin/model"
	"go_serveadmin/util"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (UserController) LoginUser(c *gin.Context) {
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
	fmt.Println(c.Request.Header.Get("Authorization"))
	token := c.Request.Header.Get("Authorization")
	// token = token[6:]
	fmt.Println(token, "------------------------->")
	// fmt.Println(util.ConfirmToken(token, mapss))
	auser, _ := util.ConfirmToken(token, mapss)
	fmt.Println(auser, "------------------------->")
	util.Success(c, auser)
	// util.Success(c, "auser")
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
