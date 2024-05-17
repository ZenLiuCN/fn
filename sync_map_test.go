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
func TestSyncMapValidate(t *testing.T) {
	m := SyncMap[int, int]{}
	w := new(sync.WaitGroup)
	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 500; i++ {
			m.Store(i, i)
			t.Logf("store %d", i)
		}
	}()
	w.Add(1)
	go func() {
		defer w.Done()
		for i := 500; i < 1000; i++ {
			m.Store(i, i)
			t.Logf("store %d", i)
		}
	}()
	w.Wait()
	w.Add(1)
	go func() {
		defer w.Done()
		var i int
		for i = 0; i < 1000; i++ {
			if v, ok := m.Load(i); ok && v != i {
				t.Errorf("%d not match at %d", v, i)
				break
			}
		}
		t.Logf("last read %d", i)
	}()
	w.Wait()
}
