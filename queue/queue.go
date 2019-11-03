package queue

type Queue []int

// 必须使用指针接收者，才能真正改变q
func (q *Queue) Push (v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}