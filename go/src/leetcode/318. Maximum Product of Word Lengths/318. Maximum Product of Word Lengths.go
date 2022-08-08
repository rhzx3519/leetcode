/**
@author ZhengHao Lou
Date    2021/11/17
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-product-of-word-lengths/
 */
const N = 26
func maxProduct(words []string) int {
	n := len(words)
	d := make([]int, n)
	for i := range words {
		for _, c := range words[i] {
			d[i] |= 1 << int(c - 'a')
		}
	}
	fmt.Println(d)
	var ans int
	for i := range words {
		for j := i+1; j < n; j++ {
			if d[i] & d[j] != 0 {
				continue
			}
			ans = max(ans, len(words[i]) * len(words[j]))
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
	fmt.Println(maxProduct([]string{"abcw","baz","foo","bar","xtfn","abcdef"}))
	fmt.Println(maxProduct([]string{"a","ab","abc","d","cd","bcd","abcd"}))
	fmt.Println(maxProduct([]string{"a","aa","aaa","aaaa"}))
}