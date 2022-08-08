/**
@author ZhengHao Lou
Date    2022/02/28
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-number-of-achievable-transfer-requests/
思路：枚举所有的可能方案
time: O(max(m,n) * 2^m)
space: O(
*/
func maximumRequests(n int, requests [][]int) int {
	m := len(requests)
	mask := 1<<m - 1
	var ans int

	for subset := mask; subset != 0; subset = (subset - 1) & mask {
		ind := make([]int, n)
		var x int
		for i := range requests {
			if subset>>i&1 == 1 {
				ind[requests[i][0]]--
				ind[requests[i][1]]++
				x++
			}
		}
		var valid = true
		for i := range ind {
			if ind[i] != 0 {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		if x > ans {
			ans = x
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	maximumRequests(5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}})
	maximumRequests(3, [][]int{{0, 0}, {1, 2}, {2, 1}})
}
