package fn

//List just an Object like slice, with some ability of slices.
// There is no fluent methods,should hold the return value as new instance.
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

func (s List[T]) Chunks(n int) [][]T {
	return SliceChunkBy(s, n)
}
func (s List[T]) Filter(fn func(T) bool) List[T] {
	return SliceFilter(s, fn)
}

func (s List[T]) IndexBy(fn func(T) bool) int {
	return SliceIndexBy(s, fn)
}

func (s List[T]) LastIndexBy(fn func(T) bool) int {
	return SliceLastIndexBy(s, fn)
}

func (s List[T]) DistinctBy(fn func(T, T) bool) List[T] {
	return SliceDistinctBy(s, fn)
}

//SliceToMap slice to map without check duplication
func SliceToMap[K key, V any](
	slice []V,
	key func(V) K,
	post func(V),
) (m map[K]V) {
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

//SliceToMapNoDup slice to map with check duplication,panic with [fn/ErrDuplicate]
func SliceToMapNoDup[K key, V any](
	slice []V,
	key func(V) K,
	post func(V),
) (m map[K]V) {
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

//SliceMap a slice
func SliceMap[A, B any](a []A, fn func(A) B) (r []B) {
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

//SliceFold slice
func SliceFold[A, B any](s []A, init B, act func(B, A) B) B {
	if s == nil {
		return init
	}
	for _, e := range s {
		init = act(init, e)
	}
	return init
}

//SliceFoldIndexed with index
func SliceFoldIndexed[A, B any](s []A, init B, act func(B, int, A) B) B {
	if s == nil {
		return init
	}
	for i, e := range s {
		init = act(init, i, e)
	}
	return init
}

//SliceFoldChunk fold next sub slice
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

//SliceDistinctBy remove duplication with an equality function
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
