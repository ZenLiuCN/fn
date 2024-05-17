package fn

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestSyncTTLMap(t *testing.T) {
	m := NewSyncTTLMap[int, int](time.Millisecond * 100)

	tar := map[int]int64{}
	w := new(sync.WaitGroup)
	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 1)
			tar[i] = m.StoreWithTTL(i, i, time.Millisecond*250)
		}
		m.Run()
	}()
	w.Wait()
	n := 0
	tk := time.NewTicker(time.Millisecond * 250)
	for {
		select {
		case now := <-tk.C:
			nano := now.UnixNano()
			v, x, ok := m.LoadWithNanos(n)
			if ok {
				assert.Less(t, x, nano, "ddl and now %d", n)
				assert.LessOrEqual(t, tar[n], x, "current and target", n)
				assert.Equal(t, v, n, "value", n)
				t.Logf("exists [%d] %d => %d", n, nano, tar[n])
			} else {
				t.Logf("clean [%d] %d => %d", n, nano, tar[n])
				assert.Less(t, tar[n], nano, "cleaned", n)
			}
			n++
			if n > 100 {
				return
			}

		}
	}
}
