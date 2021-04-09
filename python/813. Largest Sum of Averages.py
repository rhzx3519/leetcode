class Solution(object):
    def largestSumOfAverages(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: float
        """
        N = len(A)
        dp = [[0]*(K+1) for _ in range(N+1)]
        s = [0]*(N+1)
        for i in range(1, N+1):
            s[i] = s[i-1] + A[i-1]
            dp[i][1] = float(s[i])/i

        for i in range(1, N+1):
            for k in range(2, K+1):
                for j in range(i):
                    dp[i][k] = max(dp[i][k], dp[j][k-1] + (s[i]-s[j])/float(i-j))
        return dp[N][K]
        

# dp[i][k]表示前i个数构成k个子数组时的最大平均值, 那么对于所有 0 <= j <= i-1
# dp[i][k] = Math.max(dp[i][k], dp[j][k-1] + (A[j+1] + ... + A[i+1]) / (i-j))

