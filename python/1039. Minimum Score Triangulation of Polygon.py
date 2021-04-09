class Solution(object):
    def minScoreTriangulation(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        N = len(A)
        dp = [[0]*N for _ in range(N)]
        for i in range(N-3, -1, -1):
            for j in range(i+2, N):
                for m in range(i+1, j):
                    if dp[i][j] == 0:
                        dp[i][j] = A[i]*A[m]*A[j] + dp[i][m] + dp[m][j]
                    else:
                        dp[i][j] = min(dp[i][j], A[i]*A[m]*A[j] + dp[i][m] + dp[m][j])
        return dp[0][N-1]