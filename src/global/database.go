package global

// 数据库初始化
import (
	"config"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	EngineMap           map[string]*xorm.Engine
	errInvalidMysqlNode = errors.New("config mysql node is nil")
)

func InitMysql(configSql []config.MysqlConfig) error {
	var err error
	EngineMap = make(map[string]*xorm.Engine)
	length := len(configSql)
	for i := 0; i < length; i++ {
		node := configSql[i].Name
		if len(node) == 0 || node != "master" {
			return errInvalidMysqlNode
		}
		host := configSql[i].Host
		dbName := configSql[i].DbName
		password := configSql[i].Password
		username := configSql[i].Username
		timeout := configSql[i].Timeout
		data := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&timeout=%s",
			username, password, host, dbName, timeout)
		EngineMap[node], err = xorm.NewEngine("mysql", data)
		if err != nil {
			return err
		}
		err = EngineMap[node].Ping()
		if err != nil {
			return err
		}
		EngineMap[node].ShowSQL(configSql[i].ShowSql)
	}
	return nil
}

// 获取master库xorm Session
func GetMaster() *xorm.Session {
	sess := EngineMap["master"].NewSession()
	return sess
}
