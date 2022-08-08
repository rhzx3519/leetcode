/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

/**
https://leetcode-cn.com/problems/balance-a-binary-search-tree/
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var arr = []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		arr = append(arr, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return build(arr)
}

func build(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	n := len(arr)
	root := &TreeNode{
		Val: arr[n>>1],
	}
	root.Left = build(arr[:n>>1])
	root.Right = build(arr[n>>1+1:])
	return root
}