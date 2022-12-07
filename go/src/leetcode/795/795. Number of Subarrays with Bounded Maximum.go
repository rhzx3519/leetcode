/*
*
@author ZhengHao Lou
Date    2022/11/24
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/
*/
func numSubarrayBoundedMax(nums []int, left int, right int) (tot int) {
	var lastValid, l = -1, 0
	for i, x := range nums {
		if x > right {
			lastValid = i
			l = i + 1
		} else if x < left {
			tot += lastValid - l + 1
		} else {
			tot += i - l + 1
			lastValid = i
		}

	}
	fmt.Println(tot)
	return
}

func main() {
	numSubarrayBoundedMax([]int{4, 2, 1}, 2, 3) // 2
	numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3)
	numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8)
	numSubarrayBoundedMax([]int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}, 32, 69)    // 22
	numSubarrayBoundedMax([]int{1, 3, 3}, 2, 2)                                  // 0
	numSubarrayBoundedMax([]int{16, 69, 88, 85, 79, 87, 37, 33, 39, 34}, 55, 57) // 0

}
