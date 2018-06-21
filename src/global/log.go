package global

import (
	"config"
	"framework/exception"
	"framework/logger"
	"log"
	"os"
	"path"
)

const (
	sysLogName = "sys.log"
	maxLogSize = 1024 * 1024 * 1024
)

var Glogger *logger.Logger

// 日志等级 "FNST", "FINE", "DEBG", "TRAC", "INFO", "WARN", "EROR", "CRIT"
func InitLog(cfg *config.Config, pathname string) {
	defer exception.Exception()
	// 初始化
	switch cfg.Log.LogType {
	case "file":
		err := os.MkdirAll(cfg.Log.Path+"/"+pathname, 0777)
		if err != nil {
			log.Println(cfg.Log.Path+" 文件创建失败:", err.Error())
		}
		fileLog := logger.NewFileLogWriter(path.Join(cfg.Log.Path+"/"+pathname, sysLogName), true)
		fileLog = fileLog.SetRotateSize(maxLogSize)

		Glogger = &logger.Logger{
			"fileOut": &logger.Filter{
				Level:     logger.StringToLevel(cfg.Log.Level),
				LogWriter: fileLog},
		}
	default:
		Glogger = &logger.Logger{
			"stdout": &logger.Filter{
				Level:     logger.DEBUG,
				LogWriter: logger.NewConsoleLogWriter()},
		}
	}
}
