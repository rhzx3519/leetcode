package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func canMerge(trees []*TreeNode) *TreeNode {
	return nil
}

func insert(root *TreeNode, k int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: k,
		}
		return root
	}
	if root.Val > k {
		root.Left = insert(root.Left, k)
	} else if root.Val < k {
		root.Right = insert(root.Right, k)
	}
	return root
}
