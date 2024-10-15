package utils

import (
	"github.com/nsxz1114/blog/global"
	"go.uber.org/zap"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) {
	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		global.Log.Error("parse start time error:", zap.Error(err))
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	if err != nil {
		global.Log.Error("new node error:", err.Error())
		return
	}
}
func GenerateID() int64 {
	return node.Generate().Int64()
}
