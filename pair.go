package container

// Pair defines a key-value pair type largely used in map containers
type Pair[K, V any] struct {
	Key K
	Value V
}