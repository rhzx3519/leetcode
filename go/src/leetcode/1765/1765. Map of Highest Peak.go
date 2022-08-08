/**
@author ZhengHao Lou
Date    2022/01/29
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/map-of-highest-peak/
思路：BFS
后搜到的格子，高度肯定更高
time: O(mn)
space: O(mn)
*/
var offset = [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

func highestPeak(isWater [][]int) [][]int {
	que := [][]int{}
	m, n := len(isWater), len(isWater[0])
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
		for j := range grid[i] {
			grid[i][j] = -1
		}
	}

	var encode = func(i, j int) int {
		return i*n + j
	}
	var decode = func(x int) (int, int) {
		return x / n, x % n
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				que = append(que, []int{encode(i, j), 0})
				grid[i][j] = 0
			}
		}
	}

	for len(que) != 0 {
		x := que[0]
		que = que[1:]
		i, j := decode(x[0])
		h := x[1]
		grid[i][j] = h
		for _, dxy := range offset {
			ni, nj := i+dxy[0], j+dxy[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == -1 {
				grid[ni][nj] = h + 1
				que = append(que, []int{encode(ni, nj), h + 1})
			}
		}
	}

	fmt.Println(grid)
	return grid
}

func main() {
	highestPeak([][]int{{0, 1}, {0, 0}})
	highestPeak([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}})
}
