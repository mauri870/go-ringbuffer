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

func (r *LockFreeCached) Push(val int) bool {
	writeIdx := r.writeIdx.Load()
	nextWriteIdx := (writeIdx + 1) % int64(len(r.data))

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

func (r *LockFreeCached) Pop() (int, bool) {
	readIdx := r.readIdx.Load()
	if readIdx == r.writeIdxCached {
		r.writeIdxCached = r.writeIdx.Load()
		if readIdx == r.writeIdxCached {
			return 0, false
		}
	}

	nextReadIdx := (readIdx + 1) % int64(len(r.data))

	r.readIdx.Store(nextReadIdx)
	val := r.data[readIdx]
	return val, true
}
