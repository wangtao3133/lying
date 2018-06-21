package router

import (
	"bytes"
	"config"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	. "global"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"CMDB/src/controller"
	"CMDB/src/model"
)

// 验证码安全校验中间件
// 验证码在120s内连续刷新超过10次则以第10次刷新时间为准向后顺延120S限制刷新
func CheckCaptchaRequestNum(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// 获取请求参数(验证码类型)
		typ := ctx.QueryParam("type")
		if typ == "" {
			return ctx.JSON(200, controller.Base{}.Return("captcha.param_error", nil, 0))
		}
		t, err := strconv.Atoi(typ)
		if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("captcha.param_error", nil, 0))
		}
		if t != 1 && t != 2 {
			return ctx.JSON(200, controller.Base{}.Return("captcha.param_error", nil, 0))
		}
		var pre string
		if t == 1 {
			pre = "login_"
		} else {
			pre = "reg_"
		}
		// 获取请求ip
		key := pre + strconv.Itoa(StringIpToInt(ctx.RealIP()))
		// 查询redis中是否存在
		result, err := GetCaptcha().Get(key).Result()
		if err == redis.Nil {
			ctx.Set("pre", pre)
			return next(ctx)
		} else if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("redis.error", nil, 0))
		} else {
			r, err := strconv.Atoi(result)
			if err != nil {
				Glogger.Error(err.Error())
				return ctx.JSON(200, controller.Base{}.Return("system.error", nil, 0))
			}
			if r >= 50 {
				return ctx.JSON(200, controller.Base{}.Return("captcha.request_too_more", nil, 0))
			}
			ctx.Set("pre", pre)
			return next(ctx)
		}
	}
}

// 登录安全校验中间件
// 密码在120s内连续错误3次则以第3次错误时间为准向后顺延600限制登录
func CheckLoginRequestErr(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		data, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("system.error", nil, 0))
		}
		login := new(model.InputLogin)
		err = json.Unmarshal(data, login)
		if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("json.unmarshal_error", nil, 0))
		}
		ctx.Request().Body = ioutil.NopCloser(bytes.NewReader(data))
		result, err := GetLogin().Get(login.Account).Result()
		if err == redis.Nil {
			return next(ctx)
		} else if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, controller.Base{}.Return("redis.error", nil, 0))
		} else {
			r, err := strconv.Atoi(result)
			if err != nil {
				Glogger.Error(err.Error())
				return ctx.JSON(200, controller.Base{}.Return("system.error", nil, 0))
			}
			if r >= 3 {
				return ctx.JSON(200, controller.Base{}.Return("login.error_too_more", nil, 0))
			}
			return next(ctx)
		}
	}
}

// 登陆后中间校验
func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 获取token
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return c.JSON(200, controller.Base{}.Return("login.failure_or_timeout", nil, 0))
		}
		token := strings.Split(authorization, " ")
		if len(token) < 2 || token[1] == "" || token[0] != "Bearer" {
			return c.JSON(200, controller.Base{}.Return("login.failure_or_timeout", nil, 0))
		}

		// 根据token获取账号信息
		adminBean := new(model.Admin)
		ok, err := adminBean.GetByToken(token[1])
		if err != nil {
			Glogger.Error(err.Error())
			return c.JSON(200, controller.Base{}.Return("system.error", nil, 0))
		}
		if !ok {
			return c.JSON(200, controller.Base{}.Return("login.failure_or_timeout", nil, 0))
		}
		// 账号被禁用
		if adminBean.Status == StatusDisable {
			return c.JSON(200, controller.Base{}.Return("account.is_disable", nil, 0))
		}

		// 获取登录token是否过期
		result, err := GetLogin().Get(token[1]).Result()
		if err == redis.Nil {
			// 如果过期则删除数据表中的token
			_, err := adminBean.DelToken(token[1])
			if err != nil {
				Glogger.Error(err.Error())
			}
			return c.JSON(200, controller.Base{}.Return("login.failure_or_timeout", nil, 0))
		} else if err != nil {
			Glogger.Error(err.Error())
			return c.JSON(200, controller.Base{}.Return("redis.error", nil, 0))
		}

		id, err := strconv.ParseInt(result, 10, 64)
		if err != nil {
			Glogger.Error(err.Error())
			return c.JSON(200, controller.Base{}.Return("system.error", nil, 0))
		}

		err = GetLogin().Set(token[1], id, time.Minute*time.Duration(config.Conf.Expires.Login)).Err()
		if err != nil {
			Glogger.Error(err.Error())
			return c.JSON(200, controller.Base{}.Return("redis.error", nil, 0))
		}
		c.Set("token", token[1])
		return next(c)
	}
}

// 权限校验
func PowerValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		admin := controller.GetInfo(c.Get("token").(string))

		// 如果是超级管理员,则通过验证
		if admin.RoleId == RoleSA {
			return next(c)
		}

		// 获取当前请求
		method := c.Request().Method
		var rt int8
		switch method {
		case "GET":
			rt = ROUTE_GET
		case "POST":
			rt = ROUTE_POST
		case "PUT":
			rt = ROUTE_PUT
		case "DELETE":
			rt = ROUTE_DELETE
		}
		// 根据路由请求查询权限
		p := new(model.Power)
		power, has, err := p.ByRoute(c.Request().URL.Path, rt)
		if err != nil {
			return err
		}
		// 如果该路由权限未设置,则通过验证
		if !has {
			return next(c)
		}

		// 获取当前登录人所属角色权限
		rp := new(model.RolePower)
		bpi, err := rp.ByRoleId(admin.RoleId)
		if err != nil {
			return err
		}

		check := false
		for _, v := range bpi {
			if v.PowerId == power.ID {
				check = true
				break
			}
		}
		// 没有权限
		if !check {
			s := new(controller.Login)
			return c.JSON(200, s.Return("permission.denied", nil, 0))
		}
		return next(c)
	}
}

// 日志中间件
//func AccessLog(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		admin := controller.GetInfo(ctx.Get("token").(string))
//		operationBean := new(model.OperationLog)
//		operationBean.OperationAccount = admin.Account          // 操作账号
//		operationBean.OperationId = admin.Id                    // 操作人id
//		operationBean.CreateTime = time.Now().Unix()            // 操作时间
//		operationBean.OperationPath = ctx.Request().URL.Path    // 请求路由
//		operationBean.OperationIp = StringIpToInt(ctx.RealIP()) // 请求Ip
//		method := ctx.Request().Method                          // 请求方法
//		if method == "GET" {
//			operationBean.OperationQuery = ctx.QueryParams().Encode()
//		} else {
//			ctype := ctx.Request().Header.Get("Content-Type")
//			switch {
//			case strings.HasPrefix(ctype, "application/json"):
//				data, err := ioutil.ReadAll(ctx.Request().Body)
//				operationBean.Remark = string(data)
//				if err != nil {
//					Glogger.Error(err.Error())
//					return next(ctx)
//				}
//				switch string(data) {
//				case "GET":
//					operationBean.OperationType = ROUTE_GET
//				case "POST":
//					operationBean.OperationType = ROUTE_POST
//				case "PUT":
//					operationBean.OperationType = ROUTE_PUT
//				case "DELETE":
//					operationBean.OperationType = ROUTE_DELETE
//				}
//				// 重新写进body
//				ctx.Request().Body = ioutil.NopCloser(bytes.NewReader(data))
//			case strings.HasPrefix(ctype, "application/x-www-form-urlencoded"),
//				strings.HasPrefix(ctype, "multipart/form-data"):
//				data, err := ctx.FormParams()
//				if err != nil {
//					Glogger.Error(err.Error())
//					return next(ctx)
//				}
//				switch data.Encode() {
//				case "GET":
//					operationBean.OperationType = ROUTE_GET
//				case "POST":
//					operationBean.OperationType = ROUTE_POST
//				case "PUT":
//					operationBean.OperationType = ROUTE_PUT
//				case "DELETE":
//					operationBean.OperationType = ROUTE_DELETE
//				}
//			default:
//				operationBean.OperationType = 0
//			}
//		}
//		switch method {
//		case "POST": // 增加
//			operationBean.OperationType = 2
//		case "DELETE": // 删除
//			operationBean.OperationType = 4
//		case "GET": // 查看
//			operationBean.OperationType = 1
//		default: // 修改
//			operationBean.OperationType = 3
//		}
//		_, err := GetMaster().InsertOne(operationBean)
//		if err != nil {
//			Glogger.Error(err.Error())
//			return next(ctx)
//		}
//		return next(ctx)
//	}
//}
