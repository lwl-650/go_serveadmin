package router

import (
	"go_serveadmin/controller/apis"

	"github.com/gin-gonic/gin"
)

func IndexRoutersInit(r *gin.Engine) {
	index := r.Group("/")
	{
		// index.Use(Content) //路由分组中间件
		index.GET("/find", apis.AdminController{}.Findthis)
		index.POST("/setUser", apis.UserController{}.SetUser)
	}

}
