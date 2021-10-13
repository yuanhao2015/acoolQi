package jwt

import (
	"github.com/gin-gonic/gin"
	"strings"
)

//是否在放行范围内
func doSquare(c *gin.Context) bool {
	request := inSquareRequest()
	for i := 0; i < len(request); i++ {
		replace := strings.Contains(c.Request.RequestURI, request[i])
		if replace {
			return true
		}
	}
	return false
}

//放行的请求
func inSquareRequest() []string {
	var req []string
	//一下是放行的请求
	//放行登录请求
	req = append(req, "/api/v1/login")
	return req
}
