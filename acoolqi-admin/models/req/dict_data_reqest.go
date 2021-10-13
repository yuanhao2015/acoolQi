package req

import "acoolqi-admin/pkg/base"

type DiceDataQuery struct {
	base.GlobalQuery
	DictType  string `form:"dictType"`
	DictLabel string `form:"dictLabel"`
	Status    string `form:"status"`
}
