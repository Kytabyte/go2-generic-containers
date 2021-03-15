package container

import "math"
import "math/rand"
import "sort"

// avlNode is the AVLTree node structure
type avlNode[K, V any] struct {
	Key K
	Value V
	Left *avlNode[K,V]
	Right *avlNode[K,V]
	size int
	height int
}

// Height of a avlNode. nil node returns 0
func (node *avlNode[K,V]) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

// Size of a avlNode including self. nil node returns 0
func (node *avlNode[K,V]) Size() int {
	if node == nil {
		return 0
	}
	return node.size
}

// Maintain the size and height of a avlNode. nil node has no-op
func (node *avlNode[K,V]) Maintain() {
	if node == nil {
		return
	}
	node.height = 1 + max(node.Left.Height(), node.Right.Height())
	node.size = 1 + node.Left.Size() + node.Right.Size()
}

// Bal returns the height difference of a avlNode's left children and right children
func (node *avlNode[K,V]) Bal() int {
	if node == nil {
		return 0
	}
	return node.Left.Height() - node.Right.Height()
}

// AVLTree data structure with no duplicated value
type AVLTree[K, V any] struct {
	root *avlNode[K,V]
	cmp func(K,K) int
}

// NewAVLTree creates a new AVLTree given a comparator of key type
func NewAVLTree[K,V any](cmp func(K,K) int) *AVLTree[K,V] {
	return &AVLTree[K,V]{
		cmp: cmp,
	}
}

// Has key in the AVLTree
func (t *AVLTree[K,V]) Has(key K) bool {
	_, ok := t.Get(key)
	return ok
}

// MustGet returns value of given key if the key exists, otherwise
// returns zero-value of V
func (t *AVLTree[K,V]) MustGet(key K) (rVal V) {
	if v, ok := t.Get(key); ok {
		rVal = v
	}
	return
}

// Get key from AVLTree, return value and whether the key is found
// if not found, return value is the zero-value of type V
func (t *AVLTree[K,V]) Get(key K) (rVal V, ok bool) {
	node := t.root
	for node != nil {
		cmp := t.cmp(node.Key, key)
		if cmp == 0 {
			rVal, ok = node.Value, true
			return
		}
		
		if cmp > 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return
}

// GetFloor returns the entry less than or equal to the given key if exists
func (t *AVLTree[K,V]) GetFloor(key K) (rKey K, rVal V, ok bool) {
	node := t.root 
	for node != nil {
		cmp := t.cmp(node.Key, key)
		if cmp == 0 {
			rKey, rVal, ok = node.Key, node.Value, true
			return
		}

		if cmp > 0 {
			node = node.Left
		} else {
			rKey, rVal, ok = node.Key, node.Value, true
			node = node.Right
		}
	}
	return
}

// GetCeiling returns the entry greater than or equal to the given key if exists
func (t *AVLTree[K,V]) GetCeiling(key K) (rKey K, rVal V, ok bool) {
	node := t.root
	for node != nil {
		cmp := t.cmp(node.Key, key)
		if cmp == 0 {
			rKey, rVal, ok = node.Key, node.Value, true
			return
		}
		
		if cmp > 0 {
			rKey, rVal, ok = node.Key, node.Value, true
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return
}

// GetLower returns the entry less than the given key if exists
func (t *AVLTree[K,V]) GetLower(key K) (rKey K, rVal V, ok bool) {
	node := t.root
	for node != nil {
		cmp := t.cmp(node.Key, key)
		if cmp >= 0 {
			node = node.Left
		} else {
			rKey, rVal, ok = node.Key, node.Value, true
			node = node.Right
		}
	}
	return
}

// GetCeiling returns the entry greater than the given key if exists
func (t *AVLTree[K,V]) GetHigher(key K) (rKey K, rVal V, ok bool) {
	node := t.root
	for node != nil {
		cmp := t.cmp(node.Key, key)
		if cmp <= 0 {
			node = node.Right
		} else {
			rKey, rVal, ok = node.Key, node.Value, true
			node = node.Left
		}
	}
	return
}

// GetFirst returns the smallest element (according to cmp) in the AVLTree if exists
func (t *AVLTree[K,V]) GetFirst() (rKey K, rVal V, ok bool) {
	node := t.root
	for node != nil && node.Left != nil {
		node = node.Left
	}
	if node != nil {
		rKey, rVal, ok = node.Key, node.Value, true
	}
	return
}

// GetFirst returns the greatest element (according to cmp) in the AVLTree if exists
func (t *AVLTree[K,V]) GetLast() (rKey K, rVal V, ok bool) {
	node := t.root
	for node != nil && node.Right != nil {
		node = node.Right
	}
	if node != nil {
		rKey, rVal, ok = node.Key, node.Value, true
	}
	return
}

// Len return the size of the AVLTree
func (t *AVLTree[K,V]) Len() int {
	return t.root.Size()
}

// Insert a key-value pair into the AVLTree
func (t *AVLTree[K,V]) Insert(key K, value V) {
	t.root = t.insert(t.root, key, value)
}

// recursively find insert position
func (t *AVLTree[K,V]) insert(node *avlNode[K,V], key K, value V) *avlNode[K,V] {
	if node == nil {
		return &avlNode[K,V]{
			Key: key,
			Value: value,
			size: 1,
			height: 1,
		}
	}

	if cmp := t.cmp(node.Key, key); cmp == 0 {
		node.Value = value
		return node
	} else if cmp > 0 {
		node.Left = t.insert(node.Left, key, value)
	} else {
		node.Right = t.insert(node.Right, key, value)
	}

	node.Maintain()
	return t.fix(node)
}

// Remove the node with given key in the AVLTree if exists
// it has no-op if key is not in the tree
func (t *AVLTree[K,V]) Remove(key K) {
	t.root = t.remove(t.root, key)
}

// recursively find remove position
func (t *AVLTree[K,V]) remove(node *avlNode[K,V], key K) *avlNode[K,V] {
	if node == nil {
		return nil
	}

	if cmp := t.cmp(node.Key, key); cmp == 0 {
		if node.Left == nil && node.Right == nil { // no child
			return nil
		}
		if node.Left == nil { // right child only
			return node.Right
		}
		if node.Right == nil { // left child only
			return node.Left
		}
		// both sides have children
		nxtNode := t.successor(node)
		node.Key, node.Value = nxtNode.Key, nxtNode.Value
		node.Right = t.remove(node.Right, nxtNode.Key)
	} else if cmp > 0 {
		node.Left = t.remove(node.Left, key)
	} else {
		node.Right = t.remove(node.Right, key)
	}

	node.Maintain()
	return t.fix(node)
}

// find the next greater element
func (t *AVLTree[K,V]) successor(node *avlNode[K,V]) *avlNode[K,V] {
	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node
} 

// fix the tree to become a balanced binary search tree
func (t *AVLTree[K,V]) fix(node *avlNode[K,V]) *avlNode[K,V] {
	if bal := node.Bal(); -1 <= bal && bal <= 1 { // balanced
		return node
	} else if bal > 1 {
		if node.Left.Bal() < 0 { // left-right rotate
			node.Left = t.rotateLeft(node.Left)
		}
		return t.rotateRight(node) // right rotate
	} else {
		if node.Right.Bal() > 0 { // right-left rotate
			node.Right = t.rotateRight(node.Right)
		}
		return t.rotateLeft(node) // left rotate
	}
}

/*
		n							r
	  /   \						  /   \
	l	   r     ====>           n    rr
		  / \					/\		\
		 rl rr				   l  rl	 x
		      \
			   x
*/
func (t *AVLTree[K,V]) rotateLeft(node *avlNode[K,V]) *avlNode[K,V] {
	root := node.Right
	node.Right = root.Left
	root.Left = node
	node.Maintain()
	root.Maintain()
	return root
}


/*
			n							l
		  /   \						  /   \
		l	   r     ====>           ll    n
	   / \							/ 	  / \
	  ll lr				   		   x  	 lr  r
	 /
	x	   
*/
func (t *AVLTree[K,V]) rotateRight(node *avlNode[K,V]) *avlNode[K,V] {
	root := node.Left
	node.Left = root.Right
	root.Right = node
	node.Maintain()
	root.Maintain()
	return root
}

// Clear all element in the AVLTree
func (t *AVLTree[K,V]) Clear() {
	t.root = nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

/////////////////////////////
///////// Testing ///////////
/////////////////////////////

func checkBalBST(t *AVLTree[int, int]) {
	var check func(node *avlNode[int, int]) (int, int, int, bool)
	check = func(node *avlNode[int, int]) (int, int, int, bool) {
		if node == nil {
			return math.MaxInt64, math.MinInt64, 0, true
		}
		lmin, lmax, lh, lok := check(node.Left)
		rmin, rmax, rh, rok := check(node.Right)

		ok := lok && rok && lmax <= node.Key && node.Key <= rmin && abs(lh-rh) <= 1
		if !ok {
			panic(fmt.Sprintf(
				"Node %v does not match a Balanced BST property, got Left(%d,%d,%d,%t), Right(%d,%d,%d,%t)",
				node, lmin, lmax, lh, lok, rmin, rmax, rh, rok))
		}
		return min(lmin, rmin), max(lmax, rmax), 1+max(lh, rh), ok
	}

	check(t.root)
}

func checkInorderSorted(t *AVLTree[int, int]) {
	keys := []int{}

	var dfs func(node *avlNode[int, int])
	dfs = func(node *avlNode[int, int]) {
		if node != nil {
			dfs(node.Left)
			keys = append(keys, node.Key)
			dfs(node.Right)
		}
	}

	dfs(t.root)

	if (!sort.IntsAreSorted(keys)) {
		panic("Tree is not balanced.")
	}
}

func testAVLTreeElements(t *AVLTree[int, int], size int) {
	checkBalBST(t)
	checkInorderSorted(t)
	if t.Len() != size {
		panic(fmt.Sprintf("Expect tree size to be %d, but got %d.\n", size, t.Len()))
	}
}

func checkAVLTreeElement(key, val int, ok int, expKey, expVal int, expOk bool) {
	if key != expKey || val != expVal || ok != expOk {
		panic(fmt.Sprintf("Expect key, val, ok to be %d,%d,%t, but got %d,%d,%t.\n", key, val, ok, expKey,expVal, expOk))
	}
}

func checkAVLTree12345(t *AVLTree[int, int]) {
	noNeedKey, noNeedValue := 0,0
	var key, val int
	var ok bool

	val, ok = t.Get(3)
	checkAVLTreeElement(noNeedKey, val, ok, noNeedKey, 0, true)
	key, val, ok = t.GetLower(3)
	checkAVLTreeElement(key, 0, ok, 2, 0, true)
	key, val, ok = t.GetHigher(4)
	checkAVLTreeElement(key, val, ok, 5, 0, true)
	key, val, ok = t.GetLower(1)
	checkAVLTreeElement(noNeedKey, noNeedValue, ok, noNeedKey, noNeedValue, false)
	key, val, ok = t.GetHigher(5)
	checkAVLTreeElement(noNeedKey, noNeedValue, ok, noNeedKey, noNeedValue, false)
	key, val, ok = t.GetCeiling(0)
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetFloor(6)
	checkAVLTreeElement(key, val, ok, 5, 0, true)
	key, val, ok = t.GetCeiling(2)
	checkAVLTreeElement(key, val, ok, 2, 0, true)
	key, val, ok = t.GetFloor(4)
	checkAVLTreeElement(key, val, ok, 4, 0, true)
	key, val, ok = t.GetFirst()
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetLast()
	checkAVLTreeElement(key, val, ok, 5, 0, true)
}

func checkAVLTree123(t *AVLTree[int, int]) {
	noNeedKey, noNeedValue := 0,0
	var key, val int
	var ok bool

	val, ok = t.Get(3)
	checkAVLTreeElement(noNeedKey, val, ok, noNeedKey, 0, true)
	key, val, ok = t.GetLower(3)
	checkAVLTreeElement(key, 0, ok, 2, 0, true)
	key, val, ok = t.GetHigher(4)
	checkAVLTreeElement(noNeedKey, noNeedValue, ok, noNeedKey, noNeedValue, false)
	key, val, ok = t.GetCeiling(0)
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetFloor(6)
	checkAVLTreeElement(key, val, ok, 3, 0, true)
	key, val, ok = t.GetCeiling(2)
	checkAVLTreeElement(key, val, ok, 2, 0, true)
	key, val, ok = t.GetFirst()
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetLast()
	checkAVLTreeElement(key, val, ok, 3, 0, true)
}

func checkAVLTree1234(t *AVLTree[int, int]) {
	noNeedKey, noNeedValue := 0,0
	var key, val int
	var ok bool

	val, ok = t.Get(3)
	checkAVLTreeElement(noNeedKey, val, ok, noNeedKey, 0, true)
	key, val, ok = t.GetLower(3)
	checkAVLTreeElement(key, 0, ok, 2, 0, true)
	key, val, ok = t.GetHigher(4)
	checkAVLTreeElement(noNeedKey, noNeedValue, ok, noNeedKey, noNeedValue, false)
	key, val, ok = t.GetLower(1)
	checkAVLTreeElement(noNeedKey, noNeedValue, ok, noNeedKey, noNeedValue, false)
	key, val, ok = t.GetCeiling(0)
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetFloor(6)
	checkAVLTreeElement(key, val, ok, 4, 0, true)
	key, val, ok = t.GetCeiling(2)
	checkAVLTreeElement(key, val, ok, 2, 0, true)
	key, val, ok = t.GetFloor(4)
	checkAVLTreeElement(key, val, ok, 4, 0, true)
	key, val, ok = t.GetFirst()
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetLast()
	checkAVLTreeElement(key, val, ok, 4, 0, true)
}

func testAVLTree() {
	t := NewAVLTree[int, int](CmpLess[int])
	testAVLTreeElements(t,0)

	noNeedKey, noNeedValue := 0,0
	var key, val int
	var ok bool

	// test one entry
	t.Insert(1,0)
	testAVLTreeElements(t,1)
	val, ok = t.Get(1)
	checkAVLTreeElement(noNeedKey, val, ok, noNeedKey, 0, true)
	_, val, ok = t.GetLower(1)
	checkAVLTreeElement(noNeedKey, noNeedValue, ok, noNeedKey, noNeedValue, false)
	_, val, ok = t.GetHigher(1)
	checkAVLTreeElement(noNeedKey, noNeedValue, ok, noNeedKey, noNeedValue, false)
	key, val, ok = t.GetCeiling(1)
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetFloor(1)
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetFirst()
	checkAVLTreeElement(key, val, ok, 1, 0, true)
	key, val, ok = t.GetLast()
	checkAVLTreeElement(key, val, ok, 1, 0, true)

	t.Remove(1)
	testAVLTreeElements(t,0)

	// test multiple entries no rotation
	t.Insert(3,0)
	t.Insert(2,0)
	t.Insert(4,0)
	t.Insert(1,0)
	t.Insert(5,0)
	testAVLTreeElements(t,5)
	checkAVLTree12345(t)


	t.Clear()
	testAVLTreeElements(t,0)

	// test rotate left
	t.Insert(1,0)
	t.Insert(2,0)
	t.Insert(3,0)
	testAVLTreeElements(t,3)
	checkAVLTree123(t)
	t.Insert(5,0)
	t.Insert(4,0)
	checkAVLTree12345(t)
	t.Remove(5)
	checkAVLTree1234(t)
	t.Remove(4)
	checkAVLTree123(t)
	t.Remove(6) // test no-op
	checkAVLTree123(t)

	t.Clear()
	testAVLTreeElements(t,0)

	// test rotate right
	t.Insert(3,0)
	t.Insert(2,0)
	t.Insert(1,0)
	testAVLTreeElements(t,3)
	checkAVLTree123(t)
	t.Insert(5,0)
	t.Insert(4,0)
	checkAVLTree12345(t)
	t.Remove(5)
	checkAVLTree1234(t)
	t.Remove(4)
	checkAVLTree123(t)
	t.Remove(6) // test no-op
	checkAVLTree123(t)

	t.Clear()
	testAVLTreeElements(t,0)

	// test rotate left right
	t.Insert(3,0)
	t.Insert(1,0)
	t.Insert(2,0)
	testAVLTreeElements(t,3)
	checkAVLTree123(t)
	t.Insert(5,0)
	t.Insert(4,0)
	checkAVLTree12345(t)
	t.Remove(5)
	checkAVLTree1234(t)
	t.Remove(4)
	checkAVLTree123(t)
	t.Remove(6) // test no-op
	checkAVLTree123(t)

	t.Clear()
	testAVLTreeElements(t,0)

	// test rotate right left
	t.Insert(3,0)
	t.Insert(2,0)
	t.Insert(1,0)
	testAVLTreeElements(t,3)
	checkAVLTree123(t)
	t.Insert(4,0)
	checkAVLTree1234(t)
	t.Insert(5,0)
	checkAVLTree12345(t)
	t.Remove(5)
	checkAVLTree1234(t)
	t.Remove(4)
	checkAVLTree123(t)
	t.Remove(6) // test no-op
	checkAVLTree123(t)

	t.Clear()
	testAVLTreeElements(t,0)

	// random check
	nInsert := 5 + rand.Intn(10)
	insertNums := []int{}
	uniqueNums := make(map[int]struct{}) // to avoid duplicate numbers
	// insert random numbers
	for i := 0; i < nInsert; i++ {
		num := rand.Intn(math.MaxInt32)
		for _, ok = uniqueNums[num]; ok; {
			num = rand.Intn(math.MaxInt32)
		}
		insertNums = append(insertNums, num)
		uniqueNums[num] = struct{}{}
		t.Insert(num, num-1)
		testAVLTreeElements(t, i+1)
	}
	rand.Shuffle(nInsert, func(i, j int) {
		insertNums[i], insertNums[j] = insertNums[j], insertNums[i]
	})
	// remove random numbers from inserted ones
	nRemove := 1 + rand.Intn(nInsert-1)
	for i := 0; i < nRemove; i++ {
		val, ok = t.Get(insertNums[i])
		checkAVLTreeElement(noNeedKey, val, ok, noNeedKey, insertNums[i]-1, true)
		t.Remove(insertNums[i])
		testAVLTreeElements(t, nInsert-i-1)
		val, ok = t.Get(insertNums[i])
		checkAVLTreeElement(noNeedKey, noNeedValue, ok, noNeedKey, noNeedValue, false)
	}

	// random insert and remove
	for i := nRemove; i < nInsert; i++ {
		num := -rand.Intn(math.MaxInt32)
		t.Insert(num, num+1)
		t.Remove(insertNums[i])
		testAVLTreeElements(t, nInsert-nRemove)
	}
}

