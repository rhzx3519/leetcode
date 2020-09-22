class Solution(object):
    def getMaximumGold(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m, n = len(grid), len(grid[0])
        vis = [[0]*n for _ in range(m)]
        self.max_val = 0
        
        def dfs(x, y, s):
            s += grid[x][y]
            self.max_val = max(self.max_val, s)
            for dx, dy in ((1, 0), (0, 1), (-1, 0), (0, -1)):
                nx = x + dx
                ny = y + dy
                if 0<=nx<m and 0<=ny<n and not vis[nx][ny] and grid[nx][ny] > 0:
                    vis[nx][ny] = 1
                    dfs(nx, ny, s)
                    vis[nx][ny] = 0
        
        for i in range(m):
            for j in range(n):
                if grid[i][j] > 0:
                    vis[i][j] = 1
                    dfs(i, j, 0)
                    vis[i][j] = 0
        return self.max_val
            