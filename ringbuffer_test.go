package ringbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockFree(t *testing.T) {
	buf := NewLockFree(3)

	assert.NoError(t, buf.Push(1))
	assert.NoError(t, buf.Push(2))
	assert.Equal(t, ErrBufferFull, buf.Push(3))

	val, err := buf.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, val)

	val, err = buf.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = buf.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = buf.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, val)
}
