package hashmap

import (
	"sync"
	"testing"
)

const (
	TestSet = iota
	TestGet
)

type testMap interface {
	Store(key, value any)
	Load(key any) (any, bool)
}

type testGoMap struct {
	sync.RWMutex
	mp map[any]any
}

func (m *testGoMap) Store(key, value any) {
	m.Lock()
	defer m.Unlock()
	m.mp[key] = value
}

func (m *testGoMap) Load(key any) (any, bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.mp[key]
	return v, ok
}

func bencmarkMap(b *testing.B, mp testMap, op int, parallisim int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup

		for j := 0; j < parallisim; j++ {
			wg.Add(1)
			go func() {
				switch op {
				case TestSet:
					for k := 0; k < 1000000; k++ {
						mp.Store(k%parallisim, k)
					}
				case TestGet:
					for k := 0; k < 1000000; k++ {
						mp.Store(k%parallisim, k)
						if k%2 == 0 {
							mp.Load(k % parallisim)
						}
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func bencmarkHashMap(b *testing.B, mp *Map[int, int], op int, parallisim int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup

		for j := 0; j < parallisim; j++ {
			wg.Add(1)
			go func() {
				switch op {
				case TestSet:
					for k := 0; k < 1000000; k++ {
						mp.Set(k%parallisim, k)
					}
				case TestGet:
					for k := 0; k < 1000000; k++ {
						mp.Set(k%parallisim, k)
						if k%2 == 0 {
							mp.Get(k % parallisim)
						}
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkHashMap(b *testing.B) {
	b.Run("hashmap-set-parallel100", func(b *testing.B) {
		bencmarkHashMap(b, New[int, int](0), TestSet, 100)
	})
	b.Run("hashmap-set-parallel1000", func(b *testing.B) {
		bencmarkHashMap(b, New[int, int](0), TestSet, 1000)
	})
	b.Run("hashmap-set-get-parallel100", func(b *testing.B) {
		bencmarkHashMap(b, New[int, int](0), TestGet, 100)
	})
	b.Run("hashmap-set-get-parallel1000", func(b *testing.B) {
		bencmarkHashMap(b, New[int, int](0), TestGet, 1000)
	})
	b.Run("syncmap-set-parallel100", func(b *testing.B) {
		bencmarkMap(b, new(sync.Map), TestSet, 100)
	})
	b.Run("syncmap-set-parallel1000", func(b *testing.B) {
		bencmarkMap(b, new(sync.Map), TestSet, 1000)
	})
	b.Run("syncmap-set-get-parallel100", func(b *testing.B) {
		bencmarkMap(b, new(sync.Map), TestGet, 100)
	})
	b.Run("syncmap-set-get-parallel1000", func(b *testing.B) {
		bencmarkMap(b, new(sync.Map), TestGet, 1000)
	})
	b.Run("mutexmap-set-parallel100", func(b *testing.B) {
		bencmarkMap(b, &testGoMap{mp: make(map[any]any)}, TestSet, 100)
	})
	b.Run("mutexmap-set-parallel1000", func(b *testing.B) {
		bencmarkMap(b, &testGoMap{mp: make(map[any]any)}, TestSet, 1000)
	})
	b.Run("mutexmap-set-get-parallel100", func(b *testing.B) {
		bencmarkMap(b, &testGoMap{mp: make(map[any]any)}, TestGet, 100)
	})
	b.Run("mutexmap-set-get-parallel1000", func(b *testing.B) {
		bencmarkMap(b, &testGoMap{mp: make(map[any]any)}, TestGet, 1000)
	})

}
