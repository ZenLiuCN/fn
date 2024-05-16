package fn

import (
	"sync"
	"testing"
)

func BenchmarkSyncMap(b *testing.B) {
	v := new(SyncMap[int, int])
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v.Store(i, i)
	}
}
func BenchmarkOriginSyncMap(b *testing.B) {
	v := new(sync.Map)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v.Store(i, i)
	}
}
func BenchmarkSyncMapGet(b *testing.B) {
	v := new(SyncMap[int, int])
	for i := 0; i < b.N; i++ {
		v.Store(i, i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if x, _ := v.Load(i); x != i {
			b.Fail()
		}
	}
}
func BenchmarkOriginSyncMapGet(b *testing.B) {
	v := new(sync.Map)
	for i := 0; i < b.N; i++ {
		v.Store(i, i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if x, _ := v.Load(i); x.(int) != i {
			b.Fail()
		}
	}
}
func BenchmarkSyncMapDelete(b *testing.B) {
	v := new(SyncMap[int, int])
	for i := 0; i < b.N; i++ {
		v.Store(i, i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Delete(i)
	}
}
func BenchmarkOriginSyncMapDelete(b *testing.B) {
	v := new(sync.Map)
	for i := 0; i < b.N; i++ {
		v.Store(i, i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Delete(i)
	}
}
