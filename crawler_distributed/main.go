package main

import (
	"fmt"
	"learngo.com/crawler/engine"
	"learngo.com/crawler/scheduler"
	"learngo.com/crawler/zhenai/parser"
	"learngo.com/crawler_distributed/config"
	"learngo.com/crawler_distributed/persist/client"
)

func main() {

	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	// engine开始运行
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun", // 种子页面为城市列表页
		ParserFunc: parser.ParseCityList, // 城市列表页的对应Parser
	})

}