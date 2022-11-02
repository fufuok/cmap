# Banchmarks

## Benchmark 500_000 keys

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
go test -bench=^BenchmarkMap -benchmem
goos: linux
goarch: amd64
pkg: bench
cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
BenchmarkMap_Xsync_NoWarmUp/99%-reads-4                 22961880                64.04 ns/op            0 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/95%-reads-4                 17553278                82.95 ns/op            2 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/90%-reads-4                 15808310                98.18 ns/op            4 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/75%-reads-4                 10745935               118.0 ns/op             9 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/50%-reads-4                  8660215               133.6 ns/op            16 B/op          1 allocs/op
BenchmarkMap_Xsync_NoWarmUp/0%-reads-4                   7122831               168.6 ns/op            31 B/op          2 allocs/op
BenchmarkMap_CMAP_NoWarmUp/99%-reads-4                  12115374               104.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/95%-reads-4                   8256139               149.6 ns/op             1 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/90%-reads-4                   6976657               180.0 ns/op             2 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/75%-reads-4                   5394254               237.0 ns/op             5 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/50%-reads-4                   4419870               268.8 ns/op             7 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/0%-reads-4                    6944652               163.9 ns/op             4 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/99%-reads-4                6722534               179.8 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/95%-reads-4                4732837               264.9 ns/op             1 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/90%-reads-4                4137775               302.7 ns/op             3 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/75%-reads-4                3578094               350.6 ns/op             4 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/50%-reads-4                3206191               383.2 ns/op             9 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/0%-reads-4                 3705342               270.3 ns/op             8 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/99%-reads-4                13914631              1242 ns/op               0 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/95%-reads-4                 1000000              1145 ns/op               3 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/90%-reads-4                 1000000              7525 ns/op               7 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/99%-reads-4               19719538               472.0 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/95%-reads-4                3542832              1834 ns/op               1 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/90%-reads-4                1000000              1662 ns/op               4 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/99%-reads-4               3057364               476.0 ns/op            45 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/95%-reads-4               2360776               520.1 ns/op            46 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/90%-reads-4               2017683               576.9 ns/op            49 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/75%-reads-4               1885638               646.5 ns/op            50 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/50%-reads-4               1654629               696.5 ns/op            55 B/op          1 allocs/op
BenchmarkMap_Standard_NoWarmUp/0%-reads-4                1433965               762.5 ns/op            64 B/op          2 allocs/op
BenchmarkMap_Xsync_WarmUp/100%-reads-4                   8047791               149.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/99%-reads-4                    8190386               141.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/95%-reads-4                    8321588               139.7 ns/op             1 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/90%-reads-4                    8161870               140.5 ns/op             2 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/75%-reads-4                    7227441               145.5 ns/op             7 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/50%-reads-4                    7019108               147.6 ns/op            14 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/0%-reads-4                     6953701               170.0 ns/op            27 B/op          1 allocs/op
BenchmarkMap_CMAP_WarmUp/100%-reads-4                    7796668               147.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/99%-reads-4                     7463191               160.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/95%-reads-4                     5609403               205.6 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/90%-reads-4                     4712932               243.8 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/75%-reads-4                     3947751               280.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/50%-reads-4                     3633550               292.4 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/0%-reads-4                      6414294               168.6 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/100%-reads-4                 6787693               173.7 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/99%-reads-4                  4552660               254.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/95%-reads-4                  3130820               369.5 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/90%-reads-4                  2840918               417.8 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/75%-reads-4                  2572873               440.4 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/50%-reads-4                  2542430               451.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/0%-reads-4                   3922452               269.6 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/100%-reads-4                  8160154               147.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/99%-reads-4                     21225             49576 ns/op               1 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/95%-reads-4                      4381            270283 ns/op               7 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/100%-reads-4                4511062               226.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/99%-reads-4                 3677433               289.8 ns/op             8 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/95%-reads-4                 3485506               308.4 ns/op             9 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/90%-reads-4                 3477531               308.3 ns/op            10 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/75%-reads-4                 3209703               348.1 ns/op            14 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/50%-reads-4                 1684626               671.1 ns/op            37 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/0%-reads-4                  1497439               687.7 ns/op            33 B/op          1 allocs/op
BenchmarkMap_Xsync_Range-4                                   169           6955477 ns/op               2 B/op          0 allocs/op
BenchmarkMap_CMAP_Range-4                                     34          32698984 ns/op        24067286 B/op        199 allocs/op
BenchmarkMap_SafeMap_Range-4                                  25          40408270 ns/op        24020066 B/op         31 allocs/op
BenchmarkMap_HaxMap_Range-4                                  123           8804928 ns/op               2 B/op          0 allocs/op
BenchmarkMap_Standard_Range-4                                114           9769472 ns/op               3 B/op          0 allocs/op
```

```go
go test -bench=^BenchmarkMap -benchmem
goos: linux
goarch: amd64
pkg: bench
cpu: AMD EPYC 7K62 48-Core Processor
BenchmarkMap_Xsync_NoWarmUp/99%-reads-4                  5322466               242.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/95%-reads-4                  4893560               296.5 ns/op             3 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/90%-reads-4                  4304379               329.3 ns/op             5 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/75%-reads-4                  3536888               433.2 ns/op            12 B/op          0 allocs/op
BenchmarkMap_Xsync_NoWarmUp/50%-reads-4                  2332996               513.3 ns/op            22 B/op          1 allocs/op
BenchmarkMap_Xsync_NoWarmUp/0%-reads-4                   1909485               855.9 ns/op            39 B/op          2 allocs/op
BenchmarkMap_CMAP_NoWarmUp/99%-reads-4                   3784177               329.5 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/95%-reads-4                   3658107               349.7 ns/op             2 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/90%-reads-4                   3534452               370.7 ns/op             4 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/75%-reads-4                   3219830               394.4 ns/op             5 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/50%-reads-4                   2899605               438.4 ns/op            11 B/op          0 allocs/op
BenchmarkMap_CMAP_NoWarmUp/0%-reads-4                    2616900               452.6 ns/op            12 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/99%-reads-4                4374471               289.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/95%-reads-4                4137565               318.3 ns/op             1 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/90%-reads-4                3833136               335.2 ns/op             4 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/75%-reads-4                3554889               363.9 ns/op             4 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/50%-reads-4                3212438               394.8 ns/op             9 B/op          0 allocs/op
BenchmarkMap_SafeMap_NoWarmUp/0%-reads-4                 2651151               411.2 ns/op            12 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/99%-reads-4                 3608038               786.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/95%-reads-4                 1000000              4521 ns/op               3 B/op          0 allocs/op
BenchmarkMap_HaxMap_NoWarmUp/90%-reads-4                 1000000             20939 ns/op               6 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/99%-reads-4                4312915               429.4 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/95%-reads-4                1000000              1323 ns/op               2 B/op          0 allocs/op
BenchmarkMap_HashMap_NoWarmUp/90%-reads-4                1000000              5824 ns/op               4 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/99%-reads-4               2943702               440.0 ns/op            46 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/95%-reads-4               2427459               567.7 ns/op            48 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/90%-reads-4               2124974               626.8 ns/op            49 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/75%-reads-4               1827462               695.9 ns/op            49 B/op          0 allocs/op
BenchmarkMap_Standard_NoWarmUp/50%-reads-4               1648512               858.8 ns/op            61 B/op          1 allocs/op
BenchmarkMap_Standard_NoWarmUp/0%-reads-4                1375279               780.8 ns/op            58 B/op          2 allocs/op
BenchmarkMap_Xsync_WarmUp/100%-reads-4                   2032300               595.7 ns/op             0 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/99%-reads-4                    1994565               614.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/95%-reads-4                    1994385               598.5 ns/op             1 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/90%-reads-4                    2020077               590.4 ns/op             2 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/75%-reads-4                    1965814               588.9 ns/op             7 B/op          0 allocs/op
BenchmarkMap_Xsync_WarmUp/50%-reads-4                    1892293               674.5 ns/op            14 B/op          1 allocs/op
BenchmarkMap_Xsync_WarmUp/0%-reads-4                     1542258               765.0 ns/op            28 B/op          2 allocs/op
BenchmarkMap_CMAP_WarmUp/100%-reads-4                    2186379               544.8 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/99%-reads-4                     2197336               549.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/95%-reads-4                     2179720               540.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/90%-reads-4                     2211510               538.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/75%-reads-4                     2228772               522.7 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/50%-reads-4                     2215503               504.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_CMAP_WarmUp/0%-reads-4                      2305758               486.7 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/100%-reads-4                 2296856               516.5 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/99%-reads-4                  2203132               519.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/95%-reads-4                  2351454               512.8 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/90%-reads-4                  2318145               505.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/75%-reads-4                  2251226               489.7 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/50%-reads-4                  2404495               472.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap_SafeMap_WarmUp/0%-reads-4                   2540269               444.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/100%-reads-4                  1894240               611.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/99%-reads-4                      4152            321617 ns/op               5 B/op          0 allocs/op
BenchmarkMap_HaxMap_WarmUp/95%-reads-4                       633           1712095 ns/op              37 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/100%-reads-4                1803691               655.0 ns/op             0 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/99%-reads-4                 1396999               801.8 ns/op            20 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/95%-reads-4                 1403974               793.6 ns/op            21 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/90%-reads-4                 1396458              1168 ns/op              22 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/75%-reads-4                 1398770               795.7 ns/op            25 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/50%-reads-4                 1420707               806.4 ns/op            43 B/op          0 allocs/op
BenchmarkMap_Standard_WarmUp/0%-reads-4                  1435394               776.3 ns/op            34 B/op          1 allocs/op
BenchmarkMap_Xsync_Range-4                                    38          28972117 ns/op               9 B/op          0 allocs/op
BenchmarkMap_CMAP_Range-4                                     22          76655087 ns/op        24068915 B/op        203 allocs/op
BenchmarkMap_SafeMap_Range-4                                  21          73158401 ns/op        24020057 B/op         31 allocs/op
BenchmarkMap_HaxMap_Range-4                                   21          49680062 ns/op              17 B/op          0 allocs/op
BenchmarkMap_Standard_Range-4                                 26          43018142 ns/op              14 B/op          0 allocs/op
```

## Benchmark Integer

```go
go test -bench=^BenchmarkRead -benchmem
goos: linux
goarch: amd64
pkg: bench
cpu: Intel(R) Xeon(R) Silver 4314 CPU @ 2.40GHz
BenchmarkReadHashMapUint-64                      3179492               353.7 ns/op             0 B/op          0 allocs/op
BenchmarkReadHashMapWithWritesUint-64            2750632               417.0 ns/op             5 B/op          0 allocs/op
BenchmarkReadHashMapString-64                    1535917               759.4 ns/op             0 B/op          0 allocs/op
BenchmarkReadHaxMapUint-64                       3118179               364.2 ns/op             0 B/op          0 allocs/op
BenchmarkReadHaxMapWithWritesUint-64             2758036               414.3 ns/op             4 B/op          0 allocs/op
BenchmarkReadXsyncMapUint-64                     2003343               578.5 ns/op             0 B/op          0 allocs/op
BenchmarkReadSkipMapUint-64                       827330              1311 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoMapUintUnsafe-64                  2780418               408.8 ns/op             0 B/op          0 allocs/op
BenchmarkReadGoMapUintMutex-64                     21457             55038 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoMapWithWritesUintMutex-64            8019            213486 ns/op               1 B/op          0 allocs/op
BenchmarkReadGoSyncMapUint-64                    1025775              1124 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoSyncMapWithWritesUint-64           966132              1204 ns/op              37 B/op          3 allocs/op
BenchmarkReadGoMapStringUnsafe-64                2019312               572.5 ns/op             0 B/op          0 allocs/op
BenchmarkReadGoMapStringMutex-64                   18313             64991 ns/op               1 B/op          0 allocs/op
```

```go
go test -bench=^BenchmarkRead -benchmem
goos: linux
goarch: amd64
pkg: bench
cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
BenchmarkReadHashMapUint-4                        363864              3327 ns/op               0 B/op          0 allocs/op
BenchmarkReadHashMapWithWritesUint-4              242193              4963 ns/op             559 B/op         69 allocs/op
BenchmarkReadHashMapString-4                      197773              6077 ns/op               0 B/op          0 allocs/op
BenchmarkReadHaxMapUint-4                         361154              3556 ns/op               0 B/op          0 allocs/op
BenchmarkReadHaxMapWithWritesUint-4               236191              5324 ns/op             486 B/op         60 allocs/op
BenchmarkReadXsyncMapUint-4                       206422              5454 ns/op               0 B/op          0 allocs/op
BenchmarkReadSkipMapUint-4                         80217             13800 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoMapUintUnsafe-4                    281607              4386 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoMapUintMutex-4                      26410             45375 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoMapWithWritesUintMutex-4             6381            174255 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoSyncMapUint-4                      123511              9804 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoSyncMapWithWritesUint-4             91726             12426 ns/op            2960 B/op        264 allocs/op
BenchmarkReadGoMapStringUnsafe-4                  162080              7298 ns/op               0 B/op          0 allocs/op
BenchmarkReadGoMapStringMutex-4                    25834             46309 ns/op               0 B/op          0 allocs/op
```



