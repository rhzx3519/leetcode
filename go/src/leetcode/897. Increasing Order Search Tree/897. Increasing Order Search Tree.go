package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	dummy := new(TreeNode)
	tail := dummy
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		tail.Right = root
		tail = root
		root.Left = nil
		inorder(root.Right)
	}

	inorder(root)
	return dummy.Right
}
