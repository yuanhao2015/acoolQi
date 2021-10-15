/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: operlog_router
* @Date: 2021-9-28 14:51
 */
package router

import (
	"acoolqi-admin/api/v1/system"
	"github.com/gin-gonic/gin"
)

func initOperlogRouter(router *gin.RouterGroup) {
	v := new(system.OperlogApi)
	group := router.Group("/monitor/operlog")
	{
		//列表
		group.GET("/list", v.List)
		//删除
		group.DELETE("/remove/:ids", v.Delete)
		//查询
		group.GET("/:id", v.Get)

	}
}
