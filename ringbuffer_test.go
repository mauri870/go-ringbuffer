package ringbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockFree(t *testing.T) {
	buf := NewLockFree(3)

	assert.True(t, buf.Push(1))
	assert.True(t, buf.Push(2))
	assert.False(t, buf.Push(3)) // buffer is full

	val, ok := buf.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = buf.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, val)

	_, ok = buf.Pop()
	assert.False(t, ok) // buffer is empty

	_, ok = buf.Pop()
	assert.False(t, ok) // buffer is empty
}
