package cache

import (
	"github.com/yuanhao2015/acoolTools"
	"acoolqi-admin/dao"
	"acoolqi-admin/models"
	"acoolqi-admin/pkg/constant"
)

// GetRedisDictKey 根据key获取缓存中的字典数据
func GetRedisDictKey(key string) []models.SysDictData {
	get, err := dao.RedisDB.GET(key)
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
	dao.RedisDB.SET(list[0].DictType, dictList)
}

// GetRedisConfigByKey 根据key从缓存中获取配置数据
func GetRedisConfigByKey(key string) *models.SysConfig {
	get, err := dao.RedisDB.GET(key)
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
	dao.RedisDB.SET(config.ConfigKey, s)
}

// RemoveList 批量根据Key删除数据
func RemoveList(list []string) {
	dao.RedisDB.DELALL(list)
}

// SetDictCache 添加字典数据
func SetDictCache(key string, value string) {
	dao.RedisDB.SET(key, value)
}

// RemoveKey 根据key删除
func RemoveKey(key string) int {
	del, err := dao.RedisDB.DEL(key)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
	}
	return del
}
