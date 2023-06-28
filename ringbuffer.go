package ringbuffer

// RingBufferer defines an interface for a RingBuffer
type RingBufferer interface {
	Pop() (int, bool)
	Push(int) bool
}
