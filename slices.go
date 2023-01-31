package fn

//ToMap slice to map without check duplication
func ToMap[K comparable, V any](
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

//ToMapNoDup slice to map with check duplication,panic with [fn/ErrDuplicate]
func ToMapNoDup[K comparable, V any](
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

//Filter filter
func Filter[A any](a []A, fn func(A) bool) (r []A) {
	for _, a2 := range a {
		if fn(a2) {
			r = append(r, a2)
		}
	}
	return
}

//Index a slice
func Index[T comparable](s []T, v T) int {
	for i, t := range s {
		if t == v {
			return i
		}
	}
	return -1
}

//IndexBack a slice
func IndexBack[T comparable](s []T, v T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if v == s[i] {
			return i
		}
	}
	return -1
}

//IndexPred with predicate
func IndexPred[T any](s []T, pred func(T) bool) int {
	for i, t := range s {
		if pred(t) {
			return i
		}
	}
	return -1
}

//IndexPredBack with predicate
func IndexPredBack[T any](s []T, pred func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if pred(s[i]) {
			return i
		}
	}
	return -1
}

//Map a slice
func Map[A, B any](a []A, fn func(A) B) (r []B) {
	for _, a2 := range a {
		r = append(r, fn(a2))
	}
	return
}

//MapMany one to many
func MapMany[A, B any](a []A, fn func(A) []B) (r []B) {
	for _, a2 := range a {
		r = append(r, fn(a2)...)
	}
	return
}

//Chunks  if n<=0 || a==nil then nil else [[n],*,[remains]]
func Chunks[A any](a []A, n int) (r [][]A) {
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

//Fold slice
func Fold[A, B any](s []A, init B, act func(B, A) B) B {
	if s == nil {
		return init
	}
	for _, e := range s {
		init = act(init, e)
	}
	return init
}

//FoldIdx with index
func FoldIdx[A, B any](s []A, init B, act func(B, int, A) B) B {
	if s == nil {
		return init
	}
	for i, e := range s {
		init = act(init, i, e)
	}
	return init
}

//FoldN fold n sub slice
func FoldN[A, B any](s []A, n int, init B, act func(B, []A) B) B {
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

//Distinct remove duplication
func Distinct[A comparable](s []A) (r []A) {
	switch {
	case s == nil:
		return s
	default:
		set := make(map[A]struct{}, len(s))
		for _, a := range s {
			if _, ok := set[a]; !ok {
				r = append(r, a)
				set[a] = EmptyValue
			}
		}
		return
	}
}
