package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/library/tree/tree_menu"
	"acoolqi-admin/pkg/library/user_util"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/service/system"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type MenuApi struct {
	menuService system.MenuService
}

// List 查询菜单数据
func (a MenuApi) List(c *gin.Context) {
	//获取当前登录用户
	info := user_util.GetUserInfo(c)
	//获取参数
	query := req.MenuQuery{}
	if c.Bind(&query) != nil {
		resp.Error(c)
		return
	}
	resp.OK(c, a.menuService.GetMenuList(query, info))
}

// GetInfo 根据id查询菜单详情
func (a MenuApi) GetInfo(c *gin.Context) {
	param := c.Param("menuId")
	menuId, err := strconv.Atoi(param)
	if err != nil {
		resp.ParamError(c, "参数绑定错误")
		return
	}
	resp.OK(c, a.menuService.GetMenuByMenuId(menuId))
}

// RoleMenuTreeSelect 加载对应角色菜单列表树
func (a MenuApi) RoleMenuTreeSelect(c *gin.Context) {
	m := make(map[string]interface{})
	param := c.Param("roleId")
	roleId, _ := strconv.ParseInt(param, 10, 64)
	//获取当前登录用户
	info := user_util.GetUserInfo(c)
	menuList := a.menuService.GetMenuTreeByUserId(info)
	//fmt.Println(menuList)
	menus := tree_menu.SystemMenus{}
	tree := menus.GetTree(menuList)
	ids := a.menuService.SelectMenuListByRoleId(roleId)
	m["checkedKeys"] = ids
	m["menus"] = tree
	c.JSON(200, resp.Success(m))
}

// TreeSelect 获取菜单下拉树列表
func (a MenuApi) TreeSelect(c *gin.Context) {
	info := user_util.GetUserInfo(c)
	menus := a.menuService.GetMenuTreeByUserId(info)
	systemMenus := tree_menu.SystemMenus{}
	tree := systemMenus.GetTree(menus)
	c.JSON(200, resp.Success(tree))
}

// Add 添加菜单数据
func (a MenuApi) Add(c *gin.Context) {
	menu := models.SysMenu{}
	if c.Bind(&menu) != nil {
		resp.ParamError(c, "参数绑定异常")
		return
	}
	if a.menuService.InsertMenu(menu) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Edit 修改菜单数据
func (a MenuApi) Edit(c *gin.Context) {
	menu := models.SysMenu{}
	if c.Bind(&menu) != nil {
		resp.ParamError(c)
		return
	}
	menu.UpdateBy = user_util.GetUserInfo(c).UserName
	menu.UpdateTime = time.Now()
	if a.menuService.Update(menu) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除菜单
func (a MenuApi) Delete(c *gin.Context) {
	param := c.Param("menuId")
	menuId, err := strconv.Atoi(param)
	if err != nil {
		resp.ParamError(c)
		return
	}
	if a.menuService.Delete(menuId) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}
