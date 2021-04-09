class Solution(object):
    def shortestBridge(self, A):
        """
        :type A: List[List[int]]
        :rtype: int
        """
        m, n = len(A), len(A[0])

        x = y = 0
        for i in range(m):
            for j in range(n):
                if A[i][j] == 1:
                    x, y = i, j
                    break

        offset = ((1,0), (-1,0), (0,1), (0,-1))
        vis = [[0]*n for _ in range(m)]
        vis[x][y] = 1
        que = [(x, y)]
        def dfs(x, y, que):
            for dx, dy in offset:
                nx = x + dx
                ny = y + dy
                if 0<=nx<m and 0<=ny<n and A[nx][ny]==1 and not vis[nx][ny]:
                    que.append((nx, ny))
                    vis[nx][ny] = 1
                    dfs(nx, ny, que)

        dfs(x, y, que)
        # print que
        step = 0
        while que:
            sz = len(que)
            while sz:
                sz -= 1
                x, y = que.pop(0)
                if A[x][y]==1 and step > 0:
                    return step - 1
                for dx, dy in offset:
                    nx = x + dx
                    ny = y + dy
                    if 0<=nx<m and 0<=ny<n and not vis[nx][ny]:
                        que.append((nx, ny))
                        vis[nx][ny] = 1
            step += 1
        return step


# [[1,1,1,1,1],
# [1,0,0,0,1],
# [1,0,1,0,1],
# [1,0,0,0,1],
# [1,1,1,1,1]]