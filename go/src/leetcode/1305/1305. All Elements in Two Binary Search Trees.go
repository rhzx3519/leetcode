/**
@author ZhengHao Lou
Date    2022/05/01
*/
package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var dfs func(root *TreeNode, nums *[]int)
	dfs = func(root *TreeNode, nums *[]int) {
		if root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left, nums)
		}
		*nums = append(*nums, root.Val)
		if root.Right != nil {
			dfs(root.Right, nums)
		}
	}
	
	var nums1, nums2 []int
	dfs(root1, &nums1)
	dfs(root2, &nums2)
	
	var ans []int
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] <= nums2[j] {
				ans = append(ans, nums1[i])
				i++
			} else {
				ans = append(ans, nums2[j])
				j++
			}
		} else if i < len(nums1) {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	
	fmt.Println(ans)
	return ans
}

func main() {
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4},
	}
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 3},
	}
	
	getAllElements(root1, root2)
}
