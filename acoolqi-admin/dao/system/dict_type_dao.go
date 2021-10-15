package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/page"
	"github.com/go-xorm/xorm"
	"github.com/yuanhao2015/acoolTools"
)

type DictTypeDao struct {
}

func (d DictTypeDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_dict_type")
}

// FindAll 查询所有字典类型数据
func (d DictTypeDao) FindAll() []*models.SysDictType {
	types := make([]*models.SysDictType, 0)
	err := d.sql(SqlDB.NewSession()).Where("status = '0'").Find(&types)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return types
}

// Find 分页查询字典类型数据
func (d DictTypeDao) Find(query req.DictTypeQuery) (*[]models.SysDictType, int64) {
	list := make([]models.SysDictType, 0)
	session := SqlDB.NewSession().Table("sys_dict_type")
	if acoolTools.StrUtils.HasNotEmpty(query.DictName) {
		session.And("dict_name like concat('%', ?, '%')", query.DictName)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.DictType) {
		session.And("AND dict_type like concat('%', ?, '%')", query.DictType)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.BeginTime) {
		session.And("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d')", query.BeginTime)
	}
	if acoolTools.StrUtils.HasNotEmpty(query.EndTime) {
		session.And("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d')", query.EndTime)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&list)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil, 0
	}
	return &list, total
}

// GetById 根据id查询字典类型数据
func (d DictTypeDao) GetById(id int64) *models.SysDictType {
	dictType := models.SysDictType{}
	_, err := SqlDB.NewSession().Where("dict_id = ?", id).Get(&dictType)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return nil
	}
	return &dictType
}

// CheckDictTypeUnique 检验字典类型是否存在
func (d DictTypeDao) CheckDictTypeUnique(dictType models.SysDictType) int64 {
	session := SqlDB.NewSession().Table("sys_dict_type")
	if dictType.DictId > 0 {
		session.And("dict_id != ?", dictType.DictId)
	}
	count, err := session.Where("dict_type = ?", dictType.DictType).Cols("dict_id").Count()
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	return count
}

// Update 修改字典
func (d DictTypeDao) Update(dictType models.SysDictType) bool {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Where("dict_id = ?", dictType.DictId).Update(&dictType)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return false
	}
	session.Commit()
	return true
}

// Insert 新增字典类型
func (d DictTypeDao) Insert(dictType models.SysDictType) int64 {
	session := SqlDB.NewSession()
	session.Begin()
	insert, err := session.Insert(&dictType)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return 0
	}
	session.Commit()
	return insert
}

// Remove 批量删除
func (d DictTypeDao) Remove(ids []int64) bool {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.In("dict_id", ids).Delete(models.SysDictType{})
	if err != nil {
		session.Rollback()
		acoolTools.Logs.ErrorLog().Println(err)
		return false
	}
	session.Commit()
	return true
}
