package system

import (
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/page"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/service/system"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type OperlogApi struct {
	operlogService system.OperlogService
}

// List 查询集合
func (a OperlogApi) List(c *gin.Context) {
	query := req.OperlogQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, i := a.operlogService.Find(query)
	resp.OK(c, page.Page{
		List:  find,
		Total: i,
		Size:  query.PageSize,
	})
}

// Delete 删除
func (a OperlogApi) Delete(c *gin.Context) {
	ids := strings.Split(c.Param("ids"), ",")
	idList := make([]int64, 0)
	for _, s := range ids {
		id, _ := strconv.ParseInt(s, 10, 64)
		idList = append(idList, id)
	}
	if a.operlogService.Remove(idList) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Get 查询
func (a OperlogApi) Get(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	resp.OK(c, a.operlogService.Get(id))
}
