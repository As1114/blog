package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type PayLoad struct {
	Account string `json:"Account"` // 账号
	Role    int    `json:"role"`    // 权限
	UserID  uint   `json:"user_id"` // 用户id
}

var MySecret []byte

type CustomClaims struct {
	PayLoad
	jwt.StandardClaims
}
