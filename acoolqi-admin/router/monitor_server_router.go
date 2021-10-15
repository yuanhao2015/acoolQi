/*
@Time : 2021-10-14 11:50
@Author : acool
@File : monitor_server_router
*/
package router

import (
	"acoolqi-admin/api/v1/monitor"
	"github.com/gin-gonic/gin"
)

func initMonitorServerRouter(router *gin.RouterGroup) {
	v := new(monitor.SysMonitorApi)
	vg := router.Group("/monitor")
	{
		//查询菜单数据
		vg.GET("/server", v.Info)

	}
}
