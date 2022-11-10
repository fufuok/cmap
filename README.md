# Concurrent Map

sync.Map with better read and write performance (supports any type of the key)

*forked from [orcaman/concurrent-map](https://github.com/orcaman/concurrent-map)*

> For cache, you can choose: [Cache/CacheOf and Map/MapOf](https://github.com/fufuok/cache)

## Changelog

- Unified initialization method: `cmap.NewOf[K, V]()`, Used `xxhash`, thanks.
- Supports both generic and non-generic types.
- Add benchmarks for commonly used Map packages, See: [🤖 Benchmarks](benchmarks)

## ⚙️ Installation

```go
go get -u github.com/fufuok/cmap
```

## ⚡️ Quickstart

- [example_map_test.go](example_map_test.go)
- [example_mapof_test.go](example_mapof_test.go)

### non-generic

```go
package main

import (
	"fmt"

	"github.com/fufuok/cmap"
)

func main() {
	m := cmap.New()
	m.Set("A", 1)
	v := m.Upsert("A", 2, func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
		if valueInMap == 1 {
			return newValue.(int) + 1
		}
		return 0
	})
	fmt.Println(v)
	fmt.Println(m.Get("A"))
	m.SetIfAbsent("B", 42)
	m.Remove("A")
	fmt.Println(m.Count())
	for item := range m.IterBuffered() {
		fmt.Println(item)
	}
	m.Clear()

	// Output:
	// 3
	// 3 true
	// 1
	// {B 42}
}
```

### generic

```go
package main

import (
	"fmt"

	"github.com/fufuok/cmap"
)

func main() {
	m := cmap.NewOf[string, int]()
	m.Set("A", 1)
	v := m.Upsert("A", 2, func(exist bool, valueInMap int, newValue int) int {
		if valueInMap == 1 {
			return newValue + 1
		}
		return 0
	})
	fmt.Println(v)
	fmt.Println(m.Get("A"))
	m.SetIfAbsent("B", 42)
	m.Remove("A")
	fmt.Println(m.Count())
	for item := range m.IterBuffered() {
		fmt.Println(item)
	}
	m.Clear()

	// output
	// 3
	// 3 true
	// 1
	// {B 42}
}
```

## LICENSE

MIT (see [LICENSE](https://github.com/orcaman/concurrent-map/blob/master/LICENSE) file)







*ff*

