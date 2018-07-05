package main

import (
	"config"
	"framework"
	"global"
	"log"
	"path/filepath"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var err error

	// 获取当前程序运行路径
	p, err := config.GetExecPath()
	if err != nil {
		log.Fatalf("Get exec path error:%v\n", err.Error())
		return
	}

	// 获取配置文件
	cf := filepath.Join(filepath.Dir(p), "etc/server.yaml")
	// 解析配置文件
	config.Conf, err = config.ParseConfigFile(cf)
	if err != nil {
		log.Fatalf("parse config file error:%v\n", err.Error())
		return
	}

	// 全局初始化
	global.InitLog(config.Conf, "server")
	defer global.Glogger.Close()

	// 初始化数据库
	err = global.InitMysql(config.Conf.Mysqls)
	if err != nil {
		// 数据库连接错误
		global.Glogger.Error("InitDb error:%v\n", err.Error())
		return
	}

	// 初始化redis
	err = global.InitRedis(config.Conf.Redis)
	if err != nil {
		// redis连接失败
		global.Glogger.Error("Redis connection failed:%v\n", err.Error())
		return
	}

	// 启动订单定时任务
	go model.Order{}.Cron()

	// 启动主机温度定时任务
	go model.Host{}.CropMonitor()

	// 启动域名到期通知定时任务
	go model.Realm{}.Cron()

	// 启动web服务
	err = framework.NewApp(config.Conf).Run()
	if err != nil {
		global.Glogger.Error("app start error:%v\n", err.Error())
		return
	}
}
