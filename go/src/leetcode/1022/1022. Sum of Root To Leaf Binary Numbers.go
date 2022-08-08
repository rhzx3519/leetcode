/**
@author ZhengHao Lou
Date    2022/05/30
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) (ans int) {

	var dfs func(node *TreeNode, s int)
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		s = s<<1 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += s
			return
		}
		if node.Left != nil {
			dfs(node.Left, s)
		}
		if node.Right != nil {
			dfs(node.Right, s)
		}
	}
	dfs(root, 0)
	return
}
