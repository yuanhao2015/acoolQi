package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/page"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/yuanhao2015/acoolTools"
)

type RoleDao struct {
}

//角色公用sql
func (d RoleDao) sqlSelectJoin() *xorm.Session {
	return SqlDB.Table([]string{models.SysRole{}.TableName(), "r"}).
		Join("LEFT", []string{"sys_user_role", "ur"}, "ur.role_id = r.role_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = ur.user_id").
		Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
}

//用户角色关系查询sql
func (d RoleDao) sqlSelectRoleAndUser() *xorm.Session {
	return SqlDB.Table([]string{models.SysRole{}.TableName(), "r"}).
		Join("LEFT", []string{"sys_user_role", "ur"}, "ur.role_id = r.role_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = ur.user_id")
}

// SelectRoleList 根据条件查询角色数据
func (d RoleDao) SelectRoleList(q *req.RoleQuery) ([]*models.SysRole, int64) {
	roles := make([]*models.SysRole, 0)
	session := d.sqlSelectJoin()
	if !acoolTools.StrUtils.HasEmpty(q.RoleName) {
		session.And("r.role_name like concat('%', ?, '%')", q.RoleName)
	}
	if !acoolTools.StrUtils.HasEmpty(q.Status) {
		session.And("r.status = ?", q.Status)
	}
	if !acoolTools.StrUtils.HasEmpty(q.RoleKey) {
		session.And("r.role_key like concat('%', ?, '%')", q.RoleKey)
	}
	if !acoolTools.StrUtils.HasEmpty(q.BeginTime) {
		session.And("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", q.BeginTime)
	}
	if !acoolTools.StrUtils.HasEmpty(q.EndTime) {
		session.And("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", q.EndTime)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(q.PageSize, page.StartSize(q.PageNum, q.PageSize)).OrderBy("r.role_sort").Find(&roles)
	if err != nil {
		return nil, 0
	}
	return roles, total
}

// SelectRoleAll 查询所有角色
func (d RoleDao) SelectRoleAll() []*models.SysRole {
	sql := d.sqlSelectJoin()
	roles := make([]*models.SysRole, 0)
	err := sql.Find(&roles)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return roles
}

// SelectRoleListByUserId 根据用户id查询用户角色id集合
func (d RoleDao) SelectRoleListByUserId(userId int64) *[]int64 {
	sqlSelectRoleAndUser := d.sqlSelectRoleAndUser()
	var roleIds []int64
	err := sqlSelectRoleAndUser.Cols("r.role_id").Where("u.user_id = ?", userId).Find(&roleIds)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &roleIds
}

// GetRolePermissionByUserId 查询用户角色集合
func (d RoleDao) GetRolePermissionByUserId(id int64) *[]string {
	var roleKeys []string
	err := d.sqlSelectJoin().Cols("r.role_key").Where("r.del_flag = '0'").And("ur.user_id = ?", id).Find(&roleKeys)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &roleKeys
}

// GetRoleListByUserId 根据用户ID查询角色
func (d RoleDao) GetRoleListByUserId(id int64) *[]models.SysRole {
	roles := make([]models.SysRole, 0)
	err := d.sqlSelectJoin().Where("r.del_flag = '0'").And("ur.user_id = ?", id).Find(&roles)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &roles
}

// SelectRoleByRoleId 根据角色id查询角色数据
func (d RoleDao) SelectRoleByRoleId(id int64) *models.SysRole {
	role := models.SysRole{}
	_, err := d.sqlSelectJoin().Where("r.role_id = ?", id).Get(&role)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &role
}

// CheckRoleNameUnique 校验角色名称是否唯一
func (d RoleDao) CheckRoleNameUnique(role models.SysRole) int64 {
	session := SqlDB.Table(role.TableName()).Where("role_name = ?", role.RoleName)
	if role.RoleId > 0 {
		session.And("role_id != ?", role.RoleId)
	}
	count, err := session.Count(&role)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
	}
	return count
}

// CheckRoleKeyUnique 校验角色权限是否唯一
func (d RoleDao) CheckRoleKeyUnique(role models.SysRole) int64 {
	session := SqlDB.Table(role.TableName()).Where("role_key = ?", role.RoleKey)
	if role.RoleId > 0 {
		session.And("role_id != ?", role.RoleId)
	}
	count, err := session.Count(&role)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
	}
	return count
}

// Add 添加角色进入数据库操作
func (d RoleDao) Add(role models.SysRole) models.SysRole {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Insert(&role)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
	}
	session.Commit()
	return role
}

// Update 修改数据
func (d RoleDao) Update(role models.SysRole) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("role_id = ?", role.RoleId).Update(&role)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return -1
	}
	session.Commit()
	fmt.Println(update)
	return update
}

// Delete 删除角色
func (d RoleDao) Delete(role models.SysRole) int64 {
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

// UpdateRoleStatus 修改角色状态
func (d RoleDao) UpdateRoleStatus(role *models.SysRole) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("role_id = ?", role.RoleId).Cols("status", "update_by", "update_time").Update(role)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return update
}

// SelectRolesByUserName 查询角色组
func (d RoleDao) SelectRolesByUserName(name string) *[]models.SysRole {
	roles := make([]models.SysRole, 0)
	session := d.sqlSelectJoin()
	err := session.Where("r.del_flag = '0'").And("u.user_name = ?", name).Find(&roles)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &roles
}
