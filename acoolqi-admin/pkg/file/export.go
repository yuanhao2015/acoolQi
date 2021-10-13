package file

import (
	"github.com/yuanhao2015/acoolTools"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// DownloadExcel 公共下载execl方法
func DownloadExcel(c *gin.Context, file *excelize.File) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+acoolTools.IdUtils.IdUUIDToRan(false)+".xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("FileName", acoolTools.IdUtils.IdUUIDToRan(false)+".xlsx")
	file.Write(c.Writer)
}
