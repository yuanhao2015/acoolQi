package system

import (
	"acoolqi-admin/models"
	"github.com/yuanhao2015/acoolTools"
)

type RoleMenuDao struct {
}

// Add 添加角色菜单关系
func (d RoleMenuDao) Add(list []models.SysRoleMenu) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&list)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
	}
	session.Commit()
	return insert
}

// Delete 删除角色和菜单关系
func (d RoleMenuDao) Delete(role models.SysRole) {
	menu := models.SysRoleMenu{
		RoleId: role.RoleId,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&menu)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
	}
	session.Commit()
}
