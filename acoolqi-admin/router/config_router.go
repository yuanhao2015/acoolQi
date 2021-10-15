package router

import (
	"acoolqi-admin/api/v1/system"
	"github.com/gin-gonic/gin"
)

func initConfigRouter(router *gin.RouterGroup) {
	v := new(system.ConfigApi)
	group := router.Group("/system/config")
	{
		//根据参数键名查询参数值
		group.GET("/configKey/:configKey", v.GetConfigKey)
		//查询设置列表
		group.GET("/list", v.List)
		//添加
		group.POST("/add", v.Add)
		//查询
		group.GET("/:configId", v.Get)
		//修改
		group.PUT("/edit", v.Edit)
		//批量删除
		group.DELETE("/remove/:ids", v.Delete)
		//刷新缓存
		group.DELETE("/refreshCache", v.RefreshCache)
		//导出数据
		group.GET("/export", v.Export)
	}
}
