package models

import "time"

type SysUserOnline struct {
	SessionId      int64     `xorm:"pk autoincr" json:"sessionId"` // 用户会话id
	Uuid           string    ` json:"uuid"`                        // 用户标识
	Token          string    ` json:"-"`                           // 用户token
	LoginName      string    ` json:"loginName"`                   // 登录账号
	DeptName       string    ` json:"deptName"`                    // 部门名称
	Ipaddr         string    ` json:"ipaddr"`                      // 登录IP地址
	LoginLocation  string    ` json:"loginLocation"`               // 登录地点
	Browser        string    ` json:"browser"`                     // 浏览器类型
	Os             string    ` json:"os"`                          // 操作系统
	Status         int       ` json:"status"`                      // 在线状态on_line在线off_line离线
	StartTimestamp time.Time ` json:"startTimestamp"`              // session创建时间
	LastAccessTime time.Time ` json:"lastAccessTime"`              // session最后访问时间
	ExpireTime     int64     ` json:"expireTime"`                  // 超时时间，单位为分钟
}

func (SysUserOnline) TableName() string {
	return "sys_user_online"
}
