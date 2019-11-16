package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"learngo.com/crawler_distributed/config"
	"learngo.com/crawler_distributed/persist"
	"learngo.com/crawler_distributed/rpcsupport"
)

// ItemSaver RPC的服务器端，负责接收rpc调用并将item存储到ElasticSearch中

func main() {
	err := serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

func serveRpc(host, index string) error {
	client ,err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	// 为防止过多次数的拷贝，ItemSaverService使用的是指针接收者，所以这里调用时要用&persist
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index: index,
	})
}