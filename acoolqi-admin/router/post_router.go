package router

import (
	"acoolqi-admin/api/v1/system"
	"github.com/gin-gonic/gin"
)

//初始化岗位路由
func initPostRouter(router *gin.RouterGroup) {
	v := new(system.PostApi)
	group := router.Group("/system/post")
	{
		//查询岗位数据
		group.GET("/list", v.List)
		//添加岗位
		group.POST("/add", v.Add)
		//查询岗位详情
		group.GET("/:postId", v.Get)
		//删除岗位数据
		group.DELETE("/remove/:postId", v.Delete)
		//修改岗位数据接口
		group.PUT("/edit", v.Edit)
		//导出excel
		group.GET("/export", v.Export)
	}
}
