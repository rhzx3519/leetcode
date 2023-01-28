/*
*
@author ZhengHao Lou
Date    2022/12/19
*/
package main

import (
	"fmt"
)

/*
*
https://leetcode.cn/problems/smallest-value-after-replacing-with-sum-of-prime-factors/
*/
func smallestValue(n int) int {
	var prime = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if prime[i] {
			continue
		}
		for j := i + i; j <= n; j += i {
			prime[j] = true
		}
	}
	var ans = n
	var counter = make(map[int]int)
	for {
		var c int
		for i := 2; i <= n; i++ {
			if !prime[i] && n%i == 0 {
				var t int
				for tmp := n; tmp%i == 0; tmp /= i {
					t++
				}
				c += i * t
			}
		}
		fmt.Println(n, c)
		ans = min(ans, c)
		if counter[c] > 0 {
			break
		}
		counter[c]++
		n = c
	}
	fmt.Println(prime)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(smallestValue(3))
}
