/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	if root.Val < low || root.Val > high {
		if root.Left != nil {
			return root.Left
		}
		if root.Right != nil {
			return root.Right
		}
		return nil
	}

	return root
}
