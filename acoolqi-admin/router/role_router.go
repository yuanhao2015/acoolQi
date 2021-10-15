package router

import (
	"acoolqi-admin/api/v1/system"
	"github.com/gin-gonic/gin"
)

func initRoleRouter(router *gin.RouterGroup) {
	roleApi := new(system.RoleApi)
	group := router.Group("/system/role")
	{
		//查询角色数据
		group.GET("/findList", roleApi.Find)
		//根据id查询角色数据
		group.GET("/:roleId", roleApi.GetRoleId)
		//添加角色
		group.POST("/add", roleApi.Add)
		//修改角色
		group.PUT("/edit", roleApi.Edit)
		//删除角色
		group.DELETE("/remove/:roleId", roleApi.Delete)
		//修改角色状态
		group.PUT("/changeStatus", roleApi.ChangeStatus)
		//查看分配角色列表
		group.GET("/authUser/allocatedList", roleApi.AllocatedList)
		//查询未分配用户角色列表
		group.GET("/authUser/unallocatedList", roleApi.UnallocatedList)
		//取消授权
		group.PUT("/authUser/cancel", roleApi.CancelAuthUser)
		//批量选择用户授权
		group.PUT("/authUser/selectAll", roleApi.UpdateAuthUserAll)
		//导出excel
		group.GET("/export", roleApi.Export)
	}
}
