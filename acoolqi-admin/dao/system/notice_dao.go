package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/page"
	"github.com/yuanhao2015/acoolTools"
)

type NoticeDao struct {
}

// Find 查询集合
func (d NoticeDao) Find(query req.NoticeQuery) (*[]models.SysNotice, int64) {
	notices := make([]models.SysNotice, 0)
	session := SqlDB.NewSession().Table(models.SysNotice{}.TableName())
	if acoolTools.StrUtils.HasNotEmpty(query.NoticeTitle) {
		session.And("notice_title like concat('%', ?, '%')", query.NoticeTitle)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.NoticeType) {
		session.And("notice_type = ?", query.NoticeType)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.CreateBy) {
		session.And("create_by like concat('%', ?, '%')", query.CreateBy)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&notices)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil, 0
	}
	return &notices, total
}

// Add 添加数据
func (d NoticeDao) Add(notice models.SysNotice) int64 {
	session := SqlDB.NewSession()
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
func (d NoticeDao) Remove(list []int64) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	i, err := session.In("notice_id", list).Delete(&models.SysNotice{})
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return i
}

// Get 查询
func (d NoticeDao) Get(id int64) *models.SysNotice {
	notice := models.SysNotice{}
	_, err := SqlDB.NewSession().Where("notice_id = ?", id).Get(&notice)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &notice
}

// Edit 修改数据
func (d NoticeDao) Edit(notice models.SysNotice) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("notice_id = ?", notice.NoticeId).Update(&notice)
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return update
}
