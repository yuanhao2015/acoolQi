package models

import "time"

type SysOperLog struct {
	OperId        int64     `json:"operId"`        // 日志主键
	Title         string    `json:"title"`         // 模块标题
	BusinessType  int       `json:"businessType"`  // 业务类型（0其它 1新增 2修改 3删除）
	Method        string    `json:"method"`        // 方法名称
	RequestMethod string    `json:"requestMethod"` // 请求方式
	OperatorType  int       `json:"operatorType"`  // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string    `json:"operName"`      // 操作人员
	DeptName      string    `json:"deptName"`      // 部门名称
	OperUrl       string    `json:"operUrl"`       // 请求URL
	OperIp        string    `json:"operIp"`        // 主机地址
	OperLocation  string    `json:"operLocation"`  // 操作地点
	OperParam     string    `json:"operParam"`     // 请求参数
	JsonResult    string    `json:"jsonResult"`    // 返回参数
	Status        int       `json:"status"`        // 操作状态（0正常 1异常）
	ErrorMsg      string    `json:"errorMsg"`      // 错误消息
	OperTime      time.Time `json:"operTime"`      // 操作时间
}

func (SysOperLog) TableName() string {
	return "sys_oper_log"
}
