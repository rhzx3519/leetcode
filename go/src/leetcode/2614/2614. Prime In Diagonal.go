package main

/*
*
https://leetcode.cn/problems/prime-in-diagonal/
*/
const N int = 4e6 + 1

var primes = make([]bool, N)

func init() {
	primes[1] = true
	for d := 2; d*d < N; d++ {
		if !primes[d] {
			for j := d + d; j < N; j += d {
				primes[j] = true
			}
		}
	}
}

func diagonalPrime(nums [][]int) int {
	var ans int
	m := len(nums)
	for i := 0; i < m; i++ {
		x := nums[i][i]
		if !primes[x] && x > ans {
			ans = x
		}
		x = nums[i][m-1-i]
		if !primes[x] && x > ans {
			ans = x
		}
	}
	return ans
}
