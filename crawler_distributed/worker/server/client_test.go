package main

import (
	"fmt"
	"learngo.com/crawler_distributed/config"
	"learngo.com/crawler_distributed/rpcsupport"
	"learngo.com/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url:    "https://album.zhenai.com/u/1866830740",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: []string{"花儿少年", "男士"},
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}