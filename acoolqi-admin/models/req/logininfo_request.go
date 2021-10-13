package req

import "acoolqi-admin/pkg/base"

type LogininfoQuery struct {
	base.GlobalQuery
	LoginName     string `form:"loginName"` // 登录账号
	Ipaddr        string `form:"ipaddr"`    // 登录IP地址
	Status        int64  `form:"status"`
	OrderByColumn string `form:"orderByColumn"`
	IsAsc         string `form:"isAsc"`
}
