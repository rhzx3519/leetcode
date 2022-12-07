/**
@author ZhengHao Lou
Date    2022/10/16
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/range-product-queries-of-powers/
*/
const MOD int = 1e9 + 7

func productQueries(n int, queries [][]int) (ans []int) {
	var powers []int
	for i := 0; n != 0; i++ {
		if n&1 == 1 {
			powers = append(powers, i)
		}
		n >>= 1
	}
	var sums = make([]int, len(powers)+1)
	for i := 1; i <= len(powers); i++ {
		sums[i] = powers[i-1] + sums[i-1]
	}
	for _, q := range queries {
		l, r := q[0], q[1]
		p := sums[r+1] - sums[l]
		ans = append(ans, mod(p))
	}

	fmt.Println(powers, sums, ans)
	return ans
}

func mod(p int) int {
	if p <= 31 {
		return (1 << p) % MOD
	}
	return (mod(31) * mod(p-31)) % MOD
}

func main() {
	//productQueries(15, [][]int{{0, 1}, {2, 2}, {0, 3}})
	//productQueries(2, [][]int{{0, 0}})
	//productQueries(806335498, [][]int{{11, 11}, {1, 8}})
	fmt.Println(mod(104))

}
