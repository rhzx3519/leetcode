/**
@author ZhengHao Lou
Date    2021/12/04
*/
package main

/**
https://leetcode-cn.com/problems/ransom-note/
 */
const N = 26
func canConstruct(ransomNote string, magazine string) bool {

	var mapping func(string) [N]int
	mapping = func(s string) [N]int {
		m := [N]int{}
		for i := range s {
			m[int(s[i] - 'a')]++
		}
		return m
	}

	mr, ma := mapping(ransomNote), mapping(magazine)
	for i := range mr {
		if ma[i] < mr[i] {
			return false
		}
	}
	return true
}
