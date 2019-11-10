package main

import (
	"learngo.com/functional/fib"
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 为函数实现接口
type intGen func() int
// 函数intGen作为Read方法的接收者，从而实现了reader接口
func (g intGen) Read(p []byte) (n int, err error) {
	// 获得下一个元素
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// 用string来代理，把next读入到p中
	// 采用代理的原因是自己实现起来比较麻烦，要写比较多的底层代码
	s := fmt.Sprintf("%d\n", next)
	// TODO: incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// 这里既没有起始条件，也没有递增条件，只有结束条件，可以直接省略分号
	// 这就相当于while，所以go语言中并没有while
	// scanner.Scan()每次读入一行
	for scanner.Scan() {
		// scanner.Text()是Scan()出的一行
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
