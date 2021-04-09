class Solution(object):
    def maxProductPath(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m, n = len(grid), len(grid[0])
        mod = 1e9+7
        dp = [[[0, 0] for _ in range(n)] for _ in range(m)]
        dp[0][0][0] = dp[0][0][1] = grid[0][0]
        for i in range(1, m):
            dp[i][0][0] = dp[i][0][1] = dp[i-1][0][0]*grid[i][0]
        for j in range(1, n):
            dp[0][j][0] = dp[0][j][1] = dp[0][j-1][0]*grid[0][j]
        # print dp
        for i in range(1, m):
            for j in range(1, n):
                if grid[i][j] >= 0:
                    dp[i][j][0] = grid[i][j] * max(dp[i-1][j][0], dp[i][j-1][0])
                    dp[i][j][1] = grid[i][j] * min(dp[i-1][j][1], dp[i][j-1][1])
                else:
                    dp[i][j][0] = grid[i][j] * min(dp[i-1][j][1], dp[i][j-1][1])
                    dp[i][j][1] = grid[i][j] * max(dp[i-1][j][0], dp[i][j-1][0])
        # print dp
        return int(dp[m-1][n-1][0]%mod) if dp[m-1][n-1][0]>=0 else -1
        