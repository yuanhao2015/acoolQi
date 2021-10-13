package router

import (
	"github.com/gin-gonic/gin"
	v1 "acoolqi-admin/api/v1"
)

//登录调用路由
func initLoginRouter(router *gin.RouterGroup) {
	loginApi := new(v1.LoginApi)
	loginRouter := router.Group("/")
	{
		//登录
		loginRouter.POST("/login", loginApi.Login)
		loginRouter.GET("/getInfo", loginApi.GetUserInfo)
		loginRouter.GET("/getRouters", loginApi.GetRouters)
		//退出登录
		loginRouter.POST("/logout",loginApi.Logout)
	}
}
