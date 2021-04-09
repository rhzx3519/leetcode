class Solution(object):
    def largest1BorderedSquare(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        if not grid: return 0
        ans = 0
        m, n = len(grid), len(grid[0])
        dp = [[[0, 0] for _ in range(n+1)] for _ in range(m+1)]
        print dp
        for i in range(1, m+1):
            for j in range(1, n+1):
                if grid[i-1][j-1] == 0: continue
                dp[i][j][0] = max(dp[i][j][0], dp[i][j-1][0] + 1)   # left
                dp[i][j][1] = max(dp[i][j][1], dp[i-1][j][1] + 1)   # up
                min_len = min(dp[i][j])
                for k in range(min_len):
                    if dp[i-k][j][0] >= k+1 and dp[i][j-k][1] >= k+1:
                        ans = max(ans, k+1)
        return ans*ans

# 思路：dp，dp[i][j][0]表示下标(i, j)左边连续1的个数, dp[i][j][1]表示下标(i, j)上边连续1的个数.