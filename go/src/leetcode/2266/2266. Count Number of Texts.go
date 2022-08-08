/**
@author ZhengHao Lou
Date    2022/05/09
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-number-of-texts/
思路：动态规划，跳楼梯
*/
const MOD int = 1e9 + 7

var keyboard = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func countTexts(pressedKeys string) int {
	n := len(pressedKeys)
	f := make([]int, n+1)
	f[0] = 1

	var c int
	var b byte
	for i := range pressedKeys {
		if pressedKeys[i] == b {
			c++
		} else {
			c = 1
		}
		j := int(pressedKeys[i] - '0')
		for k := 0; k < min(c, len(keyboard[j])); k++ {
			f[i+1] = (f[i+1] + f[i-k]) % MOD
		}

		b = pressedKeys[i]
	}
	fmt.Println(f)
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	countTexts("222222222222222222222222222222222222")
}
