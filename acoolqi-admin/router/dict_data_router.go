package router

import (
	"acoolqi-admin/api/v1/system"
	"github.com/gin-gonic/gin"
)

//初始化字典数据路由
func initDictDataRouter(router *gin.RouterGroup) {
	v := new(system.DictDataApi)
	group := router.Group("/system/dict/data")
	{
		//根据字典类型查询字典数据信息
		group.GET("/type/:dictType", v.GetByType)
		//查询字典数据集合
		group.GET("/list", v.List)
		//根据id查询字典数据
		group.GET("/:dictCode", v.Get)
		//添加字段数据
		group.POST("/add", v.Add)
		//修改字段数据
		group.PUT("/edit", v.Edit)
		//删除字典数据
		group.DELETE("/remove/:dictCode", v.Delete)
		//导出字典
		group.GET("/export", v.Export)
	}
}
