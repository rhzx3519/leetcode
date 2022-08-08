/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/four-divisors/
 */
func sumFourDivisors(nums []int) int {
	var s int
	for _, num := range nums {
		s += factor4(num)
	}
	return s
}

func factor4(x int) int {
	var s int
	set := make(map[int]int)
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x %i == 0 {
			s += i
			s += x/i
			set[i]++
			set[x/i]++
		}
		if len(set) > 2 {
			return 0
		}
	}
	if len(set) == 2 {
		return s + 1 + x
	}
	return 0
}

func main() {
	fmt.Println(sumFourDivisors([]int{21,4,7}))
}