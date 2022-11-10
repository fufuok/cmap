//go:build go1.18
// +build go1.18

package benchmarks

import (
	"sync"
	"testing"
	"time"
	_ "unsafe"

	"github.com/fufuok/cache"
	"github.com/smallnest/safemap"

	"github.com/fufuok/cmap"
)

func runParallel(b *testing.B, benchFn func(pb *testing.PB)) {
	b.ResetTimer()
	start := time.Now()
	b.RunParallel(benchFn)
	opsPerSec := float64(b.N) / time.Since(start).Seconds()
	b.ReportMetric(opsPerSec, "ops/s")
}

func BenchmarkInteger_CMAP_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := cmap.NewOf[int, int]()
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Set(i, i)
			}
			b.ResetTimer()
			benchmarkMapOfIntegerKeys(b, func(k int) (int, bool) {
				return m.Get(k)
			}, func(k int, v int) {
				m.Set(k, v)
			}, func(k int) {
				m.Remove(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkInteger_SafeMap_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := safemap.New[int, int]()
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Set(i, i)
			}
			b.ResetTimer()
			benchmarkMapOfIntegerKeys(b, func(k int) (int, bool) {
				return m.Get(k)
			}, func(k int, v int) {
				m.Set(k, v)
			}, func(k int) {
				m.Remove(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkInteger_Cache_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := cache.NewIntegerMapOfPresized[int, int](benchmarkNumEntries)
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Store(i, i)
			}
			b.ResetTimer()
			benchmarkMapOfIntegerKeys(b, func(k int) (int, bool) {
				return m.Load(k)
			}, func(k int, v int) {
				m.Store(k, v)
			}, func(k int) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

// This is a nice scenario for sync.Map since a lot of updates
// will hit the readOnly part of the map.
func BenchmarkInteger_Standard_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			var m sync.Map
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Store(i, i)
			}
			b.ResetTimer()
			benchmarkMapOfIntegerKeys(b, func(k int) (value int, ok bool) {
				v, ok := m.Load(k)
				if ok {
					return v.(int), ok
				} else {
					return 0, false
				}
			}, func(k int, v int) {
				m.Store(k, v)
			}, func(k int) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

func benchmarkMapOfIntegerKeys(
	b *testing.B,
	loadFn func(k int) (int, bool),
	storeFn func(k int, v int),
	deleteFn func(k int),
	readPercentage int,
) {
	runParallel(b, func(pb *testing.PB) {
		// convert percent to permille to support 99% case
		storeThreshold := 10 * readPercentage
		deleteThreshold := 10*readPercentage + ((1000 - 10*readPercentage) / 2)
		for pb.Next() {
			op := int(Fastrand() % 1000)
			i := int(Fastrand() % benchmarkNumEntries)
			if op >= deleteThreshold {
				deleteFn(i)
			} else if op >= storeThreshold {
				storeFn(i, i)
			} else {
				loadFn(i)
			}
		}
	})
}

func Fastrand() uint32 {
	return fastrand()
}

//go:noescape
//go:linkname fastrand runtime.fastrand
func fastrand() uint32
