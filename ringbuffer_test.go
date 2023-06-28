package ringbuffer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleNew() {
	buf := New[int](3)

	buf.Push(1)
	buf.Push(2)
	buf.Push(3) // buffer is full!

	val, ok := buf.Pop()
	fmt.Println(val, ok)
	// Output: 1 true
}

func TestPushPop(t *testing.T) {
	testCases := []struct {
		name string
		buf  RingBufferer[int]
	}{
		{"lock free", NewLockFree[int](3)},
		{"lock free cached", NewLockFreeCached[int](3)},
		{"lock free container ring", NewContainerRing[int](3)},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf := tc.buf
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
		})
	}
}
