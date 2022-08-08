/**
@author ZhengHao Lou
Date    2022/02/06
*/
package main

import (
	"fmt"
	"sort"
)

func smallestNumber(num int64) int64 {
	if num == 0 {
		return num
	}
	var sign int64 = 1
	if num < 0 {
		num = -num
		sign = -1
	}

	var digits []int64
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})
	var c0, i int
	for i = 0; digits[i] == 0; i++ {
		c0++
	}
	var t int64
	if sign > 0 {
		t = digits[i]
		i++
		for c0 != 0 {
			t *= 10
			c0--
		}
		for ; i < len(digits); i++ {
			t = t*10 + digits[i]
		}
	} else {
		for i := len(digits) - 1; i >= 0; i-- {
			t = t*10 + digits[i]
		}
		t *= -1
	}

	fmt.Println(t)
	return t
}

func main() {
	smallestNumber(-112)
	smallestNumber(-7605)
}
