package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 || abs(dividend) < abs(divisor) {
		return 0
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	ans := int64(0)
	a, b := abs(dividend), abs(divisor)
	for a >= b {
		k := int64(1)
		t := b
		for a > t<<1 {
			t <<= 1
			k <<= 1
		}
		ans += k
		a -= t
	}

	tmp := int64(sign) * ans
	if tmp > math.MaxInt32 || tmp < math.MinInt32 {
		return math.MaxInt32
	}
	return int(tmp)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
}
