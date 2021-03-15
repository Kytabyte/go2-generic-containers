package container

// Deque is an implementation of double ended queue
type Deque[T any] struct{
	list *List[T]
}

// NewDeque returns a new Deque object
func NewDeque[T any]() *Deque[T] {
	l := NewList[T]()
	return &Deque[T]{l}
}

// PushBack adds an element at the tail of the Deque
func (d *Deque[T]) PushBack(x T) {
	d.list.PushBack(x)
}

// PushFront adds an element at the head of the Deque 
func (d *Deque[T]) PushFront(x T) {
	d.list.PushFront(x)
}

// PopBack removes the tail element of the Deque and returns the element
// the Deque must not be empty.
func (d *Deque[T]) PopBack() T {
	x := d.list.Back()
	d.list.Remove(x)
	return x.Value
}

// PopFront removes the head element from the Deque and returns the element
// the Deque must not be empty.
func (d *Deque[T]) PopFront() T {
	x := d.list.Front()
	d.list.Remove(x)
	return x.Value
}

// Back returns the tail element of the Deque
func (d *Deque[T]) Back() T {
	return d.list.Back().Value
}

// Front returns the head element in the Deque
func (d *Deque[T]) Front() T {
	return d.list.Front().Value
}

// Clear all elements in the Deque
func (d *Deque[T]) Clear() {
	d.list.Clear()
}

// Len returns the size of the Deque
func (d *Deque[T]) Len() int {
	return d.list.Len()
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkDequeSize(d *Deque[int], expect int) {
	if d.Len() != expect {
		panic(fmt.Sprintf("size check failed, got %d, expect %d", d.Len(), expect))
	}
}

func checkDequeNum(got, expect int) {
	if got != expect {
		panic(fmt.Sprintf("answer does not match, got %d, expect %d", got, expect))
	}
}

func testDeque() {
	d := NewDeque[int]()

	// single element
	checkDequeSize(d, 0)
	d.PushBack(1)
	checkDequeSize(d, 1)
	checkDequeNum(d.Back(), 1)
	d.PopFront()
	checkDequeSize(d, 0)
	d.PushFront(1)
	checkDequeSize(d, 1)
	checkDequeNum(d.Back(), 1)
	d.PopBack()
	checkDequeSize(d, 0)

	// multiple elements
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	checkDequeSize(d, 3)
	checkDequeNum(d.Back(), 3)
	checkDequeNum(d.Front(), 1)
	d.PopBack()
	checkDequeSize(d, 2)
	checkDequeNum(d.Back(), 2)
	checkDequeNum(d.Front(), 1)
	d.PushFront(4)
	d.PushFront(5)
	d.PushFront(6)
	checkDequeSize(d, 5)
	checkDequeNum(d.Front(), 6)
	checkDequeNum(d.Back(), 2)
	d.PopFront()
	checkDequeSize(d, 4)
	checkDequeNum(d.Front(), 5)
	checkDequeNum(d.Back(), 2)

	d.Clear()
	checkDequeSize(d, 0)
}