package queue

// An FIFO queue

// 让Queue支持任何类型，把原来的int改成interface{}
type Queue []interface{}

// pushes the element into the queue
// 必须使用指针接收者，才能真正改变q
func (q *Queue) Push (v interface{}) {
	*q = append(*q, v)
}

// pop the element from the head
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// 让Queue虽然底层支持任意类型，但是Push和Pop只能是int
//func (q *Queue) Push (v int) {
//	*q = append(*q, v)
//}
//
//func (q *Queue) Pop() int {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head.(int) // 取出interface内部的int并返回
//}

// 让Push和Pop虽然参数和返回值允许任意类型，但是内部实际只能存储int
//func (q *Queue) Push (v interface{}) {
//	*q = append(*q, v.(int)) // 实际只能append int，如果传入的参数非int，会产生运行时错误
//}
//
//func (q *Queue) Pop() interface{} {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head
//}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}