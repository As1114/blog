package user_ser

import (
	"blog/service/redis_ser"
	"blog/utils/jwt"
	"time"
)

func (UserService) Logout(claims *jwt.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}
