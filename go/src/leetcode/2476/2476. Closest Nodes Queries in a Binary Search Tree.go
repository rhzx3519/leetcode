/*
*
@author ZhengHao Lou
Date    2022/11/21
*/
package main

import (
	"sort"
)

/*
*
https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		nums = append(nums, node.Val)
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	n := len(nums)
	var ans [][]int
	for _, x := range queries {
		j := sort.Search(n, func(i int) bool {
			return nums[i] >= x
		})
		var a, b = -1, -1
		if j < n {
			b = nums[j]
			if nums[j] == x {
				a = x
			}
		}
		if a == -1 && j > 0 {
			a = nums[j-1]
		}
		ans = append(ans, []int{a, b})
	}

	return ans
}
