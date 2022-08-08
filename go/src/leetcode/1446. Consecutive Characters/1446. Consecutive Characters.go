/**
@author ZhengHao Lou
Date    2021/12/01
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/consecutive-characters/
 */
func maxPower(s string) int {
	var count int
	var ans = 1
	for i := range s {
		if i > 0 {
			if s[i] != s[i-1] {
				count = 1
			} else {
				count++
			}
			ans = max(count, ans)
		} else {
			count = 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxPower("leetcode"))
	fmt.Println(maxPower("abbcccddddeeeeedcba"))
	fmt.Println(maxPower("triplepillooooow"))
	fmt.Println(maxPower("hooraaaaaaaaaaay"))
	fmt.Println(maxPower("tourist"))
}
