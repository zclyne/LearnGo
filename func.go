package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 同一类型的参数可以用逗号分隔，把类型写在最后一个参数之后
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// return a / b
		// 使用div()函数来执行除法，第二个返回值，即余数，在这里不需要使用
		// 因此使用下划线_来代替，可以防止定义变量来接收后但又没有使用该变量，而导致编译器报错的情况
		q, _ := div(a, b)
		return q, nil
	default:
		// panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 返回值可以有多个，常用的情况是返回一个正常情况的结果和一个error
// 除法，返回结果和余数，例如13 / 3 = 4 ... 1
//func div(a, b int) (int, int) {
//	return a / b, a % b
//}
// 可以给返回值取名字，如下，q表示商，r表示余数，两者都是int
func div(a, b int) (q, r int) {
	// 注意这里的q和r不用:=而是用=
	q = a / b
	r = a % b
	// 因为已经指定了返回值名字为q和r，所以这里直接return，就会自动返回q和r
	return
	// 这种方法在函数体比较长时可读性差，所以一般不推荐使用
}

// go是函数式编程语言，函数的参数、返回值或函数体内部都可以是函数
// 接收的第一个参数op本身是一个函数，它有2个int参数，返回值为int，另外还有a、b两个int类型的参数
func apply(op func(int, int) int, a, b int) int {
	// 使用反射获得op函数的函数名，首先使用.Pointer()获得指向这个函数的指针
	// 然后获得函数本身的名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args " +
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// go语言有可变参数列表，但是没有重载、默认参数、可选参数等其他语言的特性
// ...int表示可以传入任意多个int
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil { // 如果存在错误，打印错误，否则打印除法结果
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(eval(3, 4, "/"))
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(div(20, 5))

	// 传入的第一个参数pow是自定义的函数
	fmt.Println(apply(pow, 3, 4))
	// 使用匿名函数直接传到参数上，效果同上
	// 但是此处函数名将变成main.main.func1，第一个main表示包名，第二个main表示函数名main
	// 这类似于其他语言中的lambda表达式
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
