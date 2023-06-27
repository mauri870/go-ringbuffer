package ringbuffer

import "testing"

var testQueueSize = 100000
var testPushIters = 100000000
var benchResult int
var benchErr error

func BenchmarkLockFreePush(b *testing.B) {
	buf := NewLockFree(testQueueSize)
	for n := 0; n < b.N; n++ {
		benchErr = buf.Push(n)
	}
}

func BenchmarkLockFreePop(b *testing.B) {
	buf := NewLockFree(testQueueSize)
	for n := 0; n < testPushIters; n++ {
		buf.Push(n)
	}

	for n := 0; n < b.N; n++ {
		benchResult, benchErr = buf.Pop()
	}
}

func BenchmarkLockFreeCachedPush(b *testing.B) {
	buf := NewLockFreeCached(testQueueSize)
	for n := 0; n < b.N; n++ {
		benchErr = buf.Push(n)
	}
}

func BenchmarkLockFreeCachedPop(b *testing.B) {
	buf := NewLockFreeCached(testQueueSize)
	for n := 0; n < testPushIters; n++ {
		buf.Push(n)
	}

	for n := 0; n < b.N; n++ {
		benchResult, benchErr = buf.Pop()
	}
}
