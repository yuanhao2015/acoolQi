package cache

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/pkg/constant"
	"github.com/yuanhao2015/acoolTools"
)

// GetRedisDictKey 根据key获取缓存中的字典数据
func GetRedisDictKey(key string) []models.SysDictData {
	get, err := system.RedisDB.GET(key)
	if err != nil {
		acoolTools.Logs.ErrorLog().Fatalf(constant.RedisConstant{}.GetRedisError(), err.Error())
		return nil
	}
	list := models.SysDictData{}.UnmarshalDictList(get)
	return list
}

// SetRedisDict 保存字典数据
func SetRedisDict(list []models.SysDictData) {
	dictList := models.SysDictData{}.MarshalDictList(list)
	system.RedisDB.SET(list[0].DictType, dictList)
}

// GetRedisConfigByKey 根据key从缓存中获取配置数据
func GetRedisConfigByKey(key string) *models.SysConfig {
	get, err := system.RedisDB.GET(key)
	if err != nil {
		acoolTools.Logs.ErrorLog().Fatalf(constant.RedisConstant{}.GetRedisError(), err.Error())
		return nil
	}
	obj := models.SysConfig{}.UnmarshalDictObj(get)
	return obj
}

// SetRedisConfig 将配置存入缓存
func SetRedisConfig(config *models.SysConfig) {
	s := models.SysConfig{}.MarshalDictObj(*config)
	system.RedisDB.SET(config.ConfigKey, s)
}

// RemoveList 批量根据Key删除数据
func RemoveList(list []string) {
	system.RedisDB.DELALL(list)
}

// SetDictCache 添加字典数据
func SetDictCache(key string, value string) {
	system.RedisDB.SET(key, value)
}

// RemoveKey 根据key删除
func RemoveKey(key string) int {
	del, err := system.RedisDB.DEL(key)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
	}
	return del
}
