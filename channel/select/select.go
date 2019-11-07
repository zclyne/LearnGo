package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // 随机sleep一段时间，1500ms以内
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	// w := createWorker(0)
	var worker = createWorker(0)

	// 存储接收到的数据，防止因发送和接收速度不一致导致的数据丢失
	var values []int

	// 设置程序运行10秒后退出
	// time.After的返回值是一个channel，在运行10秒后，会向这个channel中发送一个数据
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second) // 每秒钟送一个数据到tick channel

	for {
		var activeWorker chan<- int // activeWorker is nil
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		// 使用select来从多个channel接收数据，哪个channel中数据到达得比较快，就从哪个中取出数据
		select {
		case n := <-c1:
			// fmt.Println("Received from c1:", n)
			// hasValue = true
			values = append(values, n)
		case n := <-c2:
			// fmt.Println("Received from c2:", n)
			// hasValue = true
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:] // 第一个元素已经被送出，将其删除
		case <- time.After(800 * time.Millisecond): // 如果800毫秒后还没有接收到数据，则打印一个timeout
			fmt.Println("timeout")
		case <- tick: // 每个tick打印一下队列的长度
			fmt.Println("queue len =", len(values))
		case <- tm: // tm接收到了数据，说明程序已经运行了10秒，关闭程序
			fmt.Println("Bye")
			return
		}
	}
}
