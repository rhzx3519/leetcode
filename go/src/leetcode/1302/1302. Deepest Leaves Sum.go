/**
@author ZhengHao Lou
Date    2022/08/17
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var ans int
	que := []*TreeNode{root}
	for len(que) != 0 {
		var s int
		for l := len(que); l > 0; l-- {
			x := que[0]
			que = que[1:]
			s += x.Val
			if x.Left != nil {
				que = append(que, x.Left)
			}
			if x.Right != nil {
				que = append(que, x.Right)
			}
		}
		ans = s
	}
	return ans
}
