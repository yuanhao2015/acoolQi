package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"github.com/yuanhao2015/acoolTools"
)

type MenuDao struct {
}

// GetMenuPermission 根据用户ID查询权限
func (d MenuDao) GetMenuPermission(id int64) *[]string {
	var perms []string
	session := SqlDB.Table([]string{"sys_menu", "m"})
	err := session.Distinct("m.perms").
		Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id").
		Join("LEFT", []string{"sys_role", "r"}, "r.role_id = ur.role_id").
		Where("m.status = '0'").And("r.status = '0'").And("ur.user_id = ?", id).Find(&perms)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &perms
}

// GetMenuAll 查询所有菜单数据
func (d MenuDao) GetMenuAll() *[]models.SysMenu {
	menus := make([]models.SysMenu, 0)
	session := SqlDB.Table([]string{models.SysMenu{}.TableName(), "m"})
	err := session.Distinct("m.menu_id").Cols("m.parent_id", "m.menu_name", "m.path", "m.component", "m.visible", "m.status", "m.perms", "m.is_frame", "m.is_cache", "m.menu_type", "m.icon", "m.order_num", "m.create_time").
		Where("m.menu_type in ('M', 'C','F')").And("m.status = 0").OrderBy("m.parent_id").OrderBy("m.order_num").Find(&menus)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &menus
}

// GetMenuByUserId 根据用户ID查询菜单
func (d MenuDao) GetMenuByUserId(id int64) *[]models.SysMenu {
	menus := make([]models.SysMenu, 0)
	session := SqlDB.Table([]string{models.SysMenu{}.TableName(), "m"})
	err := session.Distinct("m.menu_id").Cols("m.parent_id", "m.menu_name", "m.path", "m.component", "m.visible", "m.status", "m.perms", "m.is_frame", "m.is_cache", "m.menu_type", "m.icon", "m.order_num", "m.create_time").
		Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id").
		Join("LEFT", []string{"sys_role", "ro"}, "ur.role_id = ro.role_id").
		Join("LEFT", []string{"sys_user", "u"}, "ur.user_id = u.user_id").Where("u.user_id = ?", id).
		And("m.menu_type in ('M', 'C','F')").And("m.status = 0").OrderBy("m.parent_id").OrderBy("m.order_num").Find(&menus)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &menus
}

// GetMenuByRoleId 根据角色ID查询菜单树信息
func (d MenuDao) GetMenuByRoleId(id int64, strictly bool) *[]int64 {
	list := make([]int64, 0)
	session := SqlDB.NewSession().Table([]string{"sys_menu", "m"})
	session.Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id")
	session.Where("rm.role_id = ?", id)
	if strictly {
		session.And("m.menu_id not in (select m.parent_id from sys_menu m inner join sys_role_menu rm on m.menu_id = rm.menu_id and rm.role_id = ?)", id)
	}
	err := session.OrderBy("m.parent_id").OrderBy("m.order_num").Cols("m.menu_id").Find(&list)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &list
}

// GetMenuList 查询系统菜单列表
func (d MenuDao) GetMenuList(query req.MenuQuery) *[]models.SysMenu {
	list := make([]models.SysMenu, 0)
	session := SqlDB.NewSession().OrderBy("parent_id").OrderBy("order_num")
	if acoolTools.StrUtils.HasNotEmpty(query.MenuName) {
		session.And("menu_name like concat('%', ?, '%')", query.MenuName)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.Visible) {
		session.And("visible = ?", query.Visible)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	err := session.Find(&list)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &list
}

// GetMenuListByUserId 根据用户查询系统菜单列表
func (d MenuDao) GetMenuListByUserId(query req.MenuQuery) *[]models.SysMenu {
	session := SqlDB.NewSession().OrderBy("parent_id").OrderBy("order_num")
	list := make([]models.SysMenu, 0)
	session.Distinct("m.menu_id", "m.parent_id", "m.menu_name", "m.path", "m.component", "m.visible", "m.status", "ifnull(m.perms,'') as perms", "m.is_frame", "m.is_cache", "m.menu_type", "m.icon", "m.order_num", "m.create_time")
	session.Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id").
		Join("LEFT", []string{"sys_role", "ro"}, "ur.role_id = ro.role_id").
		Where("ur.user_id = ?", query.UserId)
	if acoolTools.StrUtils.HasNotEmpty(query.MenuName) {
		session.And("menu_name like concat('%', ?, '%')", query.MenuName)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.Visible) {
		session.And("visible = ?", query.Visible)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	err := session.Find(&list)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &list
}

// GetMenuByMenuId 根据菜单ID查询信息
func (d MenuDao) GetMenuByMenuId(id int) *models.SysMenu {
	menu := models.SysMenu{
		MenuId: id,
	}
	_, err := SqlDB.NewSession().Where("menu_id = ?", menu.MenuId).Get(&menu)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &menu
}

// GetMenuPerms 根据权限标识查询信息
func (d MenuDao) GetMenuPerms(perms string) *models.SysMenu {
	menu := models.SysMenu{
		Perms: perms,
	}
	_, err := SqlDB.NewSession().Where("perms = ?", menu.Perms).Get(&menu)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &menu
}

// Insert 添加菜单数据
func (d MenuDao) Insert(menu models.SysMenu) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&menu)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}

// Update 修改菜单数据
func (d MenuDao) Update(menu models.SysMenu) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("menu_id = ?", menu.MenuId).Update(&menu)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return update
}

// Delete 删除菜单操作
func (d MenuDao) Delete(id int) int64 {
	menu := models.SysMenu{
		MenuId: id,
	}
	session := SqlDB.NewSession()
	session.Begin()
	i, err := session.Delete(&menu)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return i
}
