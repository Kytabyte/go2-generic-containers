package container

import (
	"fmt"
	"reflect"
)

// list is a copy of go's container/list package but with type parameter supported.
// There is a little difference in implementation where this List's zero value is
// not ready-to-use; instead, one must call NewList() for initialization

// Node type inside a list
type Element[T any] struct {
	Value T
	prev *Element[T]
	next *Element[T]
	list *List[T]
}

// Prev return the prev element if exists, otherwise return nil
func (e *Element[T]) Prev() *Element[T] {
	if prev := e.prev; e.list != nil && prev != e.list.sentinel {
		return prev
	}
	return nil
}

// Next return the next element if exists, otherwise return nil
func (e *Element[T]) Next() *Element[T] {
	if next := e.next; e.list != nil && next != e.list.sentinel {
		return next
	}
	return nil
}

// A doubly linked list type
type List[T any] struct {
	sentinel *Element[T]
	size int
}

// New initializes an empty List
func NewList[T any]() *List[T] {
	sentinel := &Element[T]{}
	sentinel.prev = sentinel
	sentinel.next = sentinel

	l := &List[T]{
		sentinel: sentinel,
		size: 0,
	}
	sentinel.list = l

	return l
}

// Back return the last element of this list. Return nil if the list is empty
func (l *List[T]) Back() *Element[T] {
	return l.sentinel.Prev()
}

// Front return the first element of this list. Return nil if the list is empty
func (l *List[T]) Front() *Element[T] {
	return l.sentinel.Next()
}

// Clear all elements in the List.
func (l *List[T]) Clear() {
	// Clear pointers pointing to sentinel
	l.sentinel.next.prev = nil
	l.sentinel.prev.next = nil

	// Clear all elements in the list
	l.sentinel.prev = l.sentinel
	l.sentinel.next = l.sentinel
	l.size = 0
}

// Find a specific element with value t inside the list using reflect.DeepEqual method.
// Return the first occurrence inside the list, or return nil if the element 
// is not in the list
func (l *List[T]) Find(t T) *Element[T] {
	for e := l.Front(); e != nil; e = e.Next() {
		if reflect.DeepEqual(e.Value, t) {
			return e
		}
	}
	return nil
}

// InsertAfter inserts a new element with value t after the mark element
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(t T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}
	e := &Element[T]{Value: t}
	l.insertAfter(e, mark)
	return e
}

// InsertBefore inserts a new element with value t before the mark element
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertBefore(t T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}
	e := &Element[T]{Value: t}
	l.insertBefore(e, mark)
	return e
}

// Insert e after mark and returns e
func (l *List[T]) insertAfter(e, mark *Element[T]) *Element[T] {
	e.next = mark.next
	e.prev = mark
	e.next.prev = e
	e.prev.next = e
	e.list = l
	l.size++

	return e
}

// Insert e before mark and returns e
func (l *List[T]) insertBefore(e, mark *Element[T]) *Element[T] {
	e.next = mark
	e.prev = mark.prev
	e.next.prev = e
	e.prev.next = e
	e.list = l
	l.size++

	return e
}

// Len returns the length of the list
func (l *List[T]) Len() int {
	return l.size
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if e.list != l || mark.list != l || e == mark {
		return
	}

	l.Remove(e)
	l.insertAfter(e, mark)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	if e.list != l || mark.list != l || e == mark {
		return
	}

	l.Remove(e)
	l.insertBefore(e, mark)
}

// MoveToBack moves element e to the last element of the list.
// If e is not an element of lthe list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToBack(e *Element[T]) {
	l.MoveBefore(e, l.sentinel)
}

// MoveToFront moves element e to the first element of the list.
// If e is not an element of lthe list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToFront(e *Element[T]) {
	l.MoveAfter(e, l.sentinel)
}

// PushBack adds an element with value t at the back of the list
func (l *List[T]) PushBack(t T) *Element[T] {
	return l.InsertBefore(t, l.sentinel)
}

// PushBackList adds a list of elements at the back of the list
func (l *List[T]) PushBackList(other *List[T]) {
	for e := other.Front(); e != nil; e = e.Next() {
		l.PushBack(e.Value)
	}
}

// PushFront adds an element with value t at the front of the list
func (l *List[T]) PushFront(t T) *Element[T] {
	return l.InsertAfter(t, l.sentinel)
}

// PushFrontList adds a list of elements at the front of the list
func (l *List[T]) PushFrontList(other *List[T]) {
	for e := other.Front(); e != nil; e = e.Next() {
		l.PushFront(e.Value)
	}
}

// Remove remove the element from the list and return the value of the element
// if list does not contain e, nothing will happen
func (l *List[T]) Remove(e *Element[T]) T {
	if e.list == l {
		e.prev.next = e.next
		e.next.prev = e.prev
		e.next = nil
		e.prev = nil
		e.list = nil
		l.size--
	}
	return e.Value
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkSize(l *List[int], size int) {
	if l.Len() != size {
		panic(fmt.Sprintf("Wrong length, got %d, expect %d", l.Len(), size))
	}
}

func checkData(l *List[int], expect []*Element[int]) {
	checkSize(l, len(expect))

	node := l.sentinel
	if len(expect) == 0 {
		if node.next != node || node.prev != node {
			panic("Initialization failed")
		}
	} else {
		for _, e := range expect {
			if node.next != e || e.prev != node {
				panic("Node connection wrong")
			}
			if e.list != l {
				panic("Element's list doesn't match")
			}
			node = node.next
		}
	}
}

func testList() {
	l := NewList[int]()
	checkData(l, []*Element[int]{})

	e := l.PushBack(1)
	checkData(l, []*Element[int]{e})
	l.MoveToFront(e)
	checkData(l, []*Element[int]{e})
	l.MoveToBack(e)
	checkData(l, []*Element[int]{e})
	l.Remove(e)
	checkData(l, []*Element[int]{})

	e1 := l.PushBack(1)
	e2 := l.PushFront(2)
	e3 := l.PushBack(3)
	checkData(l, []*Element[int]{e2, e1, e3})
	l.MoveAfter(e2, e1)
	checkData(l, []*Element[int]{e1, e2, e3})
	l.MoveBefore(e3, e2)
	checkData(l, []*Element[int]{e1, e3, e2})
	l.MoveToFront(e3)
	checkData(l, []*Element[int]{e3, e1, e2})
	l.MoveToFront(e3)
	checkData(l, []*Element[int]{e3, e1, e2})
	l.MoveToBack(e3)
	checkData(l, []*Element[int]{e1, e2, e3})
	l.MoveToBack(e3)
	checkData(l, []*Element[int]{e1, e2, e3})

	e4 := l.InsertAfter(4, e2)
	checkData(l, []*Element[int]{e1, e2, e4, e3})
	l.Remove(e4)
	checkData(l, []*Element[int]{e1, e2, e3})
	e4 = l.InsertAfter(4, e3)
	checkData(l, []*Element[int]{e1, e2, e3, e4})
	l.Remove(e4)
	checkData(l, []*Element[int]{e1, e2, e3})

	e4 = l.InsertBefore(4, e2)
	checkData(l, []*Element[int]{e1, e4, e2, e3})
	l.Remove(e4)
	checkData(l, []*Element[int]{e1, e2, e3})
	e4 = l.InsertBefore(4, e1)
	checkData(l, []*Element[int]{e4, e1, e2, e3})
	l.Remove(e4)
	checkData(l, []*Element[int]{e1, e2, e3})

	l.Clear()
	checkData(l, []*Element[int]{})
}