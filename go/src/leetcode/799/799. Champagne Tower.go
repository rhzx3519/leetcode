/**
@author ZhengHao Lou
Date    2022/11/20
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/champagne-tower/
*/
const N = 100

func champagneTower(poured int, query_row int, query_glass int) float64 {
	if query_row == 0 {
		return min(1, float64(poured))
	}
	var last = []float64{float64(poured)}
	for i := 1; i < N; i++ {
		var tmp = make([]float64, i+1)
		for j := 0; j <= i; j++ {
			if j-1 >= 0 && last[j-1] > 1 {
				tmp[j] += (last[j-1] - 1) * 0.5
			}
			if j <= i-1 && last[j] > 1 {
				tmp[j] += (last[j] - 1) * 0.5
			}

			if i == query_row && j == query_glass {
				return min(1, tmp[j])
			}
		}
		last = tmp
	}

	return 0
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(champagneTower(4, 2, 0))
}
