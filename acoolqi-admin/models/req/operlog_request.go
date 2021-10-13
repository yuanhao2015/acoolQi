/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: monitor_operlog
* @Date: 2021-9-28 11:13
 */
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
