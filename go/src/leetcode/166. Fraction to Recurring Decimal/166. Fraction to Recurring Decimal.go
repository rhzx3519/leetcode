/**
@author ZhengHao Lou
Date    2021/10/03
*/
package main

import (
	"fmt"
	"strconv"
)

/**
https://leetcode-cn.com/problems/fraction-to-recurring-decimal/
 */
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	bytes := []byte{}
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		bytes = append(bytes, '-')
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	if numerator >= denominator && numerator % denominator == 0 {
		return fmt.Sprintf("%v%v", string(bytes), numerator / denominator)
	}
	d := numerator / denominator	// 整数部分
	bytes = append(bytes, []byte(strconv.Itoa(d))...)
	bytes = append(bytes, '.')
	r := numerator % denominator

	mapper := make(map[int]int)
	for r != 0 && mapper[r] == 0 {
		mapper[r] = len(bytes)
		r *= 10
		bytes = append(bytes, byte('0' + r / denominator))
		r %= denominator
	}

	if r > 0 {
		i := mapper[r]
		bytes = append(bytes[:i], append([]byte{'('}, bytes[i:]...)...)
		bytes = append(bytes, ')')
	}
	return string(bytes)
}


func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	fmt.Println(fractionToDecimal(1, 6))
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(1, 5))
}
