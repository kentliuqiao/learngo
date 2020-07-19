package queue

// Queue An FIFO queue
type Queue []int

// Push pushes element into the queque
//		e.g. q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop pop element from the head
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty returns true if the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
