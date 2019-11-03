package main

import "fmt"

// go语言中的数组是值类型，所以在传参时会被拷贝一份后再传入
// 由于它是值类型，所以在传参时需要指定数组的长度
// 常见的go代码中允许arr []int这样的语法，是因为这表示切片，和数组不同
// 这个特性与其他大部分语言都不同
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Printf("index is %d, value is %d\n", i, v)
	}
	arr[0] = 100 // 这一行没有效果，因为arr是值类型
}

// 传入数组的指针
func printArrayPtr(arr *[5]int) {
	// 注意这里虽然arr是指针，但是也不需要写成*arr
	for i, v := range arr {
		fmt.Printf("index is %d, value is %d\n", i, v)
	}
	// 同样不需要写成(*arr)[0]
	arr[0] = 100 // 这一行会生效，因为现在传入的是指针而不是值
}

// 由于数组的这些特性，在go中一般不用数组，而是使用切片

func main() {
	// 定义数组时也与其他语言相反，数组长度放在前面，类型放在后面
	var arr1 [5]int // [0 0 0 0 0]
	// 使用:=来创建数组时，要指定初值
	arr2 := [3]int{1, 3, 5}
	// 让编译器自己判断出数组长度，要在方括号中加入...
	// 注意...不能省略，因为[]表示切片，是另一个不同的概念
	arr3 := [...]int{2, 4, 6, 8, 10}

	// 二维数组
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 数组的遍历
	// 方法1，直接循环
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// 方法2，用range关键字，获得下标
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	// 方法3，用range关键字，同时获得下标和值
	for i, v := range arr3 {
		fmt.Printf("index is %d, value is %d\n", i, v)
	}
	// 方法4，用range关键字，但是只需要值而不需要下标，则用下划线_代替i来省略变量
	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(arr3)
	printArray(arr3)

	printArrayPtr(&arr3)
	printArrayPtr(&arr3)
}
