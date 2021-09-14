package main

import "fmt"

func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	for i := 0; i < m; i++ {
		j := findPeak(mat[i])
		if (i > 0 && mat[i][j] < mat[i-1][j]) || (i < m-1 && mat[i][j] < mat[i+1][j]) {
			continue
		}
		return []int{i, j}
	}

	return []int{-1, -1}
}

// 如果不存在peak，则返回len(nums) || 0
func findPeak(nums []int) int {
	n := len(nums)
	l, r := 0, n
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m+1] < nums[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(findPeak([]int{6,3,2,1}))
	//fmt.Println(findPeakGrid([][]int{{1,4},{3,2}}))
	//fmt.Println(findPeakGrid([][]int{{10,20,15},{21,30,14},{7,16,32}}))
}