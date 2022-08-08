/**
@author ZhengHao Lou
Date    2021/11/11
*/
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var dfs func(node *TreeNode, s int) *TreeNode
	dfs = func(node *TreeNode, s int) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			if node.Val + s < limit {
				return nil
			}
			return node
		}

		node.Left, node.Right = dfs(node.Left, s + node.Val), dfs(node.Right, s + node.Val)
		if node.Left == nil && node.Right == nil {
			return nil
		}
		return node
	}

	return dfs(root, 0)
}


