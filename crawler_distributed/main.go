package main

import (
	"fmt"
	"learngo.com/crawler/engine"
	"learngo.com/crawler/scheduler"
	"learngo.com/crawler/zhenai/parser"
	"learngo.com/crawler_distributed/config"
	itemsaver "learngo.com/crawler_distributed/persist/client"
	worker "learngo.com/crawler_distributed/worker/client"
)

func main() {

	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	// engine开始运行
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
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