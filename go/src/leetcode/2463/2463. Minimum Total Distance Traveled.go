/**
@author ZhengHao Lou
Date    2022/11/07
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/minimum-total-distance-traveled/
思路：动态规划
f[i][j] 表示第i个工厂维修前j个robot的最小移动距离
f[i][j] = f[i-1][j-k] + [j-k+1, j]|robot[j-k+1] - factory[i]|
*/
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	m := len(robot)
	f := make([]int64, m+1)
	for i := range f {
		f[i] = 1e18
	}
	f[0] = 0

	for i := range factory {
		pos, limit := factory[i][0], factory[i][1]
		for j := m; j != 0; j-- {
			for k, cost := 1, int64(0); k <= min(j, limit); k++ {
				cost += int64(abs(robot[j-k] - pos))
				if f[j-k]+int64(cost) < f[j] {
					f[j] = f[j-k] + int64(cost)
				}
			}
		}
	}

	fmt.Println(f)
	return int64(f[m])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	minimumTotalDistance([]int{0, 4, 6}, [][]int{{2, 2}, {6, 2}})
	minimumTotalDistance([]int{1, -1}, [][]int{{-2, 1}, {2, 1}})
	minimumTotalDistance([]int{-333539942, 359275673, 89966494, 949684497, -733065249, 241002388, 325009248, 403868412,
		-390719486, -670541382, 563735045, 119743141, 323190444, 534058139, -684109467, 425503766, 761908175},
		[][]int{{-590277115, 0}, {-80676932, 3}, {396659814, 0}, {480747884, 9}, {118956496, 10}})
}
