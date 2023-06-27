package ringbuffer

import "sync/atomic"

type LockFreeCached struct {
	data           []int
	readIdx        atomic.Int64
	readIdxCached  int64
	writeIdx       atomic.Int64
	writeIdxCached int64
}

func NewLockFreeCached(cap int) LockFreeCached {
	return LockFreeCached{data: make([]int, cap)}
}

func (r *LockFreeCached) Push(val int) error {
	writeIdx := r.writeIdx.Load()
	nextWriteIdx := writeIdx + 1

	if nextWriteIdx == int64(len(r.data)) {
		nextWriteIdx = 0
	}

	if nextWriteIdx == r.readIdxCached {
		r.readIdxCached = r.readIdx.Load()
		if nextWriteIdx == r.readIdxCached {
			return ErrBufferFull
		}
	}

	r.data[int(writeIdx)] = val
	r.writeIdx.Store(nextWriteIdx)
	return nil
}

func (r *LockFreeCached) Pop() (int, error) {
	readIdx := r.readIdx.Load()
	if readIdx == r.writeIdxCached {
		r.writeIdxCached = r.writeIdx.Load()
		if readIdx == r.writeIdxCached {
			return 0, ErrBufferEmpty
		}
	}

	nextReadIdx := readIdx + 1
	if nextReadIdx == int64(len(r.data)) {
		nextReadIdx = 0
	}

	r.readIdx.Store(nextReadIdx)
	val := r.data[readIdx]
	return val, nil
}
