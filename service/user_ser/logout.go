package user_ser

import (
	"blog/service/redis_ser"
	"blog/utils"
	"time"
)

func Logout(claims *utils.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}
