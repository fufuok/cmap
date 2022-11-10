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
BenchmarkInteger_CMAP_WarmUp/reads=99%-12               75669090                15.66 ns/op       63850319 ops/s
BenchmarkInteger_CMAP_WarmUp/reads=90%-12               49213774                24.32 ns/op       41118049 ops/s
BenchmarkInteger_CMAP_WarmUp/reads=75%-12               39052770                30.63 ns/op       32644506 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=99%-12            42396834                27.95 ns/op       35782077 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=90%-12            24108681                48.55 ns/op       20598723 ops/s
BenchmarkInteger_SafeMap_WarmUp/reads=75%-12            18924130                61.53 ns/op       16251938 ops/s
BenchmarkInteger_Cache_WarmUp/reads=99%-12              177968762                6.389 ns/op     156519378 ops/s
BenchmarkInteger_Cache_WarmUp/reads=90%-12              160716399                7.368 ns/op     135728234 ops/s
BenchmarkInteger_Cache_WarmUp/reads=75%-12              130527783                8.983 ns/op     111325683 ops/s
BenchmarkInteger_Standard_WarmUp/reads=99%-12           42490860                23.96 ns/op       41732313 ops/s
BenchmarkInteger_Standard_WarmUp/reads=90%-12           29298574               150.1 ns/op         6661742 ops/s
BenchmarkInteger_Standard_WarmUp/reads=75%-12           15083128               137.0 ns/op         7299573 ops/s
```







*ff*





