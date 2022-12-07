/**
@author ZhengHao Lou
Date    2022/10/31
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/magical-string/
*/
func magicalString(n int) (c1 int) {
	c1++
	if n == 1 {
		return
	}
	var nums = make([]int, n)
	nums[0] = 1
	nums[1] = 2
	var slow, fast = 1, 1
	var x = 2
	for ; slow < n; slow++ {
		if nums[slow] == 1 {
			c1++
		}
		for k := 0; k < nums[slow] && fast < n; k, fast = k+1, fast+1 {
			nums[fast] = x
		}
		if x == 1 {
			x = 2
		} else {
			x = 1
		}
	}

	fmt.Println(nums)
	return
}

func main() {
	magicalString(11)
}
