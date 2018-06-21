package controller

import (
	"fmt"
	"github.com/labstack/echo"
	. "global"
	"mime"
	"path"
	"strconv"
	"CMDB/src/model"
)

type Backups struct {
	Base
}

// 备份上传文件
func (b *Backups) Upload(ctx echo.Context) error {
	// 文件
	backups, err := ctx.FormFile("backups")
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, b.Return("upload.not_file", nil, 0))
	}

	// 判断文件名是否已存在
	has, err := model.Backups{}.ExistByTitle(backups.Filename)
	if err != nil {
		return ctx.JSON(200, b.Return("db.error", nil, 0))
	}
	if has {
		return ctx.JSON(200, b.Return("upload.is_exist", nil, 0))
	}
	// 给文件名赋值
	w := new(model.Backups)
	w.Title = backups.Filename
	// 获取文件二进制内容
	file, err := backups.Open()
	con := make([]byte, backups.Size)
	if len(con) > 0 {
		_, err = file.Read(con)
		if err != nil {
			Glogger.Error(err.Error())
			return ctx.JSON(200, b.Return("system.error", nil, 0))
		}
	}
	file.Close()
	// 获取文件的类型
	typeStr := path.Ext(backups.Filename)
	// 限制文件类型
	if typeStr != ".txt" && typeStr != ".doc" && typeStr != ".xls" && typeStr != ".docx" && typeStr != ".xlsx" {
		return ctx.JSON(200, b.Return("upload.type_error", nil, 0))
	}
	w.FileType = mime.TypeByExtension(typeStr)
	// 限制文件大小为3m
	if backups.Size > THREE_MB {
		return ctx.JSON(200, b.Return("upload.data_too_long", nil, 0))
	}
	w.Content = con
	// 产品id
	pId := ctx.FormValue("id")
	if len(pId) < 1 {
		return ctx.JSON(200, b.Return("id.invalid", nil, 0))
	}
	id, err := strconv.Atoi(pId)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, b.Return("system.error", nil, 0))
	}
	w.DeviceId = id
	// 产品类型
	ty := ctx.FormValue("type")
	typeId, err := strconv.Atoi(ty)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, b.Return("system.error", nil, 0))
	}
	w.DeviceType = int8(typeId)

	// 判断产品id是否存在
	switch typeId {
	case ProductRouter: // 路由器
		has, err = model.Route{}.ExistById(w.DeviceId)
	case ProductSwitchboard: // 交换机
		has, err = model.Switchboard{}.ExistById(w.DeviceId)
	case ProductFirewall: // 防火墙
		has, err = model.Firewall{}.ExistById(w.DeviceId)
	default: // 类型不合法
		return ctx.JSON(200, b.Return("type.invalid", nil, 0))
	}
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, b.Return("db.error", nil, 0))
	}
	if !has {
		return ctx.JSON(200, b.Return("id.not_found", nil, 0))
	}

	// 添加上传备份文件
	ok, err := model.Backups{}.Create(w)
	if err != nil {
		return ctx.JSON(200, b.Return("db.error", nil, 0))
	}
	if !ok {
		return ctx.JSON(200, b.Return("upload.error", nil, 0))
	}
	return ctx.JSON(200, b.Return("success", nil, 0))
}

// 备份下载
func (b *Backups) BackupsDownLoad(ctx echo.Context) error {
	w := new(model.InputBackupsDown)
	key := b.ValidParam(w, ctx)
	if key != "valid.success" {
		return ctx.JSON(200, b.Return(key, nil, 0))
	}

	ty, err := strconv.Atoi(w.Type)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, b.Return("system.error", nil, 0))
	}
	if ty > 4 || ty < 2 {
		return ctx.JSON(200, b.Return("type.error", nil, 0))
	}

	// 备份是否存在
	backupsBean := new(model.Backups)
	has, err := backupsBean.GetById(w.Id)
	if err != nil {
		Glogger.Error(err.Error())
		return ctx.JSON(200, b.Return("db.error", nil, 0))
	}
	if !has {
		return ctx.JSON(200, b.Return("id.not_found", nil, 0))
	}
	ctx.Response().Header().Add("Content-Type", backupsBean.FileType)
	ctx.Response().Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", backupsBean.Title))
	return ctx.Blob(200, backupsBean.FileType, backupsBean.Content)
}
