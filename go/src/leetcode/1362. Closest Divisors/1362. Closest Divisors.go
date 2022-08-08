/**
@author ZhengHao Lou
Date    2021/12/03
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/closest-divisors/
 */
func closestDivisors(num int) []int {
	var ans = []int{}
	var t = math.MaxInt32 >> 1
	for _, i := range []int{num+1, num+2} {
		for j := int(math.Sqrt(float64(i))); j >= 1; j-- {
			if i % j == 0 {
				if abs(j - i/j) < t {
					t = abs(j - i/j)
					ans = []int{j, i/j}
				}
				break
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(closestDivisors(999))
}
