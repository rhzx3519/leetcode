/**
@author ZhengHao Lou
Date    2021/11/18
*/
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func findTilt(root *TreeNode) int {
	var ans int

	var dfs func(node *TreeNode) int

	mem := make(map[*TreeNode]int)
	dfs = func(node *TreeNode) int {
		if mem[node] != 0 {
			return mem[node]
		}
		if node == nil {
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		ans += abs(l - r)
		s := node.Val + l + r
		mem[node] = s
		return s
	}

	dfs(root)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}