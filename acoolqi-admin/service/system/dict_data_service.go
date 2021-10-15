package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/cache"
)

type DictDataService struct {
	dictDataDao system.DictDataDao
}

// GetByType 根据字典类型查询字典数据
func (s DictDataService) GetByType(param string) []models.SysDictData {
	//先从缓存中拉数据
	key := cache.GetRedisDictKey(param)
	if key != nil {
		return key
	} else {
		//缓存中为空，从数据库中取数据
		return s.dictDataDao.GetByType(param)
	}
}

// GetList 查询字段数据集合
func (s DictDataService) GetList(query req.DiceDataQuery) (*[]models.SysDictData, int64) {
	return s.dictDataDao.GetList(query)
}

// GetByCode 根据code查询字典数据
func (s DictDataService) GetByCode(code int64) *models.SysDictData {
	return s.dictDataDao.GetByCode(code)
}

// Insert 新增字典数据
func (s DictDataService) Insert(data models.SysDictData) bool {
	insert := s.dictDataDao.Insert(data)
	if insert > 0 {
		//刷新缓存数据
		byType := s.GetNoCacheByType(data.DictType)
		cache.SetDictCache(data.DictType, data.MarshalDictList(byType))
	}
	return insert > 0
}

// Remove 删除数据
func (s DictDataService) Remove(codes []int64) bool {
	dictType := s.GetByCode(codes[0]).DictType
	m := models.SysDictData{}
	remove := s.dictDataDao.Remove(codes)
	if remove {
		//刷新缓存
		code := s.GetNoCacheByType(dictType)
		cache.SetDictCache(dictType, m.MarshalDictList(code))
	}
	return remove
}

// GetNoCacheByType 根据字典类型查询字典数据
func (s DictDataService) GetNoCacheByType(param string) []models.SysDictData {
	return s.dictDataDao.GetByType(param)
}

// Update 修改部门
func (s DictDataService) Update(dictdata models.SysDictData) int64 {
	return s.dictDataDao.Update(dictdata)
}
