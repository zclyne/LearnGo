package scheduler

import "learngo.com/crawler/engine"

// 具有request队列和worker队列的scheduler

type QueuedScheduler struct {
	// request队列，接收新输入的request
	requestChan chan engine.Request

	// 每一个worker对外表现为一个channel engine.Request
	// 需要创建一个channel的channel，不断地接收空闲worker
	workerChan chan chan engine.Request
}

// 每一个worker都有一个自己的chan，所以这个方法直接返回一个新的chan
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// 用于告知外界有一个worker已经ready，可以接收request
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

// 因为在Run()中创建了s的两个channel，也就是改变了s的内容
// 所以要用指针接收者
func (s *QueuedScheduler) Run() {
	// 创建scheduler的channel
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		// 创建request和worker队列，用于存储下接收到的request和空闲worker
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			// 在两个队列中同时都有内容时，表示既有需要处理的request，又有空闲的worker
			// 则可以把request分配给空闲worker去处理
			// 但是不要直接在if里做发送给队列的操作，因为可能会卡死，既然使用了select，则所有channel操作都要在select中执行
			// 这里采用之前的定义active的方式
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
				// 在这里先不把队首元素删除
			}

			// 这里要用select来进行选择，因为requestChan和workerChan中的数据到达先后顺序并不确定，不能串行执行
			// 受到request或worker时，就把它加进队列中
			select {
			case r := <- s.requestChan:
				requestQ = append(requestQ, r)
			case w := <- s.workerChan:
				workerQ = append(workerQ, w)
			// 如果activeRequest或activeWorker中有一个为nil，则不会走进下面这个case
			case activeWorker <- activeRequest: // 确实走入这个case之后，再从队列首部删除active的request和worker
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}