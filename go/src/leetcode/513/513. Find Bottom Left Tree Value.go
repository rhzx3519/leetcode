/**
@author ZhengHao Lou
Date    2022/06/22
*/
package main

/**
https://leetcode.cn/problems/find-bottom-left-tree-value/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) (val int) {
	que := []*TreeNode{root}
	for len(que) != 0 {
		val = que[0].Val
		for l := len(que); l > 0; l-- {
			node := que[0]
			que = que[1:]
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return
}
