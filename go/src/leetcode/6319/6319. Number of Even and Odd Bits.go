package main

import "fmt"

/*
*
https://leetcode.cn/problems/number-of-even-and-odd-bits/
*/
func evenOddBit(n int) []int {
	var ans = make([]int, 2)
	for i := 0; n != 0; n, i = n>>1, i+1 {
		if n&1 == 1 {
			ans[i&1]++
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	evenOddBit(17)
	evenOddBit(2)
}
