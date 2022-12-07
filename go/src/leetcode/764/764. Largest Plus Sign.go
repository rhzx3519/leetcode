/**
@author ZhengHao Lou
Date    2022/11/09
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/largest-plus-sign/
*/
func orderOfLargestPlusSign(n int, mines [][]int) int {
	grid := make([][]int, n)
	xl, xr := make([][]int, n), make([][]int, n)
	yl, yr := make([][]int, n), make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
		xl[i] = make([]int, n)
		xr[i] = make([]int, n)
		yl[i] = make([]int, n)
		yr[i] = make([]int, n)
	}
	for _, m := range mines {
		grid[m[0]][m[1]] = 1
	}

	for i := 0; i < n; i++ {
		var c1, c2 int
		for j := 0; j < n; j++ {
			xl[i][j] = c1
			if grid[i][j] == 0 {
				c1++
			} else {
				c1 = 0
			}
			yl[j][i] = c2
			if grid[j][i] == 0 {
				c2++
			} else {
				c2 = 0
			}
		}

		c1, c2 = 0, 0
		for j := n - 1; j >= 0; j-- {
			xr[i][j] = c1
			if grid[i][j] == 0 {
				c1++
			} else {
				c1 = 0
			}
			yr[j][i] = c2
			if grid[j][i] == 0 {
				c2++
			} else {
				c2 = 0
			}
		}
	}

	var mx int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			var c = n + 1
			if xl[i][j] < c {
				c = xl[i][j]
			}
			if xr[i][j] < c {
				c = xr[i][j]
			}
			if yl[i][j] < c {
				c = yl[i][j]
			}
			if yr[i][j] < c {
				c = yr[i][j]
			}

			if c+1 > mx {
				mx = c + 1
			}
		}
	}

	fmt.Println(xl, xr)
	fmt.Println(yl, yr)
	fmt.Println(mx)
	return mx
}

func main() {
	orderOfLargestPlusSign(5, [][]int{{4, 2}})
}
