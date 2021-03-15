package container

import "fmt"

// ArrayStack is an implementation of Stack using slice
type ArrayStack[T any] struct {
	arr []T
}

// NewArrayStack returns a new ArrayStack object
func NewArrayStack[T any]() *ArrayStack[T] { 
	return &ArrayStack[T]{arr: make([]T, 0)} 
}

// Push a new element to the ArrayStack
func (s *ArrayStack[T]) Push(x T) { 
	s.arr = append(s.arr, x) 
}

// Pop removes the element at the top of the ArrayStack and return the element.
// the ArrayStack must not be empty.
func (s *ArrayStack[T]) Pop() T {
	if s.Len() == 0 {
		panic("stack is empty")
	}
	n := len(s.arr)
	x := s.arr[n-1]
	s.arr = s.arr[:n-1]
	return x
}

// Top returns the element at the top of the ArrayStack.
// the ArrayStack must not be empty.
func (s *ArrayStack[T]) Top() T { 
	if s.Len() == 0 {
		panic("stack is empty")
	}
	return s.arr[s.Len()-1] 
}

// Clear all elements from the ArrayStack.
func (s *ArrayStack[T]) Clear() { 
	s.arr = s.arr[:0] 
}

// Len returns the size of the ArrayStack.
func (s *ArrayStack[T]) Len() int { 
	return len(s.arr) 
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkArrayStackSize(s *ArrayStack[int], expect int) {
	if s.Len() != expect {
		panic(fmt.Sprintf("size check failed, got %d, expect %d", s.Len(), expect))
	}
}

func checkArrayStackNum(got, expect int) {
	if got != expect {
		panic(fmt.Sprintf("answer does not match, got %d, expect %d", got, expect))
	}
}

func testArrayStack() {
	s := NewArrayStack[int]()
	checkArrayStackSize(s, 0)
	s.Push(1)
	checkArrayStackSize(s, 1)
	checkArrayStackNum(s.Top(), 1)
	s.Pop()
	checkArrayStackSize(s, 0)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	checkArrayStackSize(s, 3)
	checkArrayStackNum(s.Top(), 3)
	s.Pop()
	checkArrayStackSize(s, 2)
	checkArrayStackNum(s.Top(), 2)
	s.Pop()
	checkArrayStackSize(s, 1)
	checkArrayStackNum(s.Top(), 1)

	s.Clear()
	checkArrayStackSize(s, 0)
}