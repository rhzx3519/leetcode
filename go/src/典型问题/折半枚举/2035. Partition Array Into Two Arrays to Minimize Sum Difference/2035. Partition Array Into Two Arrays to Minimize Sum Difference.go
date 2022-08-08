/**
@author ZhengHao Lou
Date    2021/10/11
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

/**
https://leetcode-cn.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/
题解: https://leetcode-cn.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/solution/zhe-ban-mei-ju-pai-xu-er-fen-by-endlessc-04fn/
思路：折半枚举
分割成两个数组做差，等价于长度为2n数组的数组，取n个数乘以正号，剩余n个数乘以负号之后的数组和，
将数组分割为前n个和后n个，枚举前n个数字取正号的所有情况，按照正号个数分组排序；
枚举后n个数字取正号的所有情况，寻找和最接近的前半部分状态，它们之间的差即为答案
 */
func minimumDifference(nums []int) int {
	n := len(nums) >> 1
	positives := make([][]int, n+1)
	a := nums[:n]
	for i := 0; i < 1<<n; i++ {
		var sum, count int
		for j := 0; j < n; j++ {
			if (1<<j) & i != 0 {	// positive number
				sum += a[j]
				count++
			} else {
				sum -= a[j]
			}
		}
		positives[count] = append(positives[count], sum)
	}

	for i := 0; i <= n; i++ {
		sort.Ints(positives[i])
	}

	var ans = math.MaxInt64
	a = nums[n:]
	for i := 0; i < 1<<n; i++ {
		var sum, count int
		for j := 0; j < n; j++ {
			if (1<<j)&i != 0 { // positive number
				sum += a[j]
				count++
			} else {
				sum -= a[j]
			}
		}
		t := positives[count]
		j := sort.SearchInts(t, sum)
		if j < len(t) {
			ans = min(ans, t[j] - sum)
		}
		if j > 0 {
			ans = min(ans, sum - t[j-1])
		}
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

func main() {
	minimumDifference([]int{3,9,7,3})
	minimumDifference([]int{-36,36})
	minimumDifference([]int{2,-1,0,4,-2,-9})
}
