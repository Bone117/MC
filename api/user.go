package api

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/model/common/response"
	Req "server/model/request"
	Res "server/model/response"
	"time"

	"server/utils"

	"go.uber.org/zap"

	"github.com/mojocn/base64Captcha"

	"github.com/gin-gonic/gin"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

func (b *BaseApi) Captcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.LOG.Error("验证码获取失败！")
		response.FailWithMessage("验证码获取失败", ctx)
	} else {
		response.OkWithDetailed(Res.CaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.CONFIG.Captcha.KeyLong,
		}, "验证码获取成功", ctx)
	}
}

func (b *BaseApi) Register(ctx *gin.Context) {
	var r Req.Register
	_ = ctx.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}

	var authorities []model.Authority
	if len(r.AuthorityIds) == 0 {
		authorities = append(authorities, model.Authority{
			AuthorityId: "444",
		})
	} else {
		for _, v := range r.AuthorityIds {
			authorities = append(authorities, model.Authority{
				AuthorityId: v,
			})
		}
	}

	user := &model.User{Username: r.Username, NickName: r.NickName, Password: r.Password, Phone: r.Phone, Email: r.Email, Authorities: authorities}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(Res.UserResponse{User: userReturn}, "注册失败", ctx)
	} else {
		response.OkWithDetailed(Res.UserResponse{User: userReturn}, "注册成功", ctx)
		//b.tokenNext(ctx, *user)
	}
}

func (b *BaseApi) File(ctx *gin.Context) {
	filePath := ctx.Param("file_path")
	ctx.File("./uploads/file/" + filePath)
	//user := &model.User{Username: r.Username, NickName: r.NickName, Password: r.Password, Phone: r.Phone, Authorities: authorities}
	//userReturn, err := userService.Register(*user)
	//if err != nil {
	//	global.LOG.Error("注册失败!", zap.Error(err))
	//	response.FailWithDetailed(Res.UserResponse{User: userReturn}, "注册失败", ctx)
	//} else {
	//	response.OkWithDetailed(Res.UserResponse{User: userReturn}, "注册成功", ctx)
	//	//b.tokenNext(ctx, *user)
	//}
}

func (b *BaseApi) Login(ctx *gin.Context) {
	// 获取参数
	var l Req.Login
	_ = ctx.ShouldBindJSON(&l)
	//// 数据验证
	//if err := utils.Verify(l, utils.LoginVerify); err != nil {
	//	response.FailWithMessage(err.Error(), ctx)
	//	return
	//}
	u := &model.User{Username: l.Username, Password: l.Password}
	if user, err := userService.Login(u); err != nil {
		global.LOG.Error("登录失败！用户名不存在或密码错误！", zap.Error(err))
		response.FailWithMessage("用户名不存在或密码错误", ctx)
	} else {
		// 发放token
		b.tokenNext(ctx, *user)
	}
}

func (b *BaseApi) tokenNext(ctx *gin.Context, user model.User) {
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(Req.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		AuthorityId: user.Authorities[0].AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", ctx)
		return
	}

	response.OkWithDetailed(Res.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", ctx)
}

func (b *BaseApi) ChangePassword(ctx *gin.Context) {
	var changeUser Req.ChangePasswordStruct
	_ = ctx.ShouldBindJSON(&changeUser)
	if err := utils.Verify(changeUser, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	u := &model.User{Username: changeUser.Username, Password: changeUser.Password}
	if _, err := userService.ChangePassword(u, changeUser.NewPassword); err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败,原密码错误!", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

func (b *BaseApi) ForgotPassword(ctx *gin.Context) {
	var user Req.ForgotPassword
	_ = ctx.ShouldBindJSON(&user)
	keyInfo := map[string]interface{}{
		"username": user.Username,
		"nickName": user.NickName,
		"phone":    user.Phone,
		"email":    user.Email,
	}
	if u, err := userService.GetUserInfoByKeys(keyInfo); err != nil {
		global.LOG.Error("信息错误!", zap.Error(err))
		response.FailWithMessage("信息有误或该用户不存在", ctx)
	} else {
		if err = userService.ResetPassword(u.ID); err != nil {
			global.LOG.Error("重置失败!", zap.Error(err))
			response.FailWithMessage("重置失败"+err.Error(), ctx)
		} else {
			response.OkWithMessage("重置成功密码为:123456！", ctx)
		}
	}
}

func (b *BaseApi) ResetPassword(ctx *gin.Context) {
	var user model.User
	_ = ctx.ShouldBindJSON(&user)
	if err := userService.ResetPassword(user.ID); err != nil {
		global.LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("重置成功", ctx)
	}
}

func (b *BaseApi) SetUserAuthority(ctx *gin.Context) {
	var sua Req.SetUserAuthorities
	_ = ctx.ShouldBindJSON(&sua)
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), ctx)
		return
	}
	userID := utils.GetUserID(ctx)
	if err := userService.SetUserAuthorities(userID, sua.AuthorityIds); err != nil {
		global.LOG.Error("修改失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
	}
}

func (b *BaseApi) SetUserInfo(ctx *gin.Context) {
	var user Req.ChangeUserInfo
	_ = ctx.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err := userService.SetUserAuthorities(user.ID, user.AuthorityIds)
	if err != nil {
		global.LOG.Error("更新角色权限失败!", zap.Error(err))
		response.FailWithMessage("更新角色权限失败", ctx)
	}
	updateUser := model.User{
		MODEL: global.MODEL{
			ID: user.ID,
		},
		NickName: user.NickName,
		Phone:    user.Phone,
		Email:    user.Email,
	}
	if user.Password != "" {
		updateUser.Password = utils.BcryptHash(user.Password)
	}
	if err := userService.SetUserInfo(updateUser); err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

func (b *BaseApi) SetSelfInfo(ctx *gin.Context) {
	var user Req.ChangeUserInfo
	_ = ctx.BindJSON(&user)
	user.ID = utils.GetUserID(ctx)
	if err := userService.SetUserInfo(model.User{
		MODEL: global.MODEL{
			ID: user.ID,
		},
		NickName: user.NickName,
		Phone:    user.Phone,
		Email:    user.Email,
	}); err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

func (b *BaseApi) DeleteUser(ctx *gin.Context) {
	var reqId request.GetById
	_ = ctx.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	jwtId := utils.GetUserID(ctx)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败，不能删除自己", ctx)
		return
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}

}

func (b *BaseApi) GetUserInfo(ctx *gin.Context) {
	uuid := utils.GetUserUuid(ctx)
	fmt.Println(uuid)
	if ReqUser, err := userService.GetUserInfo(uuid); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", ctx)
	}
}

func (b *BaseApi) GetUserList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	var list interface{}
	var total int64
	var err error
	if _, ok := pageInfo.Keyword["authorityId"]; ok {
		list, total, err = userService.GetUserListByAuthorityID(pageInfo)
	} else {
		list, total, err = userService.GetUserInfoList(pageInfo)
	}
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

func (b *BaseApi) GetStage(ctx *gin.Context) {
	var currentTime request.GetStage
	//_ = ctx.ShouldBindQuery(&currentTime)
	_ = ctx.ShouldBindJSON(&currentTime)
	currentT, _ := time.ParseInLocation("2006-01-02 15:04:05", currentTime.CurrentTime, time.Local)
	//println(currentT)
	if stage, err := stageService.GetStage(currentT); err != nil {
		global.LOG.Error("比赛时间获取失败!", zap.Error(err))
		response.FailWithDetailed(stage, "当前暂无比赛", ctx)
	} else {
		response.OkWithDetailed(stage, "比赛时间获取成功", ctx)
	}
}
