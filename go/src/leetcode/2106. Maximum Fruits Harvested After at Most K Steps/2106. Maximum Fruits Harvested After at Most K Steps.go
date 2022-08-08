/**
@author ZhengHao Lou
Date    2021/12/14
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-fruits-harvested-after-at-most-k-steps/
思路：前缀和 + 二分
位置在startPos，总共走k步，可以表示为
1. 向左走x步，再返回startPos，然后往右走y步，总共走了k = 2*x + y步
2. 向右走x步，再返回startPos，然后往左走y步，总共走了k = 2*x + y步
枚举向左/向右走x步，然后折返之后走y步的情况，求出这个区间内能获得的水果数量，取所有区间的最大值

time: O(klogn)
space: O(n)
n为fruits数组的长度
 */

const N int = 2e5 + 1
func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	sums := make([]int, n + 1)
	pos := []int{}
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + fruits[i-1][1]
		pos = append(pos, fruits[i-1][0])
	}
	//L, R := pos[0], pos[len(pos) - 1]

	fmt.Println(pos)
	fmt.Println(sums, len(sums))
	var ans int
	for x := k; x >= 0; x-- {
		var y = k - x<<1
		var l, r int
		// left -> right
		// [startPos - x, startPos + y]

		l = lowerBound(pos, startPos - x)
		r = upperBound(pos, startPos + y)

		t1 := sums[r] - sums[l]
		//fmt.Printf("x: %v, y: %v, 2*x+y: %v, (%v, %v), (%v, %v), %v\n", x, y, x<<1 + y, startPos - x, startPos + y, l, r, t1)
		ans = max(ans, t1)

		// right -> left
		// [startPos - y, startPos + x]
		l = lowerBound(pos, startPos - y)
		r = upperBound(pos, startPos + x)
		t2 := sums[r] - sums[l]
		fmt.Printf("x: %v, y: %v, 2*x+y: %v, (%v, %v), (%v, %v), %v\n", x, y, x<<1 + y,
			startPos - x, startPos + y, l, r, t1)

		ans = max(ans, t2)
	}
	fmt.Println(ans)
	return ans
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		var m = l + (r - l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(nums []int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		var m = l + (r - l)>>1
		if nums[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
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


func main() {
	//maxTotalFruits([][]int{{2,8},{6,3},{8,6}}, 5, 4)
	//maxTotalFruits([][]int{{0,9},{4,1},{5,7},{6,2},{7,4},{10,9}}, 5, 4)
	//maxTotalFruits([][]int{{0,3},{6,4},{8,5}}, 3, 2)	// 0
	//maxTotalFruits([][]int{{0,10000}}, 200000, 200000)	 // 10000
	maxTotalFruits([][]int{{0,7},{7,4},{9,10},{12,6},{14,8},{16,5},{17,8},{19,4},{20,1},{21,3},{24,3},{25,3},
		{26,1},{28,10},{30,9},{31,6},{32,1},{37,5},{40,9}}, 21, 30)	 // 71
		// [9,31] sp=21, x=10, y=10,

}