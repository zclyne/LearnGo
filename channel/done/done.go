package main

import (
	"fmt"
	"sync"
)

// 使用channel等待任务结束

// 用第二个channel done来表示是否已经通信完毕
func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		// done的发送要再开一个goroutine并行地发送，否则会导致卡死
		// 因为done的接收是在小写字母和大写字母都发送完毕之后，而发送的小写字母被接收后，done中的内容还没有被取走
		// 把done放在并行的goroutine中可以解决这个问题
		// 如果这里只需要打印一次小写字母，就不需要单独把done放在goroutine中
		//go func() {
		//	done <- true
		//}()
		// wg.Done()
		w.done()
	}
}

type worker struct {
	in chan int
	// done chan bool
	// wg *sync.WaitGroup // 这里必须使用指针，因为waitgroup必须保持只有1个
	done func() // 用函数来包装一下，而不是直接传入waitgroup，抽象程度更高
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		// done: make (chan bool),
		// wg: wg,
		done: func() { // 在定义worker的时候实现done
			wg.Done()
		},
	}
	// go doWork(id, w.in, w.done)
	//go doWork(id, w.in, wg)
	go doWork(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker

	// go中对于多个并发任务的等待提供了WaitGroup来支持
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		//workers[i] = createWorker(i)
		workers[i] = createWorker(i, &wg)
	}

	// 共有20个任务
	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i

		// 从done channel中接收到数据之后，这一轮for才会结束，这样就不需要再使用sleep来防止退出了
		// 但是这种方法导致所有worker顺序执行，goroutine就没有意义了
		// 所以要在最后接收done中的数据
		// <-workers[i].done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// 等待waitgroup完成
	wg.Wait()

	//// 接收所有worker的done信息
	//for _, worker := range workers {
	//	// 因为每个worker中发送了两次消息，所以done中也会有两个数据，所以要接收两次
	//	<- worker.done
	//	<- worker.done
	//}
}

func main() {
	chanDemo()
}
