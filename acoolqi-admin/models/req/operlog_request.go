package req

import "acoolqi-admin/pkg/base"

type OperlogQuery struct {
	base.GlobalQuery
	Title         string `form:"title"`
	OperName      string `form:"operName"`
	BusinessType  int    `form:"businessType"`
	Status        int    `form:"status"`
	OrderByColumn string `form:"orderByColumn"`
	IsAsc         string `form:"isAsc"`
}
