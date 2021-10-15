package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/excels"
	"acoolqi-admin/pkg/file"
	"acoolqi-admin/pkg/library/user_util"
	"acoolqi-admin/pkg/page"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/service/system"
	"github.com/gin-gonic/gin"
	"github.com/yuanhao2015/acoolTools"
	"strconv"
	"strings"
	"time"
)

type DictDataApi struct {
	dictDataService system.DictDataService
}

// GetByType 根据字典类型查询字典数据d
func (a DictDataApi) GetByType(c *gin.Context) {
	param := c.Param("dictType")
	if !acoolTools.StrUtils.HasEmpty(param) {
		byType := a.dictDataService.GetByType(param)
		c.JSON(200, resp.Success(byType))
	}
}

// List 查询字典数据集合
func (a DictDataApi) List(c *gin.Context) {
	query := req.DiceDataQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	list, i := a.dictDataService.GetList(query)
	resp.OK(c, page.Page{
		List:  list,
		Total: i,
		Size:  query.PageSize,
	})
}

// Get 根据id查询字典数据
func (a DictDataApi) Get(c *gin.Context) {
	param := c.Param("dictCode")
	dictCode, _ := strconv.ParseInt(param, 10, 64)
	resp.OK(c, a.dictDataService.GetByCode(dictCode))
}

// Add 添加字典数据
func (a DictDataApi) Add(c *gin.Context) {
	data := models.SysDictData{}
	if c.Bind(&data) != nil {
		resp.ParamError(c)
		return
	}
	data.CreateBy = user_util.GetUserInfo(c).UserName
	if a.dictDataService.Insert(data) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除数据
func (a DictDataApi) Delete(c *gin.Context) {
	param := c.Param("dictCode")
	split := strings.Split(param, ",")
	dictCodeList := make([]int64, 0)
	for _, s := range split {
		diceCode, _ := strconv.ParseInt(s, 10, 64)
		dictCodeList = append(dictCodeList, diceCode)
	}
	if a.dictDataService.Remove(dictCodeList) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Export 导出excel
func (a DictDataApi) Export(c *gin.Context) {
	query := req.DiceDataQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	items := make([]interface{}, 0)
	list, _ := a.dictDataService.GetList(query)
	for _, data := range *list {
		items = append(items, data)
	}
	_, files := excels.ExportExcel(items, "字典数据表")
	file.DownloadExcel(c, files)
}

// Edit 修改部门
func (a DictDataApi) Edit(c *gin.Context) {
	dict := models.SysDictData{}
	if c.Bind(&dict) != nil {
		resp.ParamError(c)
		return
	}
	dict.UpdateBy = user_util.GetUserInfo(c).UserName
	dict.UpdateTime = time.Now()
	if a.dictDataService.Update(dict) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}
