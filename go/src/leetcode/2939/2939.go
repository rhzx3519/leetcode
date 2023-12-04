package main

import (
	"fmt"
	"math/bits"
)

/*
*
https://leetcode.cn/problems/maximum-xor-product/
*/
const mod int = 1e9 + 7

func maximumXorProduct(A int64, B int64, n int) int {
	if A < B {
		return maximumXorProduct(B, A, n)
	}
	a, b := int(A), int(B)

	mask := 1<<n - 1
	ax := a &^ mask // 第n为左边无法被 x 影响的部分
	bx := b &^ mask
	a &= mask // 低于n位可以被 x 影响的部分
	b &= mask

	left := a ^ b      // 可分配： a XOR x 和 b XOR x 一个是0，一个是1
	one := mask ^ left // 无需分配: a XOR x 和 b XOR x 均为1
	ax |= one          // 先加到异或结果中
	bx |= one

	if left > 0 && ax == bx {
		// 最高位分配给a, 其余分配给b
		highBit := 1 << (bits.Len(uint(left)) - 1)
		ax |= highBit
		left ^= highBit
	}
	// 如果a的最高位比b大，无论如何分配，a都比b大，所以把1全部分配给b
	bx |= left

	return ax % mod * (bx % mod) % mod
}

func main() {
	fmt.Println(maximumXorProduct(12, 5, 4))
	fmt.Println(maximumXorProduct(6, 7, 5))
	fmt.Println(maximumXorProduct(1, 6, 3))
}
