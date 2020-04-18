class Solution(object):
    def orangesRotting(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        if not grid: return 0
        m = len(grid)
        n = len(grid[0])
        
        step = 0
        while True:
            f = False    
            is_all = True

            for i in range(m):
                for j in range(n):
                    if grid[i][j]==1:
                        is_all = False
                        for dx, dy in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                            nx = i + dx
                            ny = j + dy
                            if 0<=nx<m and 0<=ny<n and grid[nx][ny]==2:
                                grid[i][j] = 1|2
                                break
            if is_all: return step
            step += 1

            for i in range(m):
                for j in range(n):
                    if grid[i][j]==3:
                        grid[i][j] = 2
                        f = True
                    if grid[i][j]==1:
                        is_all = False
                        
            if not f:
                return -1


            
