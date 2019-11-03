package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	// 直接创建slice
	var s []int // zero value for slice is nil
	// now s == nil，表示还没有分配数组
	// 开始向s中添加元素
	for i:= 0; i < 100; i++ {
		printSlice(s) // 即使s是nil时也不会崩溃，len和cap都是0
		// 在不断添加元素的过程中，cap以2的指数形式增长
		// 即1, 2, 4, 8, 16, 32, 64, 128
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	// 先创建一个array [2, 4, 6, 8]，然后s1去view这个array
	s1 := []int{2, 4, 6, 8}
	printSlice(s1) // len = 4, cap = 4

	// 创建一个有初始大小、但是没有初始内容的slice
	// 使用make函数，传入一个参数时表示len，传入2个参数时分别为len和cap
	s2 := make([]int, 16) // len = 16, cap = 16
	printSlice(s2)
	s3 := make([]int, 10, 32) // len = 10, cap = 32
	printSlice(s3)
	// s2、s3的len个元素全部为0，因为0是int的zero value

	fmt.Println("Copying slice")
	// 第一个参数为destination，第二个参数为source
	copy(s2, s1)
	printSlice(s2) // 前4个元素为2 4 6 8，之后全0，len = 16, cap = 16

	fmt.Println("Deleting elements from slice")
	// 删除s2中的下标为3的元素8
	// append函数第二个参数开始为任意多个元素，但是不能直接传入切片s2[4:]
	// 所以要用s2[4:]...，表示把s2[4:]中的每个元素单独传入
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2) // 前3个元素为2 4 6，之后全0，len = 15，cap = 16

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2) // len = 14, cap = 15

	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)
	printSlice(s2) // len = 13, cap = 15
}
