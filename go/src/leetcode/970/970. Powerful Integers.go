package main

/*
*
https://leetcode.cn/problems/powerful-integers/
*/
func powerfulIntegers(x int, y int, bound int) []int {
	var mp = map[int]int{}
	for i := 0; powInt(x, i) < bound; i++ {
		for j := 0; powInt(x, i)+powInt(y, j) <= bound; j++ {
			mp[powInt(x, i)+powInt(y, j)]++
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	var ans []int
	for k := range mp {
		ans = append(ans, k)
	}
	return ans
}

const MOD int = 1e9 + 7

func powInt(x, n int) int {
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
	powerfulIntegers(2, 1, 10)
}
