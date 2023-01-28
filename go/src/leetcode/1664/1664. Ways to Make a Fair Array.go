/**
@author ZhengHao Lou
Date    2023/01/28
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/ways-to-make-a-fair-array/
*/
func waysToMakeFair(nums []int) (tot int) {
	n := len(nums)
	even, odd := make([]int, n+1), make([]int, n+1)
	for i := range nums {
		if i&1 == 0 {
			even[i+1] = even[i] + nums[i]
			odd[i+1] = odd[i]
		} else {
			even[i+1] = even[i]
			odd[i+1] = odd[i] + nums[i]
		}
	}
	
	for i := range nums {
		var a, b int
		if i&1 == 0 {
			a = even[i] + odd[n] - odd[i+1]
			b = odd[i] + even[n] - even[i+1]
		} else {
			b = even[i] + odd[n] - odd[i+1]
			a = odd[i] + even[n] - even[i+1]
		}
		fmt.Println(a, b, i)
		if a == b {
			tot++
		}
	}
	fmt.Println(even, odd, tot)
	return
}

func main() {
	waysToMakeFair([]int{2, 1, 6, 4})
}
