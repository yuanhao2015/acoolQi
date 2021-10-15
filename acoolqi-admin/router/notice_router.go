package router

import (
	"acoolqi-admin/api/v1/system"
	"github.com/gin-gonic/gin"
)

func initNoticeRouter(router *gin.RouterGroup) {
	v := new(system.NoticeApi)
	group := router.Group("/system/notice")
	{
		group.GET("/list", v.List)
		//添加公告
		group.POST("/add", v.Add)
		//删除
		group.DELETE("/remove/:ids", v.Delete)
		//查询
		group.GET("/:id", v.Get)
		//修改
		group.PUT("/edit", v.Edit)
	}
}
