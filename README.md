# 读写性能更优的 sync.Map (泛型版本)

*forked from orcaman/concurrent-map*

- go1.18+ 使用 `go get github.com/fufuok/cmap@v1.18.0`
- go1.18- 使用 `go get github.com/fufuok/cmap@v1.17.0`

**建议: 总体看, 直接用 `xsync.NewMapOf()` [github.com/puzpuzpuz/xsync](github.com/puzpuzpuz/xsync) 更佳, 符合日常使用场景, 且支持可比较泛型键. 下面有测试**

> 如果有 Cache 需求, 可以选用: [CacheOf and MapOf](https://github.com/fufuok/cache)

## 改动

- 增加 `m.GetValue()` 用于单纯的取值, `m.SetOrGet()` 检查存在时获取, 不存在时设置
- 增加与 `sync.Map`, `xsync.NewMap()` 对比基准测试: 读, 读多写少, 写

## 性能测试

Code: [Banchmarks](benchmarks)

- 可以根据您的场景, 修改 [benchmarks/map_bench_test.go] 的测试参数, 下面是参考值

- Number of entries used in benchmark: `500_000`

- ```go
  {"100%-reads", 100}, // 100% loads,    0% stores,    0% deletes
  {"99%-reads", 99},   //  99% loads,  0.5% stores,  0.5% deletes
  {"95%-reads", 95},   //  95% loads,  2.5% stores,  2.5% deletes
  {"90%-reads", 90},   //  90% loads,    5% stores,    5% deletes
  {"75%-reads", 75},   //  75% loads, 12.5% stores, 12.5% deletes
  {"50%-reads", 50},   //  50% loads,   25% stores,   25% deletes
  {"0%-reads", 0},     //   0% loads,   50% stores,   50% deletes
  ```

```go
go test -bench=^BenchmarkMap -benchmem
goos: linux
goarch: amd64
pkg: bench
cpu: Intel(R) Xeon(R) Silver 4314 CPU @ 2.40GHz
BenchmarkMap_Xsync_NoWarmUp/99%-reads-64                110687433               12.49 ns/op            0 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/95%-reads-64                88322406                17.27 ns/op            1 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/90%-reads-64                60254306                19.02 ns/op            3 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/75%-reads-64                47258060                24.51 ns/op            7 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/50%-reads-64                37846016                29.71 ns/op           14 B/op          1 allocs/op
BenchmarkMap_Xsync_NoWarmUp/0%-reads-64                 21427845                53.49 ns/op           29 B/op          2 allocs/op
BenchmarkMap_CMAP_NoWarmUp/99%-reads-64                 12194355                88.52 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/95%-reads-64                 10678164               103.7 ns/op             1 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/90%-reads-64                  9453764               115.5 ns/op             2 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/75%-reads-64                  7884636               147.0 ns/op             4 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/50%-reads-64                  5907775               169.7 ns/op             5 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/0%-reads-64                   8712856               136.8 ns/op             3 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/99%-reads-64              19244058                54.46 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/95%-reads-64              15166338                66.34 ns/op            1 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/90%-reads-64              13815511                79.05 ns/op            2 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/75%-reads-64              11180854                97.32 ns/op            2 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/50%-reads-64               9022095               116.0 ns/op             3 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/0%-reads-64               12543771                94.35 ns/op            2 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/99%-reads-64               49126554               322.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/95%-reads-64               11755160              1698 ns/op               4 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/90%-reads-64                2689261              1555 ns/op               6 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/99%-reads-64              59778726                70.58 ns/op            0 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/95%-reads-64              28713607               503.3 ns/op             1 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/90%-reads-64               9664908               883.1 ns/op             3 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/99%-reads-64              1000000              1256 ns/op              34 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/95%-reads-64              1000000              1802 ns/op              49 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/90%-reads-64               757347              2299 ns/op              52 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/75%-reads-64               492739              2053 ns/op              61 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/50%-reads-64               581881              2393 ns/op              67 B/op          1 allocs/op
BenchmarkMap_Standard_NoWarmUp/0%-reads-64                450807              2734 ns/op              74 B/op          2 allocs/op
BenchmarkMap_Xsync_WarmUp/100%-reads-64                 72828616                15.13 ns/op            0 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/99%-reads-64                  83636658                12.63 ns/op            0 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/95%-reads-64                  109142733               11.13 ns/op            1 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/90%-reads-64                  86160418                11.95 ns/op            2 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/75%-reads-64                  61728336                18.89 ns/op            7 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/50%-reads-64                  40054539                29.05 ns/op           14 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/0%-reads-64                   22612006                51.49 ns/op           28 B/op          1 allocs/op
BenchmarkMap_CMAP_WarmUp/100%-reads-64                  19063858                56.10 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/99%-reads-64                   12072297                92.42 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/95%-reads-64                   10185296               109.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/90%-reads-64                    8919643               125.0 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/75%-reads-64                    9807985               155.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/50%-reads-64                    5923474               192.6 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/0%-reads-64                     7987324               140.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/100%-reads-64               28904026                41.24 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/99%-reads-64                17206824                60.07 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/95%-reads-64                23110501                74.46 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/90%-reads-64                13308271                80.85 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/75%-reads-64                11731944                98.30 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/50%-reads-64                 9324200               122.5 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/0%-reads-64                 16906158                91.44 ns/op            0 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/100%-reads-64                84342962                14.05 ns/op            0 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/99%-reads-64                   204526              5227 ns/op               2 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/95%-reads-64                    39470             32335 ns/op              11 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/100%-reads-64                491080              2456 ns/op               0 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/99%-reads-64                 500222              2310 ns/op               0 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/95%-reads-64                 565968              2648 ns/op              52 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/90%-reads-64                 485007              2373 ns/op               2 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/75%-reads-64                 506257              2420 ns/op               5 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/50%-reads-64                 458026              2452 ns/op              11 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/0%-reads-64                  418957              2582 ns/op              22 B/op          1 allocs/op
BenchmarkMap_Xsync_Range-64                                 1690            728412 ns/op               3 B/op          0 allocs/op
BenchmarkMap_CMAP_Range-64                                    14          94184852 ns/op        24067662 B/op        208 allocs/op
BenchmarkMap_SafeMap_Range-64                                 34          71351911 ns/op        24262000 B/op        412 allocs/op
BenchmarkMap_HaxMap_Range-64                                1179            853686 ns/op               4 B/op          0 allocs/op
BenchmarkMap_Standard_Range-64                               876           1328893 ns/op               6 B/op          0 allocs/op
```

```go
go test -run=^$ -bench=^BenchmarkMap -benchmem
goos: windows
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkMap_Xsync_NoWarmUp/99%-reads-12                64710608                20.33 ns/op            0 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/95%-reads-12                53458306                25.36 ns/op            1 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/90%-reads-12                42397732                28.43 ns/op            3 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/75%-reads-12                39661684                36.12 ns/op            7 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/50%-reads-12                29988229                43.22 ns/op           14 B/op          1 allocs/op
BenchmarkMap_Xsync_NoWarmUp/0%-reads-12                 23644387                56.20 ns/op           28 B/op          2 allocs/op
BenchmarkMap_CMAP_NoWarmUp/99%-reads-12                 36163094                35.32 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/95%-reads-12                 30738255                39.12 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/90%-reads-12                 29988453                41.85 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/75%-reads-12                 25092056                48.78 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/50%-reads-12                 21845985                55.97 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/0%-reads-12                  22355026                51.69 ns/op            1 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/99%-reads-12              31525105                38.27 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/95%-reads-12              24108003                51.54 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/90%-reads-12              18629055                60.09 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/75%-reads-12              16177629                75.59 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/50%-reads-12              14132227                88.33 ns/op            1 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/0%-reads-12               18080948                69.31 ns/op            1 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/99%-reads-12               45537685               103.6 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/95%-reads-12               17076453               802.4 ns/op             1 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/90%-reads-12                9040446              1706 ns/op               3 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/99%-reads-12              37258054                52.16 ns/op            0 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/95%-reads-12              26412634               328.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/90%-reads-12              10161258               469.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/99%-reads-12             32354431               205.6 ns/op            25 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/95%-reads-12              7684450               282.0 ns/op            36 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/90%-reads-12              5882802               323.0 ns/op            38 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/75%-reads-12              4196446               441.0 ns/op            45 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/50%-reads-12              3584572               397.7 ns/op            44 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/0%-reads-12               3227061               406.9 ns/op            44 B/op          1 allocs/op
BenchmarkMap_Xsync_WarmUp/100%-reads-12                 26602819                41.82 ns/op            0 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/99%-reads-12                  27340144                42.23 ns/op            0 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/95%-reads-12                  26160547                43.09 ns/op            1 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/90%-reads-12                  27323150                43.50 ns/op            2 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/75%-reads-12                  23644341                46.19 ns/op            6 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/50%-reads-12                  24121766                50.25 ns/op           13 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/0%-reads-12                   20207196                58.87 ns/op           28 B/op          1 allocs/op
BenchmarkMap_CMAP_WarmUp/100%-reads-12                  23688448                48.25 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/99%-reads-12                   24590314                47.63 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/95%-reads-12                   24590869                48.50 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/90%-reads-12                   23644574                50.85 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/75%-reads-12                   20839302                55.83 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/50%-reads-12                   19516192                60.61 ns/op            0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/0%-reads-12                    20155837                53.12 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/100%-reads-12               25765946                46.52 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/99%-reads-12                23643782                49.82 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/95%-reads-12                19839825                61.30 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/90%-reads-12                17317291                68.36 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/75%-reads-12                13971813                83.34 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/50%-reads-12                12807253                93.12 ns/op            0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/0%-reads-12                 16842577                69.89 ns/op            0 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/100%-reads-12                28593690                47.12 ns/op            0 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/99%-reads-12                    56652             24877 ns/op               1 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/95%-reads-12                    18652            128776 ns/op               5 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/100%-reads-12              16443584                64.52 ns/op            0 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/99%-reads-12               15949532                72.13 ns/op            2 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/95%-reads-12               13338102                78.88 ns/op            3 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/90%-reads-12               16004941                71.90 ns/op            3 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/75%-reads-12               12688712                82.75 ns/op            7 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/50%-reads-12               11426654                90.88 ns/op           12 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/0%-reads-12                10297602               103.1 ns/op            22 B/op          1 allocs/op
BenchmarkMap_Xsync_Range-12                                  280           4545280 ns/op               3 B/op          0 allocs/op
BenchmarkMap_CMAP_Range-12                                    80          15866595 ns/op        24067569 B/op        200 allocs/op
BenchmarkMap_SafeMap_Range-12                                 81          15316763 ns/op        24046433 B/op         79 allocs/op
BenchmarkMap_HaxMap_Range-12                                 320           3632630 ns/op               3 B/op          0 allocs/op
BenchmarkMap_Standard_Range-12                               193           6159392 ns/op               5 B/op          0 allocs/op
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
