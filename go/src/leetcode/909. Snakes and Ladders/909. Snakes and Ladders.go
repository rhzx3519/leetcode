package main

import "fmt"

func snakesAndLadders(board [][]int) int {
	N := len(board)
	N2 := N*N
	visited := make([]bool, N2)
	que := []int{0}
	visited[0] = true
	var step int
	for len(que) > 0 {
		for sz := len(que); sz > 0; sz-- {
			i := que[0]
			que = que[1:]

			if i == N2-1 {
				return step
			}
			for k := 1; k <= 6 && i+k < N2; k++ {
				ni := i + k
				if visited[ni] {
					continue
				}
				visited[ni] = true
				var x, y int
				x = N - 1 - ni/N
				if (N - 1 - x)&1 == 1 {
					y = N - 1 - ni%N
				} else {
					y = ni%N
				}
				if board[x][y] != -1 {
					ni = board[x][y]-1
				}
				que = append(que, ni)
			}
		}
		step++
	}
	return -1
}

func main() {
	fmt.Println(snakesAndLadders([][]int{
		{-1,-1,-1,-1,-1,-1},
		{-1,-1,-1,-1,-1,-1},
		{-1,-1,-1,-1,-1,-1},
		{-1,35,-1,-1,13,-1},
		{-1,-1,-1,-1,-1,-1},
		{-1,15,-1,-1,-1,-1},
	}))

	fmt.Println(snakesAndLadders([][]int{
		{-1,-1,-1},
		{-1,9,8},
		{-1,8,9},
	}))

	fmt.Println(snakesAndLadders([][]int{
		{-1,-1,19,10,-1},
		{2,-1,-1,6,-1},
		{-1,17,-1,19,-1},
		{25,-1,20,-1,-1},
		{-1,-1,-1,-1,15},
	}))

}