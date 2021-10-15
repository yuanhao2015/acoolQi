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

type OperlogDao struct {
}

// Find 查询集合
func (d OperlogDao) Find(query req.OperlogQuery) (*[]models.SysOperLog, int64) {
	operlog := make([]models.SysOperLog, 0)
	session := SqlDB.NewSession().Table(models.SysOperLog{}.TableName())
	if query.Status != 0 {
		session.And("status = ?", query.Status)
	}
	if !acoolTools.StrUtils.HasEmpty(query.Title) {
		session.And("title = ?", query.Title)
	}
	if !acoolTools.StrUtils.HasEmpty(query.OperName) {
		session.And("oper_name = ?", query.OperName)
	}
	if query.BusinessType != 0 {
		session.And("business_type = ?", query.BusinessType)
	}
	if !acoolTools.StrUtils.HasEmpty(query.BeginTime) {
		session.And("date_format(oper_time,'%y%m%d') >= date_format(?,'%y%m%d')", query.BeginTime)
	}
	if !acoolTools.StrUtils.HasEmpty(query.EndTime) {
		session.And("date_format(oper_time,'%y%m%d') <= date_format(?,'%y%m%d')", query.EndTime)
	}
	if !acoolTools.StrUtils.HasEmpty(query.OrderByColumn) {
		isasc := strings.ReplaceAll(query.IsAsc, "ending", "")
		session.OrderBy(fmt.Sprintf("%s %s", common.UnMarshal(query.OrderByColumn), isasc))
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&operlog)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil, 0
	}

	return &operlog, total
}

// Add 添加数据
func (d OperlogDao) Add(operlog models.SysOperLog) int64 {
	fmt.Println(operlog)
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&operlog)
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return insert
}

// Remove 批量删除
func (d OperlogDao) Remove(list []int64) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	i, err := session.In("oper_id", list).Delete(&models.SysOperLog{})
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return i
}

// Get 查询
func (d OperlogDao) Get(id int64) *models.SysOperLog {
	operlog := models.SysOperLog{}
	_, err := SqlDB.NewSession().Where("oper_id = ?", id).Get(&operlog)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &operlog
}

// Edit 修改数据
func (d OperlogDao) Edit(operlog models.SysOperLog) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("oper_id = ?", operlog.OperId).Update(&operlog)
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return update
}
