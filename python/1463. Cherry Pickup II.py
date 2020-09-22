class Solution(object):
    def cherryPickup(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        if not grid:
            return 0
        m, n = len(grid), len(grid[0])
        dp = [[[float('-inf')]*n for _ in range(n)] for _ in range(m)]
        dp[0][0][-1] = grid[0][0] + grid[0][-1]
        # print dp
        offset = (-1, 0, 1)
        for r in range(1, m):
            for i in range(n):
                for j in range(n):
                    for k1 in offset:
                        for k2 in offset:
                            x, y = i+k1, j+k2
                            if x!=y and 0<=x<n and 0<=y<n:
                                dp[r][x][y] = max(dp[r][x][y], dp[r-1][i][j] + grid[r][x] + grid[r][y])
        return max([max(row) for row in dp[-1]])

# 思路：dp[r][i][j], 表示第r行，左右两个机器人分别在i, j 位置时的最大樱桃数量


class Solution(object):
    def cherryPickup(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m, n = len(grid), len(grid[0])
        MIN_VAL = float('-inf')
        offset = (-1, 0, 1)
        mem = {}

        def dfs(r, i, j):
            if (r, i, j) in mem:
                return mem[(r, i, j)]
            if i<0 or i>n-1 or j<0 or j>n-1:
                return MIN_VAL
            if r==m-1:
                return grid[r][i] + grid[r][j] if i!=j else MIN_VAL
            if i==j:
                return MIN_VAL

            ans = MIN_VAL
            for k1 in offset:
                for k2 in offset:
                    ans = max(ans, dfs(r+1, i+k1, j+k2) + grid[r][i] + grid[r][j])
            mem[(r, i, j)] = ans
            return ans

        return dfs(0, 0, n-1)

# 思路：记忆化搜索


