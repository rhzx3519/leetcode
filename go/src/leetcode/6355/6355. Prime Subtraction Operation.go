package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/prime-subtraction-operation/

		nums[i] - p > nums[i-1]
		p < nums[i] - nums[i-1]
	}
*/
func primeSubOperation(nums []int) bool {
	var pre int
	for _, x := range nums {
		if x <= pre {
			return false
		}
		pre = x - primes[sort.SearchInts(primes, x-pre)-1]
	}

	return true
}

const N = 1001

var primes = []int{0}

func init() {
	notPrime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		if notPrime[i] {
			continue
		}
		primes = append(primes, i)
		for j := 2; i*j <= N; j++ {
			notPrime[i*j] = true
		}
	}
}

func main() {
	fmt.Println(primes)
}
