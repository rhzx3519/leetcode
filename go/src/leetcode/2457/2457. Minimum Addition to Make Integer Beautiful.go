/**
@author ZhengHao Lou
Date    2022/11/02
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/minimum-addition-to-make-integer-beautiful/
思路: sum(n+x) <= target
定义sum(x)为数字x的各位数字之和
进位之后的数的y, sum(y) <= x,
例如467, 个位数字进位之后=470, 十位数字进位之后=500, 百位数字进位之后=1000
*/
func makeIntegerBeautiful(n int64, target int) int64 {
	var tail int64
	for tail = 1; ; tail *= 10 {
		m := n + (tail-n%tail)%tail
		var s int
		for x := m; x != 0; x /= 10 {
			s += int(x % 10)
		}
		if s <= target {
			return m - n
		}
	}
}

func main() {
	fmt.Println(makeIntegerBeautiful(16, 6))
	fmt.Println(makeIntegerBeautiful(467, 6))
	fmt.Println(makeIntegerBeautiful(1, 1))
}
