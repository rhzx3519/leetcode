/**
@author ZhengHao Lou
Date    2021/10/18
*/
package main

import (
	"strconv"
	"strings"
)

/**
https://leetcode-cn.com/problems/check-if-numbers-are-ascending-in-a-sentence/
 */
func areNumbersAscending(s string) bool {
	var last int
	tokens := strings.Split(s, " ")
	for _, t := range tokens {
		i, err := strconv.Atoi(t)
		if err == nil {
			if i <= last {
				return false
			}
			last = i
		}
	}
	return true
}
