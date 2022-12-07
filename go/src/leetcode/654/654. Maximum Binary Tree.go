/**
@author ZhengHao Lou
Date    2022/08/30
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var idx, mx int
	for i := range nums {
		if nums[i] > mx {
			mx = nums[i]
			idx = i
		}
	}
	root := &TreeNode{
		Val:   nums[idx],
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}
	return root
}
