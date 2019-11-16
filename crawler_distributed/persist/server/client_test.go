package main

import (
	"learngo.com/crawler/engine"
	"learngo.com/crawler/model"
	"learngo.com/crawler_distributed/config"
	"learngo.com/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(2 * time.Second) // 为了防止server还没完全启动时client就连接上去，这里要睡眠1s

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Id: "1866830740",
		Url: "https://album.zhenai.com/u/1866830740",
		Type: "zhenai",
		Payload: model.Profile{
			Name:          "花儿少年",
			Gender:        "男士",
			Age:           23,
			Height:        175,
			Weight:        67,
			Income:        "5-8千",
			Marriage:      "未婚",
			Education:     "大专",
			AncestralHome: "重庆",
			Constellation: "射手座",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s, err: %s", result, err)
	}
}