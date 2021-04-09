class Solution(object):
    def lenLongestFibSubseq(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        n = len(A)
        ans = 0
        dp = [[2]*n for _ in range(n)]
        mem = {A[i]: i for i in range(n)}
        for i in range(2, n):
            for j in range(i-1, -1, -1):
                if (A[i]-A[j]) in mem and mem[A[i]-A[j]] < j:
                    dp[j][i] = max(dp[j][i], dp[mem[A[i]-A[j]]][j] + 1)
                    ans = max(ans, dp[j][i])
        return ans