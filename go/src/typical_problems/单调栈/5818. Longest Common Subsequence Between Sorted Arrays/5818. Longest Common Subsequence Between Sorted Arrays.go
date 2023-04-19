package main

import "fmt"

func longestCommomSubsequence(arrays [][]int) []int {
	n := len(arrays)
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, 101)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(arrays[i]); j++ {
			grid[i][arrays[i][j]] = 1
		}
	}

	ans := []int{}

	for num := 1; num <= 100; num++ {
		has := true
		for i := range grid {
			if grid[i][num] != 1 {
				has = false
				break
			}
		}
		if has {
			ans = append(ans, num)
		}
	}
	return ans
}

func main() {
	fmt.Println(longestCommomSubsequence([][]int{{1,3,4},{1,4,7,9}}))
}
