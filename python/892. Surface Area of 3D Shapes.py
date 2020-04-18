class Solution(object):
    def surfaceArea(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        if not grid: return 0
        m, n = len(grid), len(grid[0])
        
        def foo(x, y):
            ans = 0
            if grid[x][y]==0: return 0
            for dx, dy in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                nx = x + dx
                ny = y + dy
                if 0<=nx<m and 0<=ny<n:
                    ans += 0 if grid[x][y] <= grid[nx][ny] else grid[x][y] - grid[nx][ny]
                else:
                    ans += grid[x][y]
            return ans + 2

        res = 0
        for i in range(m):
            for j in range(n):
                res += foo(i, j)
        return res