class Solution(object):
    def minFallingPathSum(self, A):
        """
        :type A: List[List[int]]
        :rtype: int
        """
        N = len(A)
        dp = [[float('inf')]*N for _ in range(N)]
        dp[0] = A[0][:]
        for i in range(1, N):
            for j in range(N):
                for k in (-1, 0, 1):
                    t = j + k
                    if 0<=t<N:
                        dp[i][j] = min(dp[i][j], dp[i-1][t] + A[i][j])
        return min(dp[N-1])