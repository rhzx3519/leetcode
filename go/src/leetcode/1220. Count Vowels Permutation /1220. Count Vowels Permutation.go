/**
@author ZhengHao Lou
Date    2022/01/17
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-vowels-permutation/
思路：DP，跳台阶
f[i][j]表示长度为i，以j结尾的字符串排列数量
"aeiou"分别对应01234
以字母'a'为例
f[1][0] = 1
f[i][0] = f[i-1][1] + f[i-1][2] + f[i-1][4]

time: O(n)
space: O(n)
 */
const MOD int = 1e9 + 7
func countVowelPermutation(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 5)
	}
	for i := range f[1] {
		f[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		f[i][0] = f[i-1][1] + f[i-1][2] + f[i-1][4]	// a
		f[i][1] = f[i-1][0] + f[i-1][2]  // e
		f[i][2] = f[i-1][1] + f[i-1][3]	 // i
		f[i][3] = f[i-1][2]				 // o
		f[i][4] = f[i-1][2] + f[i-1][3]	 // u
		for j := range f[i] {
			f[i][j] %= MOD
		}
	}

	var total int
	for i := range f[n] {
		total = (total + f[n][i]) % MOD
	}
	return total
}

func main() {
	fmt.Println(countVowelPermutation(2))
}