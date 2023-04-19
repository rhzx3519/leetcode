package main

/*
* https://leetcode.cn/problems/check-knight-tour-configuration/
 */
func checkValidGrid(grid [][]int) bool {
	type pair struct {
		i, j int
	}
	cnt := make(map[int]pair)
	for i := range grid {
		for j := range grid[i] {
			cnt[grid[i][j]] = pair{i, j}
		}
	}
	cnt[-1] = pair{}
	n := len(grid)
	for k := 0; k < n*n; k++ {
		cur := cnt[k]
		last := cnt[k-1]
		if last.i-1 == cur.i && (cur.j == last.j-2 || cur.j == last.j+2) ||
			last.i-2 == cur.i && (cur.j == last.j-1 || cur.j == last.j+1) ||
			last.i+1 == cur.i && (cur.j == last.j-2 || cur.j == last.j+2) ||
			last.i+2 == cur.i && (cur.j == last.j-1 || cur.j == last.j+1) ||
			last.i == cur.i && cur.j == last.j {
			continue
		}
		return false
	}
	return true
}
