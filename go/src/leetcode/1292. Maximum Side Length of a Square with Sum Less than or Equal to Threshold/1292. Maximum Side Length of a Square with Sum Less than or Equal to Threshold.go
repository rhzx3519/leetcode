/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/
 */
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	sums := make([][]int, m + 1)
	for i := range sums {
		sums[i] = make([]int, n + 1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sums[i][j] = sums[i][j-1] + mat[i-1][j-1]
		}
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			sums[i][j] += sums[i-1][j]
		}
	}

	var side int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 1; i-k+1 >= 0 && j-k+1 >= 0; k++ {
				s := sums[i+1][j+1] - sums[i-k+1][j+1] - sums[i+1][j-k+1] + sums[i-k+1][j-k+1]
				if s <= threshold && k > side {
					side = k
				}
			}
		}
	}

	fmt.Println(sums)

	return side
}

func main() {
	fmt.Println(maxSideLength([][]int{{1,1,3,2,4,3,2},{1,1,3,2,4,3,2},{1,1,3,2,4,3,2}}, 4))
}
