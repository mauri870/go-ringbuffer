package ringbuffer

import "errors"

var (
	ErrBufferFull  = errors.New("buffer is full")
	ErrBufferEmpty = errors.New("buffer is empty")
)

type RingBufferer interface {
	Pop() (int, error)
	Push(int) error
}
