/**
@author ZhengHao Lou
Date    2022/08/30
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil || val > root.Val {
		node.Left = root
		return node
	}

	tmp := root
	for tmp.Right != nil && tmp.Right.Val > val {
		tmp = tmp.Right
	}
	node.Left = tmp.Right
	tmp.Right = node
	return node
}
