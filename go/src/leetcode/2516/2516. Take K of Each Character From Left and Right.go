/*
*
@author ZhengHao Lou
Date    2022/12/26
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right/
思路：滑动窗口
*/
func takeCharacters(s string, k int) int {
	var counter = [3]int{}
	for i := range s {
		counter[int(s[i]-'a')]++
	}
	for _, v := range counter {
		if v < k {
			return -1
		}
	}

	n := len(s)
	var ans = n
	var l, r int
	for ; r < n; r++ {
		counter[int(s[r]-'a')]--
		for l < r && !enough(counter, k) {
			counter[int(s[l]-'a')]++
			l++
		}
		if enough(counter, k) {
			ans = min(ans, n-(r-l+1))
		}
	}

	return ans
}

func enough(counter [3]int, k int) bool {
	for i := range counter {
		if counter[i] < k {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2))
	fmt.Println(takeCharacters("a", 1))
	fmt.Println(takeCharacters("abc", 1))
}
