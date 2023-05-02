package main

/*
*
https://leetcode.cn/problems/first-completely-painted-row-or-column/
*/
func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	index := make([]int, m*n+1)
	for i, x := range arr {
		index[x] = i
	}
	var rMin, cMin = m * n, m * n
	for i := range mat {
		var r int
		for j := range mat[i] {
			r = max(r, index[mat[i][j]])
		}
		rMin = min(rMin, r)
	}
	for j := 0; j < n; j++ {
		var c int
		for i := 0; i < m; i++ {
			c = max(c, index[mat[i][j]])
		}
		cMin = min(cMin, c)
	}

	return min(rMin, cMin)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//firstCompleteIndex([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}})
	firstCompleteIndex([]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}})
}
