/**
@author ZhengHao Lou
Date    2022/01/22
*/
package main

/**
https://leetcode-cn.com/problems/remove-palindromic-subsequences/
 */
func removePalindromeSub(s string) int {
	bytes := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		bytes = append(bytes, s[i])
	}
	if string(bytes) == s {
		return 1
	}
	return 2
}
