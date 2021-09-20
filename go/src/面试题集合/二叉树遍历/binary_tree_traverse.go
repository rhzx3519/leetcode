package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preTraverse(root *TreeNode) {
	st := []*TreeNode{root}
	for len(st) > 0 {
		node := st[len(st) - 1]
		st = st[:len(st) - 1]

		fmt.Println(node.Val)
		if node.Left != nil {
			st = append(st, node.Left)
		}
		if node.Right != nil {
			st = append(st, node.Right)
		}
	}
}
