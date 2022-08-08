package main
/**
i 表示nums2的一种排列
dp[i][j] = min(dp[i xor (1<<j), k]) + nums2[j] xor nums1[k]

 */

const N = 15

func minimumXORSum(nums1 []int, nums2 []int) int {
	n := len(nums2)

	dp := make([][]int, 1<<N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}

	for i := 1; i < 1<<n; i++ {
		for j := 0; j < n; j++ {

		}
	}

	return dp[(1<<n)-1][n-1]
}
