package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/cache"
)

type DictTypeService struct {
	dictTypeDao system.DictTypeDao
	dictDataDao system.DictDataDao
}

// Find 分页查询字典类型数据
func (s DictTypeService) Find(query req.DictTypeQuery) (*[]models.SysDictType, int64) {
	return s.dictTypeDao.Find(query)
}

// GetById 根据id查询字典类型数据
func (s DictTypeService) GetById(id int64) *models.SysDictType {
	return s.dictTypeDao.GetById(id)
}

// CheckDictTypeUnique 检验字典类型是否存在
func (s DictTypeService) CheckDictTypeUnique(dictType models.SysDictType) bool {
	return s.dictTypeDao.CheckDictTypeUnique(dictType) > 0
}

// Update 修改字典数据
func (s DictTypeService) Update(dictType models.SysDictType) bool {
	return s.dictTypeDao.Update(dictType)
}

// Insert 新增字典类型
func (s DictTypeService) Insert(dictType models.SysDictType) bool {
	insert := s.dictTypeDao.Insert(dictType)
	if insert > 0 {
		cache.SetDictCache(dictType.DictType, "")
	}
	return insert > 0
}

// Remove 批量删除
func (s DictTypeService) Remove(ids []int64) bool {
	return s.dictTypeDao.Remove(ids)
}

// FindAll 查询所有字典类型数据
func (s DictTypeService) FindAll() []*models.SysDictType {
	return s.dictTypeDao.FindAll()
}

// RemoveAllCache 删除所有字典缓存
func (s DictTypeService) RemoveAllCache() []string {
	typeList := make([]string, 0)
	allType := s.FindAll()
	for _, dictType := range allType {
		typeList = append(typeList, dictType.DictType)
	}
	//删除缓存
	cache.RemoveList(typeList)
	return typeList
}

// LoadDictCache 将字典数据存入缓存
func (s DictTypeService) LoadDictCache() {
	typeList := make([]string, 0)
	allType := s.FindAll()
	m := models.SysDictData{}
	for _, dictType := range allType {
		typeList = append(typeList, dictType.DictType)
	}
	allData := s.dictDataDao.GetDiceDataAll()
	for _, key := range typeList {
		list := make([]models.SysDictData, 0)
		for _, data := range *allData {
			if key == data.DictType {
				list = append(list, data)
			}
		}
		system.RedisDB.SET(key, m.MarshalDictList(list))
	}
}

// RefreshCache 刷新缓存数据
func (s DictTypeService) RefreshCache() {
	typeList := s.RemoveAllCache()
	m := models.SysDictData{}
	allData := s.dictDataDao.GetDiceDataAll()
	for _, key := range typeList {
		list := make([]models.SysDictData, 0)
		for _, data := range *allData {
			if key == data.DictType {
				list = append(list, data)
			}
		}

		system.RedisDB.SET(key, m.MarshalDictList(list))
	}
}
