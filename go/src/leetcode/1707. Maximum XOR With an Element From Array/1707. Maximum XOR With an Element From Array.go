package main

import (
	"fmt"
	"math"
)

const N = 30

type node struct {
	child []*node
	minVal   int
}

func newTrie() *node {
	return &node{
		child: make([]*node, 2),
		minVal: math.MaxInt32,
	}
}

func insert(n *node, val int) {
	n.minVal = min(n.minVal, val)
	for i := N; i >= 0; i-- {
		d := (val>>i) & 1
		if n.child[d] == nil {
			n.child[d] = newTrie()
		}
		n = n.child[d]
		n.minVal = min(n.minVal, val)
	}
}

func find(n *node, x, m int) int {
	if n.minVal > m {
		return -1
	}
	ans := 0
	for i := N; i >= 0; i-- {
		d := ((x>>i)&1)
		if n.child[d^1] != nil && n.child[d^1].minVal <= m {
			ans |= 1<<i
			d ^= 1
		}
		n = n.child[d]
	}
	return ans
}

func maximizeXor(nums []int, queries [][]int) []int {
	root := newTrie()
	for _, num := range nums {
		insert(root, num)
	}
	ans := []int{}
	for _, q := range queries {
		x, m := q[0], q[1]
		ans = append(ans, find(root, x, m))
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximizeXor([]int{0,1,2,3,4}, [][]int{{3,1},{1,3},{5,6}}))
	fmt.Println(maximizeXor([]int{5,2,4,6,6,3}, [][]int{{12,4},{8,1},{6,3}}))
}