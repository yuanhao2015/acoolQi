package monitor

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/page"
	"github.com/yuanhao2015/acoolTools"
)

type UserOnlineDao struct {
}

// Find 查询集合
func (d UserOnlineDao) Find(query req.UserOnlineQuery) (*[]models.SysUserOnline, int64) {
	useronline := make([]models.SysUserOnline, 0)
	session := system.SqlDB.NewSession().Table(models.SysUserOnline{}.TableName())
	if acoolTools.StrUtils.HasNotEmpty(query.LoginName) {
		session.And("login_name like concat('%', ?, '%')", query.LoginName)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.Ipaddr) {
		session.And("ipaddr like concat('%', ?, '%')", query.Ipaddr)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&useronline)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil, 0
	}
	return &useronline, total
}

// Find 查询集合
func (d UserOnlineDao) FindAll(query req.UserOnlineBody) (*[]models.SysUserOnline, int64) {
	useronline := make([]models.SysUserOnline, 0)
	session := system.SqlDB.NewSession().Table(models.SysUserOnline{}.TableName())
	total, _ := page.GetTotal(session.Clone())
	err := session.Find(&useronline)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil, 0
	}
	return &useronline, total
}

// Add 添加数据
func (d UserOnlineDao) Add(notice models.SysUserOnline) int64 {
	session := system.SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&notice)
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return insert
}

// Remove 批量删除
func (d UserOnlineDao) Remove(list []int64) int64 {
	session := system.SqlDB.NewSession()
	session.Begin()
	i, err := session.In("session_id", list).Delete(&models.SysUserOnline{})
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return i
}

// Get 查询
func (d UserOnlineDao) Get(id int64) *models.SysUserOnline {
	notice := models.SysUserOnline{}
	_, err := system.SqlDB.NewSession().Where("session_id = ?", id).Get(&notice)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &notice
}

// GetByToken 查询
func (d UserOnlineDao) GetByToken(token string) *models.SysUserOnline {
	notice := models.SysUserOnline{}
	_, err := system.SqlDB.NewSession().Where("token = ?", token).Get(&notice)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &notice
}

// Remove 批量删除
func (d UserOnlineDao) DeleteOnlineByToken(token string) int64 {
	useronline := models.SysUserOnline{
		Token: token,
	}
	session := system.SqlDB.NewSession()
	session.Begin()
	i, err := session.Delete(&useronline)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return i
}
