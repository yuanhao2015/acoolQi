package models

import "time"

type SysLogininfor struct {
	InfoId        int64     `json:"infoId"`        // 访问ID
	LoginName     string    `json:"loginName"`     // 登录账号
	Ipaddr        string    `json:"ipaddr"`        // 登录IP地址
	LoginLocation string    `json:"loginLocation"` // 登录地点
	Browser       string    `json:"browser"`       // 浏览器类型
	Os            string    `json:"os"`            // 操作系统
	Status        int64     `json:"status"`        // 登录状态（0成功 1失败）
	Msg           string    `json:"msg"`           // 提示消息
	LoginTime     time.Time `json:"loginTime"`     // 访问时间
}

func (SysLogininfor) TableName() string {
	return "sys_logininfor"
}
