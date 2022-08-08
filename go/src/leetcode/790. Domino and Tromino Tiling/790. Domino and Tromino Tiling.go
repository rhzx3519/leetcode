/**
@author ZhengHao Lou
Date    2021/11/19
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/domino-and-tromino-tiling/
思路：动态规划
https://leetcode-cn.com/problems/domino-and-tromino-tiling/solution/duo-mi-nuo-yu-tuo-mi-nuo-ping-pu-by-leetcode/
 */

const Mod = 1e9 + 7
func numTilings(n int) int {
	f := [4]int64{1, 0, 0, 0}
	for i := 0; i < n; i++ {
		nf := [4]int64{}
		nf[0] = (f[0] + f[3])%Mod
		nf[1] = (f[0] + f[2])%Mod
		nf[2] = (f[0] + f[1])%Mod
		nf[3] = (f[0] + f[1] + f[2])%Mod
		f = nf
	}
	return int(f[0])
}

func main() {
	fmt.Println(numTilings(30))
}
