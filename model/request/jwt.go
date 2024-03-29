package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	AuthorityId string
}
