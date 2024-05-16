package fn

import (
	"bytes"
	"testing"
)

var (
	p = NewTinyPool[*bytes.Buffer](func() *bytes.Buffer {
		return new(bytes.Buffer)
	})
	x = NewByteBufferPool()
)

func TestTinyPool(t *testing.T) {
	t.Parallel()
	b := p.Get()
	b.WriteByte('1')
	p.Put(b)
}
func BenchmarkTinyPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		y := p.Get()
		y.WriteByte('1')
		y.Reset()
		p.Put(y)
	}

}
func BenchmarkBufferPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		y := x.Get()
		y.WriteByte('1')
		x.Put(y)
	}

}
