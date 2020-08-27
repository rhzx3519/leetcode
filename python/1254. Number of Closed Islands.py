class Solution(object):
    def closedIsland(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m, n = len(grid), len(grid[0])
        vis = [[0]*n for _ in range(m)]

        def dfs(x, y):
            vis[x][y] = 1

            ans = True
            for dx, dy in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                nx = x + dx
                ny = y + dy
                if nx<0 or nx>=m or ny<0 or ny>=n:
                    ans = False
                    continue
                if vis[nx][ny]==0 and grid[nx][ny]==0 and dfs(nx, ny) is False:
                    ans = False
            return ans
        
        res = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j]==0 and vis[i][j]==0 and dfs(i, j):
                    res += 1
        return res