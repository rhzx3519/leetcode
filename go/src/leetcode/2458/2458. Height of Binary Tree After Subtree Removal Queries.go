/**
@author ZhengHao Lou
Date    2022/11/07
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/height-of-binary-tree-after-subtree-removal-queries/description/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	heights := map[*TreeNode]int{nil: 0}
	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return 0
		}
		depth++
		heights[node] = 1 + max(dfs(node.Left, depth), dfs(node.Right, depth))
		return heights[node]
	}
	dfs(root, -1)
	for i := range heights {
		fmt.Println(i.Val, heights[i])
	}
	
	rest := make(map[int]int)
	var dfs2 func(node *TreeNode, depth, restH int)
	dfs2 = func(node *TreeNode, depth, restH int) {
		if node == nil {
			return
		}
		depth++
		rest[node.Val] = restH
		dfs2(node.Left, depth, max(restH, depth+heights[node.Right]))
		dfs2(node.Right, depth, max(restH, depth+heights[node.Left]))
	}
	dfs2(root, -1, 0)
	
	var ans []int
	for i := range queries {
		ans = append(ans, rest[queries[i]])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
}
