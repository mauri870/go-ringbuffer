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

func (r *LockFree) Push(val int) error {
	writeIdx := r.writeIdx.Load()
	nextWriteIdx := writeIdx + 1

	if nextWriteIdx == int64(len(r.data)) {
		nextWriteIdx = 0
	}

	if nextWriteIdx == r.readIdx.Load() {
		return ErrBufferFull
	}

	r.data[int(writeIdx)] = val
	r.writeIdx.Store(nextWriteIdx)
	return nil
}

func (r *LockFree) Pop() (int, error) {
	readIdx := r.readIdx.Load()
	if readIdx == r.writeIdx.Load() {
		return 0, ErrBufferEmpty
	}
	nextReadIdx := readIdx + 1
	if nextReadIdx == int64(len(r.data)) {
		nextReadIdx = 0
	}

	r.readIdx.Store(nextReadIdx)
	val := r.data[readIdx]
	return val, nil
}
