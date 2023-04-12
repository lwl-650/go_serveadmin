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
	aname := c.PostForm("aname")
	apass := c.PostForm("apass")
	if aname != "" && apass != "" {
		util.DB.Where("name=?", aname).First(&user)
		if user.Pass == apass {
			auser["aname"] = user.Name
			auser["apass"] = user.Pass
			token, _ := util.GenerateToken(auser)
			user.Token = token
			util.Success(c, user)
		} else {
			util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
		}
	} else {
		util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
	}
}

func (UserController) SetUser(c *gin.Context) {

	name := c.PostForm("name")
	id := util.Snow()
	// password := c.PostForm("password")
	// avatar := c.PostForm("avatar")
	// gender := c.PostForm("gender")
	// city := c.PostForm("city")
	fmt.Println(name)

	if name != "" {
		user := model.User{Name: name, Id: id}
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
