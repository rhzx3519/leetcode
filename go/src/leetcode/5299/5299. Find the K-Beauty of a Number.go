/**
@author ZhengHao Lou
Date    2022/05/15
*/
package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	var c int
	s := fmt.Sprintf("%v", num)
	for i := 0; i+k <= len(s); i++ {
		x, _ := strconv.Atoi(s[i : i+k])
		if x == 0 {
			continue
		}
		if num%x == 0 {
			c++
		}
	}
	return c
}
