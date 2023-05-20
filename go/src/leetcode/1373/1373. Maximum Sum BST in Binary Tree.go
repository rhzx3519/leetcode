package main

/*
*
https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree/description/
思路：递归
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const INF int = 1e5

func maxSumBST(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode) [4]int
	dfs = func(node *TreeNode) [4]int {
		if node == nil {
			return [4]int{1, INF, -INF, 0}
		}
		l, r := dfs(node.Left), dfs(node.Right)
		if l[0] == 1 && r[0] == 1 && l[2] < node.Val && node.Val < r[1] {
			s := node.Val + l[3] + r[3]
			ans = max(ans, s)
			return [4]int{1, min(node.Val, l[1]), max(node.Val, r[2]), s}
		} else {
			return [4]int{}
		}
	}

	dfs(root)
	return ans
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
