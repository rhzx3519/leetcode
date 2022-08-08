/**
@author ZhengHao Lou
Date    2021/12/03
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-the-smallest-divisor-given-a-threshold/
思路：二分
 */
func smallestDivisor(nums []int, threshold int) int {
	var l, r = 1, 0
	for _, num := range nums {
		r = max(r, num)
	}
	var m int
	for l < r {
		m = l + (r - l)>>1
		var s int
		for _, num := range nums {
			s += divide(num, m)
		}
		if s <= threshold {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func divide(a, b int) int {
	var c = a / b
	if a%b != 0 {
		c++
	}
	return c
}

func main() {
	fmt.Println(smallestDivisor([]int{1,2,5,9}, 6))
	fmt.Println(smallestDivisor([]int{2,3,5,7,11}, 11))
	fmt.Println(smallestDivisor([]int{19}, 5))
	fmt.Println(smallestDivisor([]int{44,22,33,11,1}, 5))
}