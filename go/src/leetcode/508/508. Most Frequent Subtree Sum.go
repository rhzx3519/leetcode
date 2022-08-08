/**
@author ZhengHao Lou
Date    2022/06/19
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	var mx int
	var counter = make(map[int]int)
	var ans []int

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		s := node.Val + dfs(node.Left) + dfs(node.Right)
		counter[s]++
		return s
	}

	dfs(root)
	for s, c := range counter {
		if c > mx {
			mx = c
			ans = []int{s}
		} else if c == mx {
			ans = append(ans, s)
		}
	}
	return ans
}
