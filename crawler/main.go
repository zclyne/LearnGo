package main

import (
	"learngo.com/crawler/engine"
	"learngo.com/crawler/persist"
	"learngo.com/crawler/scheduler"
	"learngo.com/crawler/zhenai/parser"
)

func main() {

	// 创建负责数据持久化的与ElasticSearch通信的通道，参数index指定要存入的ElasticSearch的index
	itemChan, err := persist.ItemSaver("dating_profile")
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