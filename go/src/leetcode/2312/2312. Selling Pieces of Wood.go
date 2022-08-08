/**
@author ZhengHao Lou
Date    2022/06/21
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/selling-pieces-of-wood/
*/
func sellingWood(m int, n int, prices [][]int) int64 {
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for i := range prices {
		pr[prices[i][0]][prices[i][1]] = prices[i][2]
	}
	f := make([][]int64, m+1)
	for i := 1; i <= m; i++ {
		f[i] = make([]int64, n+1)
		for j := 1; j <= n; j++ {
			f[i][j] = int64(pr[i][j])
			for k := 1; k < j; k++ { // 垂直切割
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k < i; k++ { // 水平切割
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return f[m][n]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(sellingWood(3, 5, [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}}))
}
