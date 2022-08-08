/**
@author ZhengHao Lou
Date    2022/03/19
*/
package main

import "fmt"

func maximumSubsequenceCount(text string, pattern string) int64 {
	n := len(text)
	prefix, suffix := make([]int, n+1), make([]int, n+1)
	for i := 0; i < n; i++ {
		var c int
		if text[i] == pattern[0] {
			c++
		}
		prefix[i+1] = prefix[i] + c
	}
	for i := n - 1; i >= 0; i-- {
		var c int
		if text[i] == pattern[1] {
			c++
		}
		suffix[i] = suffix[i+1] + c
	}

	//fmt.Println(prefix)
	//fmt.Println(suffix)
	var x int64
	for i := range text {
		if text[i] == pattern[1] {
			x += int64(prefix[i])
		}
	}
	x += int64(prefix[n])

	var y int64
	for i := n - 1; i >= 0; i-- {
		if text[i] == pattern[0] {
			y += int64(suffix[i+1])
		}
	}
	y += int64(suffix[0])

	fmt.Println(max(x, y))
	return max(x, y)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	maximumSubsequenceCount("abdcdbc", "ac")                         // 4
	maximumSubsequenceCount("aabb", "ab")                            // 6
	maximumSubsequenceCount("fwymvreuftzgrcrxczjacqovduqaiig", "yy") // 1
	maximumSubsequenceCount("y", "ay")                               // 1
	maximumSubsequenceCount("yy", "ay")                              // 1
}
