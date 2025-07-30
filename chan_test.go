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
	"sync"
	"testing"

	"github.com/smasher164/mem"
	"github.com/stretchr/testify/require"
	"go.yuchanns.xyz/xxchan"
)

func TestChannelFull(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	n := 10
	length := n / 2

	ptr := mem.Alloc(uint(xxchan.Sizeof[int](length)))
	t.Cleanup(func() {
		mem.Free(ptr)
	})

	ch := xxchan.Make[int](ptr, length)
	assert.NotNil(ch)
	assert.Equal(length, ch.Cap())
	assert.Equal(0, ch.Len())

	for i := range length {
		assert.True(ch.Push(i))
	}
	assert.False(ch.Push(length))
	assert.Equal(length, ch.Len())
	for i := range length {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(i, val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)

	for i := range length {
		assert.True(ch.Push(i + length))
	}
	assert.Equal(length, ch.Len())
	for i := range length {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(i+length, val)
	}
	assert.Equal(0, ch.Len())
}

func TestChannelConcurrent(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	n := 10

	ptr := mem.Alloc(uint(xxchan.Sizeof[int](n)))
	t.Cleanup(func() {
		mem.Free(ptr)
	})

	ch := xxchan.Make[int](ptr, n)

	assert.NotNil(ch)

	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())

	wg := &sync.WaitGroup{}

	c := make(chan struct{}, n)
	for i := range n {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			assert.True(ch.Push(i))
			c <- struct{}{}
		}(i)
	}

	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()

			<-c
			_, ok := ch.Pop()
			assert.True(ok)
		}()
	}

	wg.Wait()
}

func TestChannelInt32(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	n := 10

	ptr := mem.Alloc(uint(xxchan.Sizeof[int32](n)))
	t.Cleanup(func() {
		mem.Free(ptr)
	})

	ch := xxchan.Make[int32](ptr, n)

	assert.NotNil(ch)

	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())

	for i := range n {
		assert.True(ch.Push(int32(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(int32(n)))

	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(int32(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelInt(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[int](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[int](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(i))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(n))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(i, val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelInt8(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[int8](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[int8](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(int8(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(int8(n)))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(int8(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelInt16(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[int16](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[int16](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(int16(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(int16(n)))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(int16(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelInt64(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[int64](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[int64](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(int64(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(int64(n)))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(int64(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelUint(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[uint](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[uint](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(uint(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(uint(n)))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(uint(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelUint8(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[uint8](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[uint8](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(uint8(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(uint8(n)))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(uint8(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelUint16(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[uint16](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[uint16](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(uint16(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(uint16(n)))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(uint16(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelUint32(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[uint32](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[uint32](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(uint32(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(uint32(n)))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(uint32(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelUint64(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[uint64](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[uint64](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(uint64(i)))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(uint64(n)))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(uint64(i), val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelFloat32(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[float32](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[float32](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(float32(i) + 0.5))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(float32(n) + 0.5))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(float32(i)+0.5, val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelFloat64(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 10
	ptr := mem.Alloc(uint(xxchan.Sizeof[float64](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[float64](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	for i := range n {
		assert.True(ch.Push(float64(i) + 0.5))
	}
	assert.Equal(n, ch.Len())
	assert.False(ch.Push(float64(n) + 0.5))
	for i := range n {
		val, ok := ch.Pop()
		assert.True(ok)
		assert.Equal(float64(i)+0.5, val)
	}
	assert.Equal(0, ch.Len())
	_, ok := ch.Pop()
	assert.False(ok)
}

func TestChannelBool(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 2
	ptr := mem.Alloc(uint(xxchan.Sizeof[bool](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[bool](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	assert.True(ch.Push(true))
	assert.True(ch.Push(false))
	assert.False(ch.Push(true))
	val, ok := ch.Pop()
	assert.True(ok)
	assert.Equal(true, val)
	val, ok = ch.Pop()
	assert.True(ok)
	assert.Equal(false, val)
	_, ok = ch.Pop()
	assert.False(ok)
}

func TestChannelString(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 2
	ptr := mem.Alloc(uint(xxchan.Sizeof[string](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[string](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	assert.True(ch.Push("hello"))
	assert.True(ch.Push("world"))
	assert.False(ch.Push("foo"))
	val, ok := ch.Pop()
	assert.True(ok)
	assert.Equal("hello", val)
	val, ok = ch.Pop()
	assert.True(ok)
	assert.Equal("world", val)
	_, ok = ch.Pop()
	assert.False(ok)
}

func TestChannelComplex64(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 2
	ptr := mem.Alloc(uint(xxchan.Sizeof[complex64](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[complex64](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	assert.True(ch.Push(complex64(1 + 2i)))
	assert.True(ch.Push(complex64(3 + 4i)))
	assert.False(ch.Push(complex64(5 + 6i)))
	val, ok := ch.Pop()
	assert.True(ok)
	assert.Equal(complex64(1+2i), val)
	val, ok = ch.Pop()
	assert.True(ok)
	assert.Equal(complex64(3+4i), val)
	_, ok = ch.Pop()
	assert.False(ok)
}

func TestChannelComplex128(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	n := 2
	ptr := mem.Alloc(uint(xxchan.Sizeof[complex128](n)))
	t.Cleanup(func() { mem.Free(ptr) })
	ch := xxchan.Make[complex128](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())
	assert.True(ch.Push(complex128(1 + 2i)))
	assert.True(ch.Push(complex128(3 + 4i)))
	assert.False(ch.Push(complex128(5 + 6i)))
	val, ok := ch.Pop()
	assert.True(ok)
	assert.Equal(complex128(1+2i), val)
	val, ok = ch.Pop()
	assert.True(ok)
	assert.Equal(complex128(3+4i), val)
	_, ok = ch.Pop()
	assert.False(ok)
}

func TestChannelStruct(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	type MyStruct struct {
		A int
		B string
	}

	n := 2
	ptr := mem.Alloc(uint(xxchan.Sizeof[MyStruct](n)))
	t.Cleanup(func() { mem.Free(ptr) })

	ch := xxchan.Make[MyStruct](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())

	s1 := MyStruct{A: 1, B: "one"}
	s2 := MyStruct{A: 2, B: "two"}

	assert.True(ch.Push(s1))
	assert.True(ch.Push(s2))
	assert.False(ch.Push(MyStruct{}))

	val, ok := ch.Pop()
	assert.True(ok)
	assert.Equal(s1, val)

	val, ok = ch.Pop()
	assert.True(ok)
	assert.Equal(s2, val)

	_, ok = ch.Pop()
	assert.False(ok)
}

func TestChannelPointer(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	n := 2
	ptr := mem.Alloc(uint(xxchan.Sizeof[*int](n)))
	t.Cleanup(func() { mem.Free(ptr) })

	ch := xxchan.Make[*int](ptr, n)
	assert.NotNil(ch)
	assert.Equal(n, ch.Cap())
	assert.Equal(0, ch.Len())

	i1 := 1
	i2 := 2

	assert.True(ch.Push(&i1))
	assert.True(ch.Push(&i2))
	assert.False(ch.Push(nil))

	val, ok := ch.Pop()
	assert.True(ok)
	assert.Equal(&i1, val)

	val, ok = ch.Pop()
	assert.True(ok)
	assert.Equal(&i2, val)

	_, ok = ch.Pop()
	assert.False(ok)
}
