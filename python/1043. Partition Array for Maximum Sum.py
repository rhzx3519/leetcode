class Solution(object):
    def maxSumAfterPartitioning(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: int
        """
        n = len(A)
        dp = [0] * (n+1)
        for i in range(n+1):
            max_val = dp[i]
            j = i-1
            while j>=0 and (i-j)<=K:
                max_val = max(max_val, A[j])
                dp[i] = max(dp[i], dp[j] + max_val*(i-j))
                j -= 1
        return dp[n]

# dp[i] 表示前i个元素，被分隔成长度最多为K的连续子数组的最大值
# 递推公式: dp[i] = max(dp[i], dp[j] + max_val*(i-j)), i-k <= j <= i-1
# time: O(N^2), space: O(N)