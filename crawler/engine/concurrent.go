package engine

import (
	"log"
)

// 并发引擎
// 并发engine拥有一个scheduler负责调度engine发给它的requests
// scheduler把request发送给worker，由worker来执行获取网页并解析的功能

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 创建与worker通信的channel，输入为Request，输出为ParseResult
	in := make(chan Request)
	out := make(chan ParseResult)

	// 将输入channel配置到scheduler，scheduler负责调度requests给worker
	e.Scheduler.ConfigureMasterWorkerChan(in)

	// 创建worker，并将先前创建的channel与worker连接上
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	// 把种子页面发送给scheduler，开始调度
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 无穷循环
	for {
		// 从输出管道中获取各个worker执行的结果
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}

		// 把所有获得的request再次送给scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// 创建并发worker的goroutine
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			// 下面这种写法会造成循环等待，从而死锁
			// scheduler必须找到空闲的worker才能执行request
			// 但是有空闲worker的前提是这个worker已经把发给自己的上一个request做完
			// 注意Run函数中的无穷循环部分
			// 从out中接收数据的前提是上一轮循环中，result中的requests都已经提交给scheduler
			// 而成功提交的前提是必须有空闲的worker来接收
			request := <- in
			result, err := worker(request) // 调用simple.go中的worker方法来访问网页并解析结果
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}