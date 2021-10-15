package monitor

import (
	"acoolqi-admin/dao/monitor"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
)

type UserOnlineService struct {
	useronlineDao monitor.UserOnlineDao
}

// Find 查询集合数据
func (s UserOnlineService) Find(query req.UserOnlineQuery) (*[]models.SysUserOnline, int64) {
	return s.useronlineDao.Find(query)
}

// FindAll 查询集合数据(不限制条数)
func (s UserOnlineService) FindAll(query req.UserOnlineBody) (*[]models.SysUserOnline, int64) {
	return s.useronlineDao.FindAll(query)
}

// Add 添加公告
func (s UserOnlineService) Add(useronline models.SysUserOnline) bool {
	return s.useronlineDao.Add(useronline) > 0
}

// Remove 批量删除
func (s UserOnlineService) Remove(list []int64) bool {
	return s.useronlineDao.Remove(list) > 0
}

// Get 查询
func (s UserOnlineService) Get(id int64) *models.SysUserOnline {
	return s.useronlineDao.Get(id)
}

// Get 查询
func (s UserOnlineService) GetByToken(token string) *models.SysUserOnline {
	return s.useronlineDao.GetByToken(token)
}

// 强退用户

func (s UserOnlineService) ForceLogout(id int64) {

}
