package framework

import (
	"config"
	"controller"
	"framework/app"
	"github.com/go-redis/redis"
	"github.com/golang/net/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	. "global"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"router"
	"strconv"
	"time"
)

// 初始化web服务
func NewApp(cfg *config.Config) *app.App {
	a := new(app.App)

	a.Cfg = cfg
	a.WebServer = echo.New()
	a.WebServer.HideBanner = true
	// 可以支持跨域请求
	a.WebServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	//
	a.WebServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 20 << 10, // 1 KB
	}))
	// 路由
	a.WebServer.GET("/favicon.ico", func(ctx echo.Context) error {
		return ctx.JSON(200, controller.Base{}.Return("success", nil, 0))
	})
	// 程序版本
	a.WebServer.GET("/version", func(ctx echo.Context) error {
		// 获取编译时间
		bti, _ := strconv.Atoi(BUILDTIME)
		bt := time.Unix(int64(bti), 0).Format("2006-01-02 15:04:05")
		// 获取程序路径
		file, _ := exec.LookPath(os.Args[0])
		p, _ := filepath.Abs(file)
		// 获取程序文件md5
		f, err := os.Open(p + ".md5")
		if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("app.file_md5_not_found", nil, 0))
		}
		md5, err := ioutil.ReadAll(f)
		if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("app.file_md5_error", nil, 0))
		}
		// 获取程序文件sha256
		f, err = os.Open(p + ".sha256")
		if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("app.file_sha256_not_found", nil, 0))
		}
		sha256, err := ioutil.ReadAll(f)
		if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("app.file_sha256_error", nil, 0))
		}
		return ctx.JSON(200, controller.Base{}.Return("success", struct {
			Version   string `json:"version"`
			BuildTime string `json:"buildTime"`
			Md5       string `json:"md5"`
			Sha256    string `json:"sha256"`
		}{
			Version:   VERSION,
			BuildTime: bt,
			Md5:       string(md5),
			Sha256:    string(sha256),
		}, 0))
	})
	// 开启webSocket
	a.WebServer.GET("/", WSHandler)
	// 启用路由
	router.Route(a.WebServer)
	return a
}

func WSHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer func() {
			ws.Close()
		}()
		// 过滤非法连接
		token := ws.Request().URL.Query().Get("key")
		if token == "" {
			return
		}
		for {
			id := check(token)
			if id == 0 {
				// 如果当前token to id 存在于map中
				if ids, ok := MapToConn.Load(token); ok {
					// 删除该map中 id 对应的 conn
					MapToConn.Delete(ids)
				}
				// 删除token to id
				MapToConn.Delete(token)
				return
			}
			// 如果校验成功
			MapToConn.Store(token, id)
			MapToConn.Store(id, ws)
			time.Sleep(time.Second * 5)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func check(token string) int64 {
	// redis里获取登录信息
	result, err := GetLogin().Get(token).Result()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		Glogger.Error(err.Error())
		return 0
	}
	id, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		Glogger.Error(err.Error())
		return 0
	}
	return id
}
