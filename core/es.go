package core

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/nsxz1114/blog/global"
)

func InitEs() *elasticsearch.TypedClient {
	dsn := global.Config.Es.Dsn()
	cfg := elasticsearch.Config{
		Addresses: []string{
			global.Config.Es.Dsn(),
		},
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		global.Log.Fatal(fmt.Sprintf("[%s] es连接失败", dsn))
	}
	return es
}
