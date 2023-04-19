/**
@author ZhengHao Lou
Date    2021/11/08
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/vowels-of-all-substrings/
思路：动态规划
f[i]表示加入word[i-1]后，比上一个状态f[i-1]能够增加的元音数量
f[i] = f[i-1] + i, if word[i-1] is a vowel, else,
f[i] = f[i-1]
e.g. "ab", 加入"a"后，多了"aba", "ba", "a" 3个元音数
time: O(N)
space: O(N)
 */
const vowels string = "aeiou"
func countVowels(word string) int64 {
	n := len(word)
	f := make([]int64, n + 1)
	for i := 1; i <= n; i++ {
		j := strings.IndexByte(vowels, word[i-1])
		if j != -1 {
			f[i] = f[i-1] + int64(i)
		} else {
			f[i] = f[i-1]
		}
	}
	var sum int64
	for i := range f {
		sum += f[i]
	}
	return sum
}

func main() {
	fmt.Println(countVowels("aba"))
	fmt.Println(countVowels("abc"))
	fmt.Println(countVowels("ltcd"))
	fmt.Println(countVowels("noosabasboosa"))
}