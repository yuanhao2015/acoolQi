package req

import (
	"acoolqi-admin/pkg/base"
)

// UserOnlineQuery 用户get请求数据参数
type UserOnlineQuery struct {
	base.GlobalQuery
	LoginName string ` form:"loginName"`
	Ipaddr    string ` form:"ipaddr"`
}

type UserOnlineBody struct {
	LoginName string ` json:"loginName"`
	Status    int    ` json:"status"`
}
