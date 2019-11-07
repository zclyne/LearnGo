package main

import (
	"LearnGo/crawler/engine"
	"LearnGo/crawler/zhenai/parser"
)

func main() {

	// engine开始运行
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun", // 种子页面为城市列表页
		ParserFunc: parser.ParseCityList, // 城市列表页的对应Parser
	})

}