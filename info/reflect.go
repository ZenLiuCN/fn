// Package info contains toolset for lightweight reflection like actions for types
package info

func Empty[T any]() (v T) {
	return
}
func EmptyRefer[T any]() *T {
	var x T
	return &x
}
func Refer[T any](x T) (v *T) {
	v = &x
	return
}

func DeRefer[T any](x *T) (v T) {
	return *x
}
func AssignRefer[T any](v T, x *T) *T {
	*x = v
	return x
}
func SliceOf[T any](x ...T) []T {
	return x
}

// MakeSlice with positive optional capacity and length
func MakeSlice[T any](cl ...int) []T {
	switch len(cl) {
	case 0:
		return make([]T, 0)
	case 1:
		return make([]T, 0, cl[0])
	default:
		return make([]T, cl[1], cl[0])
	}
}

// MakeMap with positive optional size
func MakeMap[K comparable, T any](size ...int) (m map[K]T) {
	if len(size) > 0 {
		return make(map[K]T, size[0])
	}
	return make(map[K]T)
}
