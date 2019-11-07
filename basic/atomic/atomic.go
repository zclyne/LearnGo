package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	// 对代码区加锁，使用匿名函数
	func() {
		// 先加锁，再执行++
		a.lock.Lock()
		// ++完毕后，释放锁
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
