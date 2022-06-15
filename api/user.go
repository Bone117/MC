package api

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/model/common/response"
	Req "server/model/request"
	Res "server/model/response"

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
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, model.Authority{
			AuthorityId: v,
		})
	}

	user := &model.User{Username: r.Username, NickName: r.NickName, Password: r.Password, Authorities: authorities}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(Res.UserResponse{User: userReturn}, "注册失败", ctx)
	} else {
		response.OkWithDetailed(Res.UserResponse{User: userReturn}, "注册成功", ctx)
	}
}

func (b *BaseApi) Login(ctx *gin.Context) {
	// 获取参数
	var l Req.Login
	_ = ctx.ShouldBindJSON(&l)
	// 数据验证
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &model.User{Username: l.Username, Password: l.Password}
		if user, err := userService.Login(u); err != nil {
			global.LOG.Error("登录失败！用户名不存在或密码错误！", zap.Error(err))
			response.FailWithMessage("用户名不存在或密码错误", ctx)
		} else {
			// 发放token
			b.tokenNext(ctx, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", ctx)
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
	var user Req.ChangePasswordStruct
	_ = ctx.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	u := &model.User{Username: user.Username, Password: user.Password}
	if _, err := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败,原密码错误!", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
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
	if len(user.AuthorityIds) != 0 {
		err := userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			global.LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", ctx)
		}
	}

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
	if list, total, err := userService.GetUserInfoList(pageInfo); err != nil {
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
