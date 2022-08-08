/**
  @author ZhengHao Lou
  Date    2021/09/28
https://leetcode-cn.com/problems/path-sum-iii/
思路：前缀和
 */
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var count int

	var dfs func(node *TreeNode, s int, mapper map[int]int)
	dfs = func(node *TreeNode, s int, mapper map[int]int) {
		if node == nil {
			return
		}
		s += node.Val
		count += mapper[s - targetSum]
		mapper[s]++
		dfs(node.Left, s, mapper)
		dfs(node.Right, s, mapper)
		mapper[s]--
	}

	mapper := make(map[int]int)
	mapper[0] = 1
	dfs(root, 0, mapper)
	return count
}
