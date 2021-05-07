package tool

import (
	"fmt"
	"testing"
)

type integer int

func (i integer) CompareTo(other integer) int {
	return int(i - other)
}

func TestPQ_Push(t *testing.T) {
	pq := &PQ{}
	pq.Push(integer(0))
	pq.Push(integer(3))
	pq.Push(integer(5))

	for pq.Len() > 0 {
		fmt.Println(pq.Pop())
	}
}
