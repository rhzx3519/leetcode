/*
*
@author ZhengHao Lou
Date    2023/01/02
*/
package main

/*
*
https://leetcode.cn/problems/distinct-prime-factors-of-product-of-array/

思路：[素因数分解](https://oi-wiki.org/math/number-theory/pollard-rho/)
*/
func distinctPrimeFactors(nums []int) int {
	var counter = make(map[int]int)
	for _, x := range nums {
		factors := primeFactor(x)
		for i := range factors {
			counter[factors[i]]++
		}
	}
	return len(counter)
}

func primeFactor(x int) []int {
	var factors []int
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			for x%i == 0 {
				x /= i
			}
			factors = append(factors, i)
		}
	}
	if x != 1 {
		factors = append(factors, x)
	}
	return factors
}
