package persist

import (
	"context"
	"errors"
	"learngo.com/crawler/engine"
	"log"
	"github.com/olivere/elastic/v7"
)

// item持久化

func ItemSaver(index string) (chan engine.Item, error) {
	// 我们在内网docker上使用ElasticSearch，所以必须设置sniff为false
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := save(client, index, item)
			if err != nil { // 由于爬虫爬取的数据非常多，所以出现一个存储错误问题不大
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

// 把item存入ElasticSearch，返回error
func save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	// 开始存储，Index()方法既可以创建，也可以修改
	// ElasticSearch中的路径：index/type/id
	// index相当于sql中的数据库名称，type相当于表名称，id就是id
	// 此处id自动分配，所以不需要指定
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	// 用%+v打印结构体时，会把结构体中的字段名也打印出来
	// fmt.Printf("%+v", resp)
	return nil
}