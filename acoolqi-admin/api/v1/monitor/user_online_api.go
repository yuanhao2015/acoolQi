/*
@Time : 2021-10-14 15:32
@Author : acool
@File : user_online_api
*/
package monitor

import (
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/page"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/service/monitor"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SysUserOnlineApi struct {
	useronlineService monitor.UserOnlineService
}

// 在线用户列表
func (s *SysUserOnlineApi) List(c *gin.Context) {
	query := req.UserOnlineQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, i := s.useronlineService.Find(query)
	resp.OK(c, page.Page{
		List:  find,
		Total: i,
		Size:  query.PageSize,
	})
}

// 强退用户
func (s *SysUserOnlineApi) ForceLogout(c *gin.Context) {
	param := c.Param("sessionIds")
	if param != "" {
		resp.Error(c)
	}
	useronline := monitor.UserOnlineService{}
	id, _ := strconv.ParseInt(param, 10, 64)
	remove := useronline.Remove([]int64{id})
	if remove {
		resp.OK(c)
	}
}
