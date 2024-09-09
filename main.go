package main

import (
	"fmt"
	"github.com/axis1114/blog/core"
	"github.com/axis1114/blog/flags"
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/routers"
	"github.com/axis1114/blog/utils"
	"go.uber.org/zap"
)

// @title github.com/axis1114/blog
// @version 1.0
// @contact.name Axios
// @contact.email 1790146932@qq.com
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	core.InitConf()
	global.Log = core.InitLog()
	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()
	global.Es = core.InitEs()
	global.Etcd = core.InitEtcd()
	core.InitAddrDB()
	utils.Init(global.Config.System.StartTime, global.Config.System.MachineID)
	flags.Newflags()
	err := utils.InitTrans("en")
	if err != nil {
		global.Log.Error("fail to init trans", zap.Error(err))
		return
	}
	utils.PrintSystem()
	router := routers.InitRouter()
	err = router.Run(fmt.Sprintf(":%d", global.Config.System.Port))
	if err != nil {
		global.Log.Error("fail to start server", zap.Error(err))
		return
	}
}
