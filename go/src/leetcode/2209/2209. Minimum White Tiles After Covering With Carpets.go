/**
@author ZhengHao Lou
Date    2022/03/23
*/
package main

/**
https://leetcode-cn.com/problems/minimum-white-tiles-after-covering-with-carpets/
思路：动态规划
f[i][j]表示用i条地毯去覆盖前j块地板时的最少白色砖块数目
状态转移时，考虑第i条地毯的末尾是否覆盖j
不覆盖，f[i][j] = f[i][j-1] + int(floor[j] == '1')
覆盖，f[i][j] = f[i-1][j-carpetLen]
f[i][j]取两者的较小值
*/
func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	n := len(floor)
	f := make([][]int, numCarpets+1)
	for i := range f {
		f[i] = make([]int, n)
	}
	if floor[0] == '1' {
		f[0][0] = 1
	}
	for j := 1; j < n; j++ {
		var c int
		if floor[j] == '1' {
			c = 1
		}
		f[0][j] = f[0][j-1] + c
	}

	for i := 1; i <= numCarpets; i++ {
		for j := carpetLen; j < n; j++ {
			var c int
			if floor[j] == '1' {
				c = 1
			}
			f[i][j] = min(f[i][j-1]+c, f[i-1][j-carpetLen])
		}
	}

	return f[numCarpets][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
