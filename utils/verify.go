package utils

var (
	LoginVerify            = Rules{"username": {NotEmpty()}, "password": {NotEmpty()}}
	RegisterVerify         = Rules{"username": {NotEmpty()}, "password": {NotEmpty()}, "authorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"username": {NotEmpty()}, "password": {NotEmpty()}, "newPassword": {NotEmpty()}}
	IdVerify               = Rules{"ID": {NotEmpty()}}
	AuthorityVerify        = Rules{"authorityId": {NotEmpty()}, "authorityName": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"authorityId": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"authorityId": {NotEmpty()}}

	NoticeVerify = Rules{"Title": {NotEmpty()}, "Content": {NotEmpty()}}

	AssignVerify = Rules{"SignId": {NotEmpty()}, "JieCiId": {NotEmpty()}}

	SignVerify = Rules{"WorkName": {NotEmpty()}, "WorkFileTypeId": {NotEmpty()}, "WorkSoftware": {NotEmpty()}, "workDesc": {NotEmpty()}}

	PageInfoVerify = Rules{"page": {NotEmpty()}, "pageSize": {NotEmpty()}}
)
