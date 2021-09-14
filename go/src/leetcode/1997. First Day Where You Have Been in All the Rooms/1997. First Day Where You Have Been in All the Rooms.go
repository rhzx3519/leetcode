package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/first-day-where-you-have-been-in-all-the-rooms/
题解: https://leetcode-cn.com/problems/first-day-where-you-have-been-in-all-the-rooms/solution/qian-zhui-he-you-hua-dp-by-endlesscheng-j10b/
思路：动态规划
f[i]表示从首次访问房间i到访问房间i+1的天数
 */
const mod int = 1e9 + 7

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	sum := make([]int, n)
	for i, j := range nextVisit[:n-1] {
		f := (2 + sum[i] - sum[j] + mod) % mod	// 0 <= nextVisit[i] < i
		sum[i+1] = (sum[i] + f) % mod
	}
	return sum[n-1]
}

func main() {
	fmt.Println(firstDayBeenInAllRooms([]int{0,0}))
	fmt.Println(firstDayBeenInAllRooms([]int{0,0,2}))
	fmt.Println(firstDayBeenInAllRooms([]int{0,1,2,0}))
}

