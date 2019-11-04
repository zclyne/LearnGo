package main

import (
	"fmt"
)

func tryRecover() {
	// 在匿名函数中调用recover(), 从panic中恢复
	defer func() {
		r := recover()
		// recover的返回值是interface，即任意类型，所以要强制转换为error
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do, %v", r))
		}
	}()
	// panic(errors.New("this is an error"))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(123)
}

func main() {
	tryRecover()
}
