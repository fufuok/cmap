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
	m := cmap.NewOf[int, int]()
	m.Set(1, 1)
	v := m.Upsert(1, 2, func(exist bool, valueInMap int, newValue int) int {
		if valueInMap == 1 {
			return newValue + 1
		}
		return 0
	})
	fmt.Println(v)
	fmt.Println(m.Get(1))
	m.SetIfAbsent(2, 42)
	m.Remove(1)
	fmt.Println(m.Count())
	for item := range m.IterBuffered() {
		fmt.Println(item)
	}
	m.Clear()

	// Output:
	// 3
	// 3 true
	// 1
	// {2 42}
}
```

### custom type

```go
package main

import (
	"fmt"

	"github.com/fufuok/cmap"
)

type Person struct {
	name string
	age  int16
}

func main() {
	hasher := func(p Person) uint64 {
		return uint64(fnv32(p.name))<<32 | uint64(31*p.age)
	}
	m := cmap.NewTypedMapOf[Person, int](hasher)
	m.Set(Person{"ff", 18}, 1)
	v := m.Upsert(Person{"ff", 18}, 2, func(exist bool, valueInMap int, newValue int) int {
		if valueInMap == 1 {
			return newValue + 1
		}
		return 0
	})
	fmt.Println(v)
	fmt.Println(m.Get(Person{"ff", 18}))
	m.SetIfAbsent(Person{"uu", 20}, 42)
	m.Remove(Person{"ff", 18})
	fmt.Println(m.Count())
	for item := range m.IterBuffered() {
		fmt.Println(item)
	}
	m.Clear()

	// Output:
	// 3
	// 3 true
	// 1
	// {{uu 20} 42}
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
```



## LICENSE

MIT (see [LICENSE](https://github.com/orcaman/concurrent-map/blob/master/LICENSE) file)







*ff*

