package main

import (
	"fmt"
)

const N = 10000

var offset = [][]int{{1,0},{-1,0},{0,1},{0,-1}}

func highestPeak(isWater [][]int) [][]int {
	que := []int{}
	m, n := len(isWater), len(isWater[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				isWater[i][j] = 0
				que = append(que, i*N + j)
			} else {
				isWater[i][j] = -1
			}
		}
	}

	for len(que) > 0 {
		i, j := que[0]/N, que[0]%N
		que = que[1:]
		for _, d := range offset {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && isWater[ni][nj] == -1 {
				isWater[ni][nj]	= isWater[i][j] + 1
				que = append(que, ni*N + nj)
			}
		}
	}

	return isWater
}

// 思路：多源BFS

func main() {
	fmt.Println(highestPeak([][]int{{0,1},{0,0}}))
	fmt.Println(highestPeak([][]int{{0,0,1},{1,0,0},{0,0,0}}))
}
