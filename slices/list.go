package slices

//List just an Object like slice.
type List[T any] []T

func NewList[T any](size, cap int) List[T] {
	return make(List[T], size, cap)
}
func ListOf[T any](s []T) List[T] {
	return s
}

func (s List[T]) Len() int {
	return len(s)
}
func (s List[T]) Cap() int {
	return cap(s)
}

func (s List[T]) Last() T {
	return s[len(s)-1]
}
func (s List[T]) LastIndex() int {
	return len(s) - 1
}

func (s List[T]) Prepend(t ...T) List[T] {
	return append(t, s...)
}
func (s List[T]) Push(v ...T) List[T] {
	return append(s, v...)
}
func (s List[T]) Insert(n int, t ...T) List[T] {
	return append(append(s[:n], t...), s[n:]...)
}

func (s List[T]) Pop() (v T, list List[T]) {
	v = s.Last()
	list = s[:s.LastIndex()]
	return
}
func (s List[T]) PopN(n int) List[T] {
	return s[len(s)-n-1:]
}

func (s List[T]) Sub(f, t int) List[T] {
	return s[f:t]
}

func (s List[T]) Each(c func(T) bool) {
	for _, t := range s {
		if c(t) {
			break
		}
	}
}

func (s List[T]) EachIndex(c func(int, T) bool) {
	for i, t := range s {
		if c(i, t) {
			break
		}
	}
}
