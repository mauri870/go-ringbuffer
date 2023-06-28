# ringbuffer

ringBuffer is a Go package that provides an implementation of a ring buffer, also known as a circular buffer. A ring buffer is a data structure that allows efficient storage and retrieval of a fixed-size collection of elements in a circular fashion.

## Features

- Thread-safe implementation: A lock-free ring buffer implementation using atomic operations.
- Generics ready
- Flexible capacity: Set the capacity of the ring buffer upon initialization.
- Simple interface: Push and Pop methods
- Efficient operations: Constant time complexity O(1).
- High performance: Designed to be efficient for concurrent operations with minimal synchronization overhead.
- Memory efficient: uses less memory than container/ring
- Zero memory allocation (besides the initial buffer)

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
ContainerRing // Lock free backed by container/ring, Slow and memory hungry, used only for benchmarks
```

# Benchmarks

```bash
go test -bench=. -benchmem=true ./...
```

```
goos: linux
goarch: amd64
pkg: github.com/mauri870/go-ringbuffer
cpu: AMD Ryzen 7 5800X3D 8-Core Processor
BenchmarkLockFreePush-16                        652895372                1.793 ns/op
BenchmarkLockFreePop-16                         625940958                1.872 ns/op
BenchmarkLockFreeCachedPush-16                  656273632                1.798 ns/op
BenchmarkLockFreeCachedPop-16                   643739690                1.806 ns/op
BenchmarkLockFreeContainerRingPush-16           411274116                3.015 ns/op
BenchmarkLockFreeContainerRingPop-16            407771115                3.030 ns/op
PASS
ok      github.com/mauri870/go-ringbuffer       12.361s
```
