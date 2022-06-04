package bench

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/puzpuzpuz/xsync"

	"github.com/fufuok/cmap"
)

// Ref: https://github.com/puzpuzpuz/xsync/blob/main/map_test.go
const (
	// number of entries to use in benchmarks
	benchmarkNumEntries = 1_000_000
	// key prefix used in benchmarks
	benchmarkKeyPrefix = "what_a_looooooooooooooooooooooong_key_prefix_"
)

var benchmarkCases = []struct {
	name           string
	readPercentage int
}{
	{"100%-reads", 100}, // 100% loads,    0% stores,    0% deletes
	{"99%-reads", 99},   //  99% loads,  0.5% stores,  0.5% deletes
	{"90%-reads", 90},   //  90% loads,    5% stores,    5% deletes
	{"75%-reads", 75},   //  75% loads, 12.5% stores, 12.5% deletes
	{"50%-reads", 50},   //  50% loads,   25% stores,   25% deletes
	{"0%-reads", 0},     //   0% loads,   50% stores,   50% deletes
}

var benchmarkKeys []string

func init() {
	benchmarkKeys = make([]string, benchmarkNumEntries)
	for i := 0; i < benchmarkNumEntries; i++ {
		benchmarkKeys[i] = benchmarkKeyPrefix + strconv.Itoa(i)
	}
}

func BenchmarkMap_NoWarmUp(b *testing.B) {
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

func BenchmarkMapCMAP_NoWarmUp(b *testing.B) {
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

func BenchmarkMapStandard_NoWarmUp(b *testing.B) {
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

func BenchmarkMap_WarmUp(b *testing.B) {
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

func BenchmarkMapCMAP_WarmUp(b *testing.B) {
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

// This is a nice scenario for sync.Map since a lot of updates
// will hit the readOnly part of the map.
func BenchmarkMapStandard_WarmUp(b *testing.B) {
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

func BenchmarkMapRange(b *testing.B) {
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

func BenchmarkMapRangeCMAP(b *testing.B) {
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

func BenchmarkMapRangeStandard(b *testing.B) {
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

// go test -bench=^BenchmarkMap -benchtime=2s
// go version go1.17 linux/amd64
// cpu: Intel(R) Xeon(R) Gold 6161 CPU @ 2.20GHz
// BenchmarkMap_NoWarmUp/99%-reads-8       56163492                52.55 ns/op
// BenchmarkMap_NoWarmUp/90%-reads-8       35315628                73.24 ns/op
// BenchmarkMap_NoWarmUp/75%-reads-8       26161597                81.62 ns/op
// BenchmarkMap_NoWarmUp/50%-reads-8       24746589               106.9 ns/op
// BenchmarkMap_NoWarmUp/0%-reads-8        18289854               141.3 ns/op
// BenchmarkMapCMAP_NoWarmUp/99%-reads-8           28187413                93.63 ns/op
// BenchmarkMapCMAP_NoWarmUp/90%-reads-8           13235156               202.7 ns/op
// BenchmarkMapCMAP_NoWarmUp/75%-reads-8            9603061               262.1 ns/op
// BenchmarkMapCMAP_NoWarmUp/50%-reads-8            8177826               307.7 ns/op
// BenchmarkMapCMAP_NoWarmUp/0%-reads-8            15768786               147.7 ns/op
// BenchmarkMapStandard_NoWarmUp/99%-reads-8        4889781               649.5 ns/op
// BenchmarkMapStandard_NoWarmUp/90%-reads-8        2881819               848.7 ns/op
// BenchmarkMapStandard_NoWarmUp/75%-reads-8        2562492               959.4 ns/op
// BenchmarkMapStandard_NoWarmUp/50%-reads-8        2688487               929.0 ns/op
// BenchmarkMapStandard_NoWarmUp/0%-reads-8         2304576              1045 ns/op
// BenchmarkMap_WarmUp/100%-reads-8                26866794                92.23 ns/op
// BenchmarkMap_WarmUp/99%-reads-8                 26107092                88.13 ns/op
// BenchmarkMap_WarmUp/90%-reads-8                 29198966                80.35 ns/op
// BenchmarkMap_WarmUp/75%-reads-8                 28485249                80.61 ns/op
// BenchmarkMap_WarmUp/50%-reads-8                 27306980                83.98 ns/op
// BenchmarkMap_WarmUp/0%-reads-8                  21844280               113.1 ns/op
// BenchmarkMapCMAP_WarmUp/100%-reads-8            23667356                95.94 ns/op
// BenchmarkMapCMAP_WarmUp/99%-reads-8             17253115               138.5 ns/op
// BenchmarkMapCMAP_WarmUp/90%-reads-8              9708538               231.9 ns/op
// BenchmarkMapCMAP_WarmUp/75%-reads-8              8166346               280.1 ns/op
// BenchmarkMapCMAP_WarmUp/50%-reads-8              7348316               315.0 ns/op
// BenchmarkMapCMAP_WarmUp/0%-reads-8              15754057               129.3 ns/op
// BenchmarkMapStandard_WarmUp/100%-reads-8        10278343               197.9 ns/op
// BenchmarkMapStandard_WarmUp/99%-reads-8          7406224               272.3 ns/op
// BenchmarkMapStandard_WarmUp/90%-reads-8          7389327               273.5 ns/op
// BenchmarkMapStandard_WarmUp/75%-reads-8          4945608               436.3 ns/op
// BenchmarkMapStandard_WarmUp/50%-reads-8          3071937               792.5 ns/op
// BenchmarkMapStandard_WarmUp/0%-reads-8           2321718               903.5 ns/op
// BenchmarkMapRange-8                                  134          19187676 ns/op
// BenchmarkMapRangeCMAP-8                               32          63195121 ns/op
// BenchmarkMapRangeStandard-8                          169          14674270 ns/op
