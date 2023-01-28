/*
*
@author ZhengHao Lou
Date    2023/01/02
*/
package main

import (
	"math"
)

/*
*
https://leetcode.cn/problems/closest-prime-numbers-in-range/
*/
func closestPrimes(left int, right int) []int {
	var notPrime = make([]bool, right+1)
	notPrime[1] = true
	for i := 2; i <= right; i++ {
		if notPrime[i] {
			continue
		}
		for j := 2; i*j <= right; j++ {
			notPrime[i*j] = true
		}
	}

	var ans = []int{-1, -1}
	var gap = math.MaxInt32
	var pair []int
	for i := left; i <= right; i++ {
		if !notPrime[i] {
			pair = append(pair, i)
			if len(pair) > 2 {
				pair = pair[1:]
			}
			if len(pair) >= 2 {
				if pair[1]-pair[0] < gap {
					gap = pair[1] - pair[0]
					ans = []int{pair[0], pair[1]}
				}
			}

		}
	}
	return ans
}

func main() {
	//closestPrimes(10, 19)
	//closestPrimes(4, 6)
	closestPrimes(19, 31)
}
