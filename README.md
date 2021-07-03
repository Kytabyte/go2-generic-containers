# Go2 Generic Containers

This is a repo to implement some data structures in go's generic syntax. 
The go generic is actually under development and the proposal can be 
found [here](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md).

Since the design might be subject to change, this repo will be updated accordingly.

## Supported Containers

- [Vector](vector.go) (a wrapper of go slice)
- [LinkedList](list.go) (go's container/list with generic supported)
- Stack (impl using both [LinkedList](stack.go) and [slice](arraystack.go))
- Queue (impl using both [LinkedList](queue.go) and [slice](arrayqueue.go))
- Deque (impl using both [LinkedList](deque.go) and [slice](arraydeque.go))
- [Pair](pair.go)
- [HashSet](set.go) (a wrapper of `map[T]struct{}`)
- [PriortyQueue](priorityqueue.go)
- [OrderedMap](orderedmap.go) (hash map with insertion order preserved, e.g. can be used as LRU-cache)
- [OrderedSet](orderedset.go)
- [AVLTree](avltree.go)

## Test Script

My merged test script can be found in this [go2go-playground](https://go2goplay.golang.org/p/NIaoQEPkmGe).

Note that at this time, the `go2go playground` is the only platform that to run go-generic-style code...

## TODO

### More data structures
- Red Black Tree
- `SortedMap` and `SortedSet` (need benchmark of AVL Tree and Red Black Tree to decide which one to use)

### Iterator of Containers

A unified iterator interface for all containers is needed...

### Package structure

Since the [go2go-playground](https://go2goplay.golang.org/) is the only way to write go-generic-style code, 
writing everything in the same package will be much easier to merge them into a single file and do testing
there. The package structure is subject to change after the generic feature is released.

### Format

`go fmt` is not available at this time, as the generic syntax will be recognized as syntax error. The code format will be updated after the geenric feature is released.
