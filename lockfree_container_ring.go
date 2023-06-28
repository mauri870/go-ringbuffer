package ringbuffer

import (
	"container/ring"
	"sync/atomic"
)

type LockFreeContainerRing struct {
	buffer   *ring.Ring
	capacity uint64
	readIdx  atomic.Uint64
	writeIdx atomic.Uint64
}

func NewContainerRing(cap int) *LockFreeContainerRing {
	return &LockFreeContainerRing{
		buffer:   ring.New(cap),
		capacity: uint64(cap),
	}
}

func (r *LockFreeContainerRing) Push(val int) bool {
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

func (r *LockFreeContainerRing) Pop() (int, bool) {
	readIdx := r.readIdx.Load()
	if readIdx == r.writeIdx.Load() {
		return 0, false
	}
	nextReadIdx := (readIdx + 1) % r.capacity

	r.buffer = r.buffer.Next()
	val, ok := r.buffer.Value.(int)
	if !ok {
		return 0, ok
	}

	r.readIdx.Store(nextReadIdx)
	return val, true
}
