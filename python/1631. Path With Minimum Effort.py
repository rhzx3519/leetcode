class Solution(object):
    def minimumEffortPath(self, heights):
        """
]        :type heights: List[List[int]]
        :rtype: int
        """
        m, n = len(heights), len(heights[0])
        grid = [[float('inf')]*n for _ in range(m)]
        grid[0][0] = 0
        que = [(0, 0)]
        while que:
            x, y = que.pop(0)
            for dx, dy in ((1,0), (-1,0), (0,1), (0,-1)):
                nx = x + dx
                ny = y + dy
                if nx<0 or nx>=m or ny<0 or ny>=n:
                    continue
                t = max(grid[x][y], abs(heights[nx][ny] - heights[x][y]))
                if t >= grid[nx][ny]:
                    continue
                grid[nx][ny] = t
                que.append((nx, ny))
        return grid[-1][-1