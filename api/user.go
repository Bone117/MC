package api

import (
	"server/global"
	"server/model"
	"server/model/common/response"
	Req "server/model/request"
	Res "server/model/response"

	"server/utils"

	"go.uber.org/zap"

	"github.com/mojocn/base64Captcha"

	"github.com/gin-gonic/gin"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct {
}

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

	user := &model.User{Username: r.Username, Password: r.Password, AuthorityId: r.AuthorityId}
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
		if err, user := userService.Login(u); err != nil {
			global.LOG.Error("登录失败！用户名不存在或密码错误！", zap.Error(err))
			response.FailWithMessage("用户名不存在或密码错误", ctx)
		} else {
			// 发放token
			b.tokenNext(ctx, *user)
			//response.OkWithDetailed(user, "成功", ctx)
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
		AuthorityId: user.AuthorityId,
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
