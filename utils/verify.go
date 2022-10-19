package utils

var (
	LoginVerify            = Rules{"captchaId": {NotEmpty()}, "captcha": {NotEmpty()}, "username": {NotEmpty()}, "password": {NotEmpty()}}
	RegisterVerify         = Rules{"username": {NotEmpty()}, "password": {NotEmpty()}, "authorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"username": {NotEmpty()}, "password": {NotEmpty()}, "newPassword": {NotEmpty()}}
	IdVerify               = Rules{"ID": {NotEmpty()}}
	AuthorityVerify        = Rules{"authorityId": {NotEmpty()}, "authorityName": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"authorityId": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"authorityId": {NotEmpty()}}

	NoticeVerify = Rules{"title": {NotEmpty()}, "desc": {NotEmpty()}, "content": {NotEmpty()}}

	SignVerify = Rules{"workName": {NotEmpty()}, "workFileTypeId": {NotEmpty()}, "workSoftware": {NotEmpty()}, "workDesc": {NotEmpty()}, "majorId": {NotEmpty()}, "gradeName": {NotEmpty()}}

	PageInfoVerify = Rules{"page": {NotEmpty()}, "pageSize": {NotEmpty()}}
)
