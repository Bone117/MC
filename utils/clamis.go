package utils

import (
	"server/global"
	Req "server/model/request"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetClaims(ctx *gin.Context) (*Req.CustomClaims, error) {
	token := ctx.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

func GetUserID(ctx *gin.Context) uint {
	if claims, exists := ctx.Get("claims"); !exists {
		if cl, err := GetClaims(ctx); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*Req.CustomClaims)
		return waitUse.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*Req.CustomClaims)
		return waitUse.UUID
	}
}
