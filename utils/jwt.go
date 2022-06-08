package utils

import (
	"errors"
	"server/global"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired = errors.New("Token is expired")
	TokenInvalid = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.JWT.SigningKey),
	}
}
