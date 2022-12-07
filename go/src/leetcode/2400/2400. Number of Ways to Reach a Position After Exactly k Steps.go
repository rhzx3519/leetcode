/**
@author ZhengHao Lou
Date    2022/09/05
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/
思路：f[i][j] 表示从startPos触发，到达i还剩j步的走法数量
time: O(k^2)
*/
const (
	MOD int = 1e9 + 7
	N       = 1001
)

func numberOfWays(startPos int, endPos int, k int) int {
	startPos, endPos = startPos+N, endPos+N
	f := make([][]int, N<<2)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	f[startPos+1][1] = 1
	f[startPos-1][1] = 1

	for t := 2; t <= k; t++ {
		for i := startPos - k; i <= startPos+k; i++ {
			f[i][t] = (f[i-1][t-1] + f[i+1][t-1]) % MOD
		}
	}

	return f[endPos][k]
}

func main() {
	fmt.Println(numberOfWays(1, 2, 3))
	fmt.Println(numberOfWays(2, 5, 10))
	fmt.Println(numberOfWays(301, 308, 724))
}
