/**
@author ZhengHao Lou
Date    2021/10/16
*/
package main

import (
	"fmt"
	"strconv"
)

/**
https://leetcode-cn.com/problems/expression-add-operators/
思路：回溯查找
 */
func addOperators(num string, target int) []string {
	ans := []string{}
	n := len(num)
	var backtrace func(u int, prev int, cur int, bytes []byte)

	backtrace = func(u int, prev int, cur int, bytes []byte) {
		if u == n {
			if cur == target {
				ans = append(ans, string(bytes))
			}
			return
		}

		for i := u; i < n; i++ {
			if i != u && num[u] == '0' {
				break
			}
			next, _ := strconv.Atoi(num[u:i+1])
			if u == 0 {
				backtrace(i+1, next, next, []byte(num[u:i+1]))
			} else {
				backtrace(i+1, next, cur + next, append(bytes, append([]byte{'+'}, []byte(num[u:i+1])...)...))
				backtrace(i+1, -next, cur - next, append(bytes, append([]byte{'-'}, []byte(num[u:i+1])...)...))
				x := prev * next
				backtrace(i+1, x, cur - prev + x, append(bytes, append([]byte{'*'}, []byte(num[u:i+1])...)...))
			}
		}
	}

	backtrace(0, 0, 0, []byte{})

	return ans
}

func main() {
	fmt.Println(addOperators("123", 6))
	fmt.Println(addOperators("232", 8))
	fmt.Println(addOperators("105", 5))
}
