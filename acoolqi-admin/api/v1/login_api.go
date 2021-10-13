package v1

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/cache"
	"acoolqi-admin/pkg/library/tree/tree_menu"
	"acoolqi-admin/pkg/library/user_util"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/yuanhao2015/acoolTools"
	"time"
)

type LoginApi struct {
	loginService      service.LoginService
	roleService       service.RoleService
	permissionService service.PermissionService
	menuService       service.MenuService
}

// Login 登录
func (a LoginApi) Login(c *gin.Context) {
	loginBody := req.LoginBody{}
	if c.BindJSON(&loginBody) == nil {
		m := make(map[string]string)
		login, s := a.loginService.Login(loginBody.UserName, loginBody.Password)
		//记录登录日志中
		var logininfoService service.LogininfoService
		userAgent := c.Request.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		browser, _ := ua.Browser()
		db := models.SysLogininfor{
			InfoId:        0,
			LoginName:     loginBody.UserName,
			Ipaddr:        c.ClientIP(),
			LoginLocation: acoolTools.ClientIPUtils.GetCityByIp(c.ClientIP()),
			Browser:       browser,
			Os:            ua.OS(),
			Status:        0,
			Msg:           "",
			LoginTime:     time.Now(),
		}
		if login {
			//将token存入到redis中
			user_util.SaveRedisToken(loginBody.UserName, s)
			m["token"] = s
			db.Msg = "登录成功"
			logininfoService.Add(db)
			c.JSON(200, resp.Success(m))
		} else {
			db.Msg = s
			db.Status = 1
			logininfoService.Add(db)
			c.JSON(200, resp.ErrorResp(s))
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数绑定错误"))
	}
}

// GetUserInfo 获取用户信息
func (a LoginApi) GetUserInfo(c *gin.Context) {
	m := make(map[string]interface{})
	user := a.loginService.LoginUser(c)
	//查询用户角色集合
	roleKeys := a.permissionService.GetRolePermissionByUserId(user)
	// 权限集合
	perms := a.permissionService.GetMenuPermission(user)
	m["roles"] = roleKeys
	m["permissions"] = perms
	m["user"] = user
	c.JSON(200, resp.Success(m))
}

// GetRouters 根据用户ID查询菜单
func (a LoginApi) GetRouters(c *gin.Context) {
	//获取等钱登录用户
	user := a.loginService.LoginUser(c)
	menus := a.menuService.GetMenuTreeByUserId(user)
	systemMenus := tree_menu.SystemMenus{}
	systemMenus = *menus
	array := systemMenus.ConvertToINodeArray(menus)
	generateTree := tree_menu.GenerateTree(array, nil)
	c.JSON(200, resp.Success(generateTree))
}

// Logout 退出登录
func (a LoginApi) Logout(c *gin.Context) {
	//删除Redis缓存
	name := user_util.GetUserInfo(c).UserName
	cache.RemoveKey(name)
	//记录登录日志中
	var logininfoService service.LogininfoService
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	browser, _ := ua.Browser()
	db := models.SysLogininfor{
		InfoId:        0,
		LoginName:     name,
		Ipaddr:        c.ClientIP(),
		LoginLocation: acoolTools.ClientIPUtils.GetCityByIp(c.ClientIP()),
		Browser:       browser,
		Os:            ua.OS(),
		Status:        0,
		Msg:           "登出成功",
		LoginTime:     time.Now(),
	}
	fmt.Println(db)
	logininfoService.Add(db)
	resp.OK(c)
}
