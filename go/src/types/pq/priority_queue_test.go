package pq

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var IntComparotr = func(x, y T) int {
	i, j := x.(int), y.(int)
	if i > j {
		return 1
	} else if i < j {
		return -1
	} else {
		return 0
	}
}

func TestPriorityQueue_Offer(t *testing.T) {
	que := New([]T{3,1,2}, IntComparotr)

	assert.Equal(t, 3, que.Size())
	assert.Equal(t, 3, que.Poll())
	assert.Equal(t, 2, que.Size())
	assert.Equal(t, 2, que.Poll())
	assert.Equal(t, 1, que.Poll())
	assert.True(t, que.IsEmpty())
}

func ExamplePriorityQueue_Poll() {
	que := New([]T{3,1,2}, IntComparotr)
	for !que.IsEmpty() {
		fmt.Println(que.Poll().(int))
	}
	// output:
	// 3
	// 2
	// 1
}

func ExamplePriorityQueue_Poll_Reversed() {
	que := New([]T{3,1,2}, Reversed(IntComparotr))
	for !que.IsEmpty() {
		fmt.Println(que.Poll().(int))
	}
	// output:
	// 1
	// 2
	// 3
}