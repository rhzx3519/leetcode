/**
@author ZhengHao Lou
Date    2022/03/29
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/
思路：滑动窗口
*/
func maxConsecutiveAnswers(answerKey string, k int) int {
	var mx int
	//n := len(answerKey)
	var f, l int
	for i := range answerKey {
		switch answerKey[i] {
		case 'T':
		case 'F':
			f++
		}
		for ; f > k && l <= i; l++ {
			if answerKey[l] == 'F' {
				f--
			}
		}
		mx = max(mx, i-l+1)
	}

	var t int
	l = 0
	for i := range answerKey {
		switch answerKey[i] {
		case 'T':
			t++
		case 'F':
		}
		for ; t > k && l <= i; l++ {
			if answerKey[l] == 'T' {
				t--
			}
		}
		mx = max(mx, i-l+1)
	}
	fmt.Println(mx)
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxConsecutiveAnswers("TTFF", 2)
	maxConsecutiveAnswers("TFFT", 1)
	maxConsecutiveAnswers("TTFTTFTT", 1)
	maxConsecutiveAnswers("TFF", 1)
}
