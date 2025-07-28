package dto

type CaptchaResp struct {
    CaptchaId     string `json:"captchaId"`     // 验证码Id
    PicPath       string `json:"picPath"`       // 验证码图片
    CaptchaLength int    `json:"captchaLength"` // 验证码长度
}

type UploadResp struct {
    Id     uint   `json:"id"`
    Status string `json:"status"`
    Name   string `json:"name"`
    Url    string `json:"url"`
}
