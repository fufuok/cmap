package bench

import (
	"strings"
	"sync"
	"testing"

	"github.com/puzpuzpuz/xsync"

	"github.com/fufuok/cmap"
)

func BenchmarkReadXsync(b *testing.B) {
	var xm = xsync.NewMapOf[string]()
	xm.Store("Fufu", strings.Repeat("string", 10000))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = xm.Load("Fufu")
	}
}

func BenchmarkReadCMAP(b *testing.B) {
	var cm = cmap.New[string]()
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
	var xm = xsync.NewMapOf[string]()
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
	var cm = cmap.New[string]()
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
	var xm = xsync.NewMapOf[string]()
	v := strings.Repeat("string", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xm.Store("Fufu", v)
	}
}

func BenchmarkWriteCMAP(b *testing.B) {
	var cm = cmap.New[string]()
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

// go test -bench=. -benchtime=2s -count=3
// go version go1.17 linux/amd64
// cpu: Intel(R) Xeon(R) Gold 6161 CPU @ 2.20GHz
// BenchmarkReadXsync-8                    169672880               13.61 ns/op            0 B/op          0 allocs/op
// BenchmarkReadXsync-8                    174356899               14.48 ns/op            0 B/op          0 allocs/op
// BenchmarkReadXsync-8                    172243994               14.09 ns/op            0 B/op          0 allocs/op
// BenchmarkReadCMAP-8                     77701599                30.55 ns/op            0 B/op          0 allocs/op
// BenchmarkReadCMAP-8                     80629885                28.51 ns/op            0 B/op          0 allocs/op
// BenchmarkReadCMAP-8                     81494463                29.21 ns/op            0 B/op          0 allocs/op
// BenchmarkReadSyncMap-8                  70097936                30.65 ns/op            0 B/op          0 allocs/op
// BenchmarkReadSyncMap-8                  74468286                30.22 ns/op            0 B/op          0 allocs/op
// BenchmarkReadSyncMap-8                  79496142                31.42 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWXsync-8                   158263004               15.08 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWXsync-8                   161351006               15.17 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWXsync-8                   159669409               15.07 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWCMAP-8                    81286011                29.80 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWCMAP-8                    80469480                29.93 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWCMAP-8                    81316665                29.54 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWSyncMap-8                 74813577                31.67 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWSyncMap-8                 77983555                32.17 ns/op            0 B/op          0 allocs/op
// BenchmarkReadWSyncMap-8                 77657517                32.46 ns/op            0 B/op          0 allocs/op
// BenchmarkWriteXsync-8                   15874245               144.1 ns/op            48 B/op          3 allocs/op
// BenchmarkWriteXsync-8                   16998176               142.0 ns/op            48 B/op          3 allocs/op
// BenchmarkWriteXsync-8                   17554891               150.7 ns/op            48 B/op          3 allocs/op
// BenchmarkWriteCMAP-8                    27520351                94.70 ns/op           16 B/op          1 allocs/op
// BenchmarkWriteCMAP-8                    27236055                91.31 ns/op           16 B/op          1 allocs/op
// BenchmarkWriteCMAP-8                    26136130                91.94 ns/op           16 B/op          1 allocs/op
// BenchmarkWriteSyncMap-8                 16701786               164.8 ns/op            32 B/op          2 allocs/op
// BenchmarkWriteSyncMap-8                 16115397               152.2 ns/op            32 B/op          2 allocs/op
// BenchmarkWriteSyncMap-8                 16425662               150.6 ns/op            32 B/op          2 allocs/op
