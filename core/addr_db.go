package core

import (
	"github.com/axis1114/blog/global"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"go.uber.org/zap"
)

func InitAddrDB() {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		global.Log.Error("InitAddrDB err:", zap.Error(err))
	}
	global.AddrDB = db
}
