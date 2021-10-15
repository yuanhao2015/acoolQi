package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
)

type DeptService struct {
	deptDao system.DeptDao
	roleDao system.RoleDao
}

// TreeSelect 根据条件查询部门树
func (s DeptService) TreeSelect(query req.DeptQuery) *[]models.SysDept {
	treeSelect := s.deptDao.TreeSelect(query)
	return treeSelect
}

// SelectDeptListByRoleId 根据角色ID查询部门树信息
func (s DeptService) SelectDeptListByRoleId(id int64) *[]int64 {
	role := s.roleDao.SelectRoleByRoleId(id)
	return s.deptDao.SelectDeptListByRoleId(id, role.DeptCheckStrictly)
}

// GetList 查询部门列表
func (s DeptService) GetList(query req.DeptQuery) *[]models.SysDept {
	return s.deptDao.GetList(query)
}

// GetDeptById 根据部门编号获取详细信息
func (s DeptService) GetDeptById(id int) *models.SysDept {
	return s.deptDao.GetDeptById(id)
}

// Insert 添加部门数据
func (s DeptService) Insert(dept models.SysDept) int64 {
	return s.deptDao.Insert(dept)
}

// CheckDeptNameUnique 校验部门名称是否唯一
func (s DeptService) CheckDeptNameUnique(dept models.SysDept) bool {
	if s.deptDao.CheckDeptNameUnique(dept) > 0 {
		return true
	}
	return false
}

// Remove 删除部门
func (s DeptService) Remove(id int) int64 {
	return s.deptDao.Delete(id)
}

// HasChildByDeptId 是否存在部门子节点
func (s DeptService) HasChildByDeptId(id int) int64 {
	return s.deptDao.HasChildByDeptId(id)
}

// CheckDeptExistUser 查询部门是否存在用户
func (s DeptService) CheckDeptExistUser(id int) int64 {
	return s.deptDao.CheckDeptExistUser(id)
}

// Update 修改部门
func (s DeptService) Update(menu models.SysDept) int64 {
	return s.deptDao.Update(menu)
}
