package client

import (
	"learngo.com/crawler/engine"
	"learngo.com/crawler_distributed/config"
	"learngo.com/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)

			// Call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if result == "ok" {
				log.Printf("Saved item #%d with RPC", itemCount)
			}
			if err != nil { // 由于爬虫爬取的数据非常多，所以出现一个存储错误问题不大
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
			itemCount++
		}
	}()
	return out, nil
}