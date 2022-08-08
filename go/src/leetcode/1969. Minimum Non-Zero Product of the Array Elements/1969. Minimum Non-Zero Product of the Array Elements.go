/**
@author ZhengHao Lou
Date    2021/12/04
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimum-non-zero-product-of-the-array-elements/
 */
const mod int = 1e9 + 7

func minNonZeroProduct(p int) int {
	return (1<<p - 1) % mod * pow(1<<p-2, 1<<(p-1)-1) % mod
}

func pow(x, n int) int {
	res := 1
	for x %= mod; n > 0; n >>= 1 {
		res = res * x % mod // 由于 n 的二进制全是 1，所以无需判断 n 的奇偶性
		x = x * x % mod
	}
	return res
}

func main() {
	fmt.Println(minNonZeroProduct(1))
}