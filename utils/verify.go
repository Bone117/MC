package utils

var (
	LoginVerify    = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify = Rules{"username": {NotEmpty()}, "password": {NotEmpty()}, "authorityId": {NotEmpty()}}
)
