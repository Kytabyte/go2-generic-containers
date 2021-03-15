package container

import (
	"fmt"
	"reflect"
)

// Vector is a generic type of go slice with more methods provided
type Vector[T any] []T

// NewVector returns a new Vector object
func NewVector[T any]() *Vector[T] {
	return &Vector[T]{}
}

// Get the ith element of the Vector
func (v Vector[T]) Get(i int) T {
	return v[i]
}

// Append a new element at the tail of the Vector
func (v *Vector[T]) Append(x T) {
	*v = append(*v, x)
}

// Pop the tail element from the the Vector
// Vector must not be empty
func (v *Vector[T]) Pop() T {
	old := *v
	n := len(old)
	x := old[n-1]
	*v = old[:n-1]
	return x
}

// Insert a new element at the given index i
// index i must be satisfied 0 <= i <= v.Len()
func (v *Vector[T]) Insert(i int, t T) {
	old := *v
	*v = append(old[:i], append([]T{t}, old[i:]...)...)
}

// Remove the element at given index
// index i must be satisfied 0 <= i < v.Len()
func (v *Vector[T]) Remove(i int) T {
	old := *v
	x := old[i]
	*v = append(old[:i], old[i+1:]...)
	return x
}

// Len returns the size of the Vector
func (v Vector[T]) Len() int {
	return len(v)
}

// Reverse the Vector
func (v Vector[T]) Reverse() {
	for i,j := 0, len(v)-1; i < j; i,j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
}

// Index the given element. If the element doesn't exist, return -1
func (v Vector[T]) Index(x T) int {
	for i, e := range v {
		if reflect.DeepEqual(e, x) {
			return i
		}
	}
	return -1
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkElements(vector *Vector[int], expect []int) {
	if vector.Len() != len(expect) {
		panic(fmt.Sprintf("Expect length is %d, but got %d.\n", len(expect), vector.Len()))
	}
	for i := range expect {
		if expect[i] != vector.Get(i) {
			panic(fmt.Sprintf("Expect the %dth element is %d, but got %d.\n", i, expect[i], vector.Get(i)))
		}
	}
}

func checkElement(elem, expect int) {
	if elem != expect {
		panic(fmt.Sprintf("Expect element to be %d, but got %d.\n", expect, elem))
	}
}

func testVector() {
	v := NewVector[int]()
	v.Append(1)
	checkElements(v, []int{1})
	num := v.Pop()
	checkElement(num, 1)
	checkElements(v, []int{})

	v.Append(1)
	v.Append(2)
	v.Append(3)
	checkElements(v, []int{1,2,3})
	v.Reverse()
	checkElements(v, []int{3,2,1})
	v.Insert(1, 4)
	checkElements(v, []int{3,4,2,1})
	v.Remove(2)
	checkElements(v, []int{3,4,1})
	index := v.Index(1)
	checkElement(index, 2)
	index = v.Index(2)
	checkElement(index, -1)
}
