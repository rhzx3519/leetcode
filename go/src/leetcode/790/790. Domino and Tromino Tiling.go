/**
@author ZhengHao Lou
Date    2022/11/12
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/domino-and-tromino-tiling/
思路：动态规划

*/
const MOD int64 = 1e9 + 7

func numTilings(n int) int {
	f := [4]int64{0, 0, 0, 1}
	for i := 0; i < n; i++ {
		var nf = [4]int64{}
		nf[0] = f[3]
		nf[1] = (f[0] + f[2]) % MOD
		nf[2] = (f[0] + f[1]) % MOD
		nf[3] = (f[0] + f[1] + f[2] + f[3]) % MOD
		f = nf
	}
	return int(f[3])
}

func main() {
	fmt.Println(numTilings(3))
	//fmt.Println(numTilings(1))
}
