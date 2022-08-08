/**
@author ZhengHao Lou
Date    2022/07/11
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-the-number-of-ideal-arrays/
思路：f[i][j]表示长度为i，结尾为j的理想数组的数目
*/
const MOD int = 1e9 + 7

func idealArrays(n int, maxValue int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, maxValue+1)
	}

	div := map[int][]int{}
	for i := 1; i <= maxValue; i++ {
		f[1][i] = 1
		for k := 1; i/k != 0; k++ {
			if i%k == 0 {
				div[i] = append(div[i], i/k)
			}
		}
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= maxValue; j++ {
			for _, d := range div[j] {
				f[i][j] = (f[i][j] + f[i-1][d]) % MOD
			}
		}
	}

	var ans int
	for j := 1; j <= maxValue; j++ {
		ans = (ans + f[n][j]) % MOD
	}

	fmt.Println(f, ans)
	return ans
}

func main() {
	idealArrays(2, 5)
	idealArrays(5, 3)
	idealArrays(5878, 2900)
}
