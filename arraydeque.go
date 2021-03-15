package container

// ArrayDeque is an implementation of deque using slices (instead of linked list)
type ArrayDeque[T any] struct {
	left []T
	right []T
}

// NewArrayDeque return a new ArrayDeque object
func NewArrayDeque[T any]() *ArrayDeque[T] { 
	return &ArrayDeque[T]{
		left: make([]T, 0),
		right: make([]T, 0),
	} 
}

// PushBack adds an element at the tail of the ArrayDeque
func (d *ArrayDeque[T]) PushBack(x T) {
	d.right = append(d.right, x)
}

// PushFront adds an element at the head of the ArrayDeque
func (d *ArrayDeque[T]) PushFront(x T) {
	d.left = append(d.left, x)
}

// PopBack removes the tail element of the ArrayDeque and returns the element
// the ArrayDeque must not be empty.
func (d *ArrayDeque[T]) PopBack() (x T) {
	if d.Len() == 0 {
		panic("deque is empty")
	}
	if len(d.right) != 0 {
		n := len(d.right)
		x = d.right[n-1]
		d.right = d.right[:n-1]
	} else {
		x = d.left[0]
		d.left = d.left[1:]
	}
	return
}

// PopFront removes the head element from the ArrayDeque and returns the element
// the ArrayDeque must not be empty.
func (d *ArrayDeque[T]) PopFront() (x T) {
	if d.Len() == 0 {
		panic("deque is empty")
	}
	if len(d.left) != 0 {
		n := len(d.left)
		x = d.left[n-1]
		d.left = d.left[:n-1]
	} else {
		x = d.right[0]
		d.right = d.right[1:]
	}
	return
}

// Back returns the tail element of the ArrayDeque
func (d *ArrayDeque[T]) Back() (x T) {
	if d.Len() == 0 {
		panic("deque is empty")
	}
	if len(d.right) != 0 {
		n := len(d.right)
		x = d.right[n-1]
	} else {
		x = d.left[0]
	}
	return
}

// Front returns the head element in the ArrayDeque
func (d *ArrayDeque[T]) Front() (x T) {
	if d.Len() == 0 {
		panic("deque is empty")
	}
	if len(d.left) != 0 {
		n := len(d.left)
		x = d.left[n-1]
	} else {
		x = d.right[0]
	}
	return
}

// Clear all elements in the ArrayDeque
func (d *ArrayDeque[T]) Clear() {
	d.left = d.left[:0]
	d.right = d.right[:0]
}

// Len returns the size of the ArrayDeque
func (d *ArrayDeque[T]) Len() int {
	return len(d.left) + len(d.right)
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkArrayDequeSize(d *ArrayDeque[int], expect int) {
	if d.Len() != expect {
		panic(fmt.Sprintf("size check failed, got %d, expect %d", d.Len(), expect))
	}
}

func checkArrayDequeNum(got, expect int) {
	if got != expect {
		panic(fmt.Sprintf("answer does not match, got %d, expect %d", got, expect))
	}
}

func testArrayDeque() {
	d := NewArrayDeque[int]()

	// single element
	checkArrayDequeSize(d, 0)
	d.PushBack(1)
	checkArrayDequeSize(d, 1)
	checkArrayDequeNum(d.Back(), 1)
	d.PopFront()
	checkArrayDequeSize(d, 0)
	d.PushFront(1)
	checkArrayDequeSize(d, 1)
	checkArrayDequeNum(d.Back(), 1)
	d.PopBack()
	checkArrayDequeSize(d, 0)

	// multiple elements
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	checkArrayDequeSize(d, 3)
	checkArrayDequeNum(d.Back(), 3)
	checkArrayDequeNum(d.Front(), 1)
	d.PopBack()
	checkArrayDequeSize(d, 2)
	checkArrayDequeNum(d.Back(), 2)
	checkArrayDequeNum(d.Front(), 1)
	d.PushFront(4)
	d.PushFront(5)
	d.PushFront(6)
	checkArrayDequeSize(d, 5)
	checkArrayDequeNum(d.Front(), 6)
	checkArrayDequeNum(d.Back(), 2)
	d.PopFront()
	checkArrayDequeSize(d, 4)
	checkArrayDequeNum(d.Front(), 5)
	checkArrayDequeNum(d.Back(), 2)

	d.Clear()
	checkArrayDequeSize(d, 0)
}