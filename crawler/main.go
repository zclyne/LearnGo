package main

import (
	"LearnGo/crawler/engine"
	"LearnGo/crawler/scheduler"
	"LearnGo/crawler/zhenai/parser"
)

func main() {

	// engine开始运行
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
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