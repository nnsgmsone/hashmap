# `hashmap`
[![GoDoc](https://godoc.org/github.com/nnsgmsone/hashmap?status.svg)](https://godoc.org/github.com/nnsgmsone/hashmap)

A high performance concurrent map for Go, all functions are thread safe, support for generics.

## Getting Started

```go
// create a new hashmap with key type int and value type int
m := New[int, int](0)

// set the value for key
m.Set(1, 1)
m.Set(1, 2)
m.Set(2, 1)

// get count of key-value for the hashmap
n = m.Len()

// get value by key
v, ok = m.Get(1)
v, ok = m.Get(2)

// delete the value for a key.
m.Del(1)
```

## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/nnsgmsone/hashmap
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkHashMap/hashmap-set-parallel100-20         	       1	1586248494 ns/op	  116872 B/op	     	561 allocs/op
BenchmarkHashMap/hashmap-set-parallel1000-20        	       1	11420826547 ns/op	  836576 B/op	    	4695 allocs/op
BenchmarkHashMap/hashmap-set-get-parallel100-20     	       1	3255109261 ns/op	   47776 B/op	     	243 allocs/op
BenchmarkHashMap/hashmap-set-get-parallel1000-20    	       1	17511975838 ns/op	  620832 B/op	    	2966 allocs/op
BenchmarkHashMap/syncmap-set-parallel100-20         	       1	19082770010 ns/op	2400015608 B/op		199976074 allocs/op
BenchmarkHashMap/syncmap-set-parallel1000-20        	       1	204721266296 ns/op	29956965176 B/op	2743807276 allocs/op
BenchmarkHashMap/syncmap-set-get-parallel100-20     	       1	1636063296 ns/op	2400001528 B/op		199975746 allocs/op
BenchmarkHashMap/syncmap-set-get-parallel1000-20    	       1	19790039965 ns/op	32929943928 B/op	3115775177 allocs/op
BenchmarkHashMap/mutexmap-set-parallel100-20        	       1	15495332887 ns/op	799877384 B/op		99974948 allocs/op
BenchmarkHashMap/mutexmap-set-parallel1000-20       	       1	178083314115 ns/op	13960544056 B/op	1743848208 allocs/op
BenchmarkHashMap/mutexmap-set-get-parallel100-20    	       1	9547300918 ns/op	799986024 B/op		99976220 allocs/op
BenchmarkHashMap/mutexmap-set-get-parallel1000-20   	       1	143467062246 ns/op	16978766920 B/op	2116287321 allocs/op
```

## License

`hashmap` source code is available under the MIT [License](/LICENSE).
