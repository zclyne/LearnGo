package client

import (
	"learngo.com/crawler/engine"
	"learngo.com/crawler_distributed/config"
	"learngo.com/crawler_distributed/worker"
	"net/rpc"
)

// 参数是一个Client的channel，从channel中获取client，并把worker派给它执行
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		// 序列化request，使其能够传输给rpc做参数
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		// 调用rpc来获取结果
		c := <- clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		// 对rpc获得的结果做反序列化
		return worker.DeserializeResult(sResult), nil
	}
}