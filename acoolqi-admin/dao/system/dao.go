package system

import (
	"acoolqi-admin/config"
	"acoolqi-admin/models"
	"acoolqi-admin/pkg/common"
	redisTool "acoolqi-admin/pkg/redistool"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yuanhao2015/acoolTools"
	"log"
	"time"
)

// X 全局DB
var (
	SqlDB   *xorm.Engine
	RedisDB *redisTool.RedisClient
)

func init() {
	var err error
	//配置mysql数据库
	mysql := config.GetMysqlCfg()
	jdbc := mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.Database + "?charset=utf8&parseTime=True&loc=Local"
	SqlDB, _ = xorm.NewEngine(mysql.ShowType, jdbc)

	if err != nil {
		log.Fatalf("db error: %#v\n", err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatalf("db connect error: %#v\n", err.Error())
	}
	SqlDB.SetMaxIdleConns(10)
	SqlDB.SetMaxOpenConns(100)
	_ = SqlDB.Sync2(
	//new(model.User),
	)
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				log.Fatalf("db connect error: %#v\n", err.Error())
			}
		}
	}(SqlDB)
	SqlDB.ShowSQL(mysql.Debug)
	//初始化redis开始
	redisCfg := config.GetRedisCfg()
	redisOpt := common.RedisConnOpt{
		true,
		redisCfg.RedisHost,
		int32(redisCfg.Port),
		redisCfg.RedisPwd,
		int32(redisCfg.RedisDB),
		240,
	}

	RedisDB = redisTool.NewRedis(redisOpt)
	//配置redis结束
	//缓存初始化数据
	saveCache()
}

//初始化缓存数据
func saveCache() {
	initDict()
	initConfig()
}

func initDict() {
	//查询字典类型数据

	dictTypeDao := new(DictTypeDao)
	typeAll := dictTypeDao.FindAll()
	//所有字典数据
	d := new(DictDataDao)
	listData := d.GetDiceDataAll()
	m := new(models.SysDictData)
	for _, dictType := range typeAll {
		dictData := make([]models.SysDictData, 0)
		for _, data := range *listData {
			if dictType.DictType == data.DictType {
				dictData = append(dictData, data)
			}
		}
		RedisDB.SET(dictType.DictType, m.MarshalDictList(dictData))
	}
}

func initConfig() {
	//查询配置数据存入到缓存中
	configDao := new(ConfigDao)
	configSession := configDao.Sql(SqlDB.NewSession())
	configs := make([]*models.SysConfig, 0)
	err := configSession.Find(&configs)
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
		return
	}
	m2 := new(models.SysConfig)
	for _, sysConfig := range configs {
		RedisDB.SET(sysConfig.ConfigKey, m2.MarshalDictObj(*sysConfig))
	}
}
