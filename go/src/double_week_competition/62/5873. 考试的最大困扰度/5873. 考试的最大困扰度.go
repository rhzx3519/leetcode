/**
@author ZhengHao Lou
Date    2021/10/02
*/
package main

import "fmt"

/**
滑动窗口
 */
func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(slidewindow(answerKey, k, 'T'), slidewindow(answerKey, k, 'F'))
}

func slidewindow(s string, k int, ch byte) int {
	var cnt int
	var l int
	var maxLen int
	for r := range s {
		if s[r] != ch {
			cnt++
		}

		for l <= r && cnt > k {
			if s[l] != ch {
				cnt--
			}
			l++
		}
		maxLen = max(maxLen, r - l + 1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxConsecutiveAnswers("TTFF", 2))
	fmt.Println(maxConsecutiveAnswers("TFFT", 1))
	fmt.Println(maxConsecutiveAnswers("TTFTTFTT", 1))
	fmt.Println(maxConsecutiveAnswers("FFTFTTTFFF", 1))
}
