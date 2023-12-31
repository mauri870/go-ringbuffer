package ringbuffer

import "testing"

var testQueueSize = 100000
var testPushIters = 100000000
var benchResult int
var benchOk bool

func benchmarkBufferPush(b *testing.B, buf RingBufferer[int]) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Push(i)
	}
}

func benchmarkBufferPop(b *testing.B, buf RingBufferer[int]) {
	for n := 0; n < testPushIters; n++ {
		buf.Push(n)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResult, benchOk = buf.Pop()
	}
}

func BenchmarkLockFreePush(b *testing.B) { benchmarkBufferPush(b, NewLockFree[int](testQueueSize)) }
func BenchmarkLockFreePop(b *testing.B)  { benchmarkBufferPop(b, NewLockFree[int](testQueueSize)) }
func BenchmarkLockFreeCachedPush(b *testing.B) {
	benchmarkBufferPush(b, NewLockFreeCached[int](testQueueSize))
}
func BenchmarkLockFreeCachedPop(b *testing.B) {
	benchmarkBufferPop(b, NewLockFreeCached[int](testQueueSize))
}

func BenchmarkLockFreeContainerRingPush(b *testing.B) {
	benchmarkBufferPush(b, NewContainerRing[int](testQueueSize))
}
func BenchmarkLockFreeContainerRingPop(b *testing.B) {
	benchmarkBufferPop(b, NewContainerRing[int](testQueueSize))
}
