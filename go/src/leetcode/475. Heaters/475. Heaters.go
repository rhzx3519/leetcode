/**
@author ZhengHao Lou
Date    2021/12/20
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

/**
https://leetcode-cn.com/problems/heaters/
思路：贪心
1. 对于每个house，不是被前面的heater覆盖，就是被后面的覆盖
2. 选择前后heater中的离得近的，取所有距离的最大值
time: O(mlogn)
space: O(1)
 */
const INF = math.MaxInt32 >> 1
func findRadius(houses []int, heaters []int) int {
	nh := len(heaters)
	var ans int
	sort.Ints(heaters)
	for _, x := range houses {
		i := lowerBound(heaters, x)
		var t1, t2 = INF, INF
		if i > 0 {
			t1 = x - heaters[i-1]
		}
		if i < nh {
			t2 = heaters[i] - x
		}
		ans = max(ans, min(t1, t2))
		//fmt.Println(x, t1, t2)
	}
	fmt.Println(ans)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	findRadius([]int{1,2,3}, []int{2})
	findRadius([]int{1,2,3,4}, []int{1,4})
	findRadius([]int{1,5}, []int{2})
}