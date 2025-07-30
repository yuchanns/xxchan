# xxchan

[![Go Reference](https://pkg.go.dev/badge/go.yuchanns.xyz/xxchan)](https://pkg.go.dev/go.yuchanns.xyz/xxchan)

A garbage collection-free channel implementation for Go that operates on user-allocated memory.

## Overview

`xxchan` provides a `Channel[T]` type that functions similarly to Go's built-in channels but without internal memory allocation. It uses a user-provided memory block to store data, making it suitable for applications where garbage collection overhead needs to be minimized.

## Features

- **Zero garbage collection**: Operates on pre-allocated memory blocks
- **Thread-safe**: Supports concurrent access from multiple goroutines
- **Generic**: Works with any Go type using generics
- **Simple API**: Provides basic channel operations (Push, Pop, Len, Cap)
- **Ring buffer**: Efficient memory usage with O(1) operations

<details>
<summary>Benchmark</summary>

```bash
❯ go test -bench=BenchmarkAll -benchmem -count=6 -run=^$ -v
goos: darwin
goarch: arm64
pkg: go.yuchanns.xyz/xxchan
cpu: Apple M1 Max
BenchmarkAll
BenchmarkAll/XXChan/Push
BenchmarkAll/XXChan/Push/cap-10
BenchmarkAll/XXChan/Push/cap-10-10      41307733                29.36 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-10      40870600                29.67 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-10      40142505                29.69 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-10      40990839                29.53 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-10      40424229                29.75 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-10      40561951                29.55 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100
BenchmarkAll/XXChan/Push/cap-100-10     40971303                29.35 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-10     40208356                29.50 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-10     40453302                29.28 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-10     40130200                29.40 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-10     40281579                29.76 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-10     40613947                29.54 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000
BenchmarkAll/XXChan/Push/cap-1000-10            40477123                29.35 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-10            40142562                29.37 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-10            40425081                28.95 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-10            40512709                29.67 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-10            40719720                29.69 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-10            40573266                29.36 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000
BenchmarkAll/XXChan/Push/cap-10000-10           40412656                29.73 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-10           40389590                29.72 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-10           40580698                29.68 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-10           40631422                29.73 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-10           40505359                29.27 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-10           40733024                29.70 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push
BenchmarkAll/Builtin/Push/cap-10
BenchmarkAll/Builtin/Push/cap-10-10             45300612                26.14 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-10             46183968                25.96 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-10             45603316                25.94 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-10             46196043                26.17 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-10             45608660                26.22 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-10             45462868                26.13 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100
BenchmarkAll/Builtin/Push/cap-100-10            46338979                25.54 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-10            46107738                25.49 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-10            46467051                25.51 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-10            46327574                25.65 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-10            46364565                25.28 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-10            46820283                25.12 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000
BenchmarkAll/Builtin/Push/cap-1000-10           46590706                25.63 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-10           46371882                25.54 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-10           46350090                25.26 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-10           47157993                25.15 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-10           47724074                25.12 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-10           46177600                25.40 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000
BenchmarkAll/Builtin/Push/cap-10000-10          47401633                25.58 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-10          46626538                24.82 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-10          42185068                25.45 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-10          46926019                25.62 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-10          47030926                25.63 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-10          46865083                25.31 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop
BenchmarkAll/XXChan/Pop/cap-10
BenchmarkAll/XXChan/Pop/cap-10-10               59240242                19.65 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-10               61409475                19.50 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-10               60945154                19.66 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-10               59723652                19.99 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-10               59253040                20.14 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-10               60282195                20.00 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100
BenchmarkAll/XXChan/Pop/cap-100-10              59529345                19.95 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-10              59826874                19.82 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-10              60053798                19.94 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-10              59715975                19.62 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-10              59260233                19.84 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-10              59713622                19.76 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000
BenchmarkAll/XXChan/Pop/cap-1000-10             59772368                19.49 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-10             59857962                20.02 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-10             59461993                19.76 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-10             60522127                19.95 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-10             60136434                20.05 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-10             59377647                20.01 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000
BenchmarkAll/XXChan/Pop/cap-10000-10            59233298                20.01 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-10            60308325                19.87 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-10            59578605                19.93 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-10            59606101                20.04 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-10            61017078                19.95 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-10            60250792                19.80 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop
BenchmarkAll/Builtin/Pop/cap-10
BenchmarkAll/Builtin/Pop/cap-10-10              45709496                26.09 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-10              45556068                26.17 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-10              46061943                26.21 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-10              45826959                26.08 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-10              46122062                27.70 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-10              45531435                26.16 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100
BenchmarkAll/Builtin/Pop/cap-100-10             47495438                25.25 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-10             46792442                25.47 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-10             47261072                24.88 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-10             47685434                25.43 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-10             46851894                25.41 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-10             47796226                25.27 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000
BenchmarkAll/Builtin/Pop/cap-1000-10            47420286                24.96 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-10            47418957                25.18 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-10            47292892                25.06 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-10            46853114                25.32 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-10            47383305                25.39 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-10            47450991                24.85 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000
BenchmarkAll/Builtin/Pop/cap-10000-10           47080670                24.72 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-10           47167414                25.22 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-10           46930530                24.66 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-10           47493088                25.07 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-10           47435203                24.90 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-10           47217992                24.64 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed
BenchmarkAll/XXChan/Mixed/cap-10
BenchmarkAll/XXChan/Mixed/cap-10-10             100000000               10.29 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-10             100000000               10.24 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-10             100000000               10.36 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-10             100000000               10.43 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-10             100000000               10.16 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-10             100000000               10.32 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100
BenchmarkAll/XXChan/Mixed/cap-100-10            100000000               10.33 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-10            100000000               10.44 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-10            100000000               10.37 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-10            100000000               10.33 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-10            100000000               10.29 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-10            100000000               10.37 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000
BenchmarkAll/XXChan/Mixed/cap-1000-10           100000000               10.14 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-10           100000000               10.25 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-10           100000000               10.26 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-10           100000000               10.40 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-10           100000000               10.07 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-10           100000000               10.29 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000
BenchmarkAll/XXChan/Mixed/cap-10000-10          100000000               10.43 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-10          100000000               10.42 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-10          100000000               10.32 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-10          100000000               10.16 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-10          100000000               10.37 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-10          100000000               10.41 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed
BenchmarkAll/Builtin/Mixed/cap-10
BenchmarkAll/Builtin/Mixed/cap-10-10            90709234                12.96 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-10            90069804                13.07 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-10            89351643                13.26 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-10            89813636                13.24 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-10            90658693                13.19 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-10            91336948                13.21 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100
BenchmarkAll/Builtin/Mixed/cap-100-10           93760071                12.79 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-10           92766170                12.85 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-10           93621693                12.74 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-10           92706150                12.69 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-10           93283886                12.75 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-10           92852910                12.83 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000
BenchmarkAll/Builtin/Mixed/cap-1000-10          93237375                12.79 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-10          95101917                12.78 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-10          93882638                12.83 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-10          93752438                12.80 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-10          93015439                12.82 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-10          92301777                12.73 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000
BenchmarkAll/Builtin/Mixed/cap-10000-10         93597961                12.79 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-10         93086098                12.81 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-10         93692356                12.83 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-10         93168606                12.74 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-10         92438648                12.84 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-10         93101143                12.61 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation
BenchmarkAll/XXChan/Creation/cap-10
BenchmarkAll/XXChan/Creation/cap-10-10            361184              3037 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-10            351102              3004 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-10            339495              2969 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-10            354002              3054 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-10            371895              3068 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-10            350430              3009 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100
BenchmarkAll/XXChan/Creation/cap-100-10           431234              3039 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-10           380083              3008 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-10           385603              3038 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-10           408903              3024 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-10           383323              3122 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-10           370524              3158 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000
BenchmarkAll/XXChan/Creation/cap-1000-10          417492              3161 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-10          327760              3133 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-10          420300              3172 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-10          379503              3142 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-10          413883              3148 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-10          361140              3155 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000
BenchmarkAll/XXChan/Creation/cap-10000-10         294318              4049 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-10         297874              4083 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-10         304057              4042 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-10         318649              4025 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-10         298657              4055 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-10         283863              4028 ns/op               0 B/op          0 allocs/op
BenchmarkAll/Builtin/Creation
BenchmarkAll/Builtin/Creation/cap-10
BenchmarkAll/Builtin/Creation/cap-10-10         16586995                60.56 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-10         22746699                52.49 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-10         16674964                66.61 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-10         16254196                62.38 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-10         22525911                55.71 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-10         23313691                53.27 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100
BenchmarkAll/Builtin/Creation/cap-100-10         7151137               158.3 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-10         6831030               166.5 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-10         6899384               160.5 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-10         6960453               144.7 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-10         6531861               163.5 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-10         7064332               168.7 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000
BenchmarkAll/Builtin/Creation/cap-1000-10        1342399               940.9 ns/op          8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-10        1406841               853.4 ns/op          8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-10        1366630               865.2 ns/op          8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-10        1433420               827.3 ns/op          8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-10        1000000              1005 ns/op            8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-10        1531189               807.8 ns/op          8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000
BenchmarkAll/Builtin/Creation/cap-10000-10        250425              4777 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-10        240328              5211 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-10        214100              6214 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-10        248594              4868 ns/op           81920 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-10        221650              4831 ns/op           81920 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-10        214250              4689 ns/op           81920 B/op          1 allocs/op
BenchmarkAll/XXChan/Concurrent
BenchmarkAll/XXChan/Concurrent-10               83892108                16.03 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-10               78921309                14.16 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-10               82054827                14.67 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-10               83747055                15.65 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-10               67477646                15.28 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-10               78525037                14.58 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent
BenchmarkAll/Builtin/Concurrent-10              67770608                19.29 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-10              64118120                19.28 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-10              67221396                19.25 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-10              67406890                19.06 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-10              66272403                19.00 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-10              67154780                18.52 ns/op            0 B/op          0 allocs/op
PASS
ok      go.yuchanns.xyz/xxchan  242.563s
```

</details>

## Installation

```bash
go get go.yuchanns.xyz/xxchan
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "unsafe"
    
    "go.yuchanns.xyz/xxchan"
)

func main() {
    // Calculate required memory size for 10 integers
    size := xxchan.Sizeof[int](10)
    
    // Allocate memory
    buf := make([]byte, size)
    
    // Create channel
    ch := xxchan.Make[int](unsafe.Pointer(&buf[0]), 10)
    
    // Push values
    ch.Push(1)
    ch.Push(2)
    ch.Push(3)
    
    fmt.Printf("Length: %d, Capacity: %d\n", ch.Len(), ch.Cap())
    // Output: Length: 3, Capacity: 10
    
    // Pop values
    for ch.Len() > 0 {
        if val, ok := ch.Pop(); ok {
            fmt.Printf("Popped: %d\n", val)
        }
    }
}
```

### Using with Memory Pools

```go
package main

import (
    "sync"
    "unsafe"
    
    "go.yuchanns.xyz/xxchan"
)

var bufferPool = sync.Pool{
    New: func() interface{} {
        size := xxchan.Sizeof[int](100)
        return make([]byte, size)
    },
}

func createChannel() *xxchan.Channel[int] {
    buf := bufferPool.Get().([]byte)
    return xxchan.Make[int](unsafe.Pointer(&buf[0]), 100)
}

func releaseChannel(ch *xxchan.Channel[int]) {
    // Calculate buffer from channel pointer
    size := xxchan.Sizeof[int](100)
    buf := (*[1]byte)(unsafe.Pointer(ch))[:size:size]
    bufferPool.Put(buf)
}
```

### Stack Allocation

```go
func processData() {
    // Allocate on stack (for small channels)
    var buf [1024]byte // Adjust size as needed
    
    capacity := 50 // Ensure buf is large enough
    if len(buf) >= xxchan.Sizeof[int](capacity) {
        ch := xxchan.Make[int](unsafe.Pointer(&buf[0]), capacity)
        
        // Use channel...
        ch.Push(42)
        val, ok := ch.Pop()
        _ = val
        _ = ok
    }
}
```

## Memory Management

Users are responsible for:
- Allocating sufficient memory using `Sizeof[T](n)`
- Ensuring memory remains valid during the channel's lifetime
- Proper memory alignment (handled automatically by `Make`)

The channel does not allocate or free memory internally.

## Thread Safety

All operations are thread-safe and can be called concurrently from multiple goroutines. The implementation uses atomic operations with spin-lock synchronization.

## Performance Characteristics

- **Time Complexity**: All operations are O(1)
- **Space Complexity**: O(n) where n is the channel capacity
- **Synchronization**: Spin-lock with microsecond delays

## Limitations

- Non-blocking operations only (no blocking Push/Pop like Go channels)
- Requires `unsafe` package usage
- Manual memory management
- Fixed capacity (cannot be resized after creation)

## License

Licensed under the Apache License, Version 2.0.

## Contributing

Contributions are welcome. Please ensure that any changes maintain the zero-allocation guarantee and thread-safety properties.
