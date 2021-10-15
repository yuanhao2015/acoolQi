package system

import (
	"acoolqi-admin/config"
	rds "acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/cache"
	"acoolqi-admin/pkg/jwt"
	"acoolqi-admin/pkg/library/tree/tree_menu"
	"acoolqi-admin/pkg/library/user_util"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/pkg/uuid"
	"acoolqi-admin/service/monitor"
	"acoolqi-admin/service/system"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/yuanhao2015/acoolTools"
	"time"
)

type LoginApi struct {
	loginService      system.LoginService
	roleService       system.RoleService
	permissionService system.PermissionService
	menuService       system.MenuService
}

// Login 登录
func (a LoginApi) Login(c *gin.Context) {
	loginBody := req.LoginBody{}
	if c.BindJSON(&loginBody) == nil {
		m := make(map[string]string)
		login, s := a.loginService.Login(loginBody.UserName, loginBody.Password)
		//记录登录日志中
		var logininfoService system.LogininfoService
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
			//判断是否一个账号只能在一个地方登录
			appServer := config.GetServerCfg()
			lock := appServer.Lock
			if lock == "0" {
				keys, err := rds.RedisDB.KEYS("users-" + loginBody.UserName + "*")
				if err != nil {
					db.Msg = "获取" + loginBody.UserName + "  key失败"
					db.Status = 1
					c.JSON(200, resp.ErrorResp(db.Msg))
					return
				}
				_, err = rds.RedisDB.DELALL(keys)
				if err != nil {
					db.Msg = "清理" + loginBody.UserName + "key失败"
					db.Status = 1
					c.JSON(200, resp.ErrorResp(db.Msg))
					return
				}
			}
			//将token存入到redis中
			redisKeys, uuid := uuid.CreateRedisUuid(loginBody.UserName)
			user_util.SaveRedisToken(redisKeys, s)
			//登录信息存放user_online中
			useronline := monitor.UserOnlineService{}
			j := jwt.NewJWT()
			claims, err := j.ParseToken(s)
			if err != nil {
				acoolTools.Logs.ErrorLog().Println(err)
			}
			useronline.Add(models.SysUserOnline{
				Uuid:           uuid,
				Token:          s,
				LoginName:      claims.UserInfo.NickName,
				DeptName:       claims.DeptInfo.DeptName,
				Ipaddr:         c.ClientIP(),
				LoginLocation:  acoolTools.ClientIPUtils.GetCityByIp(c.ClientIP()),
				Browser:        browser,
				Os:             ua.OS(),
				Status:         0,
				StartTimestamp: time.Now(),
				ExpireTime:     0,
			})
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
	var logininfoService system.LogininfoService
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
