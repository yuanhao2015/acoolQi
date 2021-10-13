package v1

import (
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/page"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/service"
	"github.com/gin-gonic/gin"
)

type LogininfoApi struct {
	logininfoService service.LogininfoService
}

// List 查询集合
func (a LogininfoApi) List(c *gin.Context) {
	query := req.LogininfoQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, i := a.logininfoService.Find(query)
	resp.OK(c, page.Page{
		List:  find,
		Total: i,
		Size:  query.PageSize,
	})
}
