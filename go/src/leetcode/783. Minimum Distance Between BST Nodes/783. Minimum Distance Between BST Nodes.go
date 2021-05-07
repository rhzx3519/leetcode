package _83__Minimum_Distance_Between_BST_Nodes

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func minDiffInBST(root *TreeNode) int {
	var last = -1
	var minDiff = math.MaxInt32
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		if last != -1 {
			if root.Val - last < minDiff {
				minDiff = root.Val - last
			}
		}
		last = root.Val
		traverse(root.Right)
	}

	traverse(root)
	return minDiff
}