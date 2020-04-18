class Solution(object):
    def maxIncreaseKeepingSkyline(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        if not grid: return 0

        m = len(grid)
        n = len(grid[0])
        row = [0]*m
        column = [0]*n
        for i in range(m):
            for j in range(n):
                row[i] = max(row[i], grid[i][j])

        for i in range(n):
            for j in range(m):
                column[i] = max(column[i], grid[j][i])
            
        res = 0
        for i in range(m):
            for j in range(n):
                res += min(row[i], column[j]) - grid[i][j]
        return res

# 统计每行每列的最大值        