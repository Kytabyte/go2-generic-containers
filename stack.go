package container

import "fmt"

// Stack is an implementation of Stack
type Stack[T any] struct {
	list *List[T] // hide all methods of List
}

// NewStack returns a new Stack object
func NewStack[T any]() *Stack[T] {
	l := NewList[T]()
	return &Stack[T]{l}
}

// Push a new element to the Stack
func (s *Stack[T]) Push(x T) {
	s.list.PushBack(x)
}

// Pop removes the element at the top of the Stack and return the element.
// the Stack must not be empty.
func (s *Stack[T]) Pop() T {
	x := s.list.Back()
	s.list.Remove(x)
	return x.Value
}

// Top returns the element at the top of the Stack.
// the Stack must not be empty.
func (s *Stack[T]) Top() T {
	return s.list.Back().Value
}

// Clear all elements from the Stack.
func (s *Stack[T]) Clear() {
	s.list.Clear()
}

// Len returns the size of the Stack.
func (s *Stack[T]) Len() int {
	return s.list.Len()
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkStackSize(s *Stack[int], expect int) {
	if s.Len() != expect {
		panic(fmt.Sprintf("size check failed, got %d, expect %d", s.Len(), expect))
	}
}

func checkStackNum(got, expect int) {
	if got != expect {
		panic(fmt.Sprintf("answer does not match, got %d, expect %d", got, expect))
	}
}

func testStack() {
	s := NewStack[int]()
	checkStackSize(s, 0)
	s.Push(1)
	checkStackSize(s, 1)
	checkStackNum(s.Top(), 1)
	s.Pop()
	checkStackSize(s, 0)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	checkStackSize(s, 3)
	checkStackNum(s.Top(), 3)
	s.Pop()
	checkStackSize(s, 2)
	checkStackNum(s.Top(), 2)
	s.Pop()
	checkStackSize(s, 1)
	checkStackNum(s.Top(), 1)

	s.Clear()
	checkStackSize(s, 0)
}