package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var Conf *Config

// 整个config文件对应的结构体
type Config struct {
	Addr    string        `yaml:"addr"`    // http请求地址
	Log     LogConfig     `yaml:"log"`     // 日志
	Spread  SpreadConfig  `yaml:"spread"`  // 推广链接
	Mysqls  []MysqlConfig `yaml:"mysql"`   // mysql
	Redis   []RedisConfig `yaml:"redis"`   // redis
	Expires Expires       `yaml:"expires"` // 登录过期时间配置
	Routine RoutineConfig `yaml:"routine"` // 协程
	ExePath string        `yaml:"-"`
}

// 推广网址配置结构体
type SpreadConfig struct {
	Url string `yaml:"url"` // 网址
}

// 协程配置结构体
type RoutineConfig struct {
	Runtime      int `yaml:"runtime"`      // 定时任务多长时间执行一次，以秒为单位
	GoroutineNum int `yaml:"goroutineNum"` // 可以开的协程数量
	IpmiRuntime  int `yaml:"ipmiRuntime"`  // ipmi获取主机温度定时任务轮询时长，以分钟为单位
}

// 日志配置结构体
type LogConfig struct {
	Level   string `yaml:"level"` // 日志层级
	LogType string `yaml:"type"`  // 日志类型
	Path    string `yaml:"path"`  // 错误日志的存放路径
}

// mysql配置结构体
type MysqlConfig struct {
	Name     string `yaml:"name"`     // 分布式数据库名称
	Host     string `yaml:"host"`     // 数据库主机地址
	Username string `yaml:"username"` // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	DbName   string `yaml:"dbname"`   // 数据库名
	ShowSql  bool   `yaml:"showsql"`  // 是否打印sql到控制台
	Timeout  string `yaml:"timeout"`  // 超时的时间限制
}

// redis配置结构体
type RedisConfig struct {
	Name     string `yaml:"name"`     // 分布式redis名称
	Host     string `yaml:"host"`     // 主机地址
	Password string `yaml:"password"` // 密码
	DB       int    `yaml:"db"`       // 库名
}

type Expires struct {
	Captcha int `yaml:"captcha"` // 验证码过期时间(分钟)
	Login   int `yaml:"login"`   // 登录token过期时间(分钟)
}

// 加载配置文件
func GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return p, nil
}

// 从文件解析配置
func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return ParseConfigData(data)
}

// 从数据解析配置
func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	path, _ := GetExecPath()
	cfg.ExePath = filepath.Dir(path)

	if !strings.HasPrefix(cfg.Log.Path, "/") {
		cfg.Log.Path = filepath.Join(cfg.ExePath, cfg.Log.Path)
	}
	return &cfg, nil
}
