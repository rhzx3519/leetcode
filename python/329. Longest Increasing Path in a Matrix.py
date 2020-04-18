class Solution(object):
    def longestIncreasingPath(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: int
        """
        if not matrix: return 0
        m = len(matrix)
        n = len(matrix[0])
        look_up = [[0]*n for _ in range(m)]

        def dfs(x, y):
            if look_up[x][y]!=0:
                return look_up[x][y]
            
            res = 1
            for dx, dy in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                nx = x + dx
                ny = y + dy
                if nx>=0 and nx<m and ny>=0 and ny<n and matrix[nx][ny]>matrix[x][y]:
                    res = max(res, dfs(nx, ny)+1)
            look_up[x][y] = res
            return look_up[x][y]
        
        res = 1
        for i in range(m):
            for j in range(n):
                res = max(res, dfs(i, j))
        return res