package main

import (
	"learngo.com/queue"
	"fmt"
)

func main() {
	// 创建queue，原始值为1
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop()) // 1
	fmt.Println(q.Pop()) // 2
	fmt.Println(q.IsEmpty()) // false
	q.Pop()
	fmt.Println(q.IsEmpty()) // true
	// 进行了一系列Push()和Pop()之后，q和原本的q不是同一个slice

	// Queue可以支持任何类型
	q.Push("abc")
	fmt.Println(q.Pop())
}
