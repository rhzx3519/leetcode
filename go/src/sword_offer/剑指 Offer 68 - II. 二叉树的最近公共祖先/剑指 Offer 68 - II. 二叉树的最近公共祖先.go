package main

/**
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	} else if r == nil {
		return l
	}
	return root
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
