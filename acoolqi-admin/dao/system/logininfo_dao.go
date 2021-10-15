package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/common"
	"acoolqi-admin/pkg/page"
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"strings"
)

type LogininfoDao struct {
}

// Find 查询集合
func (d LogininfoDao) Find(query req.LogininfoQuery) (*[]models.SysLogininfor, int64) {
	logininfo := make([]models.SysLogininfor, 0)
	session := SqlDB.NewSession().Table(models.SysLogininfor{}.TableName())

	if query.Status != 0 {
		session.And("status = ?", query.Status)
	}
	if !acoolTools.StrUtils.HasEmpty(query.LoginName) {
		session.And("login_name = ?", query.LoginName)
	}
	if !acoolTools.StrUtils.HasEmpty(query.Ipaddr) {
		session.And("ipaddr = ?", query.Ipaddr)
	}
	if !acoolTools.StrUtils.HasEmpty(query.BeginTime) {
		session.And("date_format(login_time,'%y%m%d') >= date_format(?,'%y%m%d')", query.BeginTime)
	}
	if !acoolTools.StrUtils.HasEmpty(query.EndTime) {
		session.And("date_format(login_time,'%y%m%d') <= date_format(?,'%y%m%d')", query.EndTime)
	}
	if !acoolTools.StrUtils.HasEmpty(query.OrderByColumn) {
		isasc := strings.ReplaceAll(query.IsAsc, "ending", "")
		session.OrderBy(fmt.Sprintf("%s %s", common.UnMarshal(query.OrderByColumn), isasc))
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&logininfo)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil, 0
	}
	return &logininfo, total
}

// Add 添加数据
func (d LogininfoDao) Add(logininfo models.SysLogininfor) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&logininfo)
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return insert
}

// Remove 批量删除
func (d LogininfoDao) Remove(list []int64) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	i, err := session.In("info_id", list).Delete(&models.SysLogininfor{})
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return i
}

// Get 查询
func (d LogininfoDao) Get(id int64) *models.SysLogininfor {
	logininfo := models.SysLogininfor{}
	_, err := SqlDB.NewSession().Where("info_id = ?", id).Get(&logininfo)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &logininfo
}

// Edit 修改数据
func (d LogininfoDao) Edit(logininfo models.SysLogininfor) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("info_id = ?", logininfo.InfoId).Update(&logininfo)
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return update
}
