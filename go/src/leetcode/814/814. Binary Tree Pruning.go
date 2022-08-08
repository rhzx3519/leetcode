/**
@author ZhengHao Lou
Date    2022/07/21
*/
package main

/**
https://leetcode.cn/problems/binary-tree-pruning/
思路：后续遍历
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}
