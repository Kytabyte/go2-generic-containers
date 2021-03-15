package container

// TODO: this file should be in a separated package (e.g. cmp)
// and functions shoudl be named as e.g. cmp.Less, cmp.SliceLess

// Ordered interface copied from Go generics design draft 
// https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md
type Ordered interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		string 
}

// CmpLess compares if the first element is less than the second element
func CmpLess[T Ordered](x, y T) int {
	if x == y {
		return 0
	}
	if x > y {
		return 1
	}
	return -1
}

// CmpLess compares if the first element is greater than the second element
func CmpGreater[T Ordered](x, y T) int {
	return CmpLess[T](y, x)
}

// CmpSliceLess compares two slices element-by-element
func CmpSliceLess[T Ordered](x, y []T) int {
	for i, j := 0, 0; i < len(x) || j < len(y); i, j = i+1, j+1 {
		if i == len(x) {
			return -1
		}
		if j == len(y) {
			return 1
		}
		if x[i] < y[j] {
			return -1
		}
		if x[i] > y[j] {
			return 1
		}
	}
	return 0
}

// CmpSliceGreater compares two slices element-by-element
func CmpSliceGreater[T Ordered](x, y []T) int {
	return CmpSliceLess[T](y, x)
}