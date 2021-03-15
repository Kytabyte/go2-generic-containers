package container

// Queue is an implementation of queue
type Queue[T any] struct {
	list *List[T]
}

// NewQueue returns a new Queue object
func NewQueue[T any]() *Queue[T] {
	l := NewList[T]()
	return &Queue[T]{l}
}

// Push adds a new element into the Queue
func (q *Queue[T]) Push(x T) {
	q.list.PushBack(x)
}

// Pop removes the element at the top of the Queue and returns the element
// the Queue must not be empty.
func (q *Queue[T]) Pop() T {
	x := q.list.Front()
	q.list.Remove(x)
	return x.Value
}

// Top returns the element at the top of the Queue
// the Queue must not be empty.
func (q *Queue[T]) Top() T {
	return q.list.Front().Value
}

// Clear all elemnts in the Queue
func (q *Queue[T]) Clear() {
	q.list.Clear()
}

// Len returns the size of the Queue
func (q *Queue[T]) Len() int {
	return q.list.Len()
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkQueueSize(q *Queue[int], expect int) {
	if q.Len() != expect {
		panic(fmt.Sprintf("size check failed, got %d, expect %d", q.Len(), expect))
	}
}

func checkQueueNum(got, expect int) {
	if got != expect {
		panic(fmt.Sprintf("answer does not match, got %d, expect %d", got, expect))
	}
}

func testQueue() {
	q := NewQueue[int]()
	checkQueueSize(q, 0)
	q.Push(1)
	checkQueueSize(q, 1)
	checkQueueNum(q.Top(), 1)
	q.Pop()
	checkQueueSize(q, 0)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	checkQueueSize(q, 3)
	checkQueueNum(q.Top(), 1)
	q.Pop()
	checkQueueSize(q, 2)
	checkQueueNum(q.Top(), 2)
	q.Pop()
	checkQueueSize(q, 1)
	checkQueueNum(q.Top(), 3)

	q.Clear()
	checkQueueSize(q, 0)
}