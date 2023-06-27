package ringbuffer

import "testing"

var testQueueSize = 100000
var testPushIters = 100000000
var benchResult int
var benchOk bool

func BenchmarkLockFreePush(b *testing.B) {
	buf := NewLockFree(testQueueSize)
	for n := 0; n < b.N; n++ {
		benchOk = buf.Push(n)
	}
}

func BenchmarkLockFreePop(b *testing.B) {
	buf := NewLockFree(testQueueSize)
	for n := 0; n < testPushIters; n++ {
		buf.Push(n)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		benchResult, benchOk = buf.Pop()
	}
}

func BenchmarkLockFreeCachedPush(b *testing.B) {
	buf := NewLockFreeCached(testQueueSize)
	for n := 0; n < b.N; n++ {
		benchOk = buf.Push(n)
	}
}

func BenchmarkLockFreeCachedPop(b *testing.B) {
	buf := NewLockFreeCached(testQueueSize)
	for n := 0; n < testPushIters; n++ {
		buf.Push(n)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		benchResult, benchOk = buf.Pop()
	}
}
