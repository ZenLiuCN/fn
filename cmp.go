package fn

import (
	"golang.org/x/exp/constraints"
	"slices"
)

func NotNil[T ~*V, V any](t T) bool {
	return t != nil
}
func Positive[T constraints.Ordered](t T) bool {
	var z T
	return t > z
}
func Negative[T constraints.Ordered](t T) bool {
	var z T
	return t < z
}
func Zero[T constraints.Ordered](t T) bool {
	var z T
	return t == z
}

// Empty check map or slice is empty
func Empty[V any, K comparable, T ~[]V | ~map[K]V](t T) bool {
	return len(t) == 0
}

// NotEmpty check map or slice is empty
func NotEmpty[V any, K comparable, T ~[]V | ~map[K]V](t T) bool {
	return len(t) > 0
}

// HasSuffix check slice have end with values.
func HasSuffix[S ~[]V, V comparable](s S, v ...V) bool {
	switch ss, vs := len(s), len(v); true {
	case ss < vs:
		return false
	case vs == 0:
		return true
	default:
		return slices.Equal(s[ss-vs:], v)
	}
}

// HasPrefix check slice have begun with values.
func HasPrefix[S ~[]V, V comparable](s S, v ...V) bool {
	switch ss, vs := len(s), len(v); true {
	case ss < vs:
		return false
	case vs == 0:
		return true
	default:
		return slices.Equal(s[:vs], v)
	}
}
