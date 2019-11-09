package scheduler

import "LearnGo/crawler/engine"

// SimpleScheduler，实现了scheduler接口
// 所有worker都连接在workerChan上，而SimpleScheduler简单的把接收到的request放到这个channel上
// 最先取到这个request的worker就会开始执行它

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 在SimpleScheduler中，所有worker共享一个channel
// 所以这里直接返回s.workerChan，就是worker用来接收request的channel
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

// 由于会改变struct内部的内容，所以这里要使用指针接收者
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// submit必须另开一个goroutine来把r放入workerChan中，否则会造成死锁
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}