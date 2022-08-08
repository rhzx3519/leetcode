/**
@author ZhengHao Lou
Date    2022/08/05
*/
package main

/**
https://leetcode.cn/problems/add-one-row-to-tree/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	var parent []*TreeNode
	que := []*TreeNode{root}
	for d := 1; d < depth && len(que) != 0; d++ {
		parent = []*TreeNode{}
		for tmp := len(que); tmp != 0; tmp-- {
			node := que[0]
			que = que[1:]
			parent = append(parent, node)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}

	if len(parent) == 0 {
		newRoot := &TreeNode{Val: val, Left: root}
		return newRoot
	}
	for i := range parent {
		l := parent[i].Left
		parent[i].Left = &TreeNode{Val: val, Left: l}
		r := parent[i].Right
		parent[i].Right = &TreeNode{Val: val, Right: r}
	}
	return root
}
