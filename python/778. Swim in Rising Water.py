class UFS(object):
    def __init__(self, n):
        self.parent = [i for i in range(n)]
        self.n = n
    
    def find(self, x):
        px = self.parent[x]
        if px != x:
            self.parent[x] = self.find(px)
        return self.parent[x]
    
    def union(self, x, y):
        px = self.find(x)
        py = self.find(y)
        if px != py:
            self.parent[py] = px
    
    def isConnect(self, x, y):
        return self.find(x) == self.find(y)

class Solution(object):
    def swimInWater(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        N = len(grid)
        l = N*N

        ufs = UFS(l)
        index = [0]*l
        for i in range(N):
            for j in range(N):
                index[grid[i][j]] = i*N + j
        
        for i in range(l):
            x = index[i]/N
            y = index[i]%N
            for dx, dy in ((1,0),(-1,0),(0,1),(0,-1)):
                nx = x + dx
                ny = y + dy
                if 0<=nx<N and 0<=ny<N and grid[nx][ny]<=i:
                    ufs.union(index[i], nx*N+ny)
                
                if ufs.isConnect(0, l-1):
                    return i
        return -1












