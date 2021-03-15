package container

// ArrayQueue is an implementation of queue using slice
type ArrayQueue[T any] struct {
	arr []T
}

// NewArrayQueue returns a new ArrayQueue object
func NewArrayQueue[T any]() *ArrayQueue[T] { 
	return &ArrayQueue[T]{arr: make([]T, 0)} 
}

// Push adds a new element into the ArrayQueue
func (q *ArrayQueue[T]) Push(x T) { 
	q.arr = append(q.arr, x) 
}

// Pop removes the element at the top of the ArrayQueue and returns the element
// the ArrayQueue must not be empty.
func (q *ArrayQueue[T]) Pop() T {
	if q.Len() == 0 {
		panic("queue is empty")
	}
	x := q.arr[0]
	q.arr = q.arr[1:]
	return x
}

// Top returns the element at the top of the ArrayQueue
// the ArrayQueue must not be empty.
func (q *ArrayQueue[T]) Top() T { 
	if q.Len() == 0 {
		panic("queue is empty")
	}
	return q.arr[0] 
}

// Clear all elemnts in the ArrayQueue
func (q *ArrayQueue[T]) Clear() { 
	q.arr = q.arr[:0] 
}

// Len returns the size of the ArrayQueue
func (q *ArrayQueue[T]) Len() int { 
	return len(q.arr) 
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkArrayQueueSize(q *ArrayQueue[int], expect int) {
	if q.Len() != expect {
		panic(fmt.Sprintf("size check failed, got %d, expect %d", q.Len(), expect))
	}
}

func checkArrayQueueNum(got, expect int) {
	if got != expect {
		panic(fmt.Sprintf("answer does not match, got %d, expect %d", got, expect))
	}
}

func testArrayQueue() {
	q := NewArrayQueue[int]()
	checkArrayQueueSize(q, 0)
	q.Push(1)
	checkArrayQueueSize(q, 1)
	checkArrayQueueNum(q.Top(), 1)
	q.Pop()
	checkArrayQueueSize(q, 0)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	checkArrayQueueSize(q, 3)
	checkArrayQueueNum(q.Top(), 1)
	q.Pop()
	checkArrayQueueSize(q, 2)
	checkArrayQueueNum(q.Top(), 2)
	q.Pop()
	checkArrayQueueSize(q, 1)
	checkArrayQueueNum(q.Top(), 3)

	q.Clear()
	checkArrayQueueSize(q, 0)
}
