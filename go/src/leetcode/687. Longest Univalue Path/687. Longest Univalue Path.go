/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

/**
https://leetcode-cn.com/problems/longest-univalue-path/
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxL int
	var dfs func(node *TreeNode, val int) int	// 返回经过节点node的最大同值路径长度
	dfs = func(node *TreeNode, val int) int {
		if node == nil {
			return 0
		}
		l, r := dfs(node.Left, node.Val), dfs(node.Right, node.Val)
		maxL = max(maxL, l + r)
		if node.Val == val {
			return max(l, r) + 1
		}
		return 0
	}

	dfs(root, root.Val)
	return maxL
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
