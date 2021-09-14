package main

import "fmt"

const MAX_BIT = 30

func findMaximumXOR(nums []int) int {
	root := newTree()
	for _, num := range nums {
		insert(root, num)
	}

	var maxVal int
	//fmt.Println(root)
	for _, num := range nums {
		t := find(root, num)
		//fmt.Println(t)
		maxVal = max(maxVal, t)
	}
	fmt.Println(maxVal)
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type node struct {
	val int
	child []*node
}

func newTree() *node {
	return &node{val: -1, child: make([]*node, 2)}
}

func insert(root *node, n int) {
	if n == 0 {
		return
	}

	for i := MAX_BIT; i >= 0; i-- {
		bit := n>>i & 1
		if root.child[bit] == nil {
			root.child[bit] = &node{val: bit, child: make([]*node, 2)}
		}
		root = root.child[bit]
	}
}

func find(root *node, n int) int {
	var t int
	for i := MAX_BIT; i >= 0; i-- {
		bit := n>>i & 1
		if root != nil {
			if root.child[bit^1] != nil {
				root = root.child[bit^1]
				t |= 1<<i
			} else {
				root = root.child[bit]
			}
		} else {
			t |= n>>i & 1
		}
	}

	return t
}

func main() {
	findMaximumXOR([]int{3,10,5,25,2,8})
	findMaximumXOR([]int{0})
	findMaximumXOR([]int{2,4})
	findMaximumXOR([]int{8,10,2})
	findMaximumXOR([]int{14,70,53,83,49,91,36,80,92,51,66,70})
}