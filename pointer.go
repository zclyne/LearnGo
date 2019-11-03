package main

import "fmt"

// go语言中有指针和引用的概念，与其他语言不同的是，go语言中的所有传递都是值传递，即会把参数拷贝一份
// 但要注意的是，如果参数是一个object，且包含一个指向某个数据的指针，则拷贝之后的函数内部的object会拥有指向同一块数据的指针

// a、b都是值传递，这个swap无效
func swap_wrong(a, b int) {
	b, a = a, b
}
// 正确的方式：使用指针传入a和b
func swap(a, b *int) {
	*b, *a = *a, *b
}
// 另一种实现方式，这种方法比使用指针更好
func swap_another(a, b int) (int , int) {
	return b, a
}

func main() {
	a, b := 1, 2

	// 错误的交换方式
	swap_wrong(a, b)
	fmt.Println(a, b)

	// 注意传入指针时变量名之前要加上&
	swap(&a, &b)
	fmt.Println(a, b)

	// 另一种交换方式，利用go语言的允许多个返回值的特性
	a, b = swap_another(a, b)
	fmt.Println(a, b)
}
