package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/cache"
)

type ConfigService struct {
	configDao system.ConfigDao
}

// GetConfigKey 根据键名查询参数配置信息
func (s ConfigService) GetConfigKey(param string) *models.SysConfig {
	//从缓存中取出数据判断是否存在，存在直接使用，不存在就从数据库查询
	key := cache.GetRedisConfigByKey(param)
	if key != nil {
		return key
	}
	configKey := s.configDao.GetConfigKey(param)
	cache.SetRedisConfig(configKey)
	return configKey
}

// Find 分页查询数据
func (s ConfigService) Find(query req.ConfigQuery) (*[]models.SysConfig, int64) {
	return s.configDao.List(query)
}

// CheckConfigKeyUnique 校验是否存在
func (s ConfigService) CheckConfigKeyUnique(config models.SysConfig) bool {
	return s.configDao.CheckConfigKeyUnique(config) > 0
}

// Add 添加数据
func (s ConfigService) Add(config models.SysConfig) int64 {
	return s.configDao.Insert(config)
}

// GetInfo 查询数据
func (s ConfigService) GetInfo(id int64) *models.SysConfig {
	return s.configDao.GetById(id)
}

// Update 修改数据
func (s ConfigService) Update(config models.SysConfig) int64 {
	return s.configDao.Update(config)
}

// Remove 批量删除
func (s ConfigService) Remove(list []int64) bool {
	return s.configDao.Remove(list)
}

// CheckConfigByIds 根据id集合查询
func (s ConfigService) CheckConfigByIds(list []int64) *[]models.SysConfig {
	return s.configDao.CheckConfigByIds(list)
}

// FindAll 查询所有数据
func (s ConfigService) FindAll() *[]models.SysConfig {
	return s.configDao.FindAll()
}
