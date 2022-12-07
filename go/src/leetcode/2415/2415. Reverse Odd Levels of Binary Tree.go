/**
@author ZhengHao Lou
Date    2022/09/19
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	que := []*TreeNode{root}
	for lv := 0; len(que) != 0; lv++ {
		var nodes []*TreeNode
		for tmp := len(que); tmp != 0; tmp-- {
			node := que[0]
			que = que[1:]
			nodes = append(nodes, node)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		if lv&1 != 0 {
			l, r := 0, len(nodes)-1
			for l < r {
				nodes[l].Val, nodes[r].Val = nodes[r].Val, nodes[l].Val
				l++
				r--
			}
		}
	}
	return root
}
