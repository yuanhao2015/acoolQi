package req

import "acoolqi-admin/pkg/base"

type PostQuery struct {
	base.GlobalQuery
	PostCode string `form:"postCode"`
	Status   string `form:"status"`
	PostName string `form:"postName"`
}
