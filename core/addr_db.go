package core

import (
	"blog/global"
	geoip2db "github.com/cc14514/go-geoip2-db"
)

func InitAddrDB() {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		global.Log.Error("InitAddrDB err:", err.Error())
	}
	global.AddrDB = db
}
