# 读写性能更优的 sync.Map (泛型版本)

*forked from orcaman/concurrent-map*

- go1.18+ 使用 `go get github.com/fufuok/cmap@v1.18.0`
- go1.18- 使用 `go get github.com/fufuok/cmap@v1.17.0`

**建议: 总体看, 直接用 `xsync.NewMapOf()` [xsync](https://github.com/fufuok/utils/tree/master/xsync) 更佳, 见下面 2 个测试**

**注意: `xsync.NewMapOf()` 仅 64 位构建通过了官方认证, 若要使用 32 位, 注意自己测试**

或直接用原生 `sync.Map` 性能也很好了

## 改动

- 增加 `m.GetValue()` 用于单纯的取值, `m.SetOrGet()` 检查存在时获取, 不存在时设置
- 增加与 `sync.Map`, `xsync.NewMap()` 对比基准测试: 读, 读多写少, 写

## 性能测试

```go
go test -run=^$ -benchmem -count=2 -bench='^BenchmarkRead|^BenchmardWrite'
goos: linux
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkReadXsync-12           129671468                9.304 ns/op            0 B/op          0 allocs/op
BenchmarkReadXsync-12           128664236                9.289 ns/op            0 B/op          0 allocs/op
BenchmarkReadCMAP-12             62272962                19.52 ns/op            0 B/op          0 allocs/op
BenchmarkReadCMAP-12             61951471                19.43 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncMap-12          60506438                20.14 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncMap-12          59779711                19.71 ns/op            0 B/op          0 allocs/op
BenchmarkReadWXsync-12          100000000                10.09 ns/op            0 B/op          0 allocs/op
BenchmarkReadWXsync-12          100000000                10.11 ns/op            0 B/op          0 allocs/op
BenchmarkReadWCMAP-12            58568583                20.05 ns/op            0 B/op          0 allocs/op
BenchmarkReadWCMAP-12            60009001                20.09 ns/op            0 B/op          0 allocs/op
BenchmarkReadWSyncMap-12         59734182                20.00 ns/op            0 B/op          0 allocs/op
BenchmarkReadWSyncMap-12         59192904                19.99 ns/op            0 B/op          0 allocs/op
```

------

```go
go test -run=^$ -benchmem -count=2 -bench=^BenchmarkMap
goos: linux
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkMap_NoWarmUp/99%-reads-12              48327076               31.24 ns/op             1 B/op          0 allocs/op
BenchmarkMap_NoWarmUp/99%-reads-12              40286436               31.51 ns/op             1 B/op          0 allocs/op
BenchmarkMap_NoWarmUp/90%-reads-12              25770649               59.91 ns/op            22 B/op          0 allocs/op
BenchmarkMap_NoWarmUp/90%-reads-12              28524430               45.01 ns/op            10 B/op          0 allocs/op
BenchmarkMap_NoWarmUp/75%-reads-12              18421100               67.56 ns/op            32 B/op          0 allocs/op
BenchmarkMap_NoWarmUp/75%-reads-12              22139414               52.30 ns/op            15 B/op          0 allocs/op
BenchmarkMap_NoWarmUp/50%-reads-12              17698430               74.60 ns/op            36 B/op          0 allocs/op
BenchmarkMap_NoWarmUp/50%-reads-12              17486898               77.29 ns/op            36 B/op          0 allocs/op
BenchmarkMap_NoWarmUp/0%-reads-12               12857574               84.32 ns/op            53 B/op          1 allocs/op
BenchmarkMap_NoWarmUp/0%-reads-12               11199740               95.67 ns/op            59 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/99%-reads-12          11457136               105.5 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/99%-reads-12          11232000               103.4 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/90%-reads-12           5148685               247.7 ns/op             3 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/90%-reads-12           5222319               248.9 ns/op             3 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/75%-reads-12           4601702               282.7 ns/op             6 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/75%-reads-12           4572554               295.7 ns/op             6 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/50%-reads-12           4226767               312.3 ns/op            15 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/50%-reads-12           4102942               315.0 ns/op            15 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/0%-reads-12           11634411               86.01 ns/op             5 B/op          0 allocs/op
BenchmarkMapCMAP_NoWarmUp/0%-reads-12           13807988               81.58 ns/op             4 B/op          0 allocs/op
BenchmarkMapStandard_NoWarmUp/99%-reads-12       3392006               394.0 ns/op            47 B/op          0 allocs/op
BenchmarkMapStandard_NoWarmUp/99%-reads-12       3531196               378.2 ns/op            47 B/op          0 allocs/op
BenchmarkMapStandard_NoWarmUp/90%-reads-12       2304710               599.3 ns/op            49 B/op          0 allocs/op
BenchmarkMapStandard_NoWarmUp/90%-reads-12       2330407               551.2 ns/op            49 B/op          0 allocs/op
BenchmarkMapStandard_NoWarmUp/75%-reads-12       1879966               627.0 ns/op            54 B/op          0 allocs/op
BenchmarkMapStandard_NoWarmUp/75%-reads-12       1886103               593.8 ns/op            52 B/op          0 allocs/op
BenchmarkMapStandard_NoWarmUp/50%-reads-12       1779472               652.9 ns/op            57 B/op          1 allocs/op
BenchmarkMapStandard_NoWarmUp/50%-reads-12       1801131               648.7 ns/op            56 B/op          1 allocs/op
BenchmarkMapStandard_NoWarmUp/0%-reads-12        1411503               789.6 ns/op            86 B/op          2 allocs/op
BenchmarkMapStandard_NoWarmUp/0%-reads-12        1607422               716.0 ns/op            77 B/op          2 allocs/op
BenchmarkMap_WarmUp/100%-reads-12               24238456               51.00 ns/op             0 B/op          0 allocs/op
BenchmarkMap_WarmUp/100%-reads-12               22185780               60.96 ns/op             0 B/op          0 allocs/op
BenchmarkMap_WarmUp/99%-reads-12                23345212               60.69 ns/op             0 B/op          0 allocs/op
BenchmarkMap_WarmUp/99%-reads-12                21724095               55.89 ns/op             0 B/op          0 allocs/op
BenchmarkMap_WarmUp/90%-reads-12                21041851               66.33 ns/op             1 B/op          0 allocs/op
BenchmarkMap_WarmUp/90%-reads-12                22371948               48.79 ns/op             1 B/op          0 allocs/op
BenchmarkMap_WarmUp/75%-reads-12                22804063               44.78 ns/op             3 B/op          0 allocs/op
BenchmarkMap_WarmUp/75%-reads-12                26133721               45.75 ns/op             3 B/op          0 allocs/op
BenchmarkMap_WarmUp/50%-reads-12                20519379               54.57 ns/op             6 B/op          0 allocs/op
BenchmarkMap_WarmUp/50%-reads-12                26183919               48.47 ns/op             6 B/op          0 allocs/op
BenchmarkMap_WarmUp/0%-reads-12                 19737588               57.82 ns/op            12 B/op          1 allocs/op
BenchmarkMap_WarmUp/0%-reads-12                 22736452               52.16 ns/op            12 B/op          1 allocs/op
BenchmarkMapCMAP_WarmUp/100%-reads-12           19694181               51.60 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/100%-reads-12           22738555               52.44 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/99%-reads-12             8034818               152.0 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/99%-reads-12             7730708               145.9 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/90%-reads-12             4043752               295.6 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/90%-reads-12             4028773               291.4 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/75%-reads-12             3191673               325.6 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/75%-reads-12             3344899               329.2 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/50%-reads-12             3449197               361.8 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/50%-reads-12             3552146               323.8 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/0%-reads-12             16581148               84.50 ns/op             0 B/op          0 allocs/op
BenchmarkMapCMAP_WarmUp/0%-reads-12             15677961               78.09 ns/op             0 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/100%-reads-12        7987387               131.5 ns/op             0 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/100%-reads-12        7768620               145.2 ns/op             0 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/99%-reads-12         2109114               480.2 ns/op            27 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/99%-reads-12         2140489               471.2 ns/op            27 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/90%-reads-12         2104882               481.7 ns/op            29 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/90%-reads-12         2078965               494.4 ns/op            29 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/75%-reads-12         1952872               564.9 ns/op            34 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/75%-reads-12         1984945               541.7 ns/op            34 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/50%-reads-12         1953594               634.6 ns/op            27 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/50%-reads-12         1993894               614.4 ns/op            27 B/op          0 allocs/op
BenchmarkMapStandard_WarmUp/0%-reads-12          1737729               740.0 ns/op            40 B/op          1 allocs/op
BenchmarkMapStandard_WarmUp/0%-reads-12          1890639               698.9 ns/op            39 B/op          1 allocs/op
BenchmarkMapRange-12                                  82            16068934 ns/op           124 B/op          1 allocs/op
BenchmarkMapRange-12                                  99            15810145 ns/op           127 B/op          1 allocs/op
BenchmarkMapRangeCMAP-12                              31            47355519 ns/op      48127209 B/op        199 allocs/op
BenchmarkMapRangeCMAP-12                              25            47870128 ns/op      48127224 B/op        200 allocs/op
BenchmarkMapRangeStandard-12                          78            14191099 ns/op            14 B/op          0 allocs/op
BenchmarkMapRangeStandard-12                          79            14703753 ns/op            14 B/op          0 allocs/op
```

# concurrent map [![Build Status](https://travis-ci.com/orcaman/concurrent-map.svg?branch=master)](https://travis-ci.com/orcaman/concurrent-map)

正如 [这里](http://golang.org/doc/faq#atomic_maps) 和 [这里](http://blog.golang.org/go-maps-in-action) 所描述的, Go语言原生的`map`类型并不支持并发读写。`concurrent-map`提供了一种高性能的解决方案:通过对内部`map`进行分片，降低锁粒度，从而达到最少的锁等待时间(锁冲突)

在Go 1.9之前，go语言标准库中并没有实现并发`map`。在Go 1.9中，引入了`sync.Map`。新的`sync.Map`与此`concurrent-map`有几个关键区别。标准库中的`sync.Map`是专为`append-only`场景设计的。因此，如果您想将`Map`用于一个类似内存数据库，那么使用我们的版本可能会受益。你可以在golang repo上读到更多，[这里](https://github.com/golang/go/issues/21035) and [这里](https://stackoverflow.com/questions/11063473/map-with-concurrent-access)
***译注:`sync.Map`在读多写少性能比较好，否则并发性能很差***

## 用法

导入包:

```go
import (
	"github.com/fufuok/cmap"
)

```

```bash
go get "github.com/fufuok/cmap"
```

## 示例

```go

	// 创建一个新的 map.
	m := cmap.New[string]()

	// 设置变量m一个键为“foo”值为“bar”键值对
	m.Set("foo", "bar")

	// 从m中获取指定键值.
	if tmp, ok := m.Get("foo"); ok {
		bar := tmp
	}

	// 删除键为“foo”的项
	m.Remove("foo")

```

更多使用示例请查看: `concurrent_map_test.go`.

运行测试:

```bash
go test "github.com/fufuok/cmap"
```

## 贡献说明

我们非常欢迎大家的贡献。如欲合并贡献，请遵循以下指引:
- 新建一个issue,并且叙述为什么这么做(解决一个bug，增加一个功能，等等)
- 根据核心团队对上述问题的反馈，提交一个PR，描述变更并链接到该问题。
- 新代码必须具有测试覆盖率。
- 如果代码是关于性能问题的，则必须在流程中包括基准测试(无论是在问题中还是在PR中)。
- 一般来说，我们希望`concurrent-map`尽可能简单，且与原生的`map`有相似的操作。当你新建issue时请注意这一点。

## 许可证
MIT (see [LICENSE](https://github.com/orcaman/concurrent-map/blob/master/LICENSE) file)
