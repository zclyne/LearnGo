package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// 用go关键字来并发执行函数，这里实际上使用的是协程(Coroutine)
		go func(i int) {
			// fmt.Printf仍然会被抢占，但这是因为Printf是一个IO操作，并不是因为协程被抢占
			// IO操作可以交出控制权
			// fmt.Printf("Hello from " + "goroutine %d\n", i)

			// 下面这种写法会导致程序卡死在某一个协程内部，因为协程是非抢占式的
			// 而main()本身也是一个goroutine，而由于程序卡死在某个线程的死循环中
			// 所以虽然main程序里只sleep了1毫秒，但它始终不会停止，因为main()得不到执行
			// 除非加入runtime.Gosched()来手动交出控制权
			for {
				a[i]++
				runtime.Gosched() // 手动交出控制权，让别人也有机会运行，但这个方法一般不会使用到
			}
		}(i) // 这里的i表示把循环变量作为参数传入
		// 这里必须用传参的方法传入i，否则会出现Index out of range报错，因为发生了数据访问冲突(race condition)
		// i形成了闭包，并发的匿名函数内部使用的i和循环变量是同一个地址上的数据
		// 当for循环结束后，i = 10，而goroutine中取到的i = 10，导致index out of range
		// 用传参的方式可以避免这个问题
	}
	// 必须在这里加上sleep，因为for中创建的goroutine与main是并发执行的，而main执行完毕后，会把所有goroutine都杀死
	// 导致程序来不及输出内容就退出了
	time.Sleep(time.Millisecond)

	// 这里打印a也可能造成竞态条件，并发的goroutine一遍在写入a，main中一边在输出
	// 这个问题要通过channel来解决
	fmt.Println(a)

	// 协程：
	// 轻量级线程
	// 非抢占式多任务处理，由协程主动交出控制权
	// 编译器/解释器/虚拟机层面的多任务，而不是操作系统层面的
	// 多个协程可能再一个或多个线程上运行

	// goroutine可能的切换点
	// I/O, select
	// channel
	// 等待锁
	// 函数调用（有时）
	// runtime.Gosched()
	// 以上只是参考，不能保证切换，也不能保证在其他时候不切换
}
