package req

import "acoolqi-admin/pkg/base"

type ConfigQuery struct {
	base.GlobalQuery
	ConfigName string `form:"configName"`
	ConfigType string `form:"configType"`
	ConfigKey  string `form:"configKey"`
}
