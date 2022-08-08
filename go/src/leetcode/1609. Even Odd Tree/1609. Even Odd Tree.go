/**
@author ZhengHao Lou
Date    2021/12/25
*/
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/**
https://leetcode-cn.com/problems/even-odd-tree/
 */
func isEvenOddTree(root *TreeNode) bool {
	var lv int
	que := []*TreeNode{root}
	for len(que) != 0 {
		l := len(que)
		var last int
		if lv&1 == 1 {
			last = 1e6 + 10
		}
		for i := 0; i < l; i++ {
			node := que[0]
			que = que[1:]
			x := node.Val
			if lv&1 == 0 {		// even level
				if x&1 == 0 || x <= last {
					return false
				}
			} else {		// odd level
				if x&1 == 1 || x >= last {
					return false
				}
			}
			last = x

			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		lv++
	}

	return true
}
