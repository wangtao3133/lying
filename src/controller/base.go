package controller

import (
	"framework/validation"
	"github.com/labstack/echo"
	. "global"
	"strings"
)

type Base struct{}

// 获取listparam的值(分页使用)
func (bc Base) GetParam(listparam *ListParams, ctx echo.Context) error {
	param := new(ListParams)
	// 获取结构体的json数据
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(200, bc.Return("input.error_type", nil, 0))
	}
	listparam.Desc = param.Desc
	listparam.OrderBy = param.OrderBy
	page := param.Page // 页码
	if page < 1 {
		page = 1
	}
	listparam.Page = page
	pageSize := param.PageSize // 每页数量
	if pageSize < 1 {
		pageSize = 15
	} else {
		if pageSize >= 100 {
			pageSize = 100
		}
	}
	listparam.PageSize = pageSize
	offset := (page - 1) * pageSize
	var limit = []int{pageSize, offset}
	listparam.Limit = limit
	return nil
}

// 验证参数
func (Base) ValidParameter(column interface{}, ctx echo.Context) int {
	// 请求参数绑定
	if err := ctx.Bind(column); err != nil {
		return 1
	}
	// 请求参数验证
	valid := validation.Validation{}
	ok, err := valid.Valid(column)
	if err != nil {
		return -3
	}
	if !ok {
		return -100
	}
	return 0
}

// 验证参数
func (Base) ValidParam(column interface{}, ctx echo.Context) string {
	if err := ctx.Bind(column); err != nil {
		return "input.error_type"
	}
	valid := validation.Validation{}
	ok, err := valid.Valid(column)
	if err != nil {
		Glogger.Error(err.Error())
		return "valid.error"
	}
	if !ok {
		return valid.Errors[0].Key + "," + valid.Errors[0].Alisa
	}
	return "valid.success"
}

// 返回值
type Response struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"` // 集合总数
}

func (Base) ReturnData(code int, data interface{}, count int64) *Response {
	r := &Response{
		Code: -10,
		Msg:  "code码不存在",
		Data: ResponseData{
			Data:  data,
			Count: count,
		},
	}
	if info, ok := ReturnCode[code]; ok {
		r.Code = code
		r.Msg = info
	}
	return r
}

func (Base) Return(key string, data interface{}, count int64) *Response {
	r := &Response{
		Data: ResponseData{
			Data:  data,
			Count: count,
		},
	}
	var alisa string
	par := strings.Split(key, ",")
	if len(par) == 2 {
		alisa = par[1]
	}

	// TODO: 添加前端需要特定code做判断用
	switch par[0] {
	case "success":
		r.Code = 0
		r.Msg = "操作成功"
	case "captcha.request_too_more":
		r.Code = 130
		r.Msg = "验证码获取次数超过限制,请稍候再试"
	case "login.failure_or_timeout":
		r.Code = -7
		r.Msg = "登录失败或超时"
	default:
		r.Code = -1
		r.Msg = ParseResponseMsg(par[0], alisa)
	}
	return r
}
