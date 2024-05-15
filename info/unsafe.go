package info

import (
	"math/bits"
	"unsafe"
)

var (
	PointerSize = IntSize
	IntSize     = BitSize / 8
	BitSize     = bits.UintSize
)

type StringUnsafeHeader struct {
	Data *[0x7fff0000]byte
	Len  int
}
type SliceUnsafeHeader struct {
	Data *[0x7fff0000]byte
	Len  int
	Cap  int
}

// StringData unsafe.StringData in older than go v1.20
func StringData(v string) []byte {
	n := len(v)
	if n == 0 {
		return nil
	}
	p := (*StringUnsafeHeader)(unsafe.Pointer(&v))
	return p.Data[:n:n]
}
func ToString(v []byte) string {
	n := len(v)
	if n == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&v))
}

/*
Reader create read function at offset of struct with unsafe.

# This should only use to assign to a global variable, It's not safe!

examples:

1. read slice underlying:

	//underlying slice is struct{ data *[size]T, len int,cap int }
	pointerReader := Reader[[]Some, uintptr](0)
	lenReader := Reader[[]Some, int](PointerSize)
	capReader := Reader[[]Some, int](PointerSize + IntSize)

2. read string underlying:

	//underlying string is struct{ data []byte len int }
	dataReader := Reader[string, []byte](0)
	lenReader := Reader[string, int](PointerSize)

2. read first field

	reader := Reader[Some, int](0)
*/
func Reader[T, F any](offset int) (fn func(*T) F) {
	return func(t *T) F {
		return *(*F)(unsafe.Pointer(uintptr(unsafe.Pointer(t)) + uintptr(offset)))
	}
}
func Writer[T, K any](offset int) func(*T, K) {
	return func(t *T, v K) {
		p := unsafe.Pointer(t)
		*(*K)(unsafe.Pointer(uintptr(p) + uintptr(offset))) = v
	}
}

//goland:noinspection GoVetUnsafePointer
func Caster[T any]() func(p uintptr) *T {
	return func(p uintptr) *T {
		return (*T)(unsafe.Pointer(p))
	}
}

// MakeAccessor see also [ReaderPointer] [WriterPointer]
func MakeAccessor[T, F any](field Field) (a Accessor) {
	a[0] = Reader[T, F](field.Offset)
	a[1] = Writer[T, F](field.Offset)
	return a
}

func SizeOf[T any]() uint {
	var v T
	return uint(unsafe.Sizeof(v))
}
