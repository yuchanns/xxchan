// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package xxchan_test

import (
	"fmt"
	"testing"

	"github.com/smasher164/mem"
	"go.yuchanns.xyz/xxchan"
)

// Benchmark configurations
var benchmarkSizes = []int{10, 100, 1000, 10000}

// benchmarkXXChanPush benchmarks xxchan Push operations
func benchmarkXXChanPush(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("cap-%d", size), func(b *testing.B) {
			ptr := mem.Alloc(uint(xxchan.Sizeof[int](size)))
			defer mem.Free(ptr)

			ch := xxchan.Make[int](ptr, size)

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				if !ch.Push(i) {
					// Channel full, pop one item to make space
					ch.Pop()
					ch.Push(i)
				}
			}
		})
	}
}

// benchmarkBuiltinChanPush benchmarks built-in channel send operations
func benchmarkBuiltinChanPush(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("cap-%d", size), func(b *testing.B) {
			ch := make(chan int, size)

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				select {
				case ch <- i:
				default:
					// Channel full, receive one item to make space
					<-ch
					ch <- i
				}
			}
		})
	}
}

// benchmarkXXChanPop benchmarks xxchan Pop operations
func benchmarkXXChanPop(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("cap-%d", size), func(b *testing.B) {
			ptr := mem.Alloc(uint(xxchan.Sizeof[int](size)))
			defer mem.Free(ptr)

			ch := xxchan.Make[int](ptr, size)

			// Pre-fill the channel
			for i := range size {
				ch.Push(i)
			}

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				if val, ok := ch.Pop(); !ok {
					// Channel empty, push one item
					ch.Push(i)
					ch.Pop()
				} else {
					// Push the value back to maintain channel state
					ch.Push(val)
				}
			}
		})
	}
}

// benchmarkBuiltinChanPop benchmarks built-in channel receive operations
func benchmarkBuiltinChanPop(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("cap-%d", size), func(b *testing.B) {
			ch := make(chan int, size)

			// Pre-fill the channel
			for i := range size {
				ch <- i
			}

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				select {
				case val := <-ch:
					// Put the value back to maintain channel state
					ch <- val
				default:
					// Channel empty, send one item
					ch <- i
					<-ch
				}
			}
		})
	}
}

// benchmarkXXChanMixed benchmarks mixed Push/Pop operations for xxchan
func benchmarkXXChanMixed(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("cap-%d", size), func(b *testing.B) {
			ptr := mem.Alloc(uint(xxchan.Sizeof[int](size)))
			defer mem.Free(ptr)

			ch := xxchan.Make[int](ptr, size)

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				if i%2 == 0 {
					ch.Push(i)
				} else {
					ch.Pop()
				}
			}
		})
	}
}

// benchmarkBuiltinChanMixed benchmarks mixed send/receive operations for built-in channels
func benchmarkBuiltinChanMixed(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("cap-%d", size), func(b *testing.B) {
			ch := make(chan int, size)

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				if i%2 == 0 {
					select {
					case ch <- i:
					default:
					}
				} else {
					select {
					case <-ch:
					default:
					}
				}
			}
		})
	}
}

// BenchmarkChannelCreation benchmarks channel creation overhead
func benchmarkXXChanCreation(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("cap-%d", size), func(b *testing.B) {
			b.ReportAllocs()

			for b.Loop() {
				ptr := mem.Alloc(uint(xxchan.Sizeof[int](size)))
				ch := xxchan.Make[int](ptr, size)
				_ = ch
				mem.Free(ptr)
			}
		})
	}
}

func benchmarkBuiltinChanCreation(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("cap-%d", size), func(b *testing.B) {
			b.ReportAllocs()

			for b.Loop() {
				ch := make(chan int, size)
				_ = ch
			}
		})
	}
}

// BenchmarkConcurrentAccess benchmarks concurrent access patterns
func benchmarkXXChanConcurrent(b *testing.B) {
	const size = 1000
	const numWorkers = 4

	ptr := mem.Alloc(uint(xxchan.Sizeof[int](size)))
	defer mem.Free(ptr)

	ch := xxchan.Make[int](ptr, size)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%2 == 0 {
				ch.Push(i)
			} else {
				ch.Pop()
			}
			i++
		}
	})
}

func benchmarkBuiltinChanConcurrent(b *testing.B) {
	const size = 1000
	ch := make(chan int, size)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%2 == 0 {
				select {
				case ch <- i:
				default:
				}
			} else {
				select {
				case <-ch:
				default:
				}
			}
			i++
		}
	})
}

// Benchmark helper to run all benchmarks
func BenchmarkAll(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(*testing.B)
	}{
		{"XXChan/Push", benchmarkXXChanPush},
		{"Builtin/Push", benchmarkBuiltinChanPush},
		{"XXChan/Pop", benchmarkXXChanPop},
		{"Builtin/Pop", benchmarkBuiltinChanPop},
		{"XXChan/Mixed", benchmarkXXChanMixed},
		{"Builtin/Mixed", benchmarkBuiltinChanMixed},
		{"XXChan/Creation", benchmarkXXChanCreation},
		{"Builtin/Creation", benchmarkBuiltinChanCreation},
		{"XXChan/Concurrent", benchmarkXXChanConcurrent},
		{"Builtin/Concurrent", benchmarkBuiltinChanConcurrent},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, bm.fn)
	}
}
