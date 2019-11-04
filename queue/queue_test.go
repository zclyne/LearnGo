package queue

import "fmt"

// 示例代码，注意要在最后加上Output:
// 表明应该得到的输出，这样这个方法就可以执行了

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	//1
	//2
	//false
	//3
	//true
}
