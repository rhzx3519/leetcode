package main

import "fmt"

/*
*
https://leetcode.cn/problems/count-collisions-of-monkeys-on-a-polygon/
*/
const MOD int = 1e9 + 7

func monkeyMove(n int) int {
	return (pow(2, n) - 2 + MOD) % MOD
}

func pow(x, n int) int {
	var res = 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % MOD
		}
		x = x * x % MOD
	}
	return res
}

func main() {
	fmt.Println(monkeyMove(3))
}
