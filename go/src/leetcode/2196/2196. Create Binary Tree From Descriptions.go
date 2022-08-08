/**
@author ZhengHao Lou
Date    2022/03/08
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	ind := make(map[int]int)
	left := make(map[int]int)
	adj := make(map[int][]int)
	for i := range descriptions {
		parent, child, isLeft := descriptions[i][0], descriptions[i][1], descriptions[i][2]
		adj[parent] = append(adj[parent], child)
		ind[child]++
		if ind[parent] == 0 {
			ind[parent] = 0
		}
		left[child] = isLeft
	}

	fmt.Println(ind)
	var root *TreeNode
	for k := range ind {
		if ind[k] == 0 {
			root = &TreeNode{Val: k}
			break
		}
	}
	que := []*TreeNode{root}
	for len(que) != 0 {
		x := que[0]
		que = que[1:]
		for _, nx := range adj[x.Val] {
			if left[nx] == 1 {
				x.Left = &TreeNode{Val: nx}
				que = append(que, x.Left)
			} else {
				x.Right = &TreeNode{Val: nx}
				que = append(que, x.Right)
			}
		}
	}

	return root
}

func main() {
	//createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}})
	createBinaryTree([][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}})
}
