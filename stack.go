package fn

type StackAny[T any] []T

func (s StackAny[T]) First() (v T) {
	if s.Empty() {
		return
	}
	return s[0]
}
func (s StackAny[T]) Nth(n int) T {
	if len(s) > n {
		return s[n]
	}
	var empty T
	return empty
}
func (s StackAny[T]) LNth(n int) T {
	if len(s) >= n {
		return s[len(s)-n]
	}
	var empty T
	return empty
}
func (s StackAny[T]) Last() (v T) {
	if s.Empty() {
		return
	}
	return s[s.Len()-1]
}
func (s StackAny[T]) Len() int {
	return len(s)
}
func (s StackAny[T]) Empty() bool {
	return len(s) == 0
}
func (s StackAny[T]) MaxIdx() int {
	return len(s) - 1
}

func (s StackAny[T]) NotEmpty() bool {
	return len(s) > 0
}
func (s StackAny[T]) Pop() (x StackAny[T], v T) {
	if len(s) == 0 {
		return
	}
	v = s[len(s)-1]
	x = s[:len(s)-1]
	return
}
func (s StackAny[T]) Push(v T) (x StackAny[T]) {
	x = append(s, v)
	return
}
func (s StackAny[T]) Clean() (x StackAny[T]) {
	x = s[0:0]
	return
}

func (s *StackAny[T]) PopSitu() (v T) {
	if len(*s) == 0 {
		return
	}
	v = s.Last()
	*s = (*s)[:len(*s)-1]
	return
}

func (s *StackAny[T]) PushSitu(v T) {
	*s = append(*s, v)
	return
}
func (s *StackAny[T]) CleanSitu() {
	*s = (*s)[0:0]
}

func (s StackAny[T]) LastIndex(v T, equals func(T, T) bool) int {
	if s.Empty() {
		return -1
	}
	n := s.Len()
	for i := n - 1; i >= 0; i-- {
		if equals(s[i], v) {
			return i
		}
	}
	return -1
}
func (s StackAny[T]) Equals(equals func(T, T) bool, x ...T) bool {
	if s.Empty() || len(x) != s.Len() {
		return false
	}

	return SliceEqualsFunc(s, x, equals)
}
func (s StackAny[T]) EndWith(equals func(T, T) bool, x ...T) bool {
	if s.Empty() || len(x) > s.Len() {
		return false
	}
	nx := len(x)
	ns := len(s)
	return SliceEqualsFunc(s[ns-nx:ns], x, equals)
}
func (s StackAny[T]) StartWith(equals func(T, T) bool, x ...T) bool {
	if s.Empty() || len(x) > s.Len() {
		return false
	}
	nx := len(x)
	return SliceEqualsFunc(s[0:nx], x, equals)
}
func (s StackAny[T]) Continues(equals func(T, T) bool, v T, skip int) (x int) {
	if s.Empty() {
		return
	}
	n := s.Len()
	if skip >= n {
		return
	}
	for i := skip; i < n; i++ {
		if equals(s[i], v) {
			x++
		} else {
			return
		}
	}
	return
}
func (s StackAny[T]) LastContinues(equals func(T, T) bool, v T, skip int) (x int) {
	if s.Empty() {
		return
	}
	n := s.Len()
	if skip >= n {
		return
	}
	for i := n - 1 - skip; i >= 0; i-- {
		if equals(s[i], v) {
			x++
		} else {
			return
		}
	}
	return
}

// MatchesKMP match pattern v with KMP algorithm
func (s StackAny[T]) MatchesKMP(equals func(T, T) bool, v ...T) (x int) {
	if s.Empty() {
		return
	}
	return MatchKMPFunc(equals, s, v...)
}

// MatchesZ match pattern v with Z Algorithm,p is the separator that will not exist in v and Stack
func (s StackAny[T]) MatchesZ(equals func(T, T) bool, p T, v ...T) (x int) {
	if s.Empty() {
		return
	}
	return MatchZFunc(equals, s, p, v...)
}

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
