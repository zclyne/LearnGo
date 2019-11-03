package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	// 切片（slice）定义为数组（array）的一个视图（view）
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println("arr[2:6] = ", s)
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] = ", s1)
	s2 := arr[:]
	fmt.Println("arr[:] = ", s2)

	updateSlice(s)
	fmt.Println("After updateSlice(s)")
	fmt.Println(s)
	// s经过update之后，原本的arr中的元素被改变，因此s1 s2中相应元素也同样被改变
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println("After updateSlice(s2)")
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)

	// 再次做slice
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	// slice可以扩展
	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	fmt.Println("s1 = ", s1) // [2 3 4 5]
	s2 = s1[3:5]
	fmt.Println("s2 = ", s2) // [5 6]
	// 这里s2能够取到s1[4]且不报错的原因是原本数组arr的元素5之后的元素是6
	// 但是直接访问s1[4]会发生报错
	// 实际上，slice包括长度（length）和容量（capacity）两个概念
	// len表示整个slice直接的长度，而cap表示从slice的第一个元素开始一直到原本的数组的最后一个元素的总长度
	// 只要不超过capacity，就可以用slice来向后扩展，但是slice不可以向前扩展
	// 直接用下标取值不可以超过len，向后扩展不可以超过cap
	fmt.Println(len(s1), cap(s1)) // 4, 6
	// fmt.Println(s1[3:7]) // 报错，因为向后扩展超过了cap(s1)

	// 向slice添加元素
	fmt.Println("向slice最后添加元素")
	s1 = arr[2:6]
	s2 = s1[3:5]
	// 由于值传递的关系，必须接收append的返回值
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println(s3) // [5 6 10]
	fmt.Println(s4) // [5 6 10 11]
	fmt.Println(s5) // [5 6 10 11 12]
	fmt.Println(arr) // [0 1 2 3 4 5 6 10]
	// 数组arr本身的长度不变，在s2 append 10之后，原本最后一个元素7被10覆盖
	// 但是之后的11、12都不会对arr造成影响，因为此时s3和s4已经不是arr的view了
	// 在添加元素时，如果超越了cap，则go会新开一个更长的数组，把原本的arr拷贝过去，s3和s4是新数组的view
	// 若没有再被使用到，则原本的数组会被垃圾回收
	fmt.Println(len(s4), cap(s4)) // 4 6
	fmt.Println(len(s5), cap(s5)) // 5 6
}
