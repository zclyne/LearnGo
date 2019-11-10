package engine

import (
	"learngo.com/crawler/model"
	"log"
)

// 并发引擎
// 并发engine拥有一个scheduler负责调度engine发给它的requests
// scheduler把request发送给worker，由worker来执行获取网页并解析的功能

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{} // 用于爬取的信息持久化的channel，与ItemSaver通信
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// 原本的WorkerReady()是Scheduler的一个方法，但是这样Scheduler太大了
// 所以单独分离出一个ReadyNotifier
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// QueuedScheduler初始化并开始运行
	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 创建worker，并将先前创建的channel与worker连接上
	for i := 0; i < e.WorkerCount; i++ {
		// e.Scheduler.WorkerChan()是scheduler内部的channel
		// 对于SimpleScheduler，所有worker共享一个chan
		// 而对于QueuedScheduler，每一个worker有自己独立的chan
		// 由于e.Scheduler中包含了ReadyNotifier，所以这里仍然可以直接传入e.Scheduler
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 把种子页面发送给scheduler，开始调度
	for _, r := range seeds {
		// 判断种子页面中是否有重复
		if isDuplicate(r.Url) {
			log.Printf("Duplicate request: " + "%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	// 无穷循环
	for {
		// 从输出管道中获取各个worker执行的结果
		result := <- out
		for _, item := range result.Items {
			// 把item转换为Profile，用于判断这个item是否是一个用户的详细信息
			// 如果是用户信息，打印log并计数
			if profile, ok := item.(model.Profile); ok {
				// 把获取到的用户送往ItemSaver
				go func(p model.Profile) { // 这里要用传参的方式把p传入，否则会出现变量作用域问题
					e.ItemChan <- p
				}(profile)
			}
		}

		// 把所有获得的request再次送给scheduler
		for _, request := range result.Requests {
			// 判断url是否重复出现
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

// 创建并发worker的goroutine
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// 告知scheduler本worker已经ready
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request) // 调用simple.go中的worker方法来访问网页并解析结果
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// URL重复性判断
var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool {
	if visitedUrls[url] { // 已经访问过这个url
		return true
	}
	// 没有访问过这个url，把它存进map
	visitedUrls[url] = true
	return false
}