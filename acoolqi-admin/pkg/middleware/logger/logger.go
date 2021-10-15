package logger

import (
	"acoolqi-admin/config"
	"acoolqi-admin/models"
	"acoolqi-admin/pkg/jwt"
	"acoolqi-admin/service/system"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/yuanhao2015/acoolTools"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.GetLoggerCfg().LogPath
	logFileName := config.GetLoggerCfg().LogName

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if !acoolTools.FileUtils.Exists(fileName) {
		create, err := os.Create(fileName)
		if err != nil {
			acoolTools.Logs.ErrorLog().Println(err)
		}
		defer create.Close()
		//acoolTools.FileUtils.Create(fileName)
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		var body string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			bf := bytes.NewBuffer(nil)
			wt := bufio.NewWriter(bf)
			_, err := io.Copy(wt, c.Request.Body)
			if err != nil {
				err = nil
			}
			rb, _ := ioutil.ReadAll(bf)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
			body = string(rb)
		}
		// 处理请求
		c.Next()
		url := c.Request.RequestURI
		for _, v := range []string{
			"logout", "login", "refreshCache",
		} {
			if strings.Index(url, v) > -1 {
				return
			}
		}

		// 结束时间
		endTime := time.Now()
		if c.Request.Method == http.MethodOptions {
			return
		}

		rt, bl := c.Get("result")
		var result = ""
		if bl {
			rb, err := json.Marshal(rt)
			if err != nil {
			} else {
				result = string(rb)
			}
		}

		em, bl := c.Get("error_msg")
		var errorMsg = ""
		if bl {
			errorMsg = em.(string)
		}
		st, bl := c.Get("status")
		var statusBus = 0
		if bl {
			statusBus = st.(int)
		}
		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
		if c.Request.Method != "OPTIONS" && statusCode != 404 && reqMethod != "GET" {
			SetDBOperLog(c, clientIP, reqUri, body, result, statusBus, errorMsg)
		}
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// SetDBOperLog 写入操作日志表
func SetDBOperLog(c *gin.Context, clientIP string, reqUri string, body string, result string, status int, errorMsg string) {
	if status == http.StatusOK {
		status = 0
	} else {
		status = 1
	}
	var claims *jwt.CustomClaims
	get, exists := c.Get("claims")
	if exists {
		claims = get.(*jwt.CustomClaims)
	}

	compile, _ := regexp.Compile(`/`)
	split := compile.Split(reqUri, -1)
	methods := split[3:6]
	if methods[2] == "data" || methods[2] == "type" {
		methods = split[3:5]
		methods = append(methods, split[6])
	}
	var businessType int
	if methods[2] == "add" {
		businessType = 1
	} else if methods[2] == "edit" {
		businessType = 2
	} else if methods[2] == "remove" {
		businessType = 3
	} else if methods[2] == "authUser" {
		businessType = 4
	} else if methods[2] == "clean" {
		businessType = 9
	} else {
		businessType = 10
	}

	var menuService system.MenuService
	perms := menuService.GetMenuPerms(strings.Join(methods, ":"))

	db := models.SysOperLog{
		OperId:        0,
		Title:         perms.MenuName,
		BusinessType:  businessType,
		Method:        strings.Join(methods, ":"),
		RequestMethod: c.Request.Method,
		OperatorType:  1,
		OperName:      claims.UserInfo.NickName,
		DeptName:      claims.DeptInfo.DeptName,
		OperUrl:       reqUri,
		OperIp:        clientIP,
		OperLocation:  acoolTools.ClientIPUtils.GetCityByIp(clientIP),
		OperParam:     body,
		JsonResult:    result,
		Status:        status,
		ErrorMsg:      errorMsg,
		OperTime:      time.Now(),
	}
	//fmt.Println(db)
	var operlogService system.OperlogService
	operlogService.Add(db)
}
