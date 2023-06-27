package ringbuffer

import "sync/atomic"

type LockFree struct {
	data     []int
	readIdx  atomic.Int64
	writeIdx atomic.Int64
}

func NewLockFree(cap int) *LockFree {
	return &LockFree{data: make([]int, cap)}
}

func (r *LockFree) Push(val int) bool {
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

func (r *LockFree) Pop() (int, bool) {
	readIdx := r.readIdx.Load()
	if readIdx == r.writeIdx.Load() {
		return 0, false
	}
	nextReadIdx := (readIdx + 1) % int64(len(r.data))

	r.readIdx.Store(nextReadIdx)
	val := r.data[readIdx]
	return val, true
}
