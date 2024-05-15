package fn

import "golang.org/x/exp/constraints"

func ConvString2BytesRef[R any](fn func(*R) string) func(*R) []byte {
	return func(r *R) []byte {
		return []byte(fn(r))
	}
}
func ConvString2Bytes[R any](fn func(R) string) func(R) []byte {
	return func(r R) []byte {
		return []byte(fn(r))
	}
}
func ConvString2RunesRef[R any](fn func(*R) string) func(*R) []rune {
	return func(r *R) []rune {
		return []rune(fn(r))
	}
}
func ConvString2Runes[R any](fn func(R) string) func(R) []rune {
	return func(r R) []rune {
		return []rune(fn(r))
	}
}

func ConvBytes2StringRef[R any](fn func(*R) []byte) func(*R) string {
	return func(r *R) string {
		return string(fn(r))
	}
}
func ConvBytes2String[R any](fn func(R) []byte) func(R) string {
	return func(r R) string {
		return string(fn(r))
	}
}
func ConvRunes2StringRef[R any](fn func(*R) []rune) func(*R) string {
	return func(r *R) string {
		return string(fn(r))
	}
}
func ConvRunes2String[R any](fn func(R) []rune) func(R) string {
	return func(r R) string {
		return string(fn(r))
	}
}

func ConvNum2Num[F, T constraints.Integer | constraints.Float](fn func() F) func() T {
	return func() (o T) {
		return T(fn())
	}
}

func ConvAnySlice[T any, S ~[]any](s S) (v []T, fail int) {
	if len(s) == 0 {
		return
	}
	fail = -1
	for i, a := range s {
		val, ok := a.(T)
		if !ok {
			fail = i
			return
		}
		v = append(v, val)
	}
	return
}

func ConvToAny[T any](fn func() T) func() any {
	return func() any {
		return fn()
	}
}
