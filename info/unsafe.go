package info

import "unsafe"

var (
	PointerSize = int(SizeOf[unsafe.Pointer]())
	IntSize     = int(SizeOf[int]())
)

func SizeOf[T any]() uint {
	var t T
	return uint(unsafe.Sizeof(t))
}

/*
Reader create read function at offset of struct with unsafe. **It's not safe!**

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
func Reader[T, F any](offset int) func(*T) F {
	return func(t *T) F {
		p := unsafe.Pointer(t)
		return *(*F)(unsafe.Pointer(uintptr(p) + uintptr(offset)))
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
