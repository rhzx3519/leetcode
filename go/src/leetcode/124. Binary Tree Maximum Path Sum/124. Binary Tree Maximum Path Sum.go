package main

import "math"

/**
https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
思路：一个树的最大路径和为最大的 经过节点node的最大路径和
 */
func maxPathSum(root *TreeNode) int {
	var ans = math.MinInt32

	// 返回包含node的最大路径和
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := max(0, dfs(node.Left))
		r := max(0, dfs(node.Right))
		ans = max(ans, l + r + node.Val)	// 经过节点node最大路径和
		return  max(l, r) + node.Val
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
