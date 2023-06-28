# ringbuffer 

ringBuffer is a Go package that provides an implementation of a ring buffer, also known as a circular buffer. A ring buffer is a data structure that allows efficient storage and retrieval of a fixed-size collection of elements in a circular fashion.

## Features

- Thread-safe implementation: A lock-free ring buffer implementation using atomic operations.
- Generics ready
- Flexible capacity: Set the capacity of the ring buffer upon initialization.
- Simple interface: Push and Pop methods
- Efficient operations: Constant time complexity O(1).
- High performance: Designed to be efficient for concurrent operations with minimal synchronization overhead.
- Faster than container/ring and more memory efficient
- Zero memory allocations (besides the initial buffer)

## Installation

Use `go get` to install the package:

```bash
go get github.com/mauri870/go-ringbuffer
```

# Usage

[Godoc](https://pkg.go.dev/github.com/mauri870/go-ringbuffer)

There are a couple of different implementations available:

```go
LockFreeCached // Fastest implementation, Lock free with a cache approach to minimize atomic memory operations
LockFree // Lock free implementation
ContainerRing // Lock free based on container/ring, Slow and memory hungry, used only for benchmarks
```
