package main

import (
	"learngo.com/functional/fib"
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	// defer关键字使得语句在函数结束时调用，参数在defer语句时计算
	// defer以栈的形式存储，所以先输出2，再输出1，最终结果为3 2 1
	// defer的作用：在函数中间可能穿插着return或panic
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred") // 虽然有panic，但是仍然能输出3 2 1
	fmt.Println(4)
}

func tryAnotherDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i) // 参数在defer语句时计算，所以这里的i并不是全部为100，而是0, 1, 2, ..., 30
		if i == 30 {
			panic("Printed too many numbers")
		}
	}
}

// 使用defer进行资源管理
func writeFile(filename string) {
	// file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL | os.O_CREATE, 0666) // 在文件已存在时，这句会报错
	// 创建自定义error
	// err = errors.New("this is a custome error") // 这句话会造成之后的强转PathError失败，从而调用panic()
	if err != nil {
		// panic(err)
		// os.OpenFile()实际返回的error是一个*PathError，可以用以下方式转换出来
		if pathError, ok := err.(*os.PathError); !ok { // 强制转换成PathError失败
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err) // open, fib.txt, file exists
		}
		// fmt.Println("Error:", err.Error())
		return
	}
	defer file.Close()

	// 用有buffer的writer包装一下file，因为直接用file来写会比较慢
	writer := bufio.NewWriter(file)
	defer writer.Flush() // 这一行必须要加，因为下面的循环只是把fib数列写到了writer的buffer中，而没有实际写到文件里
	// 必须要调用Flush()来把buffer中的内容写入文件
	// 由于defer是栈的形式，所以会先Flush()，再关闭file，顺序正确

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	// tryAnotherDefer()
	writeFile("fib.txt")
}
