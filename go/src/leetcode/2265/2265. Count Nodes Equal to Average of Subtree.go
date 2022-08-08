/**
@author ZhengHao Lou
Date    2022/05/09
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.cn/problems/count-nodes-equal-to-average-of-subtree/
*/
func averageOfSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var c int
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		l1, l2 := dfs(root.Left)
		r1, r2 := dfs(root.Right)
		l, r := l1+r1+1, l2+r2+root.Val
		// fmt.Println(root.Val, l, r)
		if root.Val == r/l {
			c++
		}
		return l, r
	}

	dfs(root)
	return c
}

func main() {
	root := &TreeNode{Val: 1}
	averageOfSubtree(root)
}
