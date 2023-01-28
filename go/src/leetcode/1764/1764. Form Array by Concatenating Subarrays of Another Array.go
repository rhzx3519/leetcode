/*
*
@author ZhengHao Lou
Date    2022/12/17
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/form-array-by-concatenating-subarrays-of-another-array/
*/
func canChoose(groups [][]int, nums []int) bool {
	var j int
	n := len(nums)
OUT:
	for i := range groups {
		m := len(groups[i])
		for {
			if j+m > n {
				return false
			}
			var k int
			for ; k < m; k++ {
				if groups[i][k] != nums[j+k] {
					break
				}
			}
			if k == m {
				j += m
				continue OUT
			}
			j++
		}

	}
	return true
}

func main() {
	fmt.Println(canChoose([][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0}))
}
