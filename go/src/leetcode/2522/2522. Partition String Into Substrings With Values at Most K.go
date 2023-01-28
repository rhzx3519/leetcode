/*
*
@author ZhengHao Lou
Date    2023/01/02
*/
package main

/*
*
https://leetcode.cn/problems/partition-string-into-substrings-with-values-at-most-k/
*/

func minimumPartition(s string, k int) int {
	var c, x = 1, 0
	for i := range s {
		v := int(s[i] - '0')
		if v > k {
			return -1
		}
		x = x*10 + v
		if x > k {
			x = v
			c++
		}
	}
	return c
}

func main() {
	minimumPartition("7614953199576414777", 16)
}
