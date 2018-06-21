package controller

import (
	"bytes"
	"encoding/base64"
	. "global"
	"image/jpeg"
	"image/png"
	"regexp"
	"strings"
)

type Account struct {
	Base
}

// 验证头像
func checkAvatar(avatar string) string {
	if avatar == "" {
		return "valid.success"
	}
	reg, err := regexp.Compile(`^(?:[A-Za-z0-99+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`)
	if err != nil {
		Glogger.Error(err.Error())
		return "db.error"
	}
	// base64正则校验
	if reg.MatchString(avatar) {
		return "avatar.invalid"
	}
	coi := strings.Index(avatar, ",")
	// base64解密
	b64, err := base64.StdEncoding.DecodeString(avatar[coi+1:])
	if err != nil {
		Glogger.Error(err.Error())
		return "base64.decode_error"
	}
	reader := bytes.NewBuffer(b64)
	// 大小校验
	if reader.Len() > 1024*1024 {
		return "avatar.data_too_long"
	}
	// 类型校验
	switch strings.TrimSuffix(avatar[5:coi], ";base64") {
	case "image/png":
		_, err = png.Decode(reader)
		if err != nil {
			return "avatar.invalid"
		}
	case "image/jpeg":
	case "image/jpg":
		_, err = jpeg.Decode(reader)
		if err != nil {
			Glogger.Error(err.Error())
			return "avatar.invalid"
		}
	default:
		return "avatar.type_error"
	}
	return "valid.success"
}
