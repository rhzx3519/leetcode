package main

import (
	"fmt"
	"sort"
)

type Coord struct {
	x, y int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	mapper := make(map[*TreeNode]Coord)
	mapper[root] = Coord{x: 0, y: 0}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		coord := mapper[node]
		r, c := coord.x, coord.y
		if node.Left != nil {
			mapper[node.Left] = Coord{x: r+1, y: c-1}
			dfs(node.Left)
		}
		if node.Right != nil {
			mapper[node.Right] = Coord{x: r+1, y: c+1}
			dfs(node.Right)
		}
	}

	dfs(root)

	cols := make(map[int][]Coord)
	for node, coord := range mapper {
		x, y := coord.x, coord.y
		cols[y]	= append(cols[y], Coord{x: x, y: node.Val})
	}

	keys := []int{}
	for c, _ := range cols {
		keys = append(keys, c)
	}
	sort.Ints(keys)

	ans := make([][]int, 0)
	for _, key := range keys {
		sort.Slice(cols[key], func(i, j int) bool {
			c1, c2 := cols[key][i], cols[key][j]
			if c1.x != c2.x {
				return c1.x < c2.x
			}
			return c1.y < c2.y
		})
		tmp := []int{}
		for _, val := range cols[key] {
			tmp = append(tmp, val.y)
		}
		ans = append(ans, tmp)
	}
	fmt.Println(ans)
	return ans
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	verticalTraversal(root)


}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}