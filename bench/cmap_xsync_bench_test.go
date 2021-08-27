package bench

import (
	"strings"
	"sync"
	"testing"

	"github.com/puzpuzpuz/xsync"

	"github.com/fufuok/cmap"
)

func BenchmarkReadXsync(b *testing.B) {
	var xm = xsync.NewMap()
	xm.Store("Fufu", strings.Repeat("string", 10000))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = xm.Load("Fufu")
	}
}

func BenchmarkReadCMAP(b *testing.B) {
	var cm = cmap.New()
	cm.Set("Fufu", strings.Repeat("string", 10000))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cm.Get("Fufu")
	}
}

func BenchmarkReadSyncMap(b *testing.B) {
	var sm sync.Map
	sm.Store("Fufu", strings.Repeat("string", 10000))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sm.Load("Fufu")
	}
}

func BenchmarkReadWXsync(b *testing.B) {
	var xm = xsync.NewMap()
	v := strings.Repeat("string", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			xm.Store("Fufu", v)
		}
		_, _ = xm.Load("Fufu")
	}
}

func BenchmarkReadWCMAP(b *testing.B) {
	var cm = cmap.New()
	v := strings.Repeat("string", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			cm.Set("Fufu", v)
		}
		_, _ = cm.Get("Fufu")
	}
}

func BenchmarkReadWSyncMap(b *testing.B) {
	var sm sync.Map
	v := strings.Repeat("string", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			sm.Store("Fufu", v)
		}
		_, _ = sm.Load("Fufu")
	}
}

func BenchmarkWriteXsync(b *testing.B) {
	var xm = xsync.NewMap()
	v := strings.Repeat("string", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xm.Store("Fufu", v)
	}
}

func BenchmarkWriteCMAP(b *testing.B) {
	var cm = cmap.New()
	v := strings.Repeat("string", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Set("Fufu", v)
	}
}

func BenchmarkWriteSyncMap(b *testing.B) {
	var sm sync.Map
	v := strings.Repeat("string", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Store("Fufu", v)
	}
}

// go version go1.17 linux/amd64
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkReadXsync-8            81850363                14.81 ns/op            0 B/op          0 allocs/op
// BenchmarkReadCMAP-8             48727599                25.17 ns/op            0 B/op          0 allocs/op
// BenchmarkReadSyncMap-8          35928573                33.19 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWXsync-8           74723988                15.47 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWCMAP-8            46075694                26.05 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWSyncMap-8         36574660                33.66 ns/op            0 B/op          0 allocs/op
// BenchmarkWriteXsync-8            8249457               145.0 ns/op            48 B/op          3 allocs/op
// BenchmarkWriteCMAP-8            14182921                87.27 ns/op           16 B/op          1 allocs/op
// BenchmarkWriteSyncMap-8          8048661               153.5 ns/op            32 B/op          2 allocs/op
