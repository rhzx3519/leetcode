/**
@author ZhengHao Lou
Date    2021/11/09
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/zuma-game/
思路：暴力搜索
 */
type ele struct {
	b string
	c int
}
func findMinStep(board string, hand string) int {
	var r = math.MaxInt32 >> 1
	vis := make(map[ele]int)
	var dfs func(b string, h string, c int)
	dfs = func(b string, h string, c int) {
		if len(b) == 0 {
			r = min(r, c)
			return
		}
		e := ele{b: b, c: c}
		if len(h) == 0 || vis[e] != 0 {
			return
		}
		vis[e]++

		bytes := make([]byte, len(b) + 1)
		for i := range b {
			for j := range h {
				copy(bytes, b[:i])
				bytes[i] = h[j]
				copy(bytes[i+1:], b[i:])
				b1 := combine(string(bytes))
				h1 := string(append(append([]byte{}, h[:j]...), h[j+1:]...))
				dfs(b1, h1, c+1)
			}
		}
	}

	dfs(board, hand, 0)
	if r == math.MaxInt32 >> 1 {
		return -1
	}
	return r
}

// 消除连续3个或3个以上的球
func combine(b string) string {
	var slow, fast int
	for ; fast < len(b); fast++ {
		if b[slow] == b[fast] {
			continue
		}
		if fast - slow > 2 {
			b = string(append(append([]byte{}, []byte(b)[:slow]...), []byte(b)[fast:]...))
			fast = 0
		}
		slow = fast
	}
	if fast - slow > 2 {
		b = string(append(append([]byte{}, []byte(b)[:slow]...), []byte(b)[fast:]...))
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(combine("WWWRRBBWW"))
	fmt.Println(findMinStep("WRRBBW", "RB"))		// -1
	fmt.Println(findMinStep("WWRRBBWW", "WRBRW"))	// 2
	fmt.Println(findMinStep("G", "GGGGG"))			// 2
	fmt.Println(findMinStep("RBYYBBRRB", "YRBGB"))	// 3
	fmt.Println(findMinStep("RRWWRRW", "WR"))		// -1
	fmt.Println(findMinStep("WWGWGW", "GWBWR"))		// 3
	fmt.Println(findMinStep("RWYWRRWRR", "YRY"))		// 3
	fmt.Println(findMinStep("RRWWRRBBRR", "WB"))		// 2
}
