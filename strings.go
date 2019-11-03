package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes中文测试！"
	// 直接用len()获得字节长度
	fmt.Println(len(s)) // 18
	fmt.Println(s)
	// 使用%X可以打印出字节的数字
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // 每个英文字母1字节，为ASCII编码，每个中文为3字节，UTF-8编码
	}
	fmt.Println()
	// 这里的是int32，也就是一个rune
	for i, ch := range s {
		// 这里打出的ch和上面那个循环不同，因为这里采用了Unicode编码
		// 在中文字符部分，i每次以3递增，这是因为每个下标对应1个byte，而中文字符是3byte
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// utf8是一个库，RuneCountInString获得字符数量
	fmt.Println("Rune count: ", utf8.RuneCountInString(s)) // 8
	// 逐个字符打印
	bytes := []byte(s)
	for len(bytes) > 0 {
		// 下面的size是字符的size，英文字符size为1，中文字符size为3
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 把string转rune数组之后再循环输出，就可以直接用下标来访问到对应的中文字符
	// 这里的rune数组是新开的，和原本的string没有关系，内部的元素已经发生了改变，因为把原本的UTF-8编码转换成了Unicode编码
	// rune相当于go的char
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

	// 只打印ch时，中文字符也能正常处理
	for _, ch := range s {
		fmt.Printf("%c ", ch)
	}

	// 其他字符串换操作，在strings包中
}
