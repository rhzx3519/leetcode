/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
https://leetcode-cn.com/problems/max-difference-you-can-get-from-changing-an-integer/
思路：暴力
 */
func maxDiff(num int) int {
	var a, b = 0, math.MaxInt32 >> 1
	d := strconv.Itoa(num)
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			t := strings.ReplaceAll(d, string(i), string(j))
			ti, _ := strconv.Atoi(t)
			if t[0] == '0' || ti == 0 {
				continue
			}
			a = max(a, ti)
			b = min(b, ti)
		}
	}
	return a - b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxDiff(555))
	fmt.Println(maxDiff(9))
	fmt.Println(maxDiff(123456))
	fmt.Println(maxDiff(10000))
	fmt.Println(maxDiff(9288))
}