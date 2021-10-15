package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/models/response"
)

type MenuService struct {
	menuDao system.MenuDao
	roleDao system.RoleDao
}

// GetMenuTreeByUserId 根据用户ID查询菜单
func (s MenuService) GetMenuTreeByUserId(user *response.UserResponse) *[]models.SysMenu {
	var menuList *[]models.SysMenu
	//判断是否是管理员
	flag := models.SysUser{}.IsAdmin(user.UserId)
	if flag {
		menuList = s.menuDao.GetMenuAll()
	} else {
		menuList = s.menuDao.GetMenuByUserId(user.UserId)
	}
	return menuList
}

// SelectMenuListByRoleId 根据角色ID查询菜单树信息
func (s MenuService) SelectMenuListByRoleId(id int64) *[]int64 {
	role := s.roleDao.SelectRoleByRoleId(id)
	return s.menuDao.GetMenuByRoleId(id, role.MenuCheckStrictly)
}

// GetMenuList 获取菜单列表
func (s MenuService) GetMenuList(query req.MenuQuery, info *response.UserResponse) *[]models.SysMenu {
	if info.IsAdmin() {
		return s.menuDao.GetMenuList(query)
	} else {
		query.UserId = info.UserId
		return s.menuDao.GetMenuListByUserId(query)
	}
}

// GetMenuByMenuId 根据菜单ID查询信息
func (s MenuService) GetMenuByMenuId(id int) *models.SysMenu {
	return s.menuDao.GetMenuByMenuId(id)
}

// InsertMenu 添加菜单数据
func (s MenuService) InsertMenu(menu models.SysMenu) int64 {
	return s.menuDao.Insert(menu)
}

// Update 修改菜单数据
func (s MenuService) Update(menu models.SysMenu) int64 {
	return s.menuDao.Update(menu)
}

// 删除菜单操作
func (s MenuService) Delete(id int) int64 {
	return s.menuDao.Delete(id)
}

// 根据权限标识查询
func (s MenuService) GetMenuPerms(perms string) *models.SysMenu {
	return s.menuDao.GetMenuPerms(perms)
}
