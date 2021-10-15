package router

import (
	"acoolqi-admin/pkg/filter"
	"acoolqi-admin/pkg/jwt"
	"acoolqi-admin/pkg/middleware"
	"acoolqi-admin/pkg/middleware/logger"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(logger.LoggerToFile())
	router.Use(middleware.Recover)
	router.Use(jwt.JWTAuth())
	router.Use(filter.DemoHandler())
	//v1版本api
	v1Router := router.Group("/api/v1")
	{
		//登录接口
		initLoginRouter(v1Router)
		//用户路由接口
		initUserRouter(v1Router)
		//部门路由注册
		initDeptRouter(v1Router)
		//初始化字典数据路由
		initDictDataRouter(v1Router)
		//注册配置路由
		initConfigRouter(v1Router)
		//注册角色路由
		initRoleRouter(v1Router)
		//注册菜单路由
		initMenuRouter(v1Router)
		//注册岗位路由
		initPostRouter(v1Router)
		//注册字典类型路由
		initDictTypeRouter(v1Router)
		//注册公告路由
		initNoticeRouter(v1Router)
		//注册操作日志路由
		initOperlogRouter(v1Router)
		//注册登录日志路由
		initLogininfoRouter(v1Router)
		//服务监控路由
		initMonitorServerRouter(v1Router)
		//在线用户路由
		initUserOnlineRouter(v1Router)
	}
	return router
}
