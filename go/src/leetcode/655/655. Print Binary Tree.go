/**
@author ZhengHao Lou
Date    2022/08/22
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/print-binary-tree/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	var height int
	que := []*TreeNode{root}

	for ; len(que) != 0; height++ {
		for tmp := len(que); tmp != 0; tmp-- {
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

	height--

	var m, n = height + 1, 1<<(height+1) - 1
	var grid = make([][]string, m)
	for i := range grid {
		grid[i] = make([]string, n)
	}

	counter := map[*TreeNode][]int{
		root: []int{0, (n - 1) >> 1},
	}

	que = []*TreeNode{root}
	for len(que) != 0 {
		for tmp := len(que); tmp != 0; tmp-- {
			node := que[0]
			que = que[1:]
			r, c := counter[node][0], counter[node][1]
			grid[r][c] = fmt.Sprintf("%v", node.Val)

			if node.Left != nil {
				counter[node.Left] = []int{r + 1, c - 1<<(height-r-1)}
				que = append(que, node.Left)
			}
			if node.Right != nil {
				counter[node.Right] = []int{r + 1, c + 1<<(height-r-1)}
				que = append(que, node.Right)
			}
		}
	}

	return grid
}

func main() {
	printTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 3,
		},
	})

}
