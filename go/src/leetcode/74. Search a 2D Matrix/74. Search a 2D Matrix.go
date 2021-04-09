package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	col, row := 0, m-1
	for col < n && row >= 0 {
		val := matrix[row][col]
		if val == target {
			return true
		} else if val > target {
			row--
		} else {
			col++
		}
	}
	return false
}
