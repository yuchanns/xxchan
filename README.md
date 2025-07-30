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

```
❯ go test -bench=BenchmarkAll -benchmem -count=6 -run=^$$ -v
goos: linux
goarch: amd64
pkg: go.yuchanns.xyz/xxchan
cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
BenchmarkAll
BenchmarkAll/XXChan/Push
BenchmarkAll/XXChan/Push/cap-10
BenchmarkAll/XXChan/Push/cap-10-12      24244605                48.91 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-12      24467304                48.77 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-12      24032360                48.87 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-12      24407862                49.33 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-12      24648841                48.83 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10-12      23947179                48.70 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100
BenchmarkAll/XXChan/Push/cap-100-12     24078735                48.73 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-12     24275188                48.97 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-12     20957192                49.79 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-12     24328038                48.85 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-12     23849096                49.10 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-100-12     23639869                49.45 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000
BenchmarkAll/XXChan/Push/cap-1000-12            24371706                49.02 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-12            23940170                48.86 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-12            23994022                48.69 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-12            24000056                48.83 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-12            24436834                48.74 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-1000-12            23956015                49.05 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000
BenchmarkAll/XXChan/Push/cap-10000-12           24147483                49.00 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-12           22864966                49.89 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-12           23900118                49.41 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-12           23872981                49.31 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-12           24420098                49.43 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Push/cap-10000-12           24406305                49.05 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push
BenchmarkAll/Builtin/Push/cap-10
BenchmarkAll/Builtin/Push/cap-10-12             26828751                42.74 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-12             27666813                42.56 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-12             26875081                44.23 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-12             27555166                43.01 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-12             27798236                43.52 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10-12             27816698                43.05 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100
BenchmarkAll/Builtin/Push/cap-100-12            27816471                43.03 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-12            27864440                43.94 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-12            27468655                43.06 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-12            28045586                42.60 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-12            26287956                42.64 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-100-12            27165517                43.18 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000
BenchmarkAll/Builtin/Push/cap-1000-12           28018944                42.64 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-12           23487864                43.93 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-12           27599203                42.74 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-12           27331263                42.72 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-12           26352759                43.95 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-1000-12           27321699                43.21 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000
BenchmarkAll/Builtin/Push/cap-10000-12          27027675                43.48 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-12          26242766                45.10 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-12          26849865                44.93 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-12          24220219                45.81 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-12          26229501                43.78 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Push/cap-10000-12          27839226                43.32 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop
BenchmarkAll/XXChan/Pop/cap-10
BenchmarkAll/XXChan/Pop/cap-10-12               32492216                36.29 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-12               32583741                36.76 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-12               31808386                37.26 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-12               31867864                37.30 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-12               31216932                36.74 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10-12               31940875                36.20 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100
BenchmarkAll/XXChan/Pop/cap-100-12              31992528                36.57 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-12              32867726                36.37 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-12              31768370                36.19 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-12              30691387                36.47 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-12              32173386                36.44 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-100-12              31391517                36.45 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000
BenchmarkAll/XXChan/Pop/cap-1000-12             32566542                36.14 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-12             32496032                36.42 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-12             28080708                36.20 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-12             31085019                36.27 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-12             31542200                36.18 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-1000-12             32152510                36.36 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000
BenchmarkAll/XXChan/Pop/cap-10000-12            32247908                36.45 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-12            32093013                36.29 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-12            32082722                36.17 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-12            31343571                36.63 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-12            32074370                36.29 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Pop/cap-10000-12            32138822                36.18 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop
BenchmarkAll/Builtin/Pop/cap-10
BenchmarkAll/Builtin/Pop/cap-10-12              26821446                43.91 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-12              26635228                44.19 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-12              27332872                43.71 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-12              26357892                44.01 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-12              26387878                43.98 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10-12              26703819                43.91 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100
BenchmarkAll/Builtin/Pop/cap-100-12             22816039                45.18 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-12             26635842                44.60 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-12             26486952                43.65 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-12             26419364                43.56 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-12             26691240                44.05 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-100-12             25971820                43.79 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000
BenchmarkAll/Builtin/Pop/cap-1000-12            26488880                43.79 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-12            22729364                44.36 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-12            26991532                43.46 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-12            27107428                44.14 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-12            24619338                44.01 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-1000-12            25809924                44.92 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000
BenchmarkAll/Builtin/Pop/cap-10000-12           26436820                43.46 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-12           26383824                43.81 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-12           26081811                43.92 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-12           26987203                44.39 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-12           26216955                43.92 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Pop/cap-10000-12           26604466                43.39 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed
BenchmarkAll/XXChan/Mixed/cap-10
BenchmarkAll/XXChan/Mixed/cap-10-12             61012465                18.79 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-12             62543203                18.74 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-12             58055218                19.37 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-12             58952264                18.80 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-12             58547782                19.45 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10-12             60478188                18.82 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100
BenchmarkAll/XXChan/Mixed/cap-100-12            59162498                19.14 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-12            60983960                18.81 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-12            61070022                18.67 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-12            61365847                18.78 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-12            63467029                18.74 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-100-12            61855890                18.73 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000
BenchmarkAll/XXChan/Mixed/cap-1000-12           61808678                19.04 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-12           61141629                18.72 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-12           58656177                19.24 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-12           61540680                18.73 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-12           59384338                19.81 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-1000-12           62607648                19.15 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000
BenchmarkAll/XXChan/Mixed/cap-10000-12          50930523                19.85 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-12          60723160                18.84 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-12          61270882                18.78 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-12          60489788                18.85 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-12          62079710                19.22 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Mixed/cap-10000-12          61623612                18.87 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed
BenchmarkAll/Builtin/Mixed/cap-10
BenchmarkAll/Builtin/Mixed/cap-10-12            54956883                21.52 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-12            55520859                21.47 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-12            52592210                21.48 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-12            51254076                21.52 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-12            53502679                21.37 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10-12            52443194                21.49 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100
BenchmarkAll/Builtin/Mixed/cap-100-12           54034988                21.57 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-12           53704011                21.81 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-12           55302504                21.64 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-12           53161225                21.44 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-12           53608077                21.47 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-100-12           53709132                21.93 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000
BenchmarkAll/Builtin/Mixed/cap-1000-12          53539777                22.12 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-12          54736066                22.90 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-12          53578413                22.30 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-12          53141764                22.19 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-12          53168060                22.69 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-1000-12          53134766                23.66 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000
BenchmarkAll/Builtin/Mixed/cap-10000-12         51518224                22.70 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-12         52581847                22.65 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-12         53237330                22.15 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-12         54222026                22.27 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-12         52900254                22.33 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Mixed/cap-10000-12         52643455                22.66 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation
BenchmarkAll/XXChan/Creation/cap-10
BenchmarkAll/XXChan/Creation/cap-10-12            347412              3442 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-12            333754              3492 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-12            340747              3645 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-12            290712              3665 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-12            326341              3630 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10-12            338204              3631 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100
BenchmarkAll/XXChan/Creation/cap-100-12           285386              3517 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-12           316986              3547 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-12           304924              3963 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-12           325920              3507 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-12           343802              3484 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-100-12           347139              3593 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000
BenchmarkAll/XXChan/Creation/cap-1000-12          252714              4427 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-12          275679              4532 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-12          265574              4489 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-12          266692              4681 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-12          259570              4662 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-1000-12          259284              4568 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000
BenchmarkAll/XXChan/Creation/cap-10000-12         214192              6041 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-12         193544              5781 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-12         208333              5803 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-12         172652              5921 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-12         209158              5936 ns/op               0 B/op          0 allocs/op
BenchmarkAll/XXChan/Creation/cap-10000-12         205629              5806 ns/op               0 B/op          0 allocs/op
BenchmarkAll/Builtin/Creation
BenchmarkAll/Builtin/Creation/cap-10
BenchmarkAll/Builtin/Creation/cap-10-12         29157597                38.21 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-12         30608451                38.13 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-12         32404761                41.60 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-12         31856515                36.40 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-12         33811329                36.70 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10-12         31783866                42.92 ns/op          192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100
BenchmarkAll/Builtin/Creation/cap-100-12         8784104               141.3 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-12         8162587               140.9 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-12         8676805               146.9 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-12         7113636               147.7 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-12         7898053               144.2 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-100-12         8593345               147.9 ns/op          1024 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000
BenchmarkAll/Builtin/Creation/cap-1000-12        1000000              1063 ns/op            8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-12        1000000              1034 ns/op            8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-12        1000000              1018 ns/op            8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-12        1000000              1044 ns/op            8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-12        1000000              1070 ns/op            8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-1000-12         945720              1066 ns/op            8192 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000
BenchmarkAll/Builtin/Creation/cap-10000-12        174210              6497 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-12        180828              6478 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-12        174102              6585 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-12        166633              6595 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-12        187765              6642 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/Builtin/Creation/cap-10000-12        173582              6632 ns/op           81921 B/op          1 allocs/op
BenchmarkAll/XXChan/Concurrent
BenchmarkAll/XXChan/Concurrent-12               45012202                24.88 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-12               47263854                24.62 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-12               47051287                24.65 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-12               43281940                25.12 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-12               47791260                25.18 ns/op            0 B/op          0 allocs/op
BenchmarkAll/XXChan/Concurrent-12               43464128                25.68 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent
BenchmarkAll/Builtin/Concurrent-12              36834192                33.34 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-12              36742482                32.13 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-12              36376017                32.55 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-12              35009757                32.66 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-12              36773704                32.73 ns/op            0 B/op          0 allocs/op
BenchmarkAll/Builtin/Concurrent-12              36169780                32.66 ns/op            0 B/op          0 allocs/op
PASS
ok      go.yuchanns.xyz/xxchan  237.995s

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
