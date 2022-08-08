/**
@author ZhengHao Lou
Date    2022/05/25
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/unique-substrings-in-wraparound-string/
*/
const N = 26

func findSubstringInWraproundString(p string) (ans int) {
	f := make([]int, N)
	var c int
	for i := range p {
		if i > 0 && (int(p[i-1]-'a')+N+1)%N == int(p[i]-'a') {
			c++
		} else {
			c = 1
		}
		j := int(p[i] - 'a')
		f[j] = max(f[j], c)
	}
	for i := range f {
		ans += f[i]
	}
	fmt.Println(f, ans)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findSubstringInWraproundString("a")
	findSubstringInWraproundString("cac")
	findSubstringInWraproundString("zab")
	findSubstringInWraproundString("zababc") // 9
}
