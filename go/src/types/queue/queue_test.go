package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Offer(t *testing.T) {
	que := NewQueue(4)
	assert.True(t, true, que.isEmpty())

	que.Offer(1)
	assert.Equal(t, false, que.isEmpty())
	que.Offer(2)
	que.Offer(3)
	assert.Equal(t, 3, que.Size())

	assert.Equal(t, 1, que.Poll())
	assert.Equal(t, 2, que.Size())

	assert.Equal(t, 2, que.Poll())
	assert.Equal(t, 3, que.Poll())

	assert.Equal(t, 0, que.Size())
	assert.True(t, true, que.isEmpty())

}
