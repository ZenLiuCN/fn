package fn

type Stack[T comparable] []T

func (s Stack[T]) First() (v T) {
	if s.Empty() {
		return
	}
	return s[0]
}
func (s Stack[T]) Nth(n int) T {
	if len(s) > n {
		return s[n]
	}
	var empty T
	return empty
}
func (s Stack[T]) LNth(n int) T {
	if len(s) >= n {
		return s[len(s)-n]
	}
	var empty T
	return empty
}
func (s Stack[T]) Last() (v T) {
	if s.Empty() {
		return
	}
	return s[s.Len()-1]
}
func (s Stack[T]) Equals(x ...T) bool {
	if s.Empty() || len(x) != s.Len() {
		return false
	}

	return SliceEquals(s, x)
}
func (s Stack[T]) EndWith(x ...T) bool {
	if s.Empty() || len(x) > s.Len() {
		return false
	}
	nx := len(x)
	ns := len(s)
	return SliceEquals(s[ns-nx:ns], x)
}
func (s Stack[T]) StartWith(x ...T) bool {
	if s.Empty() || len(x) > s.Len() {
		return false
	}
	nx := len(x)
	return SliceEquals(s[0:nx], x)
}
func (s Stack[T]) Continues(v T, skip int) (x int) {
	if s.Empty() {
		return
	}
	n := s.Len()
	if skip >= n {
		return
	}
	for i := skip; i < n; i++ {
		if s[i] == v {
			x++
		} else {
			return
		}
	}
	return
}
func (s Stack[T]) LastContinues(v T, skip int) (x int) {
	if s.Empty() {
		return
	}
	n := s.Len()
	if skip >= n {
		return
	}
	for i := n - 1 - skip; i >= 0; i-- {
		if s[i] == v {
			x++
		} else {
			return
		}
	}
	return
}
func (s Stack[T]) Match(anyone T, pattern ...T) bool {
	p := len(pattern)
	if p == 0 || (p == 1 && pattern[0] == anyone) {
		return true
	}
	if s.Empty() {
		return false
	}
	n := s.Len()
	if p > n {
		return false
	}
	x := 0
	for i := 0; i < n; i++ {
		if s[i] == pattern[x] {

		}
	}
	return false
}
func (s Stack[T]) LastIndex(v T) int {
	if s.Empty() {
		return -1
	}
	n := s.Len()
	for i := n - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}
func (s Stack[T]) Len() int {
	return len(s)
}
func (s Stack[T]) Empty() bool {
	return len(s) == 0
}
func (s Stack[T]) MaxIdx() int {
	return len(s) - 1
}

func (s Stack[T]) NotEmpty() bool {
	return len(s) > 0
}
func (s Stack[T]) Pop() (x Stack[T], v T) {
	if len(s) == 0 {
		return
	}
	v = s[len(s)-1]
	x = s[:len(s)-1]
	return
}
func (s Stack[T]) Push(v T) (x Stack[T]) {
	x = append(s, v)
	return
}
func (s Stack[T]) Clean() (x Stack[T]) {
	x = s[0:0]
	return
}

func (s *Stack[T]) PopSitu() (v T) {
	if len(*s) == 0 {
		return
	}
	v = s.Last()
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack[T]) PushSitu(v T) {
	*s = append(*s, v)
	return
}
func (s *Stack[T]) CleanSitu() {
	*s = (*s)[0:0]
}

// MatchesKMP match pattern v with KMP algorithm
func (s Stack[T]) MatchesKMP(v ...T) (x int) {
	if s.Empty() {
		return
	}
	return MatchKMP(s, v...)
}

// MatchesZ match pattern v with Z Algorithm,p is the separator that will not exist in v and Stack
func (s Stack[T]) MatchesZ(p T, v ...T) (x int) {
	if s.Empty() {
		return
	}
	return MatchZ(s, p, v...)
}
