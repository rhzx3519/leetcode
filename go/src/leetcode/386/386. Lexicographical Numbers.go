/**
@author ZhengHao Lou
Date    2022/04/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/lexicographical-numbers/
*/
func lexicalOrder(n int) []int {
	var ans = make([]int, n)
	var cur = 1
	for i := 0; i < n; i++ {
		ans[i] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			for cur%10 == 9 || cur+1 > n {
				cur /= 10
			}
			cur++
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	lexicalOrder(2)
}
