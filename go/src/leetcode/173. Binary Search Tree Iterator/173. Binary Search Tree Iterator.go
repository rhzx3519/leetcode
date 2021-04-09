package main

type BSTIterator struct {
	a []*TreeNode
	cur int
}


func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{a: make([]*TreeNode, 0)}
	st := []*TreeNode{}
	for root != nil || len(st) > 0 {
		if root != nil {
			st = append(st, root)
			root = root.Left
		} else {
			root = st[len(st)-1]
			iter.a = append(iter.a, root)
			st = st[:len(st)-1]
			root = root.Right
		}
	}
	return iter
}


func (this *BSTIterator) Next() int {
	val := this.a[this.cur].Val
	this.cur++
	return val
}


func (this *BSTIterator) HasNext() bool {
	return this.cur < len(this.a)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
