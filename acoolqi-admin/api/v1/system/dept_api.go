package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/library/tree/tree_dept"
	"acoolqi-admin/pkg/library/user_util"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/service/system"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type DeptApi struct {
	deptService system.DeptService
}

// TreeSelect 查询部门菜单树
func (a DeptApi) TreeSelect(c *gin.Context) {
	query := req.DeptQuery{}
	if c.BindQuery(&query) == nil {
		treeSelect := a.deptService.TreeSelect(query)
		list := tree_dept.DeptList{}
		c.JSON(200, resp.Success(list.GetTree(treeSelect)))
	} else {
		c.JSON(200, resp.ErrorResp("参数绑定错误"))
	}
}

// RoleDeptTreeSelect 加载对应角色部门列表树
func (a DeptApi) RoleDeptTreeSelect(c *gin.Context) {
	m := make(map[string]interface{})
	param := c.Param("roleId")
	roleId, _ := strconv.ParseInt(param, 10, 64)
	checkedKeys := a.deptService.SelectDeptListByRoleId(roleId)
	m["checkedKeys"] = checkedKeys
	treeSelect := a.deptService.TreeSelect(req.DeptQuery{})
	list := tree_dept.DeptList{}
	tree := list.GetTree(treeSelect)
	m["depts"] = tree
	resp.OK(c, m)
}

// Find 查询部门列表
func (a DeptApi) Find(c *gin.Context) {
	query := req.DeptQuery{}
	if c.BindQuery(&query) != nil {
		resp.ParamError(c)
		return
	}
	resp.OK(c, a.deptService.GetList(query))
}

// ExcludeChild 查询部门列表（排除节点)
func (a DeptApi) ExcludeChild(c *gin.Context) {
	param := c.Param("deptId")
	deptId, err := strconv.Atoi(param)
	if err != nil {
		resp.ParamError(c)
		return
	}
	list := a.deptService.GetList(req.DeptQuery{})
	var depts = *list
	deptList := make([]models.SysDept, 0)
	for _, dept := range depts {
		if dept.DeptId == deptId || strings.Contains(dept.Ancestors, strconv.Itoa(deptId)) {
			continue
		}
		deptList = append(deptList, dept)
	}
	resp.OK(c, deptList)
}

// GetInfo 根据部门编号获取详细信息
func (a DeptApi) GetInfo(c *gin.Context) {
	param := c.Param("deptId")
	deptId, err := strconv.Atoi(param)
	if err != nil {
		resp.ParamError(c)
		return
	}
	resp.OK(c, a.deptService.GetDeptById(deptId))
}

// Add 添加部门
func (a DeptApi) Add(c *gin.Context) {
	dept := models.SysDept{}
	if c.Bind(&dept) != nil {
		resp.ParamError(c)
		return
	}
	//校验部门名称是否唯一
	unique := a.deptService.CheckDeptNameUnique(dept)
	if unique {
		resp.Error(c, "新增部门'"+dept.DeptName+"'失败，部门名称已存在")
		return
	}
	info := a.deptService.GetDeptById(dept.ParentId)
	if info.Status == "1" {
		resp.Error(c, "部门停用，不允许新增")
		return
	}
	dept.DelFlag = "0"
	dept.Ancestors = info.Ancestors + "," + strconv.Itoa(dept.ParentId)
	dept.CreateBy = user_util.GetUserInfo(c).UserName
	if a.deptService.Insert(dept) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除部门
func (a DeptApi) Delete(c *gin.Context) {
	param := c.Param("deptId")
	deptId, _ := strconv.Atoi(param)
	//是否存在部门子节点
	if a.deptService.HasChildByDeptId(deptId) > 0 {
		resp.Error(c, "存在下级部门,不允许删除")
		return
	}
	if a.deptService.CheckDeptExistUser(deptId) > 0 {
		resp.Error(c, "部门存在用户,不允许删除")
		return
	}
	if a.deptService.Remove(deptId) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Edit 修改部门
func (a DeptApi) Edit(c *gin.Context) {
	dept := models.SysDept{}
	if c.Bind(&dept) != nil {
		resp.ParamError(c)
		return
	}
	dept.UpdateBy = user_util.GetUserInfo(c).UserName
	dept.UpdateTime = time.Now()
	if a.deptService.Update(dept) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}
