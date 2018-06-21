package controller

import (
	"bytes"
	"config"
	"encoding/base64"
	"errors"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/lifei6671/gocaptcha"
	"github.com/satori/go.uuid"
	."global"
	"strconv"
	"strings"
	"time"
	"CMDB/src/model"
)

type Login struct {
	Base
}

// 验证码
func (lc *Login) Captcha(ctx echo.Context) error {
	// 获取程序运行路径
	p, err := GetExecPath()
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("system.get_exec_path_error", nil, 0))
	}

	// 读取验证码字体文件
	err = gocaptcha.ReadFonts(p+"/etc/fonts", ".ttf")
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("system.get_fonts_error", nil, 0))
	}

	// 初始化一个验证码对象
	captchaImage, err := gocaptcha.NewCaptchaImage(150, 50, gocaptcha.RandLightColor())
	// 画上一条随机直线
	captchaImage.Drawline(0)
	// 画随机噪点
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexMedium)
	// 画随机文字噪点
	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexMedium)
	// 画验证码文字
	str := gocaptcha.RandText(4)
	captchaImage.DrawText(str)
	// 画边框
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x00FFFFFF))
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("captcha.init_error", nil, 0))
	}

	// 生成一个uuid字符串
	key, err := uuid.NewV4()
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("uuid.init_error", nil, 0))
	}

	// 生成返回结构体数据
	reader := new(bytes.Buffer)
	captchaImage.SaveImage(reader, gocaptcha.ImageFormatJpeg)
	encodeString := base64.StdEncoding.EncodeToString(reader.Bytes())
	data := &struct {
		Code  string `json:"code"`
		Image string `json:"image"`
	}{
		Code:  key.String(),
		Image: "data:image/jpeg;base64," + encodeString,
	}

	// 验证码key和code存入redis
	err = GetCaptcha().Set(key.String(), str, time.Minute*time.Duration(config.Conf.Expires.Captcha)).Err()
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("redis.error", nil, 0))
	}

	// 获取请求ip
	ip := ctx.Get("pre").(string) + strconv.Itoa(StringIpToInt(ctx.RealIP()))
	// 查询redis中是否存在并获取value
	ec := numLimit(GetCaptcha(), ip, 50, 2*time.Minute, 2*time.Minute)
	if ec != "valid.success" {
		return ctx.JSON(200, lc.Return(ec, nil, 0))
	}
	return ctx.JSON(200, lc.Return("success", data, 0))
}

// 登录
func (lc *Login) SignIn(ctx echo.Context) error {
	l := new(model.InputLogin)
	// 基础数据校验
	key := lc.ValidParam(l, ctx)
	if key != "valid.success" {
		return ctx.JSON(200, lc.Return(key, nil, 0))
	}

	// 验证码校验
	del := validCaptcha(l.Key, l.Code)
	if del != "valid.success" {
		return ctx.JSON(200, lc.Return(del, nil, 0))
	}

	// 密码验证
	flag := CheckPassword(l.Password)
	if !flag {
		return ctx.JSON(200, lc.Return("password.invalid", nil, 0))
	}

	// 获取账号是否存在
	adminBean := new(model.Admin)
	ok, err := adminBean.GetByAccount(l.Account)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("account.not_found", nil, 0))
	}

	if !ok {
		// ldap验证
		pass := validLdap(l.Account, l.Password)
		if pass != "valid.success" {
			return ctx.JSON(200, lc.Return(pass, nil, 0))
		} else {
			// 查询普通用户角色是否存在
			roleBean := new(model.Role)
			ok, err := roleBean.GetById(RoleOU)
			if err != nil {
				Glogger.Error(err.Error())
				return ctx.JSON(200, lc.Return("db.error", nil, 0))
			}
			if !ok {
				return ctx.JSON(200, lc.Return("role.not_found", nil, 0))
			}
			// 创建cmdb账号并关联该ldap
			row, err := model.Admin{}.Create(&model.Admin{
				Account:       l.Account,
				Password:      l.Password,
				Status:        StatusEnable,
				ExamineStatus: ExaminePass,
				RegType:       RegByLdap,
				AuthType:      1,
				RoleId:        RoleOU,
				RoleName:      roleBean.Title,
			})
			if err != nil {
				Glogger.Error(err.Error())
				return ctx.JSON(200, lc.Return("db.error", nil, 0))
			}
			if row != 1 {
				Glogger.Error(err.Error())
				return ctx.JSON(200, lc.Return("ldap.create_error", nil, 0))
			}
		}
	} else {
		// 账号是否为ldap认证
		if adminBean.AuthType != 0 {
			pass := validLdap(l.Account, l.Password)
			if pass != "valid.success" {
				if pass == "login.incorrect_account_or_password" {
					// 更新登录失败次数
					ec := numLimit(GetLogin(), l.Account, 3, 2*time.Minute, 10*time.Minute)
					if ec != "valid.success" {
						return ctx.JSON(200, lc.Return(ec, nil, 0))
					}
				}
				return ctx.JSON(200, lc.Return(pass, nil, 0))
			}
		} else {
			if Md5Password(l.Password) != adminBean.Password {
				// 更新登录失败次数
				ec := numLimit(GetLogin(), l.Account, 3, 2*time.Minute, 10*time.Minute)
				if ec != "valid.success" {
					return ctx.JSON(200, lc.Return(ec, nil, 0))
				}
				return ctx.JSON(200, lc.Return("login.incorrect_account_or_password", nil, 0))
			}
		}
	}

	// 重新获取账号
	adminBean = new(model.Admin)
	_, err = adminBean.GetByAccount(l.Account)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("account.not_found", nil, 0))
	}

	if adminBean.ExamineStatus == ExamineRefuse {
		return ctx.JSON(200, lc.Return("account.is_refuse", nil, 0)) // 账号被拒绝
	} else if adminBean.ExamineStatus == ExamineApprove {
		return ctx.JSON(200, lc.Return("account.is_approve", nil, 0)) // 账号待审核
	}

	// 账号被禁用
	if adminBean.Status == StatusDisable {
		return ctx.JSON(200, lc.Return("account.is_disable", nil, 0))
	}

	// 查询是否重复登录
	if adminBean.Token != "" {
		err := downLine(adminBean.Token)
		if err != nil {
			return ctx.JSON(200, lc.Return("db.error", nil, 0))
		}
	}

	// 生成token及redis登录数据存储
	token, err := generateLoginDataForRedis(adminBean.Id)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("system.error", nil, 0))
	}

	// 相关表更新操作
	err = upDb(adminBean, token)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("system.error", nil, 0))
	}

	// 添加登陆成功日志
	op := new(model.LoginLog)
	opp := &model.LoginLog{
		Account:   adminBean.Account,
		LoginTime: time.Now().Unix(),
		LoginIp:   StringIpToInt(ctx.RealIP()),
		Agent:     ctx.Request().Header.Get("User-Agent"),
	}
	count, err := op.AddLoginLog(opp)
	if err != nil {
		return ctx.JSON(200, lc.Return("db.error", nil, 0))
	}
	if count == 0 {
		return ctx.JSON(200, lc.Return("login.log_error", nil, 0))
	}

	// 登录成功返回数据
	bd := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	return ctx.JSON(200, lc.Return("success", bd, 0))
}

// 验证码校验
func validCaptcha(key, code string) string {
	// 验证码校验
	val, err := GetCaptcha().Get(key).Result()
	if err == redis.Nil {
		return "captcha.non_existence_or_expired"
	} else if err != nil {
		Glogger.Error(err.Error())
		return "redis.error"
	}
	// 验证码不一致
	if strings.ToLower(val) != strings.ToLower(code) {
		return "captcha.error"
	}
	// 删除验证码
	err = GetCaptcha().Del(key).Err()
	if err != nil {
		Glogger.Error(err.Error())
		return "redis.error"
	}
	return "valid.success"
}

// 操作限制
// conn redis连接库
// key redis key
// max 最多限制几次
// in 多少时间内限制
// after 多少时间后进行操作
func numLimit(conn *redis.Client, key string, max int, in, after time.Duration) string {
	result, err := conn.Get(key).Result()
	if err == redis.Nil {
		err = conn.Set(key, 1, in).Err()
		if err != nil {
			Glogger.Error(err.Error())
			return "redis.empty_key"
		}
	} else if err != nil {
		Glogger.Error(err.Error())
		return "redis.error"
	} else {
		num, err := strconv.Atoi(result)
		if err != nil {
			Glogger.Error(err.Error())
			return "system.error"
		}
		n := num + 1
		if n >= max {
			err = conn.Set(key, max, after).Err()
			if err != nil {
				Glogger.Error(err.Error())
				return "redis.error"
			}
		} else {
			exp, err := conn.TTL(key).Result()
			if err != nil {
				Glogger.Error(err.Error())
				return "redis.error"
			}
			err = conn.Set(key, n, exp).Err()
			if err != nil {
				Glogger.Error(err.Error())
				return "redis.error"
			}
		}
	}
	return "valid.success"
}

// 验证ldap
func validLdap(account, password string) string {
	// TODO:目前只配置一台ldap服务器
	// 查询ldap配置信息
	ldapBean := new(model.Ldap)
	ok, err := ldapBean.GetById(1)
	if err != nil {
		Glogger.Error(err.Error())
		return "db.error"
	}
	// 如果没有配置
	if !ok {
		return "ldap.non_config"
	}

	// 连接测试
	ok = ldapBean.Dial()
	if !ok {
		return "ldap.server_connect_error"
	}

	// 验证账号
	ok = ldapBean.Auth(account, password)
	if ok {
		return "valid.success"
	} else {
		return "login.incorrect_account_or_password"
	}
}

// 生成登录redis数据并存储
func generateLoginDataForRedis(id int64) (string, error) {
	// 生成token
	token, err := uuid.NewV4()
	if err != nil {
		Glogger.Error(err.Error())
		return "", err
	}
	t := token.String()

	// token => expTime 存入redis
	return t, GetLogin().Set(t, id, time.Minute*time.Duration(config.Conf.Expires.Login)).Err()
}

// 登录成功更新数据
func upDb(admin *model.Admin, token string) error {
	// 更新登录时间及token
	admin.LoginTime = time.Now().Unix()
	admin.Token = token
	row, err := admin.UpdateLogin()
	if err != nil {
		return err
	}
	if row != 1 {
		return errors.New(ErrorMsg["login"]["failure_or_timeout"])
	}
	return nil
}

// 踢线
func downLine(key string) error {
	_, err := model.Admin{}.DelToken(key)
	if err != nil {
		Glogger.Error(err.Error())
		return nil
	}
	err = GetLogin().Del(key).Err()
	if err == redis.Nil {
	} else if err != nil {
		Glogger.Error(err.Error())
		return err
	}
	return nil
}

// 获取登录人信息
func GetInfo(key string) *model.Admin {
	adminBean := new(model.Admin)
	adminBean.GetByToken(key)
	return adminBean
}

// 修改密码
func (lc *Login) ModifyPassword(ctx echo.Context) error {
	mp := new(model.InputModifyPassword)
	key := lc.ValidParam(mp, ctx)
	if key != "valid.success" {
		return ctx.JSON(200, lc.Return(key, nil, 0))
	}

	// 新密码与旧密码不能相同
	if mp.OldPassword == mp.NewPassword {
		return ctx.JSON(200, lc.Return("password.cannot_the_same", nil, 0))
	}

	// 新密码与重复密码不一致
	if mp.NewPassword != mp.ReplyNewPassword {
		return ctx.JSON(200, lc.Return("password.inequality", nil, 0))
	}

	// 原密码格式不正确
	if !CheckPassword(mp.OldPassword) {
		return ctx.JSON(200, lc.Return("password.invalid", nil, 0))
	}

	// 新密码格式不正确
	if !CheckPassword(mp.NewPassword) {
		return ctx.JSON(200, lc.Return("password.invalid", nil, 0))
	}

	// 验证原密码
	token := ctx.Get("token").(string)
	user := GetInfo(token)
	if Md5Password(mp.OldPassword) != user.Password {
		return ctx.JSON(200, lc.Return("password.error", nil, 0))
	}

	// 修改密码
	adminBean := new(model.Admin)
	adminBean.Id = user.Id
	adminBean.Password = Md5Password(mp.NewPassword)
	row, err := adminBean.ModifyPassword(mp.NewPassword)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("db.error", nil, 0))
	}
	if row != 1 {
		return ctx.JSON(200, lc.Return("update.error", nil, 0))
	}

	// 踢线
	err = downLine(token)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("downline.error", nil, 0))
	}
	return ctx.JSON(200, lc.Return("success", nil, 0))
}

// 获取登录账号基本信息
func (lc *Login) Info(ctx echo.Context) error {
	user := GetInfo(ctx.Get("token").(string))
	ui := model.BackUserInfo{
		Id:       user.Id,
		Avatar:   string(user.Avatar),
		Account:  user.Account,
		RoleName: user.RoleName,
		Username: user.Username,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Job:      user.Job,
		Remark:   user.Remark,
	}
	return ctx.JSON(200, lc.Return("success", ui, 0))
}

// 修改登录人信息
func (lc *Login) ModifyInfo(ctx echo.Context) error {
	mu := new(model.InputModifyUser)
	key := lc.ValidParam(mu, ctx)
	if key != "valid.success" {
		return ctx.JSON(200, lc.Return(key, nil, 0))
	}

	user := GetInfo(ctx.Get("token").(string))

	// 验证base64是否是正常图片,大小不超过1M,格式限定为(jpg,jpeg,png)
	pass := checkAvatar(mu.Avatar)
	if pass != "valid.success" {
		return ctx.JSON(200, lc.Return(pass, nil, 0))
	}

	// 更新个人资料
	adminBean := &model.Admin{
		Id:       user.Id,
		Avatar:   []byte(mu.Avatar),
		Username: mu.Username,
		Mobile:   mu.Mobile,
		Email:    mu.Email,
		Job:      mu.Job,
		Remark:   mu.Remark,
	}
	row, err := adminBean.ModifyUser()
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("db.error", nil, 0))
	}
	if row != 1 {
		return ctx.JSON(200, lc.Return("update.null", nil, 0))
	}
	return ctx.JSON(200, lc.Return("success", nil, 0))
}

// 退出登录
func (lc *Login) LogOut(ctx echo.Context) error {
	token := ctx.Get("token").(string)
	// 删除redis里面的token
	err := GetLogin().Del(token).Err()
	if err == redis.Nil {
	} else if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("success", nil, 0))
	}
	// 删除数据库里面的token
	_, err = model.Admin{}.DelToken(token)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, lc.Return("success", nil, 0))
	}
	return ctx.JSON(200, lc.Return("success", nil, 0))
}
