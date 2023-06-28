package ringbuffer

import "sync/atomic"

type LockFreeCached[T any] struct {
	data           []T
	readIdx        atomic.Uint64
	readIdxCached  uint64
	writeIdx       atomic.Uint64
	writeIdxCached uint64
}

func NewLockFreeCached[T any](cap int) *LockFreeCached[T] {
	return &LockFreeCached[T]{data: make([]T, cap)}
}

func (r *LockFreeCached[T]) Push(val T) bool {
	writeIdx := r.writeIdx.Load()
	nextWriteIdx := (writeIdx + 1) % uint64(len(r.data))

	if nextWriteIdx == r.readIdxCached {
		r.readIdxCached = r.readIdx.Load()
		if nextWriteIdx == r.readIdxCached {
			return false
		}
	}

	r.data[writeIdx] = val
	r.writeIdx.Store(nextWriteIdx)
	return true
}

func (r *LockFreeCached[T]) Pop() (T, bool) {
	readIdx := r.readIdx.Load()
	if readIdx == r.writeIdxCached {
		r.writeIdxCached = r.writeIdx.Load()
		if readIdx == r.writeIdxCached {
			return *new(T), false
		}
	}

	nextReadIdx := (readIdx + 1) % uint64(len(r.data))

	r.readIdx.Store(nextReadIdx)
	val := r.data[readIdx]
	return val, true
}
