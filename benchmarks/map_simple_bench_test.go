package benchmarks

import (
	"strings"
	"sync"
	"testing"

	"github.com/alphadose/haxmap"
	"github.com/cornelk/hashmap"
	"github.com/orcaman/concurrent-map/v2"
	"github.com/puzpuzpuz/xsync/v2"
	"github.com/smallnest/safemap"
)

var (
	testKey   = "Fufu"
	testValue = strings.Repeat("test", 1500)
)

func Benchmark_Read_Xsync(b *testing.B) {
	var m = xsync.NewMapOf[string]()
	m.Store(testKey, testValue)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Load(testKey)
	}
}

func Benchmark_Read_CMAP(b *testing.B) {
	var m = cmap.New[string]()
	m.Set(testKey, testValue)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(testKey)
	}
}

func Benchmark_Read_SafeMap(b *testing.B) {
	var m = safemap.New[string, string]()
	m.Set(testKey, testValue)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(testKey)
	}
}

func Benchmark_Read_HaxMap(b *testing.B) {
	var m = haxmap.New[string, string]()
	m.Set(testKey, testValue)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(testKey)
	}
}

func Benchmark_Read_HashMap(b *testing.B) {
	var m = hashmap.New[string, string]()
	m.Set(testKey, testValue)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(testKey)
	}
}

func Benchmark_Read_StdSyncMap(b *testing.B) {
	var m sync.Map
	m.Store(testKey, testValue)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Load(testKey)
	}
}

func Benchmark_ReadW_Xsync(b *testing.B) {
	var m = xsync.NewMapOf[string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			m.Store(testKey, v)
		}
		_, _ = m.Load(testKey)
	}
}

func Benchmark_ReadW_CMAP(b *testing.B) {
	var m = cmap.New[string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			m.Set(testKey, v)
		}
		_, _ = m.Get(testKey)
	}
}

func Benchmark_ReadW_SafeMap(b *testing.B) {
	var m = safemap.New[string, string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			m.Set(testKey, v)
		}
		_, _ = m.Get(testKey)
	}
}

func Benchmark_ReadW_HaxMap(b *testing.B) {
	var m = haxmap.New[string, string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			m.Set(testKey, v)
		}
		_, _ = m.Get(testKey)
	}
}

func Benchmark_ReadW_HashMap(b *testing.B) {
	var m = hashmap.New[string, string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			m.Set(testKey, v)
		}
		_, _ = m.Get(testKey)
	}
}

func Benchmark_ReadW_StdSyncMap(b *testing.B) {
	var m sync.Map
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			m.Store(testKey, v)
		}
		_, _ = m.Load(testKey)
	}
}

func Benchmark_Write_Xsync(b *testing.B) {
	var m = xsync.NewMapOf[string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store(testKey, v)
	}
}

func Benchmark_Write_CMAP(b *testing.B) {
	var m = cmap.New[string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(testKey, v)
	}
}

func Benchmark_Write_SafeMap(b *testing.B) {
	var m = safemap.New[string, string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(testKey, v)
	}
}

func Benchmark_Write_HaxMap(b *testing.B) {
	var m = haxmap.New[string, string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(testKey, v)
	}
}

func Benchmark_Write_HashMap(b *testing.B) {
	var m = hashmap.New[string, string]()
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(testKey, v)
	}
}

func Benchmark_Write_StdSyncMap(b *testing.B) {
	var m sync.Map
	v := testValue
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store(testKey, v)
	}
}

// go test -bench=^Benchmark_ -benchmem -count=2
// goos: linux
// goarch: amd64
// pkg: bench
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// Benchmark_Read_Xsync-4          37724912                31.85 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_Xsync-4          37749138                32.11 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_CMAP-4           47647892                24.68 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_CMAP-4           48671798                24.91 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_SafeMap-4        39895554                29.93 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_SafeMap-4        40271330                29.83 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_HaxMap-4         76125390                15.62 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_HaxMap-4         76314477                15.71 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_HashMap-4        98442661                12.17 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_HashMap-4        98248668                12.35 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_StdSyncMap-4     46558422                25.81 ns/op            0 B/op          0 allocs/op
// Benchmark_Read_StdSyncMap-4     45827604                25.78 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_Xsync-4         36525292                32.44 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_Xsync-4         36933769                32.46 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_CMAP-4          46853349                25.56 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_CMAP-4          46678730                25.55 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_SafeMap-4       39139291                30.66 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_SafeMap-4       39675928                30.43 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_HaxMap-4        67317434                17.62 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_HaxMap-4        67815920                17.64 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_HashMap-4       79321268                16.39 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_HashMap-4       73291922                16.55 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_StdSyncMap-4    41708286                28.76 ns/op            0 B/op          0 allocs/op
// Benchmark_ReadW_StdSyncMap-4    41658141                29.03 ns/op            0 B/op          0 allocs/op
// Benchmark_Write_Xsync-4          7636951               151.1 ns/op            48 B/op          3 allocs/op
// Benchmark_Write_Xsync-4          7296279               151.1 ns/op            48 B/op          3 allocs/op
// Benchmark_Write_CMAP-4          25714945                45.91 ns/op            0 B/op          0 allocs/op
// Benchmark_Write_CMAP-4          25959980                46.42 ns/op            0 B/op          0 allocs/op
// Benchmark_Write_SafeMap-4       24291164                49.58 ns/op            0 B/op          0 allocs/op
// Benchmark_Write_SafeMap-4       24012396                49.12 ns/op            0 B/op          0 allocs/op
// Benchmark_Write_HaxMap-4        16488783                68.96 ns/op           16 B/op          1 allocs/op
// Benchmark_Write_HaxMap-4        16733216                69.55 ns/op           16 B/op          1 allocs/op
// Benchmark_Write_HashMap-4       17638938                65.34 ns/op           16 B/op          1 allocs/op
// Benchmark_Write_HashMap-4       17589645                66.17 ns/op           16 B/op          1 allocs/op
// Benchmark_Write_StdSyncMap-4     6836103               157.8 ns/op            48 B/op          3 allocs/op
// Benchmark_Write_StdSyncMap-4     6914628               156.8 ns/op            48 B/op          3 allocs/op
