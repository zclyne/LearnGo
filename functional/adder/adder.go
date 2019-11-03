package main

import "fmt"

// 函数adder()的返回值也是一个函数
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	// 用adder来求出0~9所有数字之和
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + .. + %d = %d\n", i, a(i))
	}
	// 以上累加的实现使用到了闭包的概念
	// 函数体内部包含局部变量和自由变量，例如adder()的返回值的函数中，v是局部变量，sum是自由变量
	// 编译器会把自由变量连接到函数外部的某一块结构，这块结构还可以连接到另一块结构，最终形成一棵树
	// 所有连接关系结束后，形成一个闭包。闭包包括函数体以及所有由自由变量连接到的结构
	// 所以对于adder()，它返回的不仅是一段函数代码，还有对sum的引用，sum会被保存到返回的函数里面
	// 从而实现累加的功能

}
