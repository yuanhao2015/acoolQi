package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/response"
	"github.com/yuanhao2015/acoolTools"
)

type PermissionService struct {
	roleDao system.RoleDao
	menuDao system.MenuDao
}

// GetRolePermissionByUserId 查询用户角色集合
func (s PermissionService) GetRolePermissionByUserId(user *response.UserResponse) *[]string {
	admin := models.SysUser{}.IsAdmin(user.UserId)
	roleKeys := s.roleDao.GetRolePermissionByUserId(user.UserId)
	if admin && roleKeys != nil {
		*roleKeys = append(*roleKeys, "admin")
	}
	duplication := acoolTools.StrArrayUtils.ArrayDuplication(*roleKeys)
	return &duplication
}

// GetMenuPermission 获取菜单数据权限
func (s PermissionService) GetMenuPermission(user *response.UserResponse) *[]string {
	flag := models.SysUser{}.IsAdmin(user.UserId)
	//查询菜单数据权限
	permission := s.menuDao.GetMenuPermission(user.UserId)
	if flag && permission != nil {
		*permission = append(*permission, "*:*:*")
	}
	var ret []string
	duplication := acoolTools.StrArrayUtils.ArrayDuplication(*permission)
	for i := 0; i < len(duplication); i++ {
		if (i > 0 && duplication[i-1] == duplication[i]) || len(duplication[i]) == 0 {
			continue
		}
		ret = append(ret, duplication[i])
	}
	return &ret
}
