package service

import "acoolqi-admin/dao"

type UserPostService struct {
	userPostDao dao.UserPostDao
}

// CountUserPostById 删除岗位数据校验
func (s UserPostService) CountUserPostById(ids []int64) int64 {
	for _, id := range ids {
		if s.userPostDao.CountUserPostById(id) > 0 {
			return id
		}
	}
	return 0
}
