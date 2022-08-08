package main

import (
	"fmt"
	"math"
)

/**
思路：每一位的0和1的数目都是2^(p-1) - 1
 */
var mod = int64(math.Pow10(9)) + int64(7)

func minNonZeroProduct(p int) int {
	if p == 1 {
		return 1
	}
	total := 1<<p - 1

	n1 := int64(1<<p - 1 - 1) % mod
	y := int64((total - 1)>>1)
	x := pow(n1, y)
	return int(x * int64(total) % mod)
}

// mod=3, xx=2, xx*xx=4
func pow(x, y int64) int64 {
	var z = int64(1)
	for y > 0 {
		xx := x
		t := int64(1)
		for y - t<<1 >= 0 {
			xx = xx*xx % mod
			t <<= 1
		}
		z = z*xx%mod
		y -= t
	}

	//fmt.Println(x, y, z)
	return z
}

func main() {
	//for i := int64(1); i <= int64(10); i++ {
	//	pow(2, i)
	//}

	fmt.Println(minNonZeroProduct(1))
	fmt.Println(minNonZeroProduct(2))
	fmt.Println(minNonZeroProduct(3))
	fmt.Println(minNonZeroProduct(4)) // 581202553
	fmt.Println(minNonZeroProduct(5)) // 202795991
	fmt.Println(minNonZeroProduct(35)) // 112322054
	fmt.Println(minNonZeroProduct(60)) // 220750044
}
