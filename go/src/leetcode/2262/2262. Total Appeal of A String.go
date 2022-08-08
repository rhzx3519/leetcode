/**
@author ZhengHao Lou
Date    2022/05/03
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/total-appeal-of-a-string/
*/
const N = 26

func appealSum(s string) int64 {
	last := make([]int, N)
	for i := range last {
		last[i] = -1
	}
	
	var c, f int64
	for i := range s {
		j := int(s[i] - 'a')
		if last[j] == -1 { // 没有出现过
			f = f + int64(i) + 1
		} else { // 出现过
			f = f + int64(i-last[j])
		}
		last[j] = i
		c += f
	}
	fmt.Println(c)
	return c
}

func main() {
	appealSum("abbca")
	appealSum("code")
}
