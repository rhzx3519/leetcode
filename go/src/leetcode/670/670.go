package main

import "strconv"

/*
*
https://leetcode.cn/problems/maximum-swap/?envType=daily-question&envId=2024-01-22
*/
func maximumSwap(num int) (tot int) {
	tot = num
	s := []byte(strconv.Itoa(num))
	for i := range s {
		for j := range s[:i] {
			s[i], s[j] = s[j], s[i]
			k, _ := strconv.Atoi(string(s))
			tot = max(tot, k)
			s[i], s[j] = s[j], s[i]
		}
	}
	return
}
