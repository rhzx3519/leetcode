/**
@author ZhengHao Lou
Date    2021/11/08
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/sell-diminishing-valued-colored-balls/
思路：剩余球的数量是sum(inventory) - orders, 每种颜色球剩余数量的区间是[0, max(inventory)]
 */
const MOD int = 1e9 + 7

func maxProfit(inventory []int, orders int) int {
	l, r := 0, maxElement(inventory) + 1
	var m int
	for l < r {
		m = l + (r - l)>>1
		var c int64
		for _, num := range inventory {
			if num > m {
				c += int64(num - m)
			}
		}

		if c > int64(orders) {
			l = m + 1
		} else {
			r = m
		}
	}
	// 平均剩余m个球
	var val int
	for _, num := range inventory {
		if num > l {
			x := min(orders, num - l) // 取x个球
			val = (val + score(x, num)) % MOD
			orders -= x
		}
	}
	val = (val + orders*l) % MOD	// 如果选取的数量还不够，那么剩余的球都将以l的价值被获取
	fmt.Println(val)
	return val
}

// 总数为y, 取x
func score(x, y int) int {
	return (y + y - x + 1)*x>>1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxElement(inventory []int) int {
	var val int
	for _, num := range inventory {
		if num > val {
			val = num
		}
	}
	return val
}

func main() {
	//fmt.Println(score(3, 5))
	maxProfit([]int{2,5}, 4)
	maxProfit([]int{3,5}, 6)
	maxProfit([]int{2,8,4,10,6}, 20)
	maxProfit([]int{1000000000}, 1000000000)

}
