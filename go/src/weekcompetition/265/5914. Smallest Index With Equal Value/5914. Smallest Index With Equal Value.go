/**
@author ZhengHao Lou
Date    2021/11/01
*/
package main

func smallestEqual(nums []int) int {
	for i, num := range nums {
		if i % 10 == num {
			return i
		}
	}
	return -1
}
