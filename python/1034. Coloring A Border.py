import copy

class Solution(object):
    def colorBorder(self, grid, r0, c0, color):
        """
        :type grid: List[List[int]]
        :type r0: int
        :type c0: int
        :type color: int
        :rtype: List[List[int]]
        """
        m = len(grid)
        n = len(grid[0])
        vis = [[0]*n for _ in range(m)]
        mp = copy.deepcopy(grid)

        def dfs(x, y, c):
            if x<0 or x>m-1 or y<0 or y>n-1 or mp[x][y]!=c:
                return
            vis[x][y] = 1
            for dx, dy in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
                nx = x + dx
                ny = y + dy
                
                if nx<0 or nx>m-1 or ny<0 or ny>n-1 or mp[nx][ny]!=c:
                    grid[x][y]=color
                else:
                    if vis[nx][ny]:
                        continue
                    dfs(nx, ny, c)
                

        dfs(r0, c0, grid[r0][c0])
        return grid


if __name__ == '__main__':
    grid = [[1,1,1],[1,1,1],[1,1,1]]
    r0 = 1
    c0 = 1
    color = 2
    su = Solution()
    print su.colorBorder(grid, r0, c0, color)        