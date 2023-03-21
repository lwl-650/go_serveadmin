package apis

import (
	"go_serveadmin/util"

	"go_serveadmin/model"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func (AdminController) Findthis(c *gin.Context) {
	user := []model.User{}
	// util.Success(c, user)

	if util.DB.Find(&user).Error == nil {
		util.Success(c, user)
	}

}
