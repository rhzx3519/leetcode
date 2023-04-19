/**
@author ZhengHao Lou
Date    2021/10/08
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/repeated-dna-sequences/
思路：字符串Hash
time: O(n)
space: O(n)
 */
const (
	N int = 1e5 + 10
	M int = 10
	P int = 131313
)

var (
	h, p = make([]int, N), make([]int, N)
)
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	p[0] = 1
	for i := 1; i <= n; i++ {	// 将s转换为hash数组
		h[i] = h[i - 1]*P + int(s[i-1])
		p[i] = p[i - 1]*P
	}

	var ans = []string{}
	mapper := make(map[int]int)
	for i := 1; i + M - 1 <= n; i++ {
		var j = i + M - 1
		hash := h[j] - h[i-1]*p[j-i+1]	// 求s[i-1:j]的hash值
		c := mapper[hash]
		if c == 1 {
			ans = append(ans, s[i-1:j])
		}
		mapper[hash]++
	}

	return ans
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
