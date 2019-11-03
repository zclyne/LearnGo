package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	// 与if相同，go中的循环也不能有括号
	for ; n > 0; n /= 2 {
		// lsb表示最低位
		lsb := n % 2
		// 用strconv.Itoa()把数字转换为字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
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

func forever() {
	// 死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1101
		convertToBin(1234123561),
		// 注意如果括号换行，则最后一个参数之后要加上逗号
		convertToBin(0),
	)
	printFile("abc.txt")
	// ``表示跨行字符串，并且可以包含引号
	s := `abc"d"
	kkkkk
	123

	p`
	// 由于printFileContents()的参数是一个io.Reader，所以也可以像下面这样打印字符串
	// 扩展性得到增强
	printFileContents(strings.NewReader(s))

	// forever()
}
