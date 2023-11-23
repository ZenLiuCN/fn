package fn

import "io"

func Consume[V any](v V) {

}
func Not[A any](fn func(A) bool) func(A) bool {
	return func(a A) bool {
		return !fn(a)
	}
}
func Or[A any](fa func(A) bool, fb func(A) bool) func(A) bool {
	return func(a A) bool {
		return fa(a) || fb(a)
	}
}
func And[A any](fa func(A) bool, fb func(A) bool) func(A) bool {
	return func(a A) bool {
		return fa(a) && fb(a)
	}
}
func Identity[T any](v T) T {
	return v
}
func SafeClose(closer io.Closer) func() {
	return func() {
		if err := closer.Close(); err != nil {
			panic(err)
		}
	}
}
func IgnoreClose(closer io.Closer) func() {
	return func() {
		_ = closer.Close()
	}
}
