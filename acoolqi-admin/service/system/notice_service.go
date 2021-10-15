package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
)

type NoticeService struct {
	noticeDao system.NoticeDao
}

// Find 查询集合数据
func (s NoticeService) Find(query req.NoticeQuery) (*[]models.SysNotice, int64) {
	return s.noticeDao.Find(query)
}

// Add 添加公告
func (s NoticeService) Add(notice models.SysNotice) bool {
	return s.noticeDao.Add(notice) > 0
}

// Remove 批量删除
func (s NoticeService) Remove(list []int64) bool {
	return s.noticeDao.Remove(list) > 0
}

// Get 查询
func (s NoticeService) Get(id int64) *models.SysNotice {
	return s.noticeDao.Get(id)
}

// Edit 修改
func (s NoticeService) Edit(notice models.SysNotice) bool {
	return s.noticeDao.Edit(notice) > 0
}
