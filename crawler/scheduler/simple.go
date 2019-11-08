package scheduler

import "LearnGo/crawler/engine"

// SimpleScheduler，实现了scheduler接口
// 所有worker都连接在workerChan上，而SimpleScheduler简单的把接收到的request放到这个channel上
// 最先取到这个request的worker就会开始执行它

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 由于会改变struct内部的内容，所以这里要使用指针接收者
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	s.workerChan <- r
}