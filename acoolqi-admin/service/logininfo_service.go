/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: operlog_service
* @Date: 2021-9-28 11:16
 */
package service

import (
	"acoolqi-admin/dao"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
)

type LogininfoService struct {
	logininfoDao dao.LogininfoDao
}

// Find 查询集合数据
func (s LogininfoService) Find(query req.LogininfoQuery) (*[]models.SysLogininfor, int64) {
	return s.logininfoDao.Find(query)
}

// Add 添加数据
func (s LogininfoService) Add(logininfo models.SysLogininfor) bool {
	return s.logininfoDao.Add(logininfo) > 0
}

// Remove 批量删除
func (s LogininfoService) Remove(list []int64) bool {
	return s.logininfoDao.Remove(list) > 0
}

// Get 查询
func (s LogininfoService) Get(id int64) *models.SysLogininfor {
	return s.logininfoDao.Get(id)
}

// Edit 修改
func (s LogininfoService) Edit(logininfo models.SysLogininfor) bool {
	return s.logininfoDao.Edit(logininfo) > 0
}
