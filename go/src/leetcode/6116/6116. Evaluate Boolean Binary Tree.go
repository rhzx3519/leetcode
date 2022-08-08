/**
@author ZhengHao Lou
Date    2022/07/10
*/
package main

/**
https://leetcode.cn/problems/evaluate-boolean-binary-tree/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false
		}
		return true
	}
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}

	return evaluateTree(root.Left) && evaluateTree(root.Right)
}
