package fib

// 使用闭包实现斐波那契数列的生成
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}