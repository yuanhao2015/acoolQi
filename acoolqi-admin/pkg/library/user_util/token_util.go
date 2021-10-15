package user_util

import (
	"acoolqi-admin/config"
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/response"
	"acoolqi-admin/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/yuanhao2015/acoolTools"
	"strings"
)

type UserUtils struct {
}

// GetUserInfo 通过jwt获取当前登录用户
func GetUserInfo(c *gin.Context) *response.UserResponse {
	token := c.Request.Header.Get("Authorization")
	s := strings.Split(token, " ")
	j := jwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(s[1])
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
	}
	info := claims.UserInfo
	return &info
}

// CheckLockToken 校验多终端登录锁
func CheckLockToken(c *gin.Context) bool {
	server := config.GetServerCfg()
	if server.Lock == "0" {
		//获取redis中的token数据
		info := GetUserInfo(c)
		get, err := system.RedisDB.GET(info.UserName)
		if err != nil {
			acoolTools.Logs.ErrorLog().Println(err)
			return false
		}
		token := c.Request.Header.Get("Authorization")
		s := strings.Split(token, " ")
		if get == s[1] {
			return true
		} else {
			return false
		}
	}
	return true
}

// SaveRedisToken 将token存入到redis
func SaveRedisToken(key string, s string) {
	//server := config.GetServerCfg()
	//if server.Lock == "0" {
	system.RedisDB.SETEX(key, 3600, s)
	//}
}

// CheckIsAdmin 判断是否是超级管理员
func CheckIsAdmin(user *models.SysUser) bool {
	if user.UserId == 1 {
		return true
	} else {
		return false
	}
}
