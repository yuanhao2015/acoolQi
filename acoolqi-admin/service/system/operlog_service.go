package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
)

type OperlogService struct {
	operlogDao system.OperlogDao
}

// Find 查询集合数据
func (s OperlogService) Find(query req.OperlogQuery) (*[]models.SysOperLog, int64) {
	return s.operlogDao.Find(query)
}

// Add 添加数据
func (s OperlogService) Add(operlog models.SysOperLog) bool {
	return s.operlogDao.Add(operlog) > 0
}

// Remove 批量删除
func (s OperlogService) Remove(list []int64) bool {
	return s.operlogDao.Remove(list) > 0
}

// Get 查询
func (s OperlogService) Get(id int64) *models.SysOperLog {
	return s.operlogDao.Get(id)
}

// Edit 修改
func (s OperlogService) Edit(operlog models.SysOperLog) bool {
	return s.operlogDao.Edit(operlog) > 0
}
