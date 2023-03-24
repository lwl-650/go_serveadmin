package apis

import (
	"fmt"
	"go_serveadmin/model"
	"go_serveadmin/util"

	"github.com/gin-gonic/gin"
)

type UserController struct {
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
