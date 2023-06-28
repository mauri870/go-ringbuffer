package ringbuffer

// RingBufferer defines an interface for a RingBuffer
type RingBufferer[T any] interface {
	Pop() (T, bool)
	Push(T) bool
}
