//go:build go1.18
// +build go1.18

package benchmarks

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/alphadose/haxmap"
	"github.com/cornelk/hashmap"
	"github.com/fufuok/cache"
	"github.com/smallnest/safemap"

	"github.com/fufuok/cmap"
)

// Ref: https://github.com/puzpuzpuz/cache/blob/main/map_test.go
const (
	// number of entries to use in benchmarks
	benchmarkNumEntries = 200_000
	// key prefix used in benchmarks
	benchmarkKeyPrefix = "what_a_looooooooooooooooooooooong_key_prefix_"
)

var benchmarkCases = []struct {
	name           string
	readPercentage int
}{
	{"reads=99%", 99}, //  99% loads,  0.5% stores,  0.5% deletes
	{"reads=90%", 90}, //  90% loads,    5% stores,    5% deletes
	{"reads=75%", 75}, //  75% loads, 12.5% stores, 12.5% deletes
}

var (
	benchmarkKeys        []string
	benchmarkIntegerKeys []int
)

func init() {
	benchmarkKeys = make([]string, benchmarkNumEntries)
	benchmarkIntegerKeys = make([]int, benchmarkNumEntries)
	for i := 0; i < benchmarkNumEntries; i++ {
		benchmarkKeys[i] = benchmarkKeyPrefix + strconv.Itoa(i)
		benchmarkIntegerKeys[i] = i
	}
}

func BenchmarkString_CMAP_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := cmap.NewOf[string, int]()
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Get(k)
			}, func(k string, v int) {
				m.Set(k, v)
			}, func(k string) {
				m.Remove(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_SafeMap_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := safemap.New[string, int]()
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Get(k)
			}, func(k string, v int) {
				m.Set(k, v)
			}, func(k string) {
				m.Remove(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_Cache_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := cache.NewMapOf[int]()
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Load(k)
			}, func(k string, v int) {
				m.Store(k, v)
			}, func(k string) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_HaxMap_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 || bc.readPercentage < 90 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := haxmap.New[string, int]()
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Get(k)
			}, func(k string, v int) {
				m.Set(k, v)
			}, func(k string) {
				m.Del(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_HashMap_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 || bc.readPercentage < 90 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := hashmap.New[string, int]()
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Get(k)
			}, func(k string, v int) {
				m.Set(k, v)
			}, func(k string) {
				m.Del(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_Standard_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			var m sync.Map
			benchmarkMap(b, func(k string) (int, bool) {
				v, ok := m.Load(k)
				n := 0
				if v != nil {
					n = v.(int)
				}
				return n, ok
			}, func(k string, v int) {
				m.Store(k, v)
			}, func(k string) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_CMAP_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := cmap.NewOf[string, int]()
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Set(benchmarkKeyPrefix+strconv.Itoa(i), i)
			}
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Get(k)
			}, func(k string, v int) {
				m.Set(k, v)
			}, func(k string) {
				m.Remove(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_SafeMap_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := safemap.New[string, int]()
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Set(benchmarkKeyPrefix+strconv.Itoa(i), i)
			}
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Get(k)
			}, func(k string, v int) {
				m.Set(k, v)
			}, func(k string) {
				m.Remove(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_Cache_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := cache.NewMapOf[int]()
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Store(benchmarkKeyPrefix+strconv.Itoa(i), i)
			}
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Load(k)
			}, func(k string, v int) {
				m.Store(k, v)
			}, func(k string) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkString_HaxMap_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage < 95 {
			// skip, has slow write performance for write heavy uses.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := haxmap.New[string, int]()
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Set(benchmarkKeyPrefix+strconv.Itoa(i), i)
			}
			benchmarkMap(b, func(k string) (int, bool) {
				return m.Get(k)
			}, func(k string, v int) {
				m.Set(k, v)
			}, func(k string) {
				m.Del(k)
			}, bc.readPercentage)
		})
	}
}

// func BenchmarkString_HashMap_WarmUp(b *testing.B) {
// 	for _, bc := range benchmarkCases {
// 		if bc.readPercentage < 95 {
// 			// skip, has slow write performance for write heavy uses.
// 			continue
// 		}
// 		b.Run(bc.name, func(b *testing.B) {
// 			m := hashmap.New[string, int]()
// 			for i := 0; i < benchmarkNumEntries; i++ {
// 				m.Set(benchmarkKeyPrefix+strconv.Itoa(i), i)
// 			}
// 			benchmarkMap(b, func(k string) (int, bool) {
// 				return m.Get(k)
// 			}, func(k string, v int) {
// 				m.Set(k, v)
// 			}, func(k string) {
// 				m.Del(k)
// 			}, bc.readPercentage)
// 		})
// 	}
// }

// This is a nice scenario for sync.Map since a lot of updates
// will hit the readOnly part of the map.
func BenchmarkString_Standard_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			var m sync.Map
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Store(benchmarkKeyPrefix+strconv.Itoa(i), i)
			}
			benchmarkMap(b, func(k string) (int, bool) {
				v, ok := m.Load(k)
				n := 0
				if v != nil {
					n = v.(int)
				}
				return n, ok
			}, func(k string, v int) {
				m.Store(k, v)
			}, func(k string) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

func benchmarkMap(
	b *testing.B,
	loadFn func(k string) (int, bool),
	storeFn func(k string, v int),
	deleteFn func(k string),
	readPercentage int) {

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// convert percent to permille to support 99% case
		storeThreshold := 10 * readPercentage
		deleteThreshold := 10*readPercentage + ((1000 - 10*readPercentage) / 2)
		for pb.Next() {
			op := r.Intn(1000)
			i := r.Intn(benchmarkNumEntries)
			if op >= deleteThreshold {
				deleteFn(benchmarkKeys[i])
			} else if op >= storeThreshold {
				storeFn(benchmarkKeys[i], i)
			} else {
				loadFn(benchmarkKeys[i])
			}
		}
	})
}

func BenchmarkString_CMAP_Range(b *testing.B) {
	m := cmap.NewOf[string, int]()
	for i := 0; i < benchmarkNumEntries; i++ {
		m.Set(benchmarkKeys[i], i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			for range m.IterBuffered() {
				foo++
			}
			_ = foo
		}
	})
}

func BenchmarkString_SafeMap_Range(b *testing.B) {
	m := safemap.New[string, int]()
	for i := 0; i < benchmarkNumEntries; i++ {
		m.Set(benchmarkKeys[i], i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			for range m.IterBuffered() {
				foo++
			}
			_ = foo
		}
	})
}

func BenchmarkString_Cache_Range(b *testing.B) {
	m := cache.NewMapOf[int]()
	for i := 0; i < benchmarkNumEntries; i++ {
		m.Store(benchmarkKeys[i], i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			m.Range(func(key string, value int) bool {
				foo++
				return true
			})
			_ = foo
		}
	})
}

func BenchmarkString_HaxMap_Range(b *testing.B) {
	m := haxmap.New[string, int]()
	for i := 0; i < benchmarkNumEntries; i++ {
		m.Set(benchmarkKeys[i], i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			m.ForEach(func(k string, v int) bool {
				foo++
				return true
			})
			_ = foo
		}
	})
}

// func BenchmarkString_HashMap_Range(b *testing.B) {
// 	m := hashmap.New[string, int]()
// 	for i := 0; i < benchmarkNumEntries; i++ {
// 		m.Set(benchmarkKeys[i], i)
// 	}
// 	b.ResetTimer()
// 	b.RunParallel(func(pb *testing.PB) {
// 		foo := 0
// 		for pb.Next() {
// 			m.Range(func(k string, v int) bool {
// 				foo++
// 				return true
// 			})
// 			_ = foo
// 		}
// 	})
// }

func BenchmarkString_Standard_Range(b *testing.B) {
	var m sync.Map
	for i := 0; i < benchmarkNumEntries; i++ {
		m.Store(benchmarkKeys[i], i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			m.Range(func(key any, value any) bool {
				foo++
				return true
			})
			_ = foo
		}
	})
}
