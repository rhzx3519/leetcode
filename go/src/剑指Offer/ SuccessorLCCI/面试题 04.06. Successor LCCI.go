/**
@author ZhengHao Lou
Date    2022/05/16
*/
package main

/**
https://leetcode.cn/problems/successor-lcci/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) (c *TreeNode) {
	var stop bool
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if stop == true {
			c = node
			return
		}
		if node.Val == p.Val {
			stop = true
		}
		dfs(node.Right)
	}
	
	dfs(root)
	return
}
