/**
@author ZhengHao Lou
Date    2022/03/03
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/
*/
const N = 26

func minSteps(s string, t string) int {
	var sp, tp = make([]int, N), make([]int, N)
	for i := range s {
		sp[int(s[i]-'a')]++
	}
	for i := range t {
		tp[int(t[i]-'a')]++
	}

	var step int
	for i := 0; i < N; i++ {
		step += abs(sp[i] - tp[i])
	}

	fmt.Println(step)
	return step
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	minSteps("leetcode", "coats")
	minSteps("night", "thing")
}
