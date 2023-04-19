package main

import "fmt"

/*
*
思路: nums的和为s, k = s%p
子数组的和为a, 要求a%p==k, 则(s -a)%p == s%p - a%p = 0,满足去掉a之后，剩余元素和(mod p) == 0
如何求满足a%p==k的子数组？
通过前缀和, (a-k)%p==0,
*/
func minSubarray(nums []int, p int) int {
	n := len(nums)
	var ans = n
	var mem = map[int]int{0: 0}
	var s int
	for _, x := range nums {
		s = (s + x) % p
	}
	if s%p == 0 {
		return 0
	}

	k := s % p
	s = 0
	for i := 1; i <= n; i++ {
		s = (s + nums[i-1]) % p
		if j, ok := mem[(s-k+p)%p]; ok {
			ans = min(ans, i-j)
		}
		mem[s] = i
	}

	if ans == n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6))
	fmt.Println(minSubarray([]int{6, 3, 5, 2}, 9))
	fmt.Println(minSubarray([]int{1, 2, 3}, 3))
	fmt.Println(minSubarray([]int{1, 2, 3}, 7))
	fmt.Println(minSubarray([]int{1000000000, 1000000000, 1000000000}, 3))
}
