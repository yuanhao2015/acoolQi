package router

import (
	"acoolqi-admin/api/v1/system"
	"github.com/gin-gonic/gin"
)

//部门接口操作
func initDeptRouter(router *gin.RouterGroup) {
	v := new(system.DeptApi)
	group := router.Group("/system/dept")
	{
		//获取部门下拉树列表
		group.GET("/treeselect", v.TreeSelect)
		//加载对应角色部门列表树
		group.GET("/roleDeptTreeselect/:roleId", v.RoleDeptTreeSelect)
		//查询部门列表
		group.GET("/list", v.Find)
		//查询部门列表（排除节点）
		group.GET("/list/exclude/:deptId", v.ExcludeChild)
		//根据部门编号获取详细信息
		group.GET("/:deptId", v.GetInfo)
		//添加部门
		group.POST("/add", v.Add)
		//修改部门
		group.PUT("/edit", v.Edit)
		//删除部门
		group.DELETE("/remove/:deptId", v.Delete)
	}
}
