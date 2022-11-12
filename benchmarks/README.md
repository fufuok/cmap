# 🤖 Benchmarks

- Number of entries used in benchmark: `200_000`

- ```go
  {"reads=99%", 99}, //  99% loads,  0.5% stores,  0.5% deletes
  {"reads=90%", 90}, //  90% loads,    5% stores,    5% deletes
  {"reads=75%", 75}, //  75% loads, 12.5% stores, 12.5% deletes
  ```

## String keys

```go
go test -bench=^BenchmarkString_
goos: windows
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkString_CMAP_NoWarmUp/reads=99%-12              42417969                28.66 ns/op
BenchmarkString_CMAP_NoWarmUp/reads=90%-12              32355740                37.10 ns/op
BenchmarkString_CMAP_NoWarmUp/reads=75%-12              27322963                45.01 ns/op
BenchmarkString_SafeMap_NoWarmUp/reads=99%-12           27943497                36.71 ns/op
BenchmarkString_SafeMap_NoWarmUp/reads=90%-12           18915540                57.89 ns/op
BenchmarkString_SafeMap_NoWarmUp/reads=75%-12           16393486                74.96 ns/op
BenchmarkString_Cache_NoWarmUp/reads=99%-12             79572838                15.56 ns/op
BenchmarkString_Cache_NoWarmUp/reads=90%-12             55431168                22.12 ns/op
BenchmarkString_Cache_NoWarmUp/reads=75%-12             49178512                26.67 ns/op
BenchmarkString_HaxMap_NoWarmUp/reads=99%-12            55885954               116.1 ns/op
BenchmarkString_HaxMap_NoWarmUp/reads=90%-12             9531100              1605 ns/op
BenchmarkString_HashMap_NoWarmUp/reads=99%-12           37257708                32.37 ns/op
BenchmarkString_HashMap_NoWarmUp/reads=90%-12           18915658               668.0 ns/op
BenchmarkString_Standard_NoWarmUp/reads=99%-12          25614918               166.1 ns/op
BenchmarkString_Standard_NoWarmUp/reads=90%-12           6437218               292.4 ns/op
BenchmarkString_Standard_NoWarmUp/reads=75%-12           4709377               333.5 ns/op
BenchmarkString_CMAP_WarmUp/reads=99%-12                30009226                40.37 ns/op
BenchmarkString_CMAP_WarmUp/reads=90%-12                26159292                43.62 ns/op
BenchmarkString_CMAP_WarmUp/reads=75%-12                23198671                49.39 ns/op
BenchmarkString_SafeMap_WarmUp/reads=99%-12             26728795                44.69 ns/op
BenchmarkString_SafeMap_WarmUp/reads=90%-12             17076574                64.24 ns/op
BenchmarkString_SafeMap_WarmUp/reads=75%-12             14813497                78.01 ns/op
BenchmarkString_Cache_WarmUp/reads=99%-12               33230503                33.13 ns/op
BenchmarkString_Cache_WarmUp/reads=90%-12               40468738                31.63 ns/op
BenchmarkString_Cache_WarmUp/reads=75%-12               33229766                31.53 ns/op
BenchmarkString_HaxMap_WarmUp/reads=99%-12                332990              3192 ns/op
BenchmarkString_Standard_WarmUp/reads=99%-12            22830289                48.99 ns/op
BenchmarkString_Standard_WarmUp/reads=90%-12            20231482                50.75 ns/op
BenchmarkString_Standard_WarmUp/reads=75%-12            19520173                54.85 ns/op
BenchmarkString_CMAP_Range-12                                190           5590297 ns/op
BenchmarkString_SafeMap_Range-12                             226           5325449 ns/op
BenchmarkString_Cache_Range-12                              1638            632791 ns/op
BenchmarkString_HaxMap_Range-12                             1716            632481 ns/op
BenchmarkString_Standard_Range-12                            924           1358386 ns/op
```

## Integer keys

```go
go test -bench=^BenchmarkInteger_
goos: windows
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkInteger_CMAP_WarmUp/reads=99%-12       80365394                14.88 ns/op       67214011 ops/s
BenchmarkInteger_CMAP_WarmUp/reads=90%-12       51061072                23.46 ns/op       42629175 ops/s
BenchmarkInteger_CMAP_WarmUp/reads=75%-12       38628091                29.59 ns/op       33794584 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_32-12          75327201                15.18 ns/op       65886341 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_32-12          51177642                23.16 ns/op       43175612 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_32-12          40608426                29.42 ns/op       33992602 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_64-12          100000000               12.12 ns/op       82517442 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_64-12          72758290                16.31 ns/op       61304919 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_64-12          52316532                22.02 ns/op       45413294 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_128-12         112271064               10.75 ns/op       93005034 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_128-12         92783502                12.94 ns/op       77279425 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_128-12         75200691                15.49 ns/op       64552232 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_256-12         117450596                9.930 ns/op     100708511 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_256-12         100000000               11.65 ns/op       85811233 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_256-12         92559739                13.70 ns/op       72998280 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_512-12         120987655               10.97 ns/op       91176572 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_512-12         100000000               11.03 ns/op       90621828 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_512-12         100000000               12.25 ns/op       81638091 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=99%-12                    44300118                27.20 ns/op       36770373 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=90%-12                    22281928                46.94 ns/op       21303729 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=75%-12                    17700211                58.89 ns/op       16982101 ops/s
BenchmarkInteger_Cache_WarmUp/reads=99%-12                      208903616                5.439 ns/op     183868860 ops/s
BenchmarkInteger_Cache_WarmUp/reads=90%-12                      188999635                7.175 ns/op     139368514 ops/s
BenchmarkInteger_Cache_WarmUp/reads=75%-12                      127810842                9.146 ns/op     109336719 ops/s
BenchmarkInteger_Standard_WarmUp/reads=99%-12                   52292466                20.81 ns/op       48059845 ops/s
BenchmarkInteger_Standard_WarmUp/reads=90%-12                   31227331               143.2 ns/op         6981236 ops/s
BenchmarkInteger_Standard_WarmUp/reads=75%-12                   15239526               141.0 ns/op         7091457 ops/s
```

```go
go test -bench=^BenchmarkIn
goos: linux
goarch: amd64
pkg: bench
cpu: AMD EPYC 7K62 48-Core Processor
BenchmarkInteger_CMAP_WarmUp/reads=99%-2        25585442                49.00 ns/op       20408288 ops/s
BenchmarkInteger_CMAP_WarmUp/reads=90%-2        18281479                70.56 ns/op       14171650 ops/s
BenchmarkInteger_CMAP_WarmUp/reads=75%-2        14518659                82.65 ns/op       12098722 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_32-2           25510890                49.17 ns/op       20337291 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_32-2           17771215                71.29 ns/op       14028147 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_32-2           12097663                85.25 ns/op       11730422 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_64-2           25779288                49.09 ns/op       20371160 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_64-2           20969155                59.96 ns/op       16676994 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_64-2           16268608                73.03 ns/op       13693624 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_128-2          24956167                48.72 ns/op       20525810 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_128-2          22947441                55.63 ns/op       17976899 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_128-2          18864630                64.97 ns/op       15390977 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_256-2          24874155                49.88 ns/op       20047947 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_256-2          23469039                53.08 ns/op       18838796 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_256-2          21431408                57.65 ns/op       17344797 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=99%_512-2          21516894                50.64 ns/op       19748089 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=90%_512-2          23200842                53.29 ns/op       18764473 ops/s
BenchmarkInteger_CMAP_NumShards_WarmUp/reads=75%_512-2          20968974                56.32 ns/op       17757200 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=99%-2                     16491586                76.21 ns/op       13122468 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=90%-2                     13435584                95.14 ns/op       10511304 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=75%-2                     12532111                96.98 ns/op       10310961 ops/s
BenchmarkInteger_Cache_WarmUp/reads=99%-2                       32217166                38.96 ns/op       25665879 ops/s
BenchmarkInteger_Cache_WarmUp/reads=90%-2                       29863941                48.24 ns/op       20729331 ops/s
BenchmarkInteger_Cache_WarmUp/reads=75%-2                       23865232                49.05 ns/op       20387034 ops/s
BenchmarkInteger_Standard_WarmUp/reads=99%-2                     9185739               120.2 ns/op         8321800 ops/s
BenchmarkInteger_Standard_WarmUp/reads=90%-2                     7519438               149.2 ns/op         6702233 ops/s
BenchmarkInteger_Standard_WarmUp/reads=75%-2                     5696744               188.4 ns/op         5307384 ops/s
```






*ff*





