package ringbuffer

import (
	"container/ring"
	"sync/atomic"
)

type LockFreeContainerRing[T any] struct {
	buffer   *ring.Ring
	capacity uint64
	readIdx  atomic.Uint64
	writeIdx atomic.Uint64
}

func NewContainerRing[T any](cap int) *LockFreeContainerRing[T] {
	return &LockFreeContainerRing[T]{
		buffer:   ring.New(cap),
		capacity: uint64(cap),
	}
}

func (r *LockFreeContainerRing[T]) Push(val T) bool {
	writeIdx := r.writeIdx.Load()
	nextWriteIdx := (writeIdx + 1) % r.capacity

	if nextWriteIdx == r.readIdx.Load() {
		return false
	}

	r.buffer.Value = val
	r.buffer = r.buffer.Next()
	r.writeIdx.Store(nextWriteIdx)
	return true
}

func (r *LockFreeContainerRing[T]) Pop() (T, bool) {
	readIdx := r.readIdx.Load()
	if readIdx == r.writeIdx.Load() {
		return *new(T), false
	}
	nextReadIdx := (readIdx + 1) % r.capacity

	r.buffer = r.buffer.Next()
	val, ok := r.buffer.Value.(T)
	if !ok {
		return *new(T), ok
	}

	r.readIdx.Store(nextReadIdx)
	return val, true
}
