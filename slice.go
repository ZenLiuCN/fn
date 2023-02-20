package fn

// ListEx just an Object like slice, with mapping support
//
//There is no fluent methods,should hold the return value as new instance.
//FIXME: golang(2023-02-12) not treat interfaces as comparable
type ListEx[K key, T, V any] List[T]

//ToMap see SliceToMap
func (s ListEx[K, T, V]) ToMap(mapping func(T) K, post func(T)) map[K]T {
	return SliceToMap(s, mapping, post)
}

//ToMapNoDup see SliceToMapNoDup
func (s ListEx[K, T, V]) ToMapNoDup(mapping func(T) K, post func(T)) map[K]T {
	return SliceToMapNoDup(s, mapping, post)
}

//Fold see SliceFold
func (s ListEx[K, T, V]) Fold(init V, fn func(V, T) V) V {
	return SliceFold(s, init, fn)
}

//FoldIndexed see SliceFoldIndexed
func (s ListEx[K, T, V]) FoldIndexed(init V, fn func(V, int, T) V) V {
	return SliceFoldIndexed(s, init, fn)
}

//FoldChunk see SliceFoldChunk
func (s ListEx[K, T, V]) FoldChunk(init V, size int, fn func(V, []T) V) V {
	return SliceFoldChunk(s, size, init, fn)
}

//Split see SliceSplit
func (s ListEx[K, T, V]) Split(fn func(T) []V) []V {
	return SliceSplit(s, fn)
}

//Mapping  see SliceMapping
func (s ListEx[K, T, V]) Mapping(fn func(T) V) []V {

	return SliceMapping(s, fn)
}

//Len same as len(s)
func (s ListEx[K, T, V]) Len() int {
	return len(s)
}

//Cap same as cap(s)
func (s ListEx[K, T, V]) Cap() int {
	return cap(s)
}

//Last same as s[len(s)-1]
func (s ListEx[K, T, V]) Last() T {
	return s[len(s)-1]
}

//LastIndex same as len(s) - 1
func (s ListEx[K, T, V]) LastIndex() int {
	return len(s) - 1
}

//Prepend same as append(t, s...)
func (s ListEx[K, T, V]) Prepend(t ...T) ListEx[K, T, V] {
	return append(t, s...)
}

//Append same as append(s, v...)
func (s ListEx[K, T, V]) Append(v ...T) ListEx[K, T, V] {
	return append(s, v...)
}

//Insert  same as append(append(s[:n], t...), s[n:]...)
func (s ListEx[K, T, V]) Insert(n int, t ...T) ListEx[K, T, V] {
	return append(append(s[:n], t...), s[n:]...)
}

//Reverse see SliceReverse
func (s ListEx[K, T, V]) Reverse() (r ListEx[K, T, V]) {
	return SliceReverse(s)
}

//ReverseSitu see SliceReverseSitu
func (s ListEx[K, T, V]) ReverseSitu() {
	SliceReverseSitu(s)
}

//Chunks see SliceChunkBy
func (s ListEx[K, T, V]) Chunks(n int) [][]T {
	return SliceChunkBy(s, n)
}

//Filter see SliceFilter
func (s ListEx[K, T, V]) Filter(fn func(T) bool) ListEx[K, T, V] {
	return SliceFilter(s, fn)
}

// IndexBy see SliceIndexBy
func (s ListEx[K, T, V]) IndexBy(fn func(T) bool) int {
	return SliceIndexBy(s, fn)
}

//LastIndexBy see SliceLastIndexBy
func (s ListEx[K, T, V]) LastIndexBy(fn func(T) bool) int {
	return SliceLastIndexBy(s, fn)
}

//DistinctBy see SliceDistinctBy
func (s ListEx[K, T, V]) DistinctBy(fn func(T, T) bool) ListEx[K, T, V] {
	return SliceDistinctBy(s, fn)
}

//List just an Object like slice, with some ability of slices.
// There is no fluent methods,SHOULD always hold the return value .
type List[T any] []T

//NewList same as List(make([]T,size,cap)
func NewList[T any](size, cap int) List[T] {
	return make(List[T], size, cap)
}

//ListOf same as List(s)
func ListOf[T any](s []T) List[T] {
	return s
}

//Len same as len(s)
func (s List[T]) Len() int {
	return len(s)
}

//Cap same as cap(s)
func (s List[T]) Cap() int {
	return cap(s)
}

//Last same as s[len(s)-1]
func (s List[T]) Last() T {
	return s[len(s)-1]
}

//LastIndex same as len(s) - 1
func (s List[T]) LastIndex() int {
	return len(s) - 1
}

//Prepend same as append(t, s...)
func (s List[T]) Prepend(t ...T) List[T] {
	return append(t, s...)
}

//Append same as append(s, v...)
func (s List[T]) Append(v ...T) List[T] {
	return append(s, v...)
}

//Insert  same as append(append(s[:n], t...), s[n:]...)
func (s List[T]) Insert(n int, t ...T) List[T] {
	return append(append(s[:n], t...), s[n:]...)
}

//Reverse see SliceReverse
func (s List[T]) Reverse() (r List[T]) {
	return SliceReverse(s)
}

//ReverseSitu see SliceReverseSitu
func (s List[T]) ReverseSitu() {
	SliceReverseSitu(s)
}

//Chunks see SliceChunkBy
func (s List[T]) Chunks(n int) [][]T {
	return SliceChunkBy(s, n)
}

//Filter see SliceFilter
func (s List[T]) Filter(fn func(T) bool) List[T] {
	return SliceFilter(s, fn)
}

// IndexBy see SliceIndexBy
func (s List[T]) IndexBy(fn func(T) bool) int {
	return SliceIndexBy(s, fn)
}

//LastIndexBy see SliceLastIndexBy
func (s List[T]) LastIndexBy(fn func(T) bool) int {
	return SliceLastIndexBy(s, fn)
}

//DistinctBy see SliceDistinctBy
func (s List[T]) DistinctBy(fn func(T, T) bool) List[T] {
	return SliceDistinctBy(s, fn)
}

// PrependSitu see SlicePrepend
func (s *List[T]) PrependSitu(val []T) {
	*s = SlicePrepend(*s, val...)
}

// InsertBeforeSitu see SliceInsertBefore
func (s *List[T]) InsertBeforeSitu(idx int, val []T) {
	*s = SliceInsertBefore(*s, idx, val...)
}

//SliceReverse reverse the order (new slice)
func SliceReverse[T any](s []T) (r []T) {
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return
}

//SliceReverseSitu reverse the order in place
func SliceReverseSitu[T any](s []T) {
	n := len(s)
	for i := n - 1; i >= n/2; i-- {
		s[n-i] = s[i]
	}
}

//SliceToMap slice to map without check duplication
func SliceToMap[K key, V any](slice []V, key func(V) K, post func(V)) (m map[K]V) {
	m = make(map[K]V)
	if len(slice) == 0 {
		return m
	}
	if post == nil {
		post = Consume[V]
	}
	for _, v := range slice {
		m[key(v)] = v
		post(v)
	}
	return
}

//SliceToMapNoDup slice to map with check duplication,panic with fn.ErrDuplicate
func SliceToMapNoDup[K key, V any](slice []V, key func(V) K, post func(V)) (m map[K]V) {
	m = make(map[K]V)
	if len(slice) == 0 {
		return m
	}
	if post == nil {
		post = Consume[V]
	}
	for _, v := range slice {
		k := key(v)
		if _, ok := m[k]; ok {
			panic(ErrDuplicate)
		}
		m[key(v)] = v
		post(v)
	}
	return
}

//SliceFilter filter by func
func SliceFilter[A any](a []A, fn func(A) bool) (r []A) {
	for _, a2 := range a {
		if fn(a2) {
			r = append(r, a2)
		}
	}
	return
}

//SliceIndexOf a slice
func SliceIndexOf[T key](s []T, v T) int {
	for i, t := range s {
		if t == v {
			return i
		}
	}
	return -1
}

//SliceLastIndexOf a slice
func SliceLastIndexOf[T key](s []T, v T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if v == s[i] {
			return i
		}
	}
	return -1
}

//SliceIndexBy with predicate
func SliceIndexBy[T any](s []T, pred func(T) bool) int {
	for i, t := range s {
		if pred(t) {
			return i
		}
	}
	return -1
}

//SliceLastIndexBy with predicate
func SliceLastIndexBy[T any](s []T, pred func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if pred(s[i]) {
			return i
		}
	}
	return -1
}

//SliceMapping a slice
func SliceMapping[A, B any](a []A, fn func(A) B) (r []B) {
	for _, a2 := range a {
		r = append(r, fn(a2))
	}
	return
}

//SliceSplit one to many
func SliceSplit[A, B any](a []A, fn func(A) []B) (r []B) {
	for _, a2 := range a {
		r = append(r, fn(a2)...)
	}
	return
}

//SliceChunkBy  if next<=0 || a==nil then nil else [[next],*,[remains]]
func SliceChunkBy[A any](a []A, n int) (r [][]A) {
	if n <= 0 || a == nil {
		return
	}
	l := len(a)
	if l <= n {
		return append(r, a)
	}
	f := l / n
	p := f * n
	r = make([][]A, 0, f+1)
	for i := 0; i < p; i += n {
		r = append(r, a[i:i+n])
	}
	if p < l {
		r = append(r, a[p:])
	}
	return
}

//SliceFold slice mapping to one
func SliceFold[A, B any](s []A, init B, act func(B, A) B) B {
	if s == nil {
		return init
	}
	for _, e := range s {
		init = act(init, e)
	}
	return init
}

//SliceFoldIndexed slice mapping to one with index.
func SliceFoldIndexed[A, B any](s []A, init B, act func(B, int, A) B) B {
	if s == nil {
		return init
	}
	for i, e := range s {
		init = act(init, i, e)
	}
	return init
}

//SliceFoldChunk fold on sub slice.
func SliceFoldChunk[A, B any](s []A, n int, init B, act func(B, []A) B) B {
	switch {
	case s == nil || n <= 0:
		return init
	case len(s) < n:
		return act(init, s)
	default:
		f := len(s) / n
		p := f * n
		for i := 0; i < p; i += n {
			init = act(init, s[i:i+n])
		}
		if p < len(s) {
			init = act(init, s[p:])
		}
		return init
	}
}

//SliceDistinct remove duplication
//FIXME: golang(2023-02-12) not treat interfaces as comparable so compare interface==nil is impossible
func SliceDistinct[A key](s []A) (r []A) {
	switch {
	case s == nil:
		return s
	default:
		set := make(map[A]struct{}, len(s))
		for _, a := range s {
			if _, ok := set[a]; !ok {
				r = append(r, a)
				set[a] = Nothing
			}
		}
		return
	}
}

//SliceDistinctBy remove duplication with an equality function;
//FIXME: golang(2023-02-12) not treat interfaces as comparable so compare interface==nil is impossible
func SliceDistinctBy[A any](s []A, eq func(A, A) bool) (r []A) {
	switch {
	case s == nil:
		return s
	default:
		for _, o := range s {
			for _, v := range r {
				if eq(v, o) {
					goto next
				}
			}
			r = append(r, o)
		next:
		}
		return
	}
}

// SliceInsertBefore
//
//	+ idx==zero => SlicePrepend,
//	+ idx < 0 && |idx| < len(s) =>  SliceInsertBefore(s,len(s)+idx,vals...)
//	+ |idx| == len(s) =>  [s...,vals...]
//	+ idx < 0 && |idx| > len(s) =>  vals
//	+ idx > 0 && |idx| > len(s) =>  s
//	+ idx > 0 && |idx| < len(s) =>  [...,vals...,idx...]
//
//this is a simple implement without any performance improve yet.
func SliceInsertBefore[A any](s []A, idx int, val ...A) (r []A) {
do:
	if idx == 0 {
		return append(val, s...)
	} else if idx == len(s) {
		return append(s, val...)
	} else if idx > 0 && idx < len(s) {
		n := make([]A, len(s)+len(val))
		copy(n, s[:idx])
		copy(n[idx:], val)
		copy(n[idx+len(val):], s[idx:])
		return n
	} else if idx < 0 && -idx < len(s) {
		idx = len(s) + idx
		goto do
	} else if idx < 0 {
		return val
	} else {
		return s
	}
}

// SliceInsertBeforeSitu same as SliceInsertBefore but keep result in s
func SliceInsertBeforeSitu[A any](s *[]A, idx int, val ...A) {
do:
	if idx == 0 {
		*s = append(val, *s...)
	} else if idx == len(*s) {
		*s = append(*s, val...)
	} else if idx > 0 && idx < len(*s) {
		n := make([]A, len(*s)+len(val))
		copy(n, (*s)[:idx])
		copy(n[idx:], val)
		copy(n[idx+len(val):], (*s)[idx:])
		*s = n
	} else if idx < 0 && -idx < len(*s) {
		idx = len(*s) + idx
		goto do
	} else if idx < 0 {
		*s = val
	} else {

	}
}

// SlicePrepend simply prepend
func SlicePrepend[A any](s []A, val ...A) (r []A) {
	return append(val, s...)
}
