package main

import (
	"learngo.com/crawler/engine"
	"learngo.com/crawler/persist"
	"learngo.com/crawler/scheduler"
	"learngo.com/crawler/zhenai/parser"
)

func main() {

	// engine开始运行
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun", // 种子页面为城市列表页
		ParserFunc: parser.ParseCityList, // 城市列表页的对应Parser
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})

}