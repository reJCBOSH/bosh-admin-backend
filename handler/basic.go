package handler

import (
    "bosh-admin/core/ctx"
    "bosh-admin/core/log"
    "bosh-admin/dao/dto"
    "bosh-admin/global"
    "bosh-admin/service/upload"
    "bosh-admin/utils"

    "github.com/mojocn/base64Captcha"
)

type BasicHandler struct {
    ossSvc *upload.OssSvc
}

func NewBasicHandler() *BasicHandler {
    return &BasicHandler{
        ossSvc: upload.NewOssSvc(),
    }
}

func (h *BasicHandler) Health(c *ctx.Context) {
    c.Success("服务正常")
}

func (h *BasicHandler) Captcha(c *ctx.Context) {
    capConfig := global.Config.Captcha
    driverDigit := &base64Captcha.DriverDigit{
        Height:   capConfig.ImgHeight,
        Width:    capConfig.ImgWidth,
        Length:   capConfig.KeyLong,
        MaxSkew:  0.7,
        DotCount: 80,
    }
    id, b64s, answer, err := utils.GenerateCaptcha("digit", utils.DriverParam{DriverDigit: driverDigit})
    if c.HandlerError(err, "验证码获取失败") {
        return
    }
    log.Debug(answer)
    c.SuccessWithData(dto.CaptchaResp{
        CaptchaId:     id,
        PicPath:       b64s,
        CaptchaLength: capConfig.KeyLong,
    })
}

func (h *BasicHandler) Upload(c *ctx.Context) {
    file, err := c.FormFile("file")
    if c.HandlerError(err, "上传失败") {
        return
    }
    where := c.PostForm("where")
    resource, err := h.ossSvc.Upload(file, where, "api", c.ClientIP())
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(dto.UploadResp{
        Id:     resource.Id,
        Status: "success",
        Name:   resource.FileName,
        Url:    resource.FullPath,
    })
}
