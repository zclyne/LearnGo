package main

import (
	"fmt"
	"time"
)

// channel可以作为参数
func worker(id int, c chan int) {
	//for {
	//	// 关闭channel，但是在关闭之后，接收方还是会收到数据，导致Print出一堆0
	//	// 这是因为0是int类型的zero value，直到到达1毫秒
	//	// 但是使用了ok之后，就不会有这个问题
	//	// 从c中接收数据，ok表示是否还有值。如果channel被close了，则ok为false
	//	n, ok := <- c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %d\n", id, n)
	//}

	// 另一种处理channel关闭的方法。即使不关闭channel，它也会一直收下去，直到main本身退出
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

// channel作为返回值，这个函数的执行很快，创建channel，打开goroutine开始接收数据，然后立即返回
// chan<-表示返回的channel只能用来发送数据，接收数据
// 若只允许接收数据，则使用<-chan
func createWorker(id int) chan<- int {
	c := make(chan int)
	// 这里必须要开一个goroutine来进行接收
	go worker(id, c)
	return c
}

// channel是goroutine之间双向通信的通道

func chanDemo() {
	// 创建channel数组
	// var channels [10]chan int

	// 这里的channel数组定义也要用chan<-，表明只能用来发送
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		// 定义channel，c是一个channel，且里面的内容是int
		// var c chan int // 此时c == nil
		// channels[i] = make(chan int) // 此时c != nil
		channels[i] = createWorker(i)

		// 从channel接收数据，接收和发送必须在不同的goroutine中，否则会造成死锁
		// go worker(i, channels[i])
	}

	// 向channel中发送数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	// 睡眠1毫秒，防止因程序执行完毕而导致goroutine退出
	time.Sleep(time.Millisecond)
}

// 带缓冲区的channel，可以提升性能
func bufferedChannel() {
	// 第二个参数表示缓冲区大小为3
	c := make(chan int, 3)

	go worker(0, c)

	// 不设置buffer的情况下，go不允许channel只发送而没有接收，但是定义了buffer之后，可以这样操作
	// 例如这里缓冲区大小为3，可以向c中发送3个int而程序不发生死锁报错
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

// 数据已经发完之后，可以把channel给close掉，但这个操作不是必须的
// close操作永远是发送方来执行，告知接收方close掉channel
func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	// 关闭channel
	close(c)

	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	// chanDemo()

	fmt.Println("Buffered channel")
	// bufferedChannel()

	fmt.Println("Channel close and range")
	channelClose()
}
