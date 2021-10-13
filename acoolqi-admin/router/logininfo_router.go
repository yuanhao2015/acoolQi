/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: operlog_router
* @Date: 2021-9-28 14:51
 */
package router

import (
	v1 "acoolqi-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func initLogininfoRouter(router *gin.RouterGroup) {
	v := new(v1.LogininfoApi)
	group := router.Group("/monitor/logininfor")
	{
		group.GET("/list", v.List)

	}
}
