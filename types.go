package hashmap

import (
	"sync"

	"github.com/dolthub/maphash"
)

// high performace concurrent hashmap
type Map[K comparable, V any] struct {
	shards []shard[K, V]
	hasher maphash.Hasher[K]
}

// robin hood hashing
type shard[K comparable, V any] struct {
	sync.RWMutex
	count int32
	size  uint32
	// https://codecapsule.com/2013/11/17/robin-hood-hashing-backward-shift-deletion/
	shift   uint32
	maxDist uint32
	buckets []bucket[K, V]
}

type bucket[K comparable, V any] struct {
	key K
	h   uint64
	// The distance the entry is from its desired position.
	dist uint32
	val  V
}
