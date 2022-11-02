package benchmarks

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/alphadose/haxmap"
	"github.com/cornelk/hashmap"
	"github.com/orcaman/concurrent-map/v2"
	"github.com/puzpuzpuz/xsync/v2"
	"github.com/smallnest/safemap"
)

// Ref: https://github.com/puzpuzpuz/xsync/blob/main/map_test.go
const (
	// number of entries to use in benchmarks
	benchmarkNumEntries = 500_000
	// key prefix used in benchmarks
	benchmarkKeyPrefix = "what_a_looooooooooooooooooooooong_key_prefix_"
)

var benchmarkCases = []struct {
	name           string
	readPercentage int
}{
	{"100%-reads", 100}, // 100% loads,    0% stores,    0% deletes
	{"99%-reads", 99},   //  99% loads,  0.5% stores,  0.5% deletes
	{"95%-reads", 95},   //  95% loads,  2.5% stores,  2.5% deletes
	{"90%-reads", 90},   //  90% loads,    5% stores,    5% deletes
	{"75%-reads", 75},   //  75% loads, 12.5% stores, 12.5% deletes
	{"50%-reads", 50},   //  50% loads,   25% stores,   25% deletes
	{"0%-reads", 0},     //   0% loads,   50% stores,   50% deletes
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

func BenchmarkMap_Xsync_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := xsync.NewMapOf[int]()
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

func BenchmarkMap_CMAP_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := cmap.New[int]()
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

func BenchmarkMap_SafeMap_NoWarmUp(b *testing.B) {
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

func BenchmarkMap_HaxMap_NoWarmUp(b *testing.B) {
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

func BenchmarkMap_HashMap_NoWarmUp(b *testing.B) {
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

func BenchmarkMap_Standard_NoWarmUp(b *testing.B) {
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

func BenchmarkMap_Xsync_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := xsync.NewMapOf[int]()
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

func BenchmarkMap_CMAP_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := cmap.New[int]()
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

func BenchmarkMap_SafeMap_WarmUp(b *testing.B) {
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

func BenchmarkMap_HaxMap_WarmUp(b *testing.B) {
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

// func BenchmarkMap_HashMap_WarmUp(b *testing.B) {
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
func BenchmarkMap_Standard_WarmUp(b *testing.B) {
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

func BenchmarkMap_Xsync_Range(b *testing.B) {
	m := xsync.NewMapOf[int]()
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

func BenchmarkMap_CMAP_Range(b *testing.B) {
	m := cmap.New[int]()
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

func BenchmarkMap_SafeMap_Range(b *testing.B) {
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

func BenchmarkMap_HaxMap_Range(b *testing.B) {
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

// func BenchmarkMap_HashMap_Range(b *testing.B) {
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

func BenchmarkMap_Standard_Range(b *testing.B) {
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

// go test -bench=^BenchmarkMap -benchmem
// goos: linux
// goarch: amd64
// pkg: bench
// cpu: Intel(R) Xeon(R) Silver 4314 CPU @ 2.40GHz
// BenchmarkMap_Xsync_NoWarmUp/99%-reads-64                110687433               12.49 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/95%-reads-64                88322406                17.27 ns/op            1 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/90%-reads-64                60254306                19.02 ns/op            3 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/75%-reads-64                47258060                24.51 ns/op            7 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/50%-reads-64                37846016                29.71 ns/op           14 B/op          1 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/0%-reads-64                 21427845                53.49 ns/op           29 B/op          2 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/99%-reads-64                 12194355                88.52 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/95%-reads-64                 10678164               103.7 ns/op             1 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/90%-reads-64                  9453764               115.5 ns/op             2 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/75%-reads-64                  7884636               147.0 ns/op             4 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/50%-reads-64                  5907775               169.7 ns/op             5 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/0%-reads-64                   8712856               136.8 ns/op             3 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/99%-reads-64              19244058                54.46 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/95%-reads-64              15166338                66.34 ns/op            1 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/90%-reads-64              13815511                79.05 ns/op            2 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/75%-reads-64              11180854                97.32 ns/op            2 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/50%-reads-64               9022095               116.0 ns/op             3 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/0%-reads-64               12543771                94.35 ns/op            2 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/99%-reads-64               49126554               322.2 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/95%-reads-64               11755160              1698 ns/op               4 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/90%-reads-64                2689261              1555 ns/op               6 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/99%-reads-64              59778726                70.58 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/95%-reads-64              28713607               503.3 ns/op             1 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/90%-reads-64               9664908               883.1 ns/op             3 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/99%-reads-64              1000000              1256 ns/op              34 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/95%-reads-64              1000000              1802 ns/op              49 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/90%-reads-64               757347              2299 ns/op              52 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/75%-reads-64               492739              2053 ns/op              61 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/50%-reads-64               581881              2393 ns/op              67 B/op          1 allocs/op
// BenchmarkMap_Standard_NoWarmUp/0%-reads-64                450807              2734 ns/op              74 B/op          2 allocs/op
// BenchmarkMap_Xsync_WarmUp/100%-reads-64                 72828616                15.13 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/99%-reads-64                  83636658                12.63 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/95%-reads-64                  109142733               11.13 ns/op            1 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/90%-reads-64                  86160418                11.95 ns/op            2 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/75%-reads-64                  61728336                18.89 ns/op            7 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/50%-reads-64                  40054539                29.05 ns/op           14 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/0%-reads-64                   22612006                51.49 ns/op           28 B/op          1 allocs/op
// BenchmarkMap_CMAP_WarmUp/100%-reads-64                  19063858                56.10 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/99%-reads-64                   12072297                92.42 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/95%-reads-64                   10185296               109.3 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/90%-reads-64                    8919643               125.0 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/75%-reads-64                    9807985               155.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/50%-reads-64                    5923474               192.6 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/0%-reads-64                     7987324               140.3 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/100%-reads-64               28904026                41.24 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/99%-reads-64                17206824                60.07 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/95%-reads-64                23110501                74.46 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/90%-reads-64                13308271                80.85 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/75%-reads-64                11731944                98.30 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/50%-reads-64                 9324200               122.5 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/0%-reads-64                 16906158                91.44 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/100%-reads-64                84342962                14.05 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/99%-reads-64                   204526              5227 ns/op               2 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/95%-reads-64                    39470             32335 ns/op              11 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/100%-reads-64                491080              2456 ns/op               0 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/99%-reads-64                 500222              2310 ns/op               0 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/95%-reads-64                 565968              2648 ns/op              52 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/90%-reads-64                 485007              2373 ns/op               2 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/75%-reads-64                 506257              2420 ns/op               5 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/50%-reads-64                 458026              2452 ns/op              11 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/0%-reads-64                  418957              2582 ns/op              22 B/op          1 allocs/op
// BenchmarkMap_Xsync_Range-64                                 1690            728412 ns/op               3 B/op          0 allocs/op
// BenchmarkMap_CMAP_Range-64                                    14          94184852 ns/op        24067662 B/op        208 allocs/op
// BenchmarkMap_SafeMap_Range-64                                 34          71351911 ns/op        24262000 B/op        412 allocs/op
// BenchmarkMap_HaxMap_Range-64                                1179            853686 ns/op               4 B/op          0 allocs/op
// BenchmarkMap_Standard_Range-64                               876           1328893 ns/op               6 B/op          0 allocs/op

// go test -bench=^BenchmarkMap -benchmem
// goos: linux
// goarch: amd64
// pkg: bench
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkMap_Xsync_NoWarmUp/99%-reads-4                 22961880                64.04 ns/op            0 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/95%-reads-4                 17553278                82.95 ns/op            2 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/90%-reads-4                 15808310                98.18 ns/op            4 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/75%-reads-4                 10745935               118.0 ns/op             9 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/50%-reads-4                  8660215               133.6 ns/op            16 B/op          1 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/0%-reads-4                   7122831               168.6 ns/op            31 B/op          2 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/99%-reads-4                  12115374               104.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/95%-reads-4                   8256139               149.6 ns/op             1 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/90%-reads-4                   6976657               180.0 ns/op             2 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/75%-reads-4                   5394254               237.0 ns/op             5 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/50%-reads-4                   4419870               268.8 ns/op             7 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/0%-reads-4                    6944652               163.9 ns/op             4 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/99%-reads-4                6722534               179.8 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/95%-reads-4                4732837               264.9 ns/op             1 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/90%-reads-4                4137775               302.7 ns/op             3 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/75%-reads-4                3578094               350.6 ns/op             4 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/50%-reads-4                3206191               383.2 ns/op             9 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/0%-reads-4                 3705342               270.3 ns/op             8 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/99%-reads-4                13914631              1242 ns/op               0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/95%-reads-4                 1000000              1145 ns/op               3 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/90%-reads-4                 1000000              7525 ns/op               7 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/99%-reads-4               19719538               472.0 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/95%-reads-4                3542832              1834 ns/op               1 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/90%-reads-4                1000000              1662 ns/op               4 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/99%-reads-4               3057364               476.0 ns/op            45 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/95%-reads-4               2360776               520.1 ns/op            46 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/90%-reads-4               2017683               576.9 ns/op            49 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/75%-reads-4               1885638               646.5 ns/op            50 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/50%-reads-4               1654629               696.5 ns/op            55 B/op          1 allocs/op
// BenchmarkMap_Standard_NoWarmUp/0%-reads-4                1433965               762.5 ns/op            64 B/op          2 allocs/op
// BenchmarkMap_Xsync_WarmUp/100%-reads-4                   8047791               149.1 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/99%-reads-4                    8190386               141.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/95%-reads-4                    8321588               139.7 ns/op             1 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/90%-reads-4                    8161870               140.5 ns/op             2 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/75%-reads-4                    7227441               145.5 ns/op             7 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/50%-reads-4                    7019108               147.6 ns/op            14 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/0%-reads-4                     6953701               170.0 ns/op            27 B/op          1 allocs/op
// BenchmarkMap_CMAP_WarmUp/100%-reads-4                    7796668               147.2 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/99%-reads-4                     7463191               160.2 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/95%-reads-4                     5609403               205.6 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/90%-reads-4                     4712932               243.8 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/75%-reads-4                     3947751               280.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/50%-reads-4                     3633550               292.4 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/0%-reads-4                      6414294               168.6 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/100%-reads-4                 6787693               173.7 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/99%-reads-4                  4552660               254.3 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/95%-reads-4                  3130820               369.5 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/90%-reads-4                  2840918               417.8 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/75%-reads-4                  2572873               440.4 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/50%-reads-4                  2542430               451.3 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/0%-reads-4                   3922452               269.6 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/100%-reads-4                  8160154               147.2 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/99%-reads-4                     21225             49576 ns/op               1 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/95%-reads-4                      4381            270283 ns/op               7 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/100%-reads-4                4511062               226.2 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/99%-reads-4                 3677433               289.8 ns/op             8 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/95%-reads-4                 3485506               308.4 ns/op             9 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/90%-reads-4                 3477531               308.3 ns/op            10 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/75%-reads-4                 3209703               348.1 ns/op            14 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/50%-reads-4                 1684626               671.1 ns/op            37 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/0%-reads-4                  1497439               687.7 ns/op            33 B/op          1 allocs/op
// BenchmarkMap_Xsync_Range-4                                   169           6955477 ns/op               2 B/op          0 allocs/op
// BenchmarkMap_CMAP_Range-4                                     34          32698984 ns/op        24067286 B/op        199 allocs/op
// BenchmarkMap_SafeMap_Range-4                                  25          40408270 ns/op        24020066 B/op         31 allocs/op
// BenchmarkMap_HaxMap_Range-4                                  123           8804928 ns/op               2 B/op          0 allocs/op
// BenchmarkMap_Standard_Range-4                                114           9769472 ns/op               3 B/op          0 allocs/op

// go test -bench=^BenchmarkMap -benchmem
// goos: linux
// goarch: amd64
// pkg: bench
// cpu: AMD EPYC 7K62 48-Core Processor
// BenchmarkMap_Xsync_NoWarmUp/99%-reads-4                  5322466               242.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/95%-reads-4                  4893560               296.5 ns/op             3 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/90%-reads-4                  4304379               329.3 ns/op             5 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/75%-reads-4                  3536888               433.2 ns/op            12 B/op          0 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/50%-reads-4                  2332996               513.3 ns/op            22 B/op          1 allocs/op
// BenchmarkMap_Xsync_NoWarmUp/0%-reads-4                   1909485               855.9 ns/op            39 B/op          2 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/99%-reads-4                   3784177               329.5 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/95%-reads-4                   3658107               349.7 ns/op             2 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/90%-reads-4                   3534452               370.7 ns/op             4 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/75%-reads-4                   3219830               394.4 ns/op             5 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/50%-reads-4                   2899605               438.4 ns/op            11 B/op          0 allocs/op
// BenchmarkMap_CMAP_NoWarmUp/0%-reads-4                    2616900               452.6 ns/op            12 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/99%-reads-4                4374471               289.3 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/95%-reads-4                4137565               318.3 ns/op             1 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/90%-reads-4                3833136               335.2 ns/op             4 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/75%-reads-4                3554889               363.9 ns/op             4 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/50%-reads-4                3212438               394.8 ns/op             9 B/op          0 allocs/op
// BenchmarkMap_SafeMap_NoWarmUp/0%-reads-4                 2651151               411.2 ns/op            12 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/99%-reads-4                 3608038               786.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/95%-reads-4                 1000000              4521 ns/op               3 B/op          0 allocs/op
// BenchmarkMap_HaxMap_NoWarmUp/90%-reads-4                 1000000             20939 ns/op               6 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/99%-reads-4                4312915               429.4 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/95%-reads-4                1000000              1323 ns/op               2 B/op          0 allocs/op
// BenchmarkMap_HashMap_NoWarmUp/90%-reads-4                1000000              5824 ns/op               4 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/99%-reads-4               2943702               440.0 ns/op            46 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/95%-reads-4               2427459               567.7 ns/op            48 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/90%-reads-4               2124974               626.8 ns/op            49 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/75%-reads-4               1827462               695.9 ns/op            49 B/op          0 allocs/op
// BenchmarkMap_Standard_NoWarmUp/50%-reads-4               1648512               858.8 ns/op            61 B/op          1 allocs/op
// BenchmarkMap_Standard_NoWarmUp/0%-reads-4                1375279               780.8 ns/op            58 B/op          2 allocs/op
// BenchmarkMap_Xsync_WarmUp/100%-reads-4                   2032300               595.7 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/99%-reads-4                    1994565               614.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/95%-reads-4                    1994385               598.5 ns/op             1 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/90%-reads-4                    2020077               590.4 ns/op             2 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/75%-reads-4                    1965814               588.9 ns/op             7 B/op          0 allocs/op
// BenchmarkMap_Xsync_WarmUp/50%-reads-4                    1892293               674.5 ns/op            14 B/op          1 allocs/op
// BenchmarkMap_Xsync_WarmUp/0%-reads-4                     1542258               765.0 ns/op            28 B/op          2 allocs/op
// BenchmarkMap_CMAP_WarmUp/100%-reads-4                    2186379               544.8 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/99%-reads-4                     2197336               549.1 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/95%-reads-4                     2179720               540.2 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/90%-reads-4                     2211510               538.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/75%-reads-4                     2228772               522.7 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/50%-reads-4                     2215503               504.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_CMAP_WarmUp/0%-reads-4                      2305758               486.7 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/100%-reads-4                 2296856               516.5 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/99%-reads-4                  2203132               519.2 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/95%-reads-4                  2351454               512.8 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/90%-reads-4                  2318145               505.9 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/75%-reads-4                  2251226               489.7 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/50%-reads-4                  2404495               472.3 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_SafeMap_WarmUp/0%-reads-4                   2540269               444.1 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/100%-reads-4                  1894240               611.1 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/99%-reads-4                      4152            321617 ns/op               5 B/op          0 allocs/op
// BenchmarkMap_HaxMap_WarmUp/95%-reads-4                       633           1712095 ns/op              37 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/100%-reads-4                1803691               655.0 ns/op             0 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/99%-reads-4                 1396999               801.8 ns/op            20 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/95%-reads-4                 1403974               793.6 ns/op            21 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/90%-reads-4                 1396458              1168 ns/op              22 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/75%-reads-4                 1398770               795.7 ns/op            25 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/50%-reads-4                 1420707               806.4 ns/op            43 B/op          0 allocs/op
// BenchmarkMap_Standard_WarmUp/0%-reads-4                  1435394               776.3 ns/op            34 B/op          1 allocs/op
// BenchmarkMap_Xsync_Range-4                                    38          28972117 ns/op               9 B/op          0 allocs/op
// BenchmarkMap_CMAP_Range-4                                     22          76655087 ns/op        24068915 B/op        203 allocs/op
// BenchmarkMap_SafeMap_Range-4                                  21          73158401 ns/op        24020057 B/op         31 allocs/op
// BenchmarkMap_HaxMap_Range-4                                   21          49680062 ns/op              17 B/op          0 allocs/op
// BenchmarkMap_Standard_Range-4                                 26          43018142 ns/op              14 B/op          0 allocs/op
