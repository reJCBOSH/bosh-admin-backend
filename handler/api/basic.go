package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/core/log"
	"bosh-admin/dao/dto"
	"bosh-admin/global"
	"bosh-admin/utils"
	"github.com/mojocn/base64Captcha"
)

type BasicHandler struct{}

func NewBasicHandler() *BasicHandler {
	return &BasicHandler{}
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
	c.SuccessWithData(dto.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: capConfig.KeyLong,
	})
}
