package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"github.com/yuanhao2015/acoolTools"
)

type UserRoleDao struct {
}

// BatchUserRole 批量新增用户角色信息
func (d UserRoleDao) BatchUserRole(roles []models.SysUserRole) {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Table(models.SysUserRole{}.TableName()).Insert(&roles)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// RemoveUserRole 删除用户和角色关系
func (d UserRoleDao) RemoveUserRole(id int64) {
	role := models.SysUserRole{
		UserId: id,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&role)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// DeleteAuthUser 取消用户授权
func (d UserRoleDao) DeleteAuthUser(role models.SysUserRole) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	i, err := session.Delete(&role)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return i
}

// InsertAuthUsers 批量选择用户授权
func (d UserRoleDao) InsertAuthUsers(body req.UserRoleBody) int64 {
	ids := body.UserIds
	roles := make([]models.SysUserRole, 0)
	for i := 0; i < len(ids); i++ {
		role := models.SysUserRole{
			RoleId: body.RoleId,
			UserId: ids[i],
		}
		roles = append(roles, role)
	}
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&roles)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}
