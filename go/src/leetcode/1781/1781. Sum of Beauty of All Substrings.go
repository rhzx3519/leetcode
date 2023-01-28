/*
*
@author ZhengHao Lou
Date    2022/12/12
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/sum-of-beauty-of-all-substrings/
*/
const N = 26

func beautySum(s string) (tot int) {
	n := len(s)
	for i := range s {
		var freq = [N]int{}
		for j := i; j < n; j++ {
			freq[s[j]-'a']++
			var l, r = n, 0
			for _, k := range freq {
				if k < l {
					l = k
				}
				if k > r {
					r = k
				}
			}
			tot += r - l
		}
		fmt.Println(freq)
	}
	return tot
}
