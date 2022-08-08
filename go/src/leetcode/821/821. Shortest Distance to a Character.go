/**
@author ZhengHao Lou
Date    2022/04/19
*/
package main

import "fmt"

func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	for i := range s {
		ans[i] = n + 1
	}
	var arr []int
	for i := range s {
		if s[i] == c {
			arr = append(arr, i)
			ans[i] = 0
		}
	}

	var j int
	for i := range s {
		if j < len(arr) && i > arr[j] {
			j++
		}
		if j > 0 {
			ans[i] = min(ans[i], i-arr[j-1])
		}
		if j < len(arr) {
			ans[i] = min(ans[i], arr[j]-i)
		}
	}

	fmt.Println(ans)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//shortestToChar("loveleetcode", 'e')
	//shortestToChar("aaab", 'b')
	shortestToChar("abaa", 'b')
}
