# fifo

![test](https://github.com/WinPooh32/fifo/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/WinPooh32/fifo.svg)](https://pkg.go.dev/github.com/WinPooh32/fifo)

**fifo** (first in, first out) cache.

## Example

```Go
package main

import (
	"fmt"

	"github.com/WinPooh32/fifo"
)

func main() {
	const capacity = 3

	keys := []string{"a", "b", "c", "d", "e", "f"}
	data := []int{0, 1, 2, 3, 4, 5}

	cache := fifo.New[string, int](capacity)

	for i, key := range keys {
		cache.Set(key, data[i])
	}

	for _, key := range keys {
		fmt.Println(cache.Get(key))
	}
}
```

Output:

```
0 false
0 false
0 false
3 true
4 true
5 true
```

## Benchmarks

No allocs at **Get** and **Set** operations:

```
goos: linux
goarch: amd64
pkg: github.com/WinPooh32/fifo
cpu: AMD Ryzen 7 3700X 8-Core Processor             
BenchmarkCache_Get_1024-16    96886166  3.00  ns/op   0 B/op  0 allocs/op
BenchmarkCache_Set_1024-16    19376492  60.47 ns/op   0 B/op  0 allocs/op
```
