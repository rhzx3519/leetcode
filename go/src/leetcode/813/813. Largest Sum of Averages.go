/*
*
@author ZhengHao Lou
Date    2022/11/28
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/largest-sum-of-averages/
*/
func largestSumOfAverages(nums []int, K int) float64 {
	n := len(nums)
	f := make([][]float64, n+1)
	for i := range f {
		f[i] = make([]float64, K+1)
	}
	s := make([]int, n+1)
	for i := range nums {
		s[i+1] = s[i] + nums[i]
		f[i+1][1] = float64(s[i+1]) / float64(i+1)
	}
	for i := 1; i <= n; i++ {
		for k := 2; k <= K; k++ {
			for j := 0; j < i; j++ {
				f[i][k] = max(f[i][k], f[j][k-1]+float64(s[i]-s[j])/float64(i-j))
			}
		}
	}

	fmt.Println(f[n][K])
	return f[n][K]
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3)
}
