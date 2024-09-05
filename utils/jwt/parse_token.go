package jwt

import (
	"blog/global"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
)

func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error("parse token error:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		global.Log.Error("parse token error:", err)
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
