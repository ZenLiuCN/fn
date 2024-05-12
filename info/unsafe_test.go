package info

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Some struct {
	v int
}

func TestReader(t *testing.T) {
	reader := Reader[Some, int](0)
	var v Some

	for i := 0; i < 100; i++ {
		v.v = i
		assert.Equal(t, reader(&v), i)
	}
	t.Log(v.v)
}
func TestReaderSlice(t *testing.T) {
	pointerReader := Reader[[]Some, uintptr](0)
	lenReader := Reader[[]Some, int](PointerSize)
	capReader := Reader[[]Some, int](PointerSize + IntSize)
	var v []Some
	v = append(v, Some{1}, Some{2}, Some{3}, Some{4})
	assert.Equal(t, lenReader(&v), len(v))
	assert.Equal(t, capReader(&v), cap(v))
	t.Logf("%p:%d:%d", pointerReader, cap(v), len(v))
}
func TestReaderString(t *testing.T) {
	dataReader := Reader[string, []byte](0)
	lenReader := Reader[string, int](PointerSize)
	var v string = "123"
	assert.Equal(t, dataReader(&v), []byte(v))
	assert.Equal(t, lenReader(&v), 3)
	b := dataReader(&v)
	for _, b2 := range b {
		t.Log(b2)
	}
}
func BenchmarkReader(b *testing.B) {
	reader := Reader[Some, int](0)
	var v Some
	v.v = 1
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = reader(&v)
	}
}
func TestWriter(t *testing.T) {
	writer := Writer[Some, int](0)
	var v Some
	for i := 0; i < 100; i++ {
		writer(&v, i)
		assert.Equal(t, v.v, i)
	}
	t.Log(v.v)

}
func BenchmarkWriter(b *testing.B) {
	writer := Writer[Some, int](0)
	var v Some
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writer(&v, i)
	}
}
