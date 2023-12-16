package hashmap

import (
	"math/bits"
)

func (s *shard[K, V]) get(h uint64, k K) (V, bool) {
	var v V
	var dist uint32

	s.RLock()
	defer s.RUnlock()
	for i := uint32(h >> s.shift); ; i++ {
		b := &s.buckets[i]
		if b.h == h && b.key == k {
			return b.val, b.h != 0
		}
		if b.dist < dist { // not found
			return v, false
		}
		dist++
	}
}

// set sets the value for the given key.
// return true if the key already exists.
func (s *shard[K, V]) set(h uint64, k K, v V) bool {
	s.Lock()
	defer s.Unlock()
	return s.set0(h, k, v)
}

func (s *shard[K, V]) delete(h uint64, k K) {
	var dist uint32

	s.Lock()
	defer s.Unlock()
	for i := uint32(h >> s.shift); ; i++ {
		b := &s.buckets[i]
		// found, shift the following buckets backwards
		// util the next bucket is empty or has zero distance.
		// note the empty values ara guarded by the zero distance.
		if b.h == h && b.key == k {
			for j := i + 1; ; j++ {
				t := &s.buckets[j]
				if t.dist == 0 {
					s.count -= 1
					// mark h as empty for delete
					*b = bucket[K, V]{}
					return
				}
				b.h = t.h
				b.key = t.key
				b.val = t.val
				b.dist = t.dist - 1
				b = t
			}
		}
		if dist > b.dist { // not found
			return
		}
		dist++
	}
}

func (s *shard[K, V]) Len() int {
	s.RLock()
	defer s.RUnlock()
	return int(s.count)
}

func (s *shard[K, V]) set0(h uint64, k K, v V) bool {
	maybeExists := true
	n := bucket[K, V]{h: h, key: k, val: v, dist: 0}
	for i := uint32(h >> s.shift); ; i++ {
		b := &s.buckets[i]
		if maybeExists && b.h == h && b.key == k { // exists, update
			b.h = n.h
			b.val = n.val
			return true
		}
		if b.h == 0 { // empty bucket, insert here
			s.count += 1
			*b = n
			return false
		}
		if b.dist < n.dist {
			n, *b = *b, n
			maybeExists = false
		}
		// far away, swap and keep searching
		n.dist++
		// rehash if the distance is too big
		if n.dist == s.maxDist {
			s.rehash(2 * s.size)
			i = uint32(n.h>>s.shift) - 1
			n.dist = 0
			maybeExists = false
		}
	}
}

func (s *shard[K, V]) rehash(size uint32) {
	oldBuckets := s.buckets
	s.count = 0
	s.size = size
	s.shift = uint32(64 - bits.Len32(s.size-1))
	s.maxDist = maxDistForSize(size)
	s.buckets = make([]bucket[K, V], size+s.maxDist)
	for i := range oldBuckets {
		b := &oldBuckets[i]
		if b.h != 0 {
			s.set0(b.h, b.key, b.val)
		}
	}
}

func maxDistForSize(size uint32) uint32 {
	desired := uint32(bits.Len32(size))
	if desired < 4 {
		desired = 4
	}
	return desired
}
