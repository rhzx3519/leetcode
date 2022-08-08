/**
@author ZhengHao Lou
Date    2022/01/10
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/additive-number/
思路：回溯查找
time: O(n^2)
space: O(n)
 */
func isAdditiveNumber(num string) bool {
	n := len(num)

	var dfs func(idx int, seq []int) bool
	dfs = func(idx int, seq []int) bool {
		if idx == n {
			return len(seq) >= 3
		}
		var cur int
		for i := idx; i < n; i++ {
			cur = cur*10 + int(num[i] - '0')
			if i > idx && num[idx] == '0' {
				continue
			}
			if len(seq) < 2 || (len(seq) >= 2 && seq[len(seq)-2] + seq[len(seq)-1] == cur) {
				seq = append(seq, cur)
				if dfs(i+1, seq) {
					return true
				}
				seq = seq[:len(seq) - 1]
			}

		}

		return false
	}

	return dfs(0, []int{})
}

func main() {
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("1023"))
}
