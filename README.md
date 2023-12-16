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

## License

`hashmap` source code is available under the MIT [License](/LICENSE).
