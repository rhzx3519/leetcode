package main

/**
https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence/
思路：DP
dp[i][d]表示公差为d，且以nums[i]结尾的等差子序列的个数
第二维使用map来保存
 */
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	var count int
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			var t int
			if cnt, ok := dp[j][d]; ok {
				t = cnt
			}
			dp[i][d] += t + 1

			count += t
		}
	}


	//fmt.Println(dp)
	//fmt.Println(count)
	return count
}

func main() {
	numberOfArithmeticSlices([]int{7,7,7,7,7})
	numberOfArithmeticSlices([]int{2,4,6,8,10})
}
