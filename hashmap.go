package hashmap

import (
	"runtime"

	"github.com/dolthub/maphash"
)

func New[K comparable, V any](cap int) *Map[K, V] {
	var mp Map[K, V]

	// 4 is a good number for most cases.
	m := 4 * runtime.GOMAXPROCS(0)
	if cap /= m; cap <= 0 {
		cap = 1
	}
	mp.hasher = maphash.NewHasher[K]()
	mp.shards = make([]shard[K, V], m)
	for i := range mp.shards {
		mp.shards[i].rehash(uint32(cap))
	}
	return &mp
}

func (mp *Map[K, V]) Len() int {
	var count int

	for i := range mp.shards {
		count += mp.shards[i].Len()
	}
	return count
}

func (mp *Map[K, V]) Set(k K, v V) {
	var s *shard[K, V]

	h := mp.hasher.Hash(k)
	s = &mp.shards[h%uint64(len(mp.shards))]
	s.set(h, k, v)
}

func (mp *Map[K, V]) Get(k K) (V, bool) {
	var s *shard[K, V]

	h := mp.hasher.Hash(k)
	s = &mp.shards[h%uint64(len(mp.shards))]
	return s.get(h, k)
}

func (mp *Map[K, V]) Delete(k K) {
	var s *shard[K, V]

	h := mp.hasher.Hash(k)
	s = &mp.shards[h%uint64(len(mp.shards))]
	s.delete(h, k)
}
