/*
@Time : 2021-10-14 15:43
@Author : acool
@File : user_online_router
*/
package router

import (
	"acoolqi-admin/api/v1/monitor"
	"github.com/gin-gonic/gin"
)

func initUserOnlineRouter(router *gin.RouterGroup) {
	v := new(monitor.SysUserOnlineApi)
	vg := router.Group("/monitor/online")
	{
		//查询数据
		vg.GET("/list", v.List)
		//强退用户
		vg.DELETE("/remove/:sessionIds", v.ForceLogout)
	}
}
