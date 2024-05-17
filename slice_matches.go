package fn

func SliceEquals[E comparable, S ~[]E](a, b S) bool {
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
func SliceEqualsFunc[E any, S ~[]E](a, b S, equals func(E, E) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equals(b[i], a[i]) {
			return false
		}
	}
	return true
}
func MatchZ[T comparable, S ~[]T](s S, b T, p ...T) int {
	sn, pn := len(s), len(p)
	if pn == 0 { //!! ALL MATCH
		return 0
	} else if pn > sn { //!! never match
		return -1
	} else if pn == sn && SliceEquals(s, p) {
		return 0
	}
	c := make([]T, sn+pn+1)
	copy(c, p)
	c[pn] = b
	copy(c[pn+1:], s)
	z := AlgZ(c)
	for i := 0; i < sn+pn+1; i++ {
		if z[i] == pn {
			return i - pn - 1
		}
	}
	return -1
}

func AlgZ[T comparable, S ~[]T](s S) (z []int) {
	n := len(s)
	if n == 0 {
		return
	}
	z = make([]int, n)
	z[0] = n
	for i, l, r := 1, 0, 0; i < n; i++ {
		if z[i-l] < r-i+1 {
			z[i] = z[i-l]
		} else {
			z[i] = max(r-i+1, 0)
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				z[i]++
			}
			l, r = i, i+z[i]-1
		}
	}
	return
}
func MatchZFunc[T any, S ~[]T](equals func(T, T) bool, s S, b T, p ...T) int {
	sn, pn := len(s), len(p)
	if pn == 0 { //!! ALL MATCH
		return 0
	} else if pn > sn { //!! never match
		return -1
	} else if pn == sn && SliceEqualsFunc(s, p, equals) {
		return 0
	}
	c := make([]T, sn+pn+1)
	copy(c, p)
	c[pn] = b
	copy(c[pn+1:], s)
	z := AlgZFunc(equals, c)
	for i := 0; i < sn+pn+1; i++ {
		if z[i] == pn {
			return i - pn - 1
		}
	}
	return -1
}
func AlgZFunc[T any, S ~[]T](equals func(T, T) bool, s S) (z []int) {
	n := len(s)
	if n == 0 {
		return
	}
	z = make([]int, n)
	z[0] = n
	for i, l, r := 1, 0, 0; i < n; i++ {
		if z[i-l] < r-i+1 {
			z[i] = z[i-l]
		} else {
			z[i] = max(r-i+1, 0)
			for i+z[i] < n && equals(s[z[i]], s[i+z[i]]) {
				z[i]++
			}
			l, r = i, i+z[i]-1
		}
	}
	return
}
func MatchKMP[T comparable, S ~[]T](s S, p ...T) int {
	sn, pn := len(s), len(p)
	if pn == 0 { //!! ALL MATCH
		return 0
	} else if pn > sn { //!! never match
		return -1
	} else if pn == sn && SliceEquals(s, p) {
		return 0
	}
	n := AlgKMP(p)
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
func AlgKMP[T comparable, S ~[]T](s S) (k []int) {
	pn := len(s)
	if pn == 0 {
		return
	}
	k = make([]int, pn)
	k[0] = -1
	i, j := 2, 0
	for i < pn {
		switch {
		case s[i-1] == s[j]:
			k[i] = j + 1
			j++
			i++
		case j > 0:
			j = k[j]
		default:
			i++
		}
	}
	return
}

func MatchKMPFunc[T any, S ~[]T](equals func(T, T) bool, s S, p ...T) int {
	sn, pn := len(s), len(p)
	if pn == 0 { //!! ALL MATCH
		return 0
	} else if pn > sn { //!! never match
		return -1
	} else if pn == sn && SliceEqualsFunc(s, p, equals) {
		return 0
	}
	n := AlgKMPFunc(equals, p)
	i, j := 0, 0
	for i < sn {
		switch {
		case equals(p[i], s[j+i]):
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
func AlgKMPFunc[T any, S ~[]T](equals func(T, T) bool, s S) (k []int) {
	pn := len(s)
	if pn == 0 {
		return
	}
	k = make([]int, pn)
	k[0] = -1
	i, j := 2, 0
	for i < pn {
		switch {
		case equals(s[i-1], s[j]):
			k[i] = j + 1
			j++
			i++
		case j > 0:
			j = k[j]
		default:
			i++
		}
	}
	return
}
