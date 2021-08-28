package queue

// Queue An FIFO queue.
type Queue []int

// Push the element into the queue.
// 		e.g. q.Push(1)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop the element from the head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty Return if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
