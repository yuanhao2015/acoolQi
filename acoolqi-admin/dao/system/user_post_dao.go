package system

import (
	"acoolqi-admin/models"
	"github.com/yuanhao2015/acoolTools"
)

type UserPostDao struct {
}

// BatchUserPost 批量新增用户岗位信息
func (d UserPostDao) BatchUserPost(posts []models.SysUserPost) {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Table(models.SysUserPost{}.TableName()).Insert(&posts)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// RemoveUserPost 删除用户和岗位关系
func (d UserPostDao) RemoveUserPost(id int64) {
	post := models.SysUserPost{
		UserId: id,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&post)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		session.Rollback()
	}
	session.Commit()
}

// CountUserPostById 通过岗位ID查询岗位使用数量
func (d UserPostDao) CountUserPostById(id int64) int64 {
	count, err := SqlDB.NewSession().Table("sys_user_post").Cols("post_id").Where("post_id = ?", id).Count()
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return 0
	}
	return count
}
