class Solution(object):
    def regionsBySlashes(self, grid):
        """
        :type grid: List[str]
        :rtype: int
        """
        N = len(grid)
        look_up = [i for i in range(4*N*N)]
        def find(x):
            px = look_up[x]
            if px!=x:
                look_up[x] = find(px)
            return look_up[x]
    
        def union(x, y):
            px = find(x)
            py = find(y)
            if px!=py:
                look_up[py] = px

        for i in range(N):
            for j in range(N):
                k = 4*(i*N + j)
                if grid[i][j] == "/":
                    union(k, k+3)
                    union(k+1, k+2)
                elif grid[i][j] == "\\":
                    union(k, k+1)
                    union(k+2, k+3)
                else:
                    union(k, k+1)
                    union(k+1, k+2)
                    union(k+2, k+3)
                
                # 向右合并
                if j+1 < N:
                    union(k+1, 4*(i*N+j+1)+3)
                # 向下合并
                if i+1 < N:
                    union(k+2, 4*((i+1)*N+j))

        ans = 0
        for i in range(N*N*4):
            if look_up[i]==i:
                ans += 1
        return ans



