package main

import (
	"fmt"
	"math"
)

var mod = int(math.Pow10(9)) + 7

func countGoodNumbers(n int64) int {
	if n ==1 {
		return 5
	}
	var c1, _ = 5, 4
	var c3 = 20
	var ans = 1
	var n2 = n>>1
	var b = int64(1)
	for n2 > 0 {
		c3 = 20
		b = 1
		for n2 > b*2 {
			c3 = c3*c3%mod
			b <<= 1
		}
		n2 -= b
		ans = (ans * c3)%mod
	}


	if n&1 == 1 {
		ans = (ans * c1) % mod
	}

	return ans
}

func main() {
	fmt.Println(countGoodNumbers(1))
	fmt.Println(countGoodNumbers(4))
	fmt.Println(countGoodNumbers(50))
}
