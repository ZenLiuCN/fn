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

	return equals(s, x)
}
func (s Stack[T]) EndWith(x ...T) bool {
	if s.Empty() || len(x) > s.Len() {
		return false
	}
	nx := len(x)
	ns := len(s)
	return equals(s[ns-nx:ns], x)
}
func (s Stack[T]) StartWith(x ...T) bool {
	if s.Empty() || len(x) > s.Len() {
		return false
	}
	nx := len(x)
	return equals(s[0:nx], x)
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
func equals[E comparable, S ~[]E](a, b S) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return true
}
func AnyMatch[T comparable, S ~[]T](s S, anyOne T, anyMore T, p ...T) int {
	sn, pn := len(s), len(p)
	if pn == 0 || (pn == 1 && p[0] == anyOne && sn > 0) || p[0] == anyMore { //!! ALL MATCH
		return 0
	} else if pn > sn { //!! never match
		return -1
	}
	x, i, j, m, si := -1, 0, 0, 0, -1
	for i < sn {
		switch {
		case j < pn && (p[j] == anyOne || s[i] == p[j]):
			if x < 0 {
				x = i
			}
			i++
			j++
		case j < pn && p[j] == anyMore:
			si = j
			m = i
			j++
		case si != -1:
			j = si + 1
			m++
			i = m
		default:
			return -1
		}
	}
	for _, t := range p[j:] {
		if t != anyMore {
			return x
		}
	}
	if j == pn {
		return x
	}
	return -1
}
func Match[T comparable, S ~[]T](s S, p ...T) int {
	sn, pn := len(s), len(p)
	if pn == 0 { //!! ALL MATCH
		return 0
	} else if pn > sn { //!! never match
		return -1
	} else if pn == sn && equals(s, p) {
		return 0
	}
	var n []int
	{

		n = make([]int, pn)
		n[0] = -1
		i, j := 2, 0
		for i < pn {
			switch {
			case p[i-1] == p[j]:
				n[i] = j + 1
				j++
				i++
			case j > 0:
				j = n[j]
			default:
				i++
			}
		}
	}

	i, j := 0, 0
	for i < sn {
		switch {
		case p[i] == s[j+i]:
			if i == pn-1 {
				return j
			}
			i++
		case n[i] > -1:
			j += i
			i = n[i]
		default:
			j++
			i = 0
		}
	}
	return -1
}
