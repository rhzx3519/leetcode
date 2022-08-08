package main

import "fmt"

/**
https://leetcode-cn.com/problems/can-i-win/
思路：记忆化搜索
 */
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if ((1 + maxChoosableInteger)*maxChoosableInteger)>>1 < desiredTotal {
		return false
	}

	mem := make(map[int]bool)

	var dfs func(state, target int) bool
	dfs = func(state, target int) bool {
		if _, ok := mem[state]; ok {
			return mem[state]
		}

		for i := 1; i <= maxChoosableInteger; i++ {
			if state & (1 << (i - 1)) == 0 {
				continue
			}
			if i >= target || !dfs(state & ^(1 << (i - 1)), target - i) {
				mem[state] = true
				return mem[state]
			}
		}

		mem[state] = false
		return mem[state]
	}

	return dfs(1<<maxChoosableInteger - 1, desiredTotal)
}

func main() {
	var x = 1<<4 - 1
	fmt.Printf("%b\n", x & (^(1<<2)))
	fmt.Println(canIWin(10, 11))
}





