package models

import (
	"encoding/json"
	"github.com/yuanhao2015/acoolTools"
	"time"
)

type SysConfig struct {
	ConfigId    int       `excel:"name=参数主键" xorm:"pk autoincr" json:"configId"`              //主键id
	ConfigName  string    `excel:"name=参数名称" xorm:"varchar(100)" json:"configName"`           //参数名称
	ConfigKey   string    `excel:"name=参数键名" xorm:"varchar(100)" json:"configKey"`            //参数建名
	ConfigValue string    `excel:"name=参数键值" xorm:"varchar(1)" json:"configValue"`            //参数键值
	ConfigType  string    `excel:"name=系统内置,format=Y=是,N=否" xorm:"char(1)" json:"configType"` //系统内置（Y是 N否）
	CreateBy    string    `xorm:"varchar(64)" json:"createBy"`                                //创建人
	CreateTime  time.Time `xorm:"created" json:"createTime"`                                  //创建时间
	UpdateBy    string    `xorm:"varchar(64)" json:"updateBy"`                                //更新人
	UpdateTime  time.Time `json:"updateTime"`                                                 //更新时间
	Remark      string    `xorm:"varchar(500)" json:"remark"`                                 //备注
}

func (c SysConfig) TableName() string {
	return "sys_config"
}

// MarshalDictList 序列化配置数据
func (SysConfig) MarshalDictList(d []*SysConfig) string {
	marshal, err := json.Marshal(&d)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return ""
	}
	return string(marshal)
}

// UnmarshalDictList 反序列化配置数据
func (SysConfig) UnmarshalDictList(data string) []*SysConfig {
	s := make([]*SysConfig, 0)
	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return s
}

// MarshalDictObj 但实体序列化
func (SysConfig) MarshalDictObj(d SysConfig) string {
	marshal, err := json.Marshal(&d)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return ""
	}
	return string(marshal)
}

// UnmarshalDictObj  单实体反序列化
func (SysConfig) UnmarshalDictObj(data string) *SysConfig {
	s := new(SysConfig)
	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return s
}
