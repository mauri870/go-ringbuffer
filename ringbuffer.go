package ringbuffer

// RingBufferer defines an interface for a RingBuffer
type RingBufferer[T any] interface {
	Pop() (T, bool)
	Push(T) bool
}

// New creates a new ring buffer. This method returns the LockFreeCached implementation.
func New[T any](cap int) RingBufferer[T] {
	return NewLockFreeCached[T](cap)
}
