package core

import (
	"github.com/axis1114/blog/global"
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"go.uber.org/zap"
)

func InitAddrDB() *geoip2.DBReader {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		global.Log.Fatal("InitAddrDB err:", zap.Error(err))
	}
	return db
}
