package persist

import (
	"context"
	"log"
	"github.com/olivere/elastic/v7"
)

// item持久化

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			save(item)
		}
	}()
	return out
}

// 把item存入ElasticSearch，返回插入的结果的id，以及一个error
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) // 我们在内网docker上使用ElasticSearch，所以必须设置sniff为false
	if err != nil {
		return "", nil
	}

	// 开始存储，Index()方法既可以创建，也可以修改
	// ElasticSearch中的路径：index/type/id
	// index相当于sql中的数据库名称，type相当于表名称，id就是id
	// 此处id自动分配，所以不需要指定
	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", nil
	}

	// 用%+v打印结构体时，会把结构体中的字段名也打印出来
	// fmt.Printf("%+v", resp)
	return resp.Id, nil
}