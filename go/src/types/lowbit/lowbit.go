package main

import "fmt"

/*
*
https://www.quora.com/What-does-x-x-lowbit-do-in-C++
x&(-x) gives the ‘weight’ of the lowest non-zero bit in x.
*/
func minOperations(n int) int {
	ans := 1
	for n&(n-1) > 0 { // n 不是 2 的幂次
		lb := n & -n
		if n&(lb<<1) > 0 { // 多个连续 1
			n += lb
		} else {
			n -= lb // 单个 1
		}
		fmt.Printf("%b\n", lb)
		ans++
	}
	return ans
}

func example() {
	var x = 39
	fmt.Printf("%b\n", x)
	lb := x & -x
	fmt.Printf("%b\n", lb)
	fmt.Printf("%b\n", x&(lb<<1))
}

func main() {
	x := 39
	fmt.Printf("%b\n", x)
	minOperations(x)
}
