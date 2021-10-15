package system

import (
	"acoolqi-admin/dao/system"
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"bytes"
	"github.com/yuanhao2015/acoolTools"
)

type PostService struct {
	postDao system.PostDao
}

// FindAll 查询所有岗位业务方法
func (s PostService) FindAll() []*models.SysPost {
	return s.postDao.SelectAll()
}

// SelectPostListByUserId 根据用户id查询岗位id集合
func (s PostService) SelectPostListByUserId(userId int64) *[]int64 {
	return s.postDao.SelectPostListByUserId(userId)
}

// FindList 查询岗位分页列表
func (s PostService) FindList(query req.PostQuery) (*[]models.SysPost, int64) {
	return s.postDao.Find(query)
}

// CheckPostNameUnique 校验岗位名称是否存在
func (s PostService) CheckPostNameUnique(post models.SysPost) bool {
	return s.postDao.CheckPostNameUnique(post) > 0
}

// CheckPostCodeUnique 校验岗位编码是否存在
func (s PostService) CheckPostCodeUnique(post models.SysPost) bool {
	return s.postDao.CheckPostCodeUnique(post) > 0
}

// Insert 添加岗位数据
func (s PostService) Insert(post models.SysPost) bool {
	return s.postDao.Insert(post) > 0
}

// GetPostById 根据id查询岗位数据
func (s PostService) GetPostById(id int64) *models.SysPost {
	post := models.SysPost{
		PostId: id,
	}
	return s.postDao.GetPostById(post)
}

// Delete 批量删除岗位信息
func (s PostService) Delete(ids []int64) bool {
	return s.postDao.Delete(ids) > 0
}

// Update 修改岗位数据
func (s PostService) Update(post models.SysPost) bool {
	return s.postDao.Update(post)
}

// SelectPostByUserName 获取岗位数据
func (s PostService) SelectPostByUserName(name string) string {
	list := s.postDao.SelectPostByUserName(name)
	var buffer bytes.Buffer
	var postName string
	for _, post := range *list {
		buffer.WriteString(post.PostName)
		buffer.WriteString(",")
	}
	s2 := buffer.String()
	if acoolTools.StrUtils.HasNotEmpty(s2) {
		postName = s2[0:(len(s2) - 1)]
	}
	return postName
}
