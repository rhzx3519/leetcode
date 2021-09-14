package main

import "math"

/**
思路：根节点是树的最小节点，遍历树，寻找比root.Val大的最小值
 */
func findSecondMinimumValue(root *TreeNode) int {
	val := root.Val
	var ans = int64(math.MaxInt64)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val > val {
			ans = min(int64(node.Val), ans)
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	if ans == int64(math.MaxInt64) {
		return -1
	}
	return int(ans)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}