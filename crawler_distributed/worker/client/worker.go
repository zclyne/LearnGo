package client

import (
	"fmt"
	"learngo.com/crawler/engine"
	"learngo.com/crawler_distributed/config"
	"learngo.com/crawler_distributed/rpcsupport"
	"learngo.com/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		// 序列化request，使其能够传输给rpc做参数
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		// 调用rpc来获取结果
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		// 对rpc获得的结果做反序列化
		return worker.DeserializeResult(sResult), nil
	}, nil
}