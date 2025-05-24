package dto

type CaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`     // 验证码Id
	PicPath       string `json:"picPath"`       // 验证码图片
	CaptchaLength int    `json:"captchaLength"` // 验证码长度
}
