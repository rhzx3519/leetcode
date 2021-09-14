package main

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	depth := make([]int, 101)

	var f = true
	var inorder func(node *TreeNode, d int)
	inorder = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		inorder(node.Left, d+1)
		depth[node.Val] = d
		if node.Left != nil && node.Right != nil {
			if (node.Left.Val == x && node.Right.Val == y) || (node.Left.Val == y && node.Right.Val == x) {
				f = false
				return
			}
		}
		inorder(node.Right, d+1)
	}
	if !f {
		return false
	}
	inorder(root, 1)
	return depth[x] == depth[y]
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

