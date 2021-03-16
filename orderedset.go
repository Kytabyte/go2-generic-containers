package container

import "fmt"

// OrderedSet is a hash set with the preservation of insertion order
type OrderedSet[K comparable] OrderedMap[K, struct{}]

// NewOrderedSet returns a new OrderedSet object
func NewOrderedSet[K comparable]() *OrderedSet[K] {
	m := NewOrderedMap[K,struct{}]()
	return (*OrderedSet[K])(m)
}

// Has the given key in the OrderedSet
func (s *OrderedSet[K]) Has(k K) bool {
	m := (*OrderedMap[K, struct{}])(s)
	return m.Has(k)
}

// Insert a key into the OrderedSet
func (s *OrderedSet[K]) Insert(k K) {
	m := (*OrderedMap[K, struct{}])(s)
 	m.Insert(k, struct{}{})
}

// Remove the given key, returns if the key exists.
func (s *OrderedSet[K]) Remove(k K) bool {
	m := (*OrderedMap[K, struct{}])(s)
	_, ok := m.Remove(k)
	return ok
}

// Clear all elements in the OrderedSet
func (s *OrderedSet[K]) Clear() {
	m := (*OrderedMap[K, struct{}])(s)
	m.Clear()
}

// Move to an element with given key to the tail of the OrderedSet
// the given key must exist in the OrderedSet
func (s *OrderedSet[K]) MoveToBack(k K) {
	m := (*OrderedMap[K, struct{}])(s)
	m.MoveToBack(k)
}

// Move to an element with given key to the head of the OrderedSet
// the given key must exist in the OrderedSet
func (s *OrderedSet[K]) MoveToFront(k K) {
	m := (*OrderedMap[K, struct{}])(s)
	m.MoveToFront(k)
} 

// Removes the tail element of the OrderedSet and returns the key
// and if the tail element exists.
// The tail element doesn't exist if and only if the OrderedSet is empty.
func (s *OrderedSet[K]) PopBack() (k K, ok bool) {
	m := (*OrderedMap[K, struct{}])(s)
	k, _, ok = m.PopBack()
	return
}

// Removes the head element of the OrderedSet and returns the key
// and if the head element exists.
// The head element doesn't exist if and only if the OrderedSet is empty.
func (s *OrderedSet[K]) PopFront() (k K, ok bool) {
	m := (*OrderedMap[K, struct{}])(s)
	k, _, ok = m.PopFront()
	return
}

// Len returns the size of the OrderedSet
func (s *OrderedSet[K]) Len() int {
	m := (*OrderedMap[K, struct{}])(s)
	return m.Len()
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkOSetSize(s *OrderedSet[string], expect int) {
	if s.Len() != expect {
		panic(fmt.Sprintf("Len check failed, got %d, expect, %d", s.Len(), expect))
	}
}

func checkOSetOrder(s *OrderedSet[string], expKeys []string) {
	checkOSetSize(s, len(expKeys))

	m := (*OrderedMap[string, struct{}])(s)

	for i, e1 := 0, m.list.Front(); e1 != nil; i, e1 = i+1, e1.Next() {
		k := expKeys[i]
		if e2, ok := m.mp[k]; !ok || e1 != e2 {
			if !ok {
				panic(fmt.Sprintf("Key %v should be in map but missed", k))
			}
			if e1 != e2 {
				panic(fmt.Sprintf("map and list point to different elements for key %v", k))
			}
		}
		if k != e1.Value.Key {
			panic(fmt.Sprintf("element entry failed, got (k): (%v), expected (%v)", e1.Value.Key, k))
		}
	}
}

func checkOSetK(gotK string, gotOk bool, expK string, expOk bool) {
	if gotK != expK || gotOk != expOk {
		panic(fmt.Sprintf("k,ok does not match where key:(%v,%v), ok:(%v,%v)",
		 	gotK, expK, gotOk, expOk))
	}
}

func testOrderedSet() {
	var k string
	var ok bool


	s:= NewOrderedSet[string]()
	checkOSetOrder(s, []string{})
	
	// single element
	s.Insert("apple")
	checkOSetOrder(s, []string{"apple"})
	ok = s.Has("apple")
	checkOSetK(k, ok, k, true)
	s.Insert("apple")
	checkOSetOrder(s, []string{"apple"})
	s.MoveToFront("apple")
	checkOSetOrder(s, []string{"apple"})
	s.MoveToBack("apple")
	checkOSetOrder(s, []string{"apple"})
	ok = s.Remove("apple")
	checkOSetK(k, ok, k, true)
	checkOSetOrder(s, []string{})
	ok = s.Has("banana")
	checkOSetK(k, ok, k, false)

	// multiple elements
	s.Insert("apple")
	s.Insert("banana")
	s.Insert("cherry")
	checkOSetOrder(s, []string{"apple", "banana", "cherry"})
	s.MoveToBack("banana")
	checkOSetOrder(s, []string{"apple", "cherry", "banana"})
	s.MoveToFront("cherry")
	checkOSetOrder(s, []string{"cherry", "apple", "banana"})
	s.Insert("cherry")
	checkOSetOrder(s, []string{"cherry", "apple", "banana"})
	k, ok = s.PopFront()
	checkOSetK(k, ok, "cherry", true)
	checkOSetOrder(s, []string{"apple", "banana"})
	s.Insert("cherry")
	checkOSetOrder(s, []string{"apple", "banana", "cherry"})
	k, ok = s.PopBack()
	checkOSetK(k, ok, "cherry", true)
	checkOSetOrder(s, []string{"apple", "banana"})
	ok = s.Remove("banana")
	checkOSetK(k, ok, k,true)
	s.Clear()
	checkOSetOrder(s, []string{})
}