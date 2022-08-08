/**
@author ZhengHao Lou
Date    2022/01/23
*/
package main

import "fmt"

var MOD int = 1e9 + 7

func numberOfWays(corridor string) int {
	//n := len(corridor)
	var gap = []int{}
	var t int
	var s, k int
	for i := range corridor {
		switch corridor[i] {
		case 'S':
			t++
			s++
		case 'P':
		}
		if s == 2 {
			gap = append(gap, k)
			s = 0
			k = 0
		} else if s == 0 {
			k++
		}
	}
	if t == 0 || t&1 == 1 {
		return 0
	}
	if len(gap) <= 1 {
		if t == 2 {
			return 1
		}
		return 0
	}
	var ans = 1
	for i := 1; i < len(gap); i++ {
		ans = ans * (gap[i] + 1) % MOD
	}
	//fmt.Println(gap, ans)
	return ans
}

func main() {
	fmt.Println(numberOfWays("SSPPSPS"))
	fmt.Println(numberOfWays("PPSPSP"))
	fmt.Println(numberOfWays("s"))
}
