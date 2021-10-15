package router

import (
	"acoolqi-admin/api/v1/system"
	"github.com/gin-gonic/gin"
)

//用户路由
func initUserRouter(router *gin.RouterGroup) {
	userApi := new(system.UserApi)
	userRouter := router.Group("/system/user")
	{
		userRouter.GET("/list", userApi.Find)
		userRouter.GET("/getInfo/:userId", userApi.GetInfo)
		userRouter.GET("/getInfo", userApi.GetInfo)
		userRouter.GET("/authRole/:userId", userApi.AuthRole)
		//新增用户
		userRouter.POST("/add", userApi.Add)
		//修改用户
		userRouter.PUT("/edit", userApi.Edit)
		//删除用户
		userRouter.DELETE("/remove/:userId", userApi.Remove)
		//重置密码
		userRouter.PUT("/resetPwd", userApi.ResetPwd)
		userRouter.GET("/export", userApi.Export)
		userRouter.GET("/profile", userApi.Profile)
		//修改个人数据
		userRouter.PUT("/profile", userApi.UpdateProfile)
		//修改个人密码
		userRouter.PUT("/profile/updatePwd", userApi.UpdatePwd)
		//修改头像
		userRouter.POST("/profile/avatar", userApi.Avatar)
	}
}
