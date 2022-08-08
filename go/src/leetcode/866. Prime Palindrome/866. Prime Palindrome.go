/**
@author ZhengHao Lou
Date    2021/11/17
*/
package main

import (
	"math"
)

/**
https://leetcode-cn.com/problems/prime-palindrome/
思路：暴力
 */

func primePalindrome(n int) int {
	for {
		if isPrime(n) && reverse(n) == n {
			return n
		}

		if n >= 1e7 &&  n < 1e8 {
			n = 1e8
			continue
		}
		n++
	}
	return -1
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func reverse(x int) int {
	var s int
	for x != 0 {
		s = s*10 + x%10
		x /= 10
	}
	return s
}

func main() {
	primePalindrome(6)
	primePalindrome(8)
	primePalindrome(13)
	primePalindrome(1)
	primePalindrome(9989900)	// 100030001
	//fmt.Println(10277467 > 1e7)
}
