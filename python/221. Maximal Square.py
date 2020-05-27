class Solution(object):
    def maximalSquare(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix: return 0
        m, n = len(matrix), len(matrix[0])
        dp = [[0]*n for _ in range(m)]
        x = 0

        for i in range(m):
            dp[i][0] = int(matrix[i][0])
            x = max(x, dp[i][0])

        for j in range(n):
            dp[0][j] = int(matrix[0][j])
            x = max(x, dp[0][j])
        
        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][j]=='1':
                    dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
                x = max(x, dp[i][j])
        return x*x

# 思路： dp[i][j]表示 dp[i][j]表示以第i行第j列为右下角所能构成的最大正方形边长,
# dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
        
