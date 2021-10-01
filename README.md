# 读写性能更优的 sync.Map

*forked from orcaman/concurrent-map*

**最新: 总体看, 直接用 `xsync.NewMap()` 更佳, 见下面 2 个测试**, 或直接用原生 `sync.Map` 性能也很好了

**注意: `xsync.NewMap()` 仅 64 位构建通过了官方认证, 若要使用 32 位, 注意自己测试**

## 改动

- 增加 `m.GetValue()` 用于单纯的取值
- 增加与 `sync.Map`, `xsync.NewMap()` 对比基准测试: 读, 读多写少, 写

## 使用场景

- `xsync.NewMap()` 适合读多写少的场景, 性能强劲: [xsync](https://github.com/puzpuzpuz/xsync)
- `cmap.New()` 适合读写均衡或写多读少的场景

```go
go test -bench=. -benchtime=2s -count=3
go version go1.17 linux/amd64
cpu: Intel(R) Xeon(R) Gold 6161 CPU @ 2.20GHz
BenchmarkReadXsync-8                    169672880               13.61 ns/op            0 B/op          0 allocs/op
BenchmarkReadXsync-8                    174356899               14.48 ns/op            0 B/op          0 allocs/op
BenchmarkReadXsync-8                    172243994               14.09 ns/op            0 B/op          0 allocs/op
BenchmarkReadCMAP-8                     77701599                30.55 ns/op            0 B/op          0 allocs/op
BenchmarkReadCMAP-8                     80629885                28.51 ns/op            0 B/op          0 allocs/op
BenchmarkReadCMAP-8                     81494463                29.21 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncMap-8                  70097936                30.65 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncMap-8                  74468286                30.22 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncMap-8                  79496142                31.42 ns/op            0 B/op          0 allocs/op
BenchmarkReadWXsync-8                   158263004               15.08 ns/op            0 B/op          0 allocs/op
BenchmarkReadWXsync-8                   161351006               15.17 ns/op            0 B/op          0 allocs/op
BenchmarkReadWXsync-8                   159669409               15.07 ns/op            0 B/op          0 allocs/op
BenchmarkReadWCMAP-8                    81286011                29.80 ns/op            0 B/op          0 allocs/op
BenchmarkReadWCMAP-8                    80469480                29.93 ns/op            0 B/op          0 allocs/op
BenchmarkReadWCMAP-8                    81316665                29.54 ns/op            0 B/op          0 allocs/op
BenchmarkReadWSyncMap-8                 74813577                31.67 ns/op            0 B/op          0 allocs/op
BenchmarkReadWSyncMap-8                 77983555                32.17 ns/op            0 B/op          0 allocs/op
BenchmarkReadWSyncMap-8                 77657517                32.46 ns/op            0 B/op          0 allocs/op
BenchmarkWriteXsync-8                   15874245               144.1 ns/op            48 B/op          3 allocs/op
BenchmarkWriteXsync-8                   16998176               142.0 ns/op            48 B/op          3 allocs/op
BenchmarkWriteXsync-8                   17554891               150.7 ns/op            48 B/op          3 allocs/op
BenchmarkWriteCMAP-8                    27520351                94.70 ns/op           16 B/op          1 allocs/op
BenchmarkWriteCMAP-8                    27236055                91.31 ns/op           16 B/op          1 allocs/op
BenchmarkWriteCMAP-8                    26136130                91.94 ns/op           16 B/op          1 allocs/op
BenchmarkWriteSyncMap-8                 16701786               164.8 ns/op            32 B/op          2 allocs/op
BenchmarkWriteSyncMap-8                 16115397               152.2 ns/op            32 B/op          2 allocs/op
BenchmarkWriteSyncMap-8                 16425662               150.6 ns/op            32 B/op          2 allocs/op
```

------

```go
go test -bench=^BenchmarkMap -benchtime=2s
go version go1.17 linux/amd64
cpu: Intel(R) Xeon(R) Gold 6161 CPU @ 2.20GHz
BenchmarkMap_NoWarmUp/99%-reads-8       56163492                52.55 ns/op
BenchmarkMap_NoWarmUp/90%-reads-8       35315628                73.24 ns/op
BenchmarkMap_NoWarmUp/75%-reads-8       26161597                81.62 ns/op
BenchmarkMap_NoWarmUp/50%-reads-8       24746589               106.9 ns/op
BenchmarkMap_NoWarmUp/0%-reads-8        18289854               141.3 ns/op
BenchmarkMapCMAP_NoWarmUp/99%-reads-8           28187413                93.63 ns/op
BenchmarkMapCMAP_NoWarmUp/90%-reads-8           13235156               202.7 ns/op
BenchmarkMapCMAP_NoWarmUp/75%-reads-8            9603061               262.1 ns/op
BenchmarkMapCMAP_NoWarmUp/50%-reads-8            8177826               307.7 ns/op
BenchmarkMapCMAP_NoWarmUp/0%-reads-8            15768786               147.7 ns/op
BenchmarkMapStandard_NoWarmUp/99%-reads-8        4889781               649.5 ns/op
BenchmarkMapStandard_NoWarmUp/90%-reads-8        2881819               848.7 ns/op
BenchmarkMapStandard_NoWarmUp/75%-reads-8        2562492               959.4 ns/op
BenchmarkMapStandard_NoWarmUp/50%-reads-8        2688487               929.0 ns/op
BenchmarkMapStandard_NoWarmUp/0%-reads-8         2304576              1045 ns/op
BenchmarkMap_WarmUp/100%-reads-8                26866794                92.23 ns/op
BenchmarkMap_WarmUp/99%-reads-8                 26107092                88.13 ns/op
BenchmarkMap_WarmUp/90%-reads-8                 29198966                80.35 ns/op
BenchmarkMap_WarmUp/75%-reads-8                 28485249                80.61 ns/op
BenchmarkMap_WarmUp/50%-reads-8                 27306980                83.98 ns/op
BenchmarkMap_WarmUp/0%-reads-8                  21844280               113.1 ns/op
BenchmarkMapCMAP_WarmUp/100%-reads-8            23667356                95.94 ns/op
BenchmarkMapCMAP_WarmUp/99%-reads-8             17253115               138.5 ns/op
BenchmarkMapCMAP_WarmUp/90%-reads-8              9708538               231.9 ns/op
BenchmarkMapCMAP_WarmUp/75%-reads-8              8166346               280.1 ns/op
BenchmarkMapCMAP_WarmUp/50%-reads-8              7348316               315.0 ns/op
BenchmarkMapCMAP_WarmUp/0%-reads-8              15754057               129.3 ns/op
BenchmarkMapStandard_WarmUp/100%-reads-8        10278343               197.9 ns/op
BenchmarkMapStandard_WarmUp/99%-reads-8          7406224               272.3 ns/op
BenchmarkMapStandard_WarmUp/90%-reads-8          7389327               273.5 ns/op
BenchmarkMapStandard_WarmUp/75%-reads-8          4945608               436.3 ns/op
BenchmarkMapStandard_WarmUp/50%-reads-8          3071937               792.5 ns/op
BenchmarkMapStandard_WarmUp/0%-reads-8           2321718               903.5 ns/op
BenchmarkMapRange-8                                  134          19187676 ns/op
BenchmarkMapRangeCMAP-8                               32          63195121 ns/op
BenchmarkMapRangeStandard-8                          169          14674270 ns/op
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
	m := cmap.New()

	// 设置变量m一个键为“foo”值为“bar”键值对
	m.Set("foo", "bar")

	// 从m中获取指定键值.
	if tmp, ok := m.Get("foo"); ok {
		bar := tmp.(string)
	}

	// 删除键为“foo”的项
	m.Remove("foo")

```

更多使用示例请查看`concurrent_map_test.go`.

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
