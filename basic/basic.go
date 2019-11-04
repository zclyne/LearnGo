package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 在函数外部也可以定义变量
// go中没有全局变量的概念，这些定义在函数之外的变量都是包内部的变量
var aa = 3
var ss = "kkk"
// 在函数外部不可以使用:=来定义变量
// bb := true 非法

// 用括号来包裹多个变量定义
var (
	bb = true
	cc = "666"
	dd = 123
)

func variableZeroValue() {
	// 定义变量，变量名在前，类型在后
	var a int // int的零值为0
	var s string // 字符串的零值为空字符串
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	// 在创建变量的同时赋初值
	// go语言中规定若定义了一个变量，则该变量必须被使用
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	// go的编译器可以从初值推断出变量的类型
	// 在同一行内可以定义不同类型的变量，如下面的a、b为int，c为bool、d为string
	var a, b, c, d = 3, 4, true, "def"
	var s = "abc"
	fmt.Println(a, s, b, c, d)
}

func variableShorter()  {
	// 一种更简洁的变量定义方式
	a, b, c, d := 3, 4, true, "def"
	// 变量定义完成之后，就不能再次使用:=
	// b := 5是非法的，重复定义变量
	b = 5
	fmt.Println(a, b, c, d)
}

func euler() {
	// 定义一个复数，注意不能写成3 + 4 * i，因为这样写编译器会认为i是一个变量
	// c是一个complex128
	c := 3 + 4i
	// 复数运算的库为cmplx
	fmt.Println(cmplx.Abs(c))
	// 验证欧拉公式，注意这里要写1i而非i，原因同上
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	// cmplx.Exp的底数就是e，所以上下两行等价
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	// 上面两行的输出都不是绝对的0，因为float的误差
	// 下面这一行的数据是0.000 + 0.000i，因为限定了精度
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func typeConversion() {
	// go语言中只有强制类型转换，没有隐式类型转换
	var a, b int = 3, 4
	var c int
	// 由于math.Sqrt()要求参数为float64，这里必须进行强制转换
	// math.Sqrt()的返回值同样是float64，所以要再转换回int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts() {
	// 定义常量时，可以规定类型，也可以不规定
	// 若不规定类型，则类型不确定
	// 常量也可以定义在函数之外、包内部
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	// 由于a、b是常量，且在定义时类型不确定，因此此处并不需要强制转float64
	// 但如果定义常量a、b时指定了类型为int，则还是要进行转换
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
	// 在go语言中，常量名一般不取全部大写，因为go语言的大小写是具有含义的
	// 在go中，首字母大写代表public
}

func enums() {
	// go中没有专门的枚举类型，因此一般采用一组常量来定义
	// iota是go中专门为实现枚举而定义的，它表示这组常量的值是自增的
	const(
		cpp = iota
		java
		python
		golang
		javascript
	)
	// iota还可以参与运算，例如：
	// b, kb, mb, gb, tb, pb都会采用公式1 << (10 * iota)来计算
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang, javascript) // 0, 1, 2, 3
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

//func main() {
//	fmt.Println("Hello world!")
//	variableZeroValue()
//	variableInitialValue()
//	variableTypeDeduction()
//	variableShorter()
//	fmt.Println(aa, ss, bb, cc, dd)
//
//	euler()
//	typeConversion()
//	consts()
//	enums()
//}
