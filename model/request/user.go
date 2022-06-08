package request

type Register struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	AuthorityId string `json:"authorityId"`
}

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"` // 验证码ID
}
