/**
@author ZhengHao Lou
Date    2022/05/23
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/cut-off-trees-for-golf-event/
思路：砍树的顺序是固定的，按照树的高度排序，求出当前点到达下一个砍树点的最短路径，所有最短路径之和为答案
*/
var (
	dirs   = []Position{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	M, N   int
	FOREST [][]int
)

type Position struct {
	x, y int
}

func cutOffTree(forest [][]int) int {
	if forest[0][0] == 0 {
		return -1
	}
	FOREST = forest
	M, N = len(forest), len(forest[0])

	var trees = []Position{}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if forest[i][j] > 1 {
				trees = append(trees, Position{x: i, y: j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		return forest[trees[i].x][trees[i].y] <= forest[trees[j].x][trees[j].y]
	})

	trees = append(trees, Position{})
	copy(trees[1:], trees)
	trees[0] = Position{}

	var dis int
	for i := 0; i < len(trees)-1; i++ {
		d := bfs(trees[i], trees[i+1])
		if d == -1 {
			return -1
		}
		FOREST[trees[i].x][trees[i].y] = 1
		dis += d
	}
	return dis
}

func bfs(cur, target Position) int {
	vis := make([][]bool, M)
	for i := range vis {
		vis[i] = make([]bool, N)
	}
	vis[cur.x][cur.y] = true
	que := []Position{cur}
	for step := 0; len(que) != 0; step++ {
		for temp := len(que) - 1; temp >= 0; temp-- {
			p := que[0]
			que = que[1:]
			if p.x == target.x && p.y == target.y {
				return step
			}
			for _, dir := range dirs {
				nx := p.x + dir.x
				ny := p.y + dir.y
				if nx < 0 || nx >= M || ny < 0 || ny >= N || FOREST[nx][ny] == 0 || vis[nx][ny] {
					continue
				}
				vis[nx][ny] = true
				que = append(que, Position{x: nx, y: ny})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(cutOffTree([][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}))
	fmt.Println(cutOffTree([][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}}))
	fmt.Println(cutOffTree([][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}}))
}
