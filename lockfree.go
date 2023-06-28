package ringbuffer

import "sync/atomic"

type LockFree[T any] struct {
	data     []T
	readIdx  atomic.Int64
	writeIdx atomic.Int64
}

func NewLockFree[T any](cap int) *LockFree[T] {
	return &LockFree[T]{data: make([]T, cap)}
}

func (r *LockFree[T]) Push(val T) bool {
	writeIdx := r.writeIdx.Load()
	nextWriteIdx := (writeIdx + 1) % int64(len(r.data))

	if nextWriteIdx == int64(len(r.data)) {
		nextWriteIdx = 0
	}

	if nextWriteIdx == r.readIdx.Load() {
		return false
	}

	r.data[writeIdx] = val
	r.writeIdx.Store(nextWriteIdx)
	return true
}

func (r *LockFree[T]) Pop() (T, bool) {
	readIdx := r.readIdx.Load()
	if readIdx == r.writeIdx.Load() {
		return *new(T), false
	}
	nextReadIdx := (readIdx + 1) % int64(len(r.data))

	r.readIdx.Store(nextReadIdx)
	val := r.data[readIdx]
	return val, true
}
