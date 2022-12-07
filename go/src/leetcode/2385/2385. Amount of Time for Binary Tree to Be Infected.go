/**
@author ZhengHao Lou
Date    2022/08/22
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	adj := make(map[int][]int)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			adj[root.Val] = append(adj[root.Val], root.Left.Val)
			adj[root.Left.Val] = append(adj[root.Left.Val], root.Val)
			dfs(root.Left)
		}
		if root.Right != nil {
			adj[root.Val] = append(adj[root.Val], root.Right.Val)
			adj[root.Right.Val] = append(adj[root.Right.Val], root.Val)
			dfs(root.Right)
		}
	}

	dfs(root)

	fmt.Println(adj)
	var minutes = -1
	var vis = make(map[int]bool)
	vis[start] = true
	que := []int{start}
	for ; len(que) != 0; minutes++ {
		for tmp := len(que); tmp != 0; tmp-- {
			x := que[0]
			que = que[1:]
			for _, nx := range adj[x] {
				if !vis[nx] {
					vis[nx] = true
					que = append(que, nx)
				}
			}
		}
	}

	fmt.Println(minutes)
	return minutes
}

func main() {
	amountOfTime(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 6},
		},
	}, 3)
}
