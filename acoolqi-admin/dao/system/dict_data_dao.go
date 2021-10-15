package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/page"
	"github.com/go-xorm/xorm"
	"github.com/yuanhao2015/acoolTools"
)

type DictDataDao struct {
}

func (d *DictDataDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_dict_data")
}

// GetByType 根据字典类型查询字典数据
func (d *DictDataDao) GetByType(param string) []models.SysDictData {
	data := make([]models.SysDictData, 0)
	session := d.sql(SqlDB.NewSession())
	err := session.Where("status = '0' ").And("dict_type = ?", param).OrderBy("dict_sort").Asc("dict_sort").
		Find(&data)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return data
}

// GetDiceDataAll 查询所有字典数据
func (d DictDataDao) GetDiceDataAll() *[]models.SysDictData {
	session := d.sql(SqlDB.NewSession())
	data := make([]models.SysDictData, 0)
	err := session.Where("status = '0' ").OrderBy("dict_sort").Asc("dict_sort").
		Find(&data)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &data
}

// GetList 查询集合数据
func (d *DictDataDao) GetList(query req.DiceDataQuery) (*[]models.SysDictData, int64) {
	list := make([]models.SysDictData, 0)
	session := SqlDB.NewSession().Table("sys_dict_data").OrderBy("dict_sort").Asc("dict_sort")
	if acoolTools.StrUtils.HasNotEmpty(query.DictType) {
		session.And("dict_type = ?", query.DictType)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.DictLabel) {
		session.And("dict_label like concat('%', ?, '%')", query.DictLabel)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&list)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil, 0
	}
	return &list, total
}

// GetByCode 根据code查询字典数据
func (d *DictDataDao) GetByCode(code int64) *models.SysDictData {
	data := models.SysDictData{}
	session := SqlDB.NewSession()
	_, err := session.Where("dict_code = ?", code).Get(&data)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &data
}

// Insert 添加字典数据
func (d *DictDataDao) Insert(data models.SysDictData) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&data)
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	session.Commit()
	return insert
}

// Remove 删除字典数据
func (d *DictDataDao) Remove(codes []int64) bool {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.In("dict_code", codes).Delete(&models.SysDictData{})
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// Update 修改字典数据
func (d *DictDataDao) Update(dictdata models.SysDictData) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	update, err := session.Where("dict_code = ?", dictdata.DictCode).Update(&dictdata)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return update
}
