package main

import (
	"fmt"
	"strings"
)

func equalPairs(grid [][]int) int {
	cnt := make(map[string]int)
	n := len(grid)
	for r := 0; r < n; r++ {
		var tmp []string
		for c := 0; c < n; c++ {
			tmp = append(tmp, fmt.Sprintf("%v", grid[r][c]))
		}
		cnt[strings.Join(tmp, ",")]++
	}
	var ans int
	for c := 0; c < n; c++ {
		var tmp []string
		for r := 0; r < n; r++ {
			tmp = append(tmp, fmt.Sprintf("%v", grid[r][c]))
		}
		ans += cnt[strings.Join(tmp, ",")]
	}
	return ans
}

func main() {
	//equalPairs([][]int{
	//	{250, 78, 253},
	//	{334, 252, 253},
	//	{250, 253, 253}})
	equalPairs([][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7}})
}
