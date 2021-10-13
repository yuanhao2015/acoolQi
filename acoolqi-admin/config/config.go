package config

import (
	"github.com/Unknwon/goconfig"
	"github.com/yuanhao2015/acoolTools"
	"log"
	"strconv"
	"time"
)

// BUILD 开发环境
const BUILD = "dev"

//生产环境
//const BUILD = "prod"

var Cfg *goconfig.ConfigFile

func init() {
	var err error
	Cfg, err = goconfig.LoadConfigFile("./config/config-" + BUILD + ".ini")
	if err != nil {
		log.Fatal(err)
	}
	return
}

type JwtConfig struct {
	TimeOut time.Duration //超时时间
	Issuer  string        //签证签发人
}

// FilePath 文件存储配置获取
type FilePath struct {
	Path string
}

// DbCfg Mysql数据库配置
type DbCfg struct {
	Username          string
	Password          string
	Database          string
	Host              string
	Port              string
	Debug             bool //是否调试模式
	MaxIdleConnection int
	MaxOpenConnection int
	ShowType          string
}

// AppServer 应用程序配置
type AppServer struct {
	Port        string //App运行端口
	Lock        string //是否开启多终端登录 0开启 1不开启
	DemoEnabled bool   //是否是演示模式
}

// LoggerCfg 日志配置结构体
type LoggerCfg struct {
	LogPath string
	LogName string
}

// RedisCfg redis配置结构体
type RedisCfg struct {
	RedisHost string //地址
	Port      int64  //端口
	RedisPwd  string //密码
	RedisDB   int64  //数据库
	Timeout   int64  //超时时间
}

// MongoDb mongo配置结构体
type MongoDb struct {
	Url      string
	Port     string
	DB       string
	User     string
	Password string
}

//读取mysql配置
func GetMysqlCfg() (mysql DbCfg) {
	mysql.Username, _ = Cfg.GetValue("mysql", "username")
	mysql.Password, _ = Cfg.GetValue("mysql", "password")
	mysql.Database, _ = Cfg.GetValue("mysql", "database")
	mysql.Host, _ = Cfg.GetValue("mysql", "host")
	mysql.Port, _ = Cfg.GetValue("mysql", "port")
	mysql.ShowType, _ = Cfg.GetValue("mysql", "sqlType")
	mysqlDebug, _ := Cfg.GetValue("mysql", "debug")
	mysql.Debug = mysqlDebug == "0"
	value, _ := Cfg.GetValue("mysql", "MaxIdleConnection")
	mysql.MaxIdleConnection, _ = strconv.Atoi(value)
	v, _ := Cfg.GetValue("mysql", "MaxOpenConnection")
	mysql.MaxOpenConnection, _ = strconv.Atoi(v)
	return mysql
}

//读取server配置
func GetServerCfg() (server AppServer) {
	server.Port, _ = Cfg.GetValue("app", "server")
	server.Lock, _ = Cfg.GetValue("app", "lock")
	demoEnabled, _ := Cfg.GetValue("app", "demoEnabled")
	server.DemoEnabled = demoEnabled == "0"
	return server
}

//获取Logger配置
func GetLoggerCfg() (logger LoggerCfg) {
	logger.LogPath, _ = Cfg.GetValue("logger", "logPath")
	logger.LogName, _ = Cfg.GetValue("logger", "logName")
	return logger
}

//获取mongo配置
func GetMongoCfg() (mongo MongoDb) {
	mongo.Url, _ = Cfg.GetValue("mongodb", "url")
	mongo.Port, _ = Cfg.GetValue("mongodb", "port")
	mongo.DB, _ = Cfg.GetValue("mongodb", "db")
	mongo.User, _ = Cfg.GetValue("mongodb", "user")
	mongo.Password, _ = Cfg.GetValue("mongodb", "password")
	return mongo
}

//获取redis配置
func GetRedisCfg() (r RedisCfg) {
	r.RedisHost, _ = Cfg.GetValue("redis", "host")
	getValue, _ := Cfg.GetValue("redis", "port")
	r.Port, _ = strconv.ParseInt(getValue, 10, 32)
	r.RedisPwd, _ = Cfg.GetValue("redis", "password")
	db, _ := Cfg.GetValue("redis", "db")
	r.RedisDB, _ = strconv.ParseInt(db, 10, 32)
	value, _ := Cfg.GetValue("redis", "timeout")
	r.Timeout, _ = strconv.ParseInt(value, 10, 64)
	return r
}

// GetFilePath 获取文件存储位置
func GetFilePath() (g FilePath) {
	g.Path, _ = Cfg.GetValue("filePath", "path")
	return g
}

func GetJwtConfig() (j JwtConfig) {
	value, _ := Cfg.GetValue("jwt", "timeOut")
	issuer, _ := Cfg.GetValue("jwt", "issuer")
	if acoolTools.StrUtils.HasEmpty(value) {
		value = "60"
	}
	if acoolTools.StrUtils.HasEmpty(issuer) {
		value = "acool"
	}
	atoi, _ := strconv.ParseInt(value, 10, 64)
	j.TimeOut = time.Duration(atoi / 60)
	j.Issuer = issuer
	return j
}
