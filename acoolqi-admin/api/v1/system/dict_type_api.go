package system

import (
	"acoolqi-admin/models"
	"acoolqi-admin/models/req"
	"acoolqi-admin/pkg/cache"
	"acoolqi-admin/pkg/excels"
	"acoolqi-admin/pkg/file"
	"acoolqi-admin/pkg/page"
	"acoolqi-admin/pkg/resp"
	"acoolqi-admin/service/system"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type DictTypeApi struct {
	dictTypeService system.DictTypeService
	dictDataService system.DictDataService
}

// List 获取字典类型数据
func (a DictTypeApi) List(c *gin.Context) {
	query := req.DictTypeQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, i := a.dictTypeService.Find(query)
	resp.OK(c, page.Page{
		List:  find,
		Total: i,
		Size:  query.PageSize,
	})
}

// Get 根据id查询字典类型数据
func (a DictTypeApi) Get(c *gin.Context) {
	param := c.Param("dictTypeId")
	dictTypeId, _ := strconv.ParseInt(param, 0, 64)
	resp.OK(c, a.dictTypeService.GetById(dictTypeId))
}

// Edit 修改字典类型
func (a DictTypeApi) Edit(c *gin.Context) {
	dictType := models.SysDictType{}
	if c.Bind(&dictType) != nil {
		resp.ParamError(c)
		return
	}
	//检验字典类型是否存在
	if a.dictTypeService.CheckDictTypeUnique(dictType) {
		resp.Error(c, "修改字典'"+dictType.DictName+"'失败，字典类型已存在")
		return
	}
	//修改数据
	if a.dictTypeService.Update(dictType) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Add 新增字典类型
func (a DictTypeApi) Add(c *gin.Context) {
	dictType := models.SysDictType{}
	if c.Bind(&dictType) != nil {
		resp.ParamError(c)
		return
	}
	//检验字典类型是否存在
	if a.dictTypeService.CheckDictTypeUnique(dictType) {
		resp.Error(c, "新增字典'"+dictType.DictName+"'失败，字典类型已存在")
		return
	}
	//新增字典类型
	if a.dictTypeService.Insert(dictType) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Remove 批量删除字典类型
func (a DictTypeApi) Remove(c *gin.Context) {
	param := c.Param("dictId")
	split := strings.Split(param, ",")
	ids := make([]int64, 0)
	types := make([]string, 0)
	for _, s := range split {
		parseInt, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, parseInt)
	}
	//校验字典类型是否使用
	for _, id := range ids {
		dictType := a.dictTypeService.GetById(id)
		if len(a.dictDataService.GetByType(dictType.DictType)) > 0 {
			resp.Error(c, dictType.DictName+"已分配,不能删除")
			return
		}
		types = append(types, dictType.DictType)
	}
	//批量删除
	if a.dictTypeService.Remove(ids) {
		//从缓存中删除数据
		cache.RemoveList(types)
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// RefreshCache 刷新缓存
func (a DictTypeApi) RefreshCache(c *gin.Context) {
	a.dictTypeService.RefreshCache()
	resp.OK(c)
}

// Export 导出Excel
func (a DictTypeApi) Export(c *gin.Context) {
	query := req.DictTypeQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	find, _ := a.dictTypeService.Find(query)
	list := make([]interface{}, 0)
	for _, dictType := range *find {
		list = append(list, dictType)
	}
	_, files := excels.ExportExcel(list, "字典类型表")
	file.DownloadExcel(c, files)
}
