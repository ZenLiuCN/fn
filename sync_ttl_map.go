package fn

import "time"

type (
	// SyncTTLMap  full copy generic [sync.Map] with TTL support
	SyncTTLMap[K comparable, V any] struct {
		m SyncMap[K, V]
		t SyncMap[K, int64]
		*HouseKeeper
	}
	// HouseKeeper a period cleaner
	HouseKeeper struct {
		When    time.Duration
		ticker  *time.Ticker
		running chan struct{}
	}
	// Purifier consume UnixNanos
	Purifier = func(int64)
)

func NewSyncTTLMap[K comparable, V any](cleanRate time.Duration) *SyncTTLMap[K, V] {
	return &SyncTTLMap[K, V]{
		m:           SyncMap[K, V]{},
		t:           SyncMap[K, int64]{},
		HouseKeeper: &HouseKeeper{When: cleanRate},
	}
}
func NewSyncTTLMapWithoutKeeper[K comparable, V any]() *SyncTTLMap[K, V] {
	return &SyncTTLMap[K, V]{
		m: SyncMap[K, V]{},
		t: SyncMap[K, int64]{},
	}
}
func (s *SyncTTLMap[K, V]) Range(fn func(K, V) bool) {
	s.m.Range(fn)
}
func (s *SyncTTLMap[K, V]) Store(k K, v V) {
	s.m.Store(k, v)
}
func (s *SyncTTLMap[K, V]) Load(k K) (v V, ok bool) {
	var t int64
	v, t, ok = s.LoadWithNanos(k)
	if ok && t > 0 && t > time.Now().UnixNano() {
		ok = false
	}
	return
}
func (s *SyncTTLMap[K, V]) LoadAndDelete(k K) (v V, ok bool) {
	var t int64
	v, t, ok = s.LoadAndDeleteWithNanos(k)
	if ok && t > 0 && t > time.Now().UnixNano() {
		ok = false
	}
	return
}
func (s *SyncTTLMap[K, V]) LoadAndDeleteWithNanos(k K) (v V, t int64, ok bool) {
	v, ok = s.m.Load(k)
	if ok {
		defer s.m.Delete(k)
	}
	if ok {
		var ok2 bool
		t, ok2 = s.t.Load(k)
		if ok2 {
			defer s.t.Delete(k)
		}
	}
	return
}
func (s *SyncTTLMap[K, V]) Delete(k K) {
	s.m.Delete(k)
	s.t.Delete(k)
}
func (s *SyncTTLMap[K, V]) LoadWithNanos(k K) (v V, t int64, ok bool) {
	v, ok = s.m.Load(k)
	if ok {
		t, _ = s.t.Load(k)
	}
	return
}

// LoadOrStore not change the TTL
func (s *SyncTTLMap[K, V]) LoadOrStore(k K, v V) (actual V, loaded bool) {
	return s.m.LoadOrStore(k, v)
}

// Swap not change the TTL
func (s *SyncTTLMap[K, V]) Swap(k K, v V) (previous V, loaded bool) {
	return s.m.Swap(k, v)
}

// CompareAndSwap not change the TTL
func (s *SyncTTLMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return s.m.CompareAndSwap(key, old, new)
}

func (s *SyncTTLMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	if s.m.CompareAndDelete(key, old) {
		s.t.Delete(key)
		return true
	}
	return false
}
func (s *SyncTTLMap[K, V]) StoreWithTTL(key K, value V, after time.Duration) (a int64) {
	a = time.Now().Add(after).UnixNano()
	s.m.Store(key, value)
	s.t.Store(key, a)
	return a
}
func (s *SyncTTLMap[K, V]) StoreWithDDL(key K, value V, ddl time.Time) (a int64) {
	a = ddl.UnixNano()
	s.m.Store(key, value)
	s.t.Store(key, a)
	return
}
func (s *SyncTTLMap[K, V]) Run() {
	if s.HouseKeeper != nil {
		s.HouseKeeper.Run(s.Purify)
	}

}
func (s *SyncTTLMap[K, V]) Stop() {
	if s.HouseKeeper != nil {
		s.HouseKeeper.Stop()
	}
}
func (s *SyncTTLMap[K, V]) Purify(now int64) {
	s.t.Range(func(k K, value int64) bool {
		if value <= now {
			s.m.Delete(k)
			s.t.Delete(k)
		}
		return false
	})
}

func NewHouseKeeper(rate time.Duration) *HouseKeeper {
	return &HouseKeeper{When: rate}
}
func (s *HouseKeeper) Run(m ...Purifier) {
	if s.running != nil {
		return
	}
	s.running = make(chan struct{})
	s.ticker = time.NewTicker(s.When)
	go func() {
		for {
			select {
			case now := <-s.ticker.C:
				t := now.UnixNano()
				for _, sx := range m {
					sx(t)
				}
			case <-s.running:
				s.ticker.Stop()
				s.ticker = nil
				close(s.running)
				s.running = nil
				return
			}
		}
	}()
}
func (s *HouseKeeper) Stop() {
	if s.running == nil {
		return
	}
	s.running <- struct{}{}
}
