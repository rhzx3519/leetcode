package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	s := 0
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if root.Val >= low && root.Val <= high {
			s += root.Val
		}
		inorder(root.Right)
	}
	inorder(root)
	return s
}

