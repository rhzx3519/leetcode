/**
@author ZhengHao Lou
Date    2022/02/05
*/
package main

import (
	"sort"
)

func minimumSum(num int) int {
	digits := []int{}
	for i := 0; i < 4; i++ {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)
	return digits[0]*10 + digits[2] + digits[1]*10 + digits[3]
}
