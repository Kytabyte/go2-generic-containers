package container

import "container/heap"

// priorityQueue implements the heap.Interface interface
type priorityQueue[T any] struct {
	h []T // heap structure
	cmp func(T, T) int
}

func (pq priorityQueue[T]) Len() int {
	return len(pq.h)
}

func (pq priorityQueue[T]) Less(i, j int) bool {
	return pq.cmp(pq.h[i], pq.h[j]) < 0
}

func (pq priorityQueue[T]) Swap(i, j int) {
	pq.h[i], pq.h[j] = pq.h[j], pq.h[i]
}

func (pq *priorityQueue[T]) Push(x interface{}) {
	pq.h = append(pq.h, x.(T))
}

func (pq *priorityQueue[T]) Pop() interface{} {
	n := len(pq.h)
	x := pq.h[n-1]
	pq.h = pq.h[:n-1]
	return x
}

// PriorityQueue is an implementation of priority queue data structure.
// It is just a wrapper of an internal pq data structure that implements
// heap.Interface, such that Push and Pop could have type T instead of interface{}
type PriorityQueue[T any] struct {
	pq *priorityQueue[T] 
}

// NewPQ returns a new PriorityQueue object, given an compare function of type T
func NewPQ[T any](cmp func(T,T) int) *PriorityQueue[T] {
	pq := &priorityQueue[T]{
		h: make([]T, 0),
		cmp: cmp,
	}
	return &PriorityQueue[T]{
		pq: pq,
	}
}

// Push a new element into the PriorityQueue
func (pq *PriorityQueue[T]) Push(x T) {
	heap.Push(pq.pq, x)
}

// Pop removes the smallest element (based on the given cmp) and returns the element
func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(pq.pq).(T)
}

// Top returns the smallest element (based on the given cmp)
func (pq *PriorityQueue[T]) Top() T {
	return pq.pq.h[0]
}

// Len returns the size of the PriorityQueue
func (pq *PriorityQueue[T]) Len() int {
	return pq.pq.Len()
}

// PushPop is equivalent to but effectively pushes a new element into,
// and pops the smallest element from the PriorityQueue
func (pq *PriorityQueue[T]) PushPop(x T) T {
	if pq.Len() == 0 || pq.pq.cmp(x, pq.Top()) < 0 {
		return x
	}
	r := pq.Top()
	pq.pq.h[0] = x
	heap.Fix(pq.pq, 0)
	return r
}

// PopPush is equivalent to but effectively pops the smallest element from,
// and pushes the given element into the PriorityQueue
func (pq *PriorityQueue[T]) PopPush(x T) T {
	if pq.Len() == 0 {
		panic("Cannot Call PopPush() on empty pq.")
	}
	r := pq.Top()
	pq.pq.h[0] = x
	heap.Fix(pq.pq, 0)
	return r
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkPQ(pq *PriorityQueue[int], expLen int, expTop int) {
	if pq.Len() != expLen {
		panic(fmt.Sprintf("Expect pq len to be %d, but got %d.\n", expLen, pq.Len()))
	}
	if expTop != -1 && pq.Top() != expTop {
		panic(fmt.Sprintf("Expect pq top to be %d, but got %d.\n", expTop, pq.Len()))
	}
}

func checkPQElement(elem, expect int) {
	if elem != expect {
		panic(fmt.Sprintf("Expect element to be %d, but got %d.\n", expect, elem))
	}
}

// Since this is only a wrapper of go's heap package, we don't need to inspect how the heap
// is handled, checking the functionality is sufficient.
func testPQ() {
	pq := NewPQ[int](CmpLess[int])
	checkPQ(pq, 0, -1)
	pq.Push(1)
	checkPQ(pq, 1, 1)
	pop := pq.Pop()
	checkPQ(pq, 0, -1)
	checkPQElement(pop, 1)

	// multiple values
	pq.Push(2)
	pq.Push(1)
	pq.Push(3)
	checkPQ(pq, 3, 1)
	pop = pq.Pop()
	checkPQ(pq, 2, 2)
	checkPQElement(pop, 1)
	pop = pq.PushPop(1)
	checkPQ(pq, 2, 2)
	checkPQElement(pop, 1)
	pop = pq.PushPop(4)
	checkPQ(pq, 2, 3)
	checkPQElement(pop, 2)
	pop = pq.PopPush(2)
	checkPQ(pq, 2, 2)
	checkPQElement(pop, 3)
}