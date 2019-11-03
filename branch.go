package main

import (
	"fmt"
	"io/ioutil"
)

// 返回值为string
func grade(score int) string {
	// 首先初始化变量g为一个字符串
	g := ""
	// go中的switch每个case最后不需要加break，默认情况是有break的，除非使用fallthrough
	// 在go中可以不指定switch的对象，而在case之后写条件语句
	switch {
	case score < 0 || score > 100:
		// panic函数为报错，它会中断程序的执行
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	// 从文件中读取内容。go允许函数返回2个值，例如此处的ReadFile函数
	// 第一个返回值是文件中的内容，第二个返回值是可能产生的错误
	//contents, err := ioutil.ReadFile(filename)
	//// go语言的if之后的条件不需要加括号
	//if err != nil { // 存在错误
	//	fmt.Println(err)
	//} else {
	//	// 用%s来打印数组中的内容
	//	fmt.Printf("%s\n", contents)
	//}
	// go语言中if的另一种写法
	// 先运行分号之前的内容完成变量赋值，然后以分号之后的语句作为判断条件
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// 采用这种方法时，在if...else...语句块之外，将无法再访问变量contents和err
	// 所以if的条件中的赋值语句的变量作用于仅仅是当前if...else...语句块

	fmt.Println(
		grade(0),
		grade(59),
		grade(62),
		grade(78),
		grade(100),
		grade(101),
	)
}
