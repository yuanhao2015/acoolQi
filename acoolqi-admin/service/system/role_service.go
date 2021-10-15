package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"bytes"
	"github.com/yuanhao2015/acoolTools"
)

type RoleService struct {
	roleDao     system.RoleDao
	roleMenuDao system.RoleMenuDao
	userRoleDao system.UserRoleDao
}

// SelectRoleAll 查询所有角色
func (s RoleService) SelectRoleAll(query *req.RoleQuery) ([]*models.SysRole, int64) {
	if query == nil {
		all := s.roleDao.SelectRoleAll()
		return all, 0
	}
	return s.roleDao.SelectRoleList(query)
}

// SelectRoleListByUserId 根据用户id查询角色id集合
func (s RoleService) SelectRoleListByUserId(parseInt int64) *[]int64 {
	return s.roleDao.SelectRoleListByUserId(parseInt)
}

// GetRoleListByUserId 根据用户ID查询角色
func (s RoleService) GetRoleListByUserId(id int64) *[]models.SysRole {
	return s.roleDao.GetRoleListByUserId(id)
}

// FindList 分页查询角色数据
func (s RoleService) FindList(query req.RoleQuery) ([]*models.SysRole, int64) {
	return s.roleDao.SelectRoleList(&query)
}

// SelectRoleByRoleId 根据角色id查询角色数据
func (s RoleService) SelectRoleByRoleId(id int64) *models.SysRole {
	return s.roleDao.SelectRoleByRoleId(id)
}

// CheckRoleNameUnique 判断角色名城是否存在
func (s RoleService) CheckRoleNameUnique(role models.SysRole) int64 {
	return s.roleDao.CheckRoleNameUnique(role)
}

// CheckRoleKeyUnique 校验角色权限是否唯一
func (s RoleService) CheckRoleKeyUnique(role models.SysRole) int64 {
	return s.roleDao.CheckRoleKeyUnique(role)

}

// Add 添加角色数据
func (s RoleService) Add(role models.SysRole) int64 {
	role = s.roleDao.Add(role)
	return s.insertRoleMenu(role)
}

//添加角色菜单关系
func (s RoleService) insertRoleMenu(role models.SysRole) int64 {
	list := make([]models.SysRoleMenu, 0)
	for _, id := range role.MenuIds {
		menu := models.SysRoleMenu{
			RoleId: role.RoleId,
			MenuId: id,
		}
		list = append(list, menu)
	}
	return s.roleMenuDao.Add(list)
}

// Update 修改角色数据
func (s RoleService) Update(role models.SysRole) int64 {
	//删除菜单关联关系
	s.roleMenuDao.Delete(role)
	s.insertRoleMenu(role)
	//修改数据
	return s.roleDao.Update(role)
}

// Delete 删除角色
func (s RoleService) Delete(id int64) int64 {
	role := models.SysRole{
		RoleId: id,
	}
	//删除菜单角色关系
	s.roleMenuDao.Delete(role)
	//删除角色
	return s.roleDao.Delete(role)
}

// CheckRoleAllowed 校验是否可以操作
func (s RoleService) CheckRoleAllowed(id int64) (bool, string) {
	if id == 1 {
		return false, "超级管理员不允许操作"
	}
	return true, ""
}

// UpdateRoleStatus 角色状态修改
func (s RoleService) UpdateRoleStatus(role *models.SysRole) int64 {
	return s.roleDao.UpdateRoleStatus(role)
}

// DeleteAuthUser 取消授权用户
func (s RoleService) DeleteAuthUser(userRole models.SysUserRole) int64 {
	return s.userRoleDao.DeleteAuthUser(userRole)
}

// InsertAuthUsers 批量选择用户授权
func (s RoleService) InsertAuthUsers(body req.UserRoleBody) int64 {
	return s.userRoleDao.InsertAuthUsers(body)
}

// SelectRolesByUserName 查询所属角色组
func (s RoleService) SelectRolesByUserName(name string) string {
	list := s.roleDao.SelectRolesByUserName(name)
	var buffer bytes.Buffer
	var roleName string
	for _, role := range *list {
		buffer.WriteString(role.RoleName)
		buffer.WriteString(",")
	}
	s2 := buffer.String()
	if acoolTools.StrUtils.HasNotEmpty(s2) {
		roleName = s2[0:(len(s2) - 1)]
	}
	return roleName
}
