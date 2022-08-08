/**
@author ZhengHao Lou
Date    2021/12/10
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

/**
https://leetcode-cn.com/problems/shortest-completing-word/
 */
const N = 26
func shortestCompletingWord(licensePlate string, words []string) string {
	var ans string
	var minLen = math.MaxInt32 >> 1
	target := supply(licensePlate)
	for _, w := range words {
		x := supply(w)
		if is(x, target) {
			if len(w) < minLen {
				minLen = len(w)
				ans = w
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func is(x, target []int) bool {
	for i := 0; i < N; i++ {
		if x[i] < target[i] {
			return false
		}
	}
	return true
}

func supply(w string) []int {
	letters := make([]int, N)
	w = strings.ToLower(w)
	for i := range w {
		if w[i] >= 'a' && w[i] <= 'z' {
			letters[int(w[i] - 'a')]++
		}
	}
	return letters
}

func main() {
	shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"})
	shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"})
	shortestCompletingWord("Ah71752", []string{"suggest","letter","of","husband","easy","education","drug","prevent","writer","old"})
	shortestCompletingWord("OgEu755", []string{"enough","these","play","wide","wonder","box","arrive","money","tax","thus"})
	shortestCompletingWord("iMSlpe4", []string{"claim","consumer","student","camera","public","never","wonder","simple","thought","use"})
}