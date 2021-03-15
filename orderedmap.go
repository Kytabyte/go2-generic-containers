package container

import "fmt"

// OrderedMap is a hash map with the preservation of insertion order
type OrderedMap[K comparable, V any] struct {
	mp map[K]*Element[Pair[K,V]]
	list *List[Pair[K,V]]
}

// NewOrderedMap returns a new OrderedMap object
func NewOrderedMap[K comparable, V any]() *OrderedMap[K,V] {
	return &OrderedMap[K,V]{
		mp: make(map[K]*Element[Pair[K,V]]),
		list: NewList[Pair[K,V]](),
	}
}

// MustGet returns the value of the given key if the key exists,
// otherwise it returns the zero-value of V
func (m *OrderedMap[K,V]) MustGet(k K) (rVal V) {
	if v, ok := m.Get(k); ok {
		rVal = v
	}
	return
}

// Get returns the value of given key and if the key exists
// if the keys doesn't exists, returned value is the zero-value of V
func (m *OrderedMap[K,V]) Get(k K) (V, bool) {
	var v V
	if e, ok := m.mp[k]; ok {
		v = e.Value.Value
		return v, true
	}
	return v, false
}

// Has the given key in the OrderedMap
func (m *OrderedMap[K,V]) Has(k K) bool {
	_, ok := m.mp[k]
	return ok
}

// Insert a key-value pair into the OrderedMap
func (m *OrderedMap[K,V]) Insert(k K, v V) {
	if m.Has(k) {
		m.mp[k].Value = Pair[K,V]{Key: k, Value: v}
	} else {
		pair := Pair[K,V]{Key: k, Value: v}
		e := m.list.PushBack(pair)
		m.mp[k] = e
	}
}

// Remove a key-value pair with given key, returns the value and if the key exists.
// If the key doesn't exist, returned value is the zero-value of V
func (m *OrderedMap[K,V]) Remove(k K) (v V, ok bool) {
	if !m.Has(k) {
		return
	}
	e := m.mp[k]
	m.list.Remove(e)
	delete(m.mp, k)
	v, ok = e.Value.Value, true
	return
}

// Clear all elements in the OrderedMap
func (m *OrderedMap[K,V]) Clear() {
	m.list.Clear()
	for k := range m.mp {
		delete(m.mp, k)
	}
}

// Move to an element with given key to the tail of the OrderedMap
// the given key must exist in the OrderedMap
func (m *OrderedMap[K,V]) MoveToBack(k K) {
	if !m.Has(k) {
		panic(fmt.Sprintf("Key %v does not in the map", k))
	}
	e := m.mp[k]
	m.list.MoveToBack(e)
}

// Move to an element with given key to the head of the OrderedMap
// the given key must exist in the OrderedMap
func (m *OrderedMap[K,V]) MoveToFront(k K) {
	if !m.Has(k) {
		panic(fmt.Sprintf("Key %v does not in the map", k))
	}
	e := m.mp[k]
	m.list.MoveToFront(e)
} 

// Removes the tail element of the OrderedMap and returns the key-value pair
// and if the tail element exists.
// The tail element doesn't exist if and only if the OrderedMap is empty.
func (m *OrderedMap[K,V]) PopBack() (k K, v V, ok bool) {
	e := m.list.Back()
	if e == nil {
		return
	}
	k, v, ok = e.Value.Key, e.Value.Value, true
	delete(m.mp, k)
	m.list.Remove(e)
	return
}

// Removes the head element of the OrderedMap and returns the key-value pair
// and if the head element exists.
// The head element doesn't exist if and only if the OrderedMap is empty.
func (m *OrderedMap[K,V]) PopFront() (k K, v V, ok bool) {
	e := m.list.Front()
	if e == nil {
		return
	}
	k, v, ok = e.Value.Key, e.Value.Value, true
	delete(m.mp, k)
	m.list.Remove(e)
	return
}

// Len returns the size of the OrderedMap
func (m *OrderedMap[K,V]) Len() int {
	return len(m.mp)
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkOMapSize(m *OrderedMap[string, int], expect int) {
	if len(m.mp) != m.list.Len() {
		panic(fmt.Sprintf("The length of map and list do not match, got map len %d, list len %d", len(m.mp), m.list.Len()))
	}
	if m.Len() != expect {
		panic(fmt.Sprintf("Len check failed, got %d, expect, %d", m.Len(), expect))
	}
}

func checkOMapOrder(m *OrderedMap[string, int], expKeys []string, expVals []int) {
	checkOMapSize(m, len(expKeys))

	for i, e1 := 0, m.list.Front(); e1 != nil; i, e1 = i+1, e1.Next() {
		k, v := expKeys[i], expVals[i]
		if e2, ok := m.mp[k]; !ok || e1 != e2 {
			if !ok {
				panic(fmt.Sprintf("Key %v should be in map but missed", k))
			}
			if e1 != e2 {
				panic(fmt.Sprintf("map and list point to different elements for key %v", k))
			}
		}
		if k != e1.Value.Key || v != e1.Value.Value {
			panic(fmt.Sprintf("element entry failed, got (k,v): (%v,%v), expected (%v,%v)", e1.Value.Key, e1.Value.Value, k, v))
		}
	}
}

func checkOMapKV(gotK string, gotV int, gotOk bool, expK string, expV int, expOk bool) {
	if gotK != expK || gotV != expV || gotOk != expOk {
		panic(fmt.Sprintf("k,v,ok does not match where key:(%v,%v), val:(%v,%v), ok:(%v,%v)",
		 	gotK, expK, gotV, expV, gotOk, expOk))
	}
}

func testOrderedMap() {
	var k string
	var v int
	var ok bool


	m := NewOrderedMap[string, int]()
	checkOMapOrder(m, []string{}, []int{})
	
	// single element
	m.Insert("apple", 1)
	checkOMapOrder(m, []string{"apple"}, []int{1})
	ok = m.Has("apple")
	checkOMapKV(k, v, ok, k, v, true)
	v, ok = m.Get("apple")
	checkOMapKV(k, v, ok, k, 1, true)
	m.Insert("apple", 2)
	checkOMapOrder(m, []string{"apple"}, []int{2})
	v, ok = m.Get("apple")
	checkOMapKV(k, v, ok, k, 2, true)
	v, ok = m.Get("banana")
	checkOMapKV(k, v, ok, k, 0, false)
	m.MoveToFront("apple")
	checkOMapOrder(m, []string{"apple"}, []int{2})
	m.MoveToBack("apple")
	checkOMapOrder(m, []string{"apple"}, []int{2})
	v, ok = m.Remove("apple")
	checkOMapKV(k, v, ok, k, 2, true)
	checkOMapOrder(m, []string{}, []int{})
	v, ok = m.Get("apple")
	checkOMapKV(k, v, ok, k, 0, false)
	ok = m.Has("banana")
	checkOMapKV(k, v, ok, k, v, false)

	// multiple elements
	m.Insert("apple", 1)
	m.Insert("banana", 2)
	m.Insert("cherry", 3)
	checkOMapOrder(m, []string{"apple", "banana", "cherry"}, []int{1,2,3})
	m.MoveToBack("banana")
	checkOMapOrder(m, []string{"apple", "cherry", "banana"}, []int{1,3,2})
	m.MoveToFront("cherry")
	checkOMapOrder(m, []string{"cherry", "apple", "banana"}, []int{3,1,2})
	m.Insert("cherry", -1)
	checkOMapOrder(m, []string{"cherry", "apple", "banana"}, []int{-1,1,2})
	k, v, ok = m.PopFront()
	checkOMapKV(k, v, ok, "cherry", -1, true)
	checkOMapOrder(m, []string{"apple", "banana"}, []int{1,2})
	m.Insert("cherry", 3)
	checkOMapOrder(m, []string{"apple", "banana", "cherry"}, []int{1,2,3})
	k, v, ok = m.PopBack()
	checkOMapKV(k, v, ok, "cherry", 3, true)
	checkOMapOrder(m, []string{"apple", "banana"}, []int{1,2})
	v, ok = m.Remove("banana")
	checkOMapKV(k, v, ok, k, 2, true)
	m.Clear()
	checkOMapOrder(m, []string{}, []int{})
}