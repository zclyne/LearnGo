package engine

import (
	"log"
)

// engine是整个爬虫的主控部分，它从seeds接收种子页面的requests，将requests放入任务队列
// 并不断从任务队列中取出request的url交给fetcher，访问相应页面并获得页面内容文本
// 随后，engine把文本内容交给Parser，解析出页面中的内容，并获得后续的requests和items

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	// request任务队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 不断循环地从队列中取出request，执行爬虫
	for len(requests) > 0 {
		// 取出队首的request
		r := requests[0]
		requests = requests[1:]

		// 把request交由worker来进行取网页内容和解析
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...) // ...表示把slice展开并加入到参数列表中

		// 打印出获得的items
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}