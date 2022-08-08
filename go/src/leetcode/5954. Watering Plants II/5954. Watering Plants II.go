/**
@author ZhengHao Lou
Date    2021/12/13
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/watering-plants-ii/
思路：贪心
 */
func minimumRefill(plants []int, capacityA int, capacityB int) int {
	n := len(plants)
	l, r := 0, n-1
	a, b := capacityA, capacityB
	var c int
	for l < r {
		if a < plants[l] {
			a = capacityA
			c++
		}
		a -= plants[l]
		if b < plants[r] {
			b = capacityB
			c++
		}
		b -= plants[r]
		l++
		r--
	}
	if l == r {
		if max(a, b) < plants[l] {
			c++
		}
	}
	fmt.Println(c)
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	minimumRefill([]int{2,2,3,3}, 5, 5)
	minimumRefill([]int{2,2,3,3}, 3, 4)
	minimumRefill([]int{5}, 10, 8)
	minimumRefill([]int{1,2,4,4,5}, 6, 5)
	minimumRefill([]int{2,2,5,2,2}, 5, 5)
}
