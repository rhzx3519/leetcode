/**
@author ZhengHao Lou
Date    2022/11/16
*/
package main

/**
https://leetcode.cn/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) (tot int) {
	que := []*TreeNode{root}
	for len(que) != 0 {
		var values []int
		for tmp := len(que); tmp != 0; tmp-- {
			node := que[0]
			que = que[1:]
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			values = append(values, node.Val)
		}
		for i := range values {
			var mi, idx = values[i], i
			for j := i + 1; j < len(values); j++ {
				if values[j] < mi {
					mi = values[j]
					idx = j
				}
			}
			if idx != i {
				values[i], values[idx] = values[idx], values[i]
				tot++
			}
		}
	}
	return
}
