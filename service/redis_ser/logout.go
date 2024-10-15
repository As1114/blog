package redis_ser

import (
	"context"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/utils"
	"time"
)

const prefix = "logout_"

func Logout(token string, diff time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	err := global.Redis.Set(ctx, prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	keys := global.Redis.Keys(ctx, prefix+"*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
