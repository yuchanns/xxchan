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

// Package xxchan provides a garbage collection-free channel implementation
// that operates on user-allocated memory blocks.
package xxchan

import (
	"sync/atomic"
	"time"
	"unsafe"
)

// Channel is a lock-free, garbage collection-free channel implementation that operates
// on a user-provided memory block. It supports concurrent access from multiple goroutines
// and provides fundamental channel operations including Push, Pop, Len, and Cap.
//
// Unlike Go's built-in channels, Channel[T] does not allocate memory internally,
// making it suitable for scenarios where garbage collector managed memory
// cannot be used, such as in certain system-level or embedded applications.
//
// The channel uses a circular buffer internally and employs atomic operations
// with spin-lock synchronization to ensure thread safety.
//
// Example usage:
//
//	// Calculate required memory size for a channel with capacity 100
//	size := xxchan.Sizeof[int](100)
//
//	// Allocate memory (can be stack, heap, or memory pool)
//	buf := make([]byte, size)
//
//	// Create channel from the allocated memory
//	ch := xxchan.Make[int](unsafe.Pointer(&buf[0]), 100)
//
//	// Use the channel
//	ch.Push(42)
//	val, ok := ch.Pop()
type Channel[T any] struct {
	l    int32
	head int64
	tail int64
	cap  int64
	_    [0]T // Zero-sized placeholder for type information; actual buffer follows the struct
}

// alignUp rounds up n to the nearest multiple of align.
// This ensures proper memory alignment for optimal performance and correctness.
func alignUp(n, align int) int {
	return (n + align - 1) &^ (align - 1)
}

// Sizeof calculates the total memory size required for a Channel[T] with the specified capacity.
//
// The function accounts for:
//   - The size of the Channel[T] struct itself
//   - Storage space for n elements of type T
//   - Proper memory alignment requirements for type T
//
// Parameters:
//   - n: The desired capacity of the channel (maximum number of elements)
//
// Returns:
//   - The total size in bytes that should be allocated for the channel
//
// The returned size should be used when allocating memory before calling Make.
func Sizeof[T any](n int) int {
	size := int(unsafe.Sizeof(Channel[T]{})) + n*int(unsafe.Sizeof(*new(T)))
	return alignUp(size, int(unsafe.Alignof(new(T))))
}

// Make initializes a new Channel[T] using a pre-allocated memory block.
//
// The provided memory block must be at least Sizeof[T](n) bytes in size
// and properly aligned. The function does not perform memory allocation;
// it only initializes the channel structure within the given memory.
//
// Parameters:
//   - ptr: Pointer to the pre-allocated memory block
//   - n: The capacity of the channel (maximum number of elements)
//
// Returns:
//   - A pointer to the initialized Channel[T]
//
// Safety:
//   - The caller must ensure the memory block remains valid for the channel's lifetime
//   - The memory block must be at least Sizeof[T](n) bytes
//   - The pointer must be properly aligned for the target type T
//
// Example:
//
//	size := Sizeof[int](10)
//	buf := make([]byte, size)
//	ch := Make[int](unsafe.Pointer(&buf[0]), 10)
func Make[T any](ptr unsafe.Pointer, n int) *Channel[T] {
	c := (*Channel[T])(ptr)
	c.cap = int64(n)
	c.head = 0
	c.tail = 0
	c.l = 0
	return c
}

// acquireLock acquires an exclusive acquireLock on the channel using atomic compare-and-swap.
// Uses spin-waiting with microsecond delays to reduce CPU usage while waiting.
func (c *Channel[T]) acquireLock() {
	for !atomic.CompareAndSwapInt32(&c.l, 0, 1) {
		time.Sleep(time.Microsecond) // Spin-wait with brief pause
	}
}

// releaseLock releases the exclusive lock on the channel.
func (c *Channel[T]) releaseLock() {
	atomic.StoreInt32(&c.l, 0)
}

// buffer returns a slice view of the internal circular buffer.
// The buffer is located immediately after the Channel struct in memory,
// properly aligned for type T.
func (c *Channel[T]) buffer() []T {
	if c == nil {
		return nil
	}
	structSize := unsafe.Sizeof(Channel[T]{})
	alignedOffset := alignUp(int(structSize), int(unsafe.Alignof(*new(T))))
	addr := unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(alignedOffset))
	return unsafe.Slice((*T)(addr), c.cap)
}

// Push attempts to add a value to the channel.
//
// The operation is atomic and thread-safe. If the channel is at full capacity,
// the operation fails immediately without blocking.
//
// Parameters:
//   - val: The value to add to the channel
//
// Returns:
//   - true if the value was successfully added
//   - false if the channel is full
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (c *Channel[T]) Push(val T) (ok bool) {
	if c == nil {
		return
	}
	c.acquireLock()
	defer c.releaseLock()

	if c.tail-c.head >= c.cap {
		return // Channel is full
	}
	c.buffer()[c.tail%c.cap] = val
	c.tail++
	ok = true
	return
}

// Pop attempts to remove and return a value from the channel.
//
// The operation is atomic and thread-safe. If the channel is empty,
// the operation returns immediately with the zero value and false.
//
// Returns:
//   - v: The value removed from the channel, or zero value of T if empty
//   - ok: true if a value was successfully removed, false if the channel was empty
//
// The function automatically resets internal head/tail pointers when the channel
// becomes empty to prevent potential overflow in long-running applications.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
func (c *Channel[T]) Pop() (v T, ok bool) {
	if c == nil {
		return
	}
	c.acquireLock()
	defer c.releaseLock()
	if c.tail == c.head {
		return // Channel is empty
	}
	ok = true
	v = c.buffer()[c.head%c.cap]
	c.head++
	if c.head == c.tail {
		c.head, c.tail = 0, 0 // Reset to prevent overflow
	}
	return
}

// Len returns the current number of elements stored in the channel.
//
// This operation is thread-safe and provides a snapshot of the channel's
// length at the time of the call. The actual length may change immediately
// after the function returns due to concurrent operations.
//
// Returns:
//   - The number of elements currently in the channel (0 to Cap())
//
// Time Complexity: O(1)
func (c *Channel[T]) Len() int {
	if c == nil {
		return 0
	}

	c.acquireLock()
	defer c.releaseLock()

	if c.tail >= c.head {
		return int(c.tail - c.head)
	}
	return int(c.cap - c.head + c.tail)
}

// Cap returns the maximum capacity of the channel.
//
// This value is set during channel creation and remains constant throughout
// the channel's lifetime. It represents the maximum number of elements
// the channel can hold simultaneously.
//
// Returns:
//   - The maximum capacity of the channel
//
// Time Complexity: O(1)
// Note: This operation does not require locking as capacity is immutable.
func (c *Channel[T]) Cap() int {
	if c == nil {
		return 0
	}
	return int(c.cap)
}
