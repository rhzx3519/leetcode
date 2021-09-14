package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/student-attendance-record-ii/
https://leetcode-cn.com/problems/student-attendance-record-ii/solution/tong-ge-lai-shua-ti-la-yi-ti-liu-jie-dfs-s5fa/

 */
var mod = int(math.Pow10(9)) + 7

func checkRecord(n int) int {
	mem := make([][][]int, n+1)
	for i := range mem {
		mem[i] = make([][]int, 2)
		for j := range mem[i] {
			mem[i][j] = make([]int, 3)
		}
	}

	var dfs func(day, absent, late int) int
	dfs = func(day, absent, late int) int {
		if day >= n {
			return 1
		}

		if mem[day][absent][late] != 0 {
			return mem[day][absent][late]
		}

		var count int
		// 1. present随便放
		count = (count + dfs(day + 1, absent, 0)) % mod

		// 2. absent最多只能有一个
		if absent < 1 {
			count = (count + dfs(day + 1, 1, 0)) % mod
		}

		// 3. late最多连续放2个
		if late < 2 {
			count = (count + dfs(day + 1, absent, late + 1)) % mod
		}

		mem[day][absent][late] = count
		return count
	}

	return dfs(0, 0, 0)
}

func main() {
	fmt.Println(checkRecord(2))
}
