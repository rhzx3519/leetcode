package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode, mi, max int)
	dfs = func(node *TreeNode, mi, mx int) {
		if node == nil {
			return
		}
		ans = max(ans, max(node.Val-mi, mx-node.Val))
		mi, mx = min(mi, node.Val), max(mx, node.Val)
		if node.Left != nil {
			dfs(node.Left, mi, mx)
		}
		if node.Right != nil {
			dfs(node.Right, mi, mx)
		}
	}
	dfs(root, root.Val, root.Val)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

