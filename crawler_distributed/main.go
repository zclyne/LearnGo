package main

import (
	"flag"
	"learngo.com/crawler/engine"
	"learngo.com/crawler/scheduler"
	"learngo.com/crawler/zhenai/parser"
	"learngo.com/crawler_distributed/config"
	itemsaver "learngo.com/crawler_distributed/persist/client"
	"learngo.com/crawler_distributed/rpcsupport"
	worker "learngo.com/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	// engine开始运行
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun", // 种子页面为城市列表页
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList,
		), // 城市列表页的对应Parser
	})

}

// 创建rpc client的连接池
func createClientPool(hosts[] string) chan *rpc.Client {
	// 对每一个host，创建一个rpc client
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connectint to %s: %v", h, err)
		}
	}
	// 创建管道并分发
	out := make(chan *rpc.Client)
	go func() {
		for { // 外面再套一层for是为了不断地发送client，否则一轮遍历结束后，go routine就退出了，不再继续向管道内发送client
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}