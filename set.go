package container

// Set is an hash set data structure wrapped by go's map
type Set[K comparable] map[K]struct{}

// NewSet returns a new Set object
func NewSet[K comparable]() Set[K] {
	return make(Set[K])
}

// Insert a new element into the Set
func (s Set[K]) Insert(k K) {
	s[k] = struct{}{}
}

// Has the given key in the Set
func (s Set[K]) Has(k K) bool {
	_, ok := s[k]
	return ok
}

// Remove the given key from the Set.
// It has no-op if the key doesn't exists in Set
func (s Set[K]) Remove(k K) {
	delete(s, k)
}

// Len returns the size of the Set
func (s Set[K]) Len() int {
	return len(s)
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkHas[T comparable](s Set[T], entry T, expect bool) {
	if got := s.Has(entry); got != expect {
		panic(fmt.Sprintf("Wrong status of entry %v, expect %v, got %v", entry, expect, got))
	}
}

type checkSetType struct {
	first int
	second string
}

func testSet() {
	s := NewSet[string]()
	checkHas[string](s, "apple", false)
	s.Insert("apple")
	checkHas[string](s, "apple", true)
	s.Insert("banana")
	checkHas[string](s, "apple", true)
	checkHas[string](s, "banana", true)
	s.Remove("banana")
	checkHas[string](s, "apple", true)
	checkHas[string](s, "banana", false)

	s2 := NewSet[checkSetType]()
	s2.Insert(checkSetType{1, "apple"})
	checkHas[checkSetType](s2, checkSetType{1, "apple"}, true)
	s2.Remove(checkSetType{1, "apple"})
	checkHas[checkSetType](s2, checkSetType{1, "apple"}, false)
}
