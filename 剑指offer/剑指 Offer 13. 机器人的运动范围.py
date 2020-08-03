class Solution(object):
    def movingCount(self, m, n, k):
        """
        :type m: int
        :type n: int
        :type k: int
        :rtype: int
        """
        vis = [[0]*n for _ in range(m)]
        self.count = 0

        def dfs(x, y):
            self.count += 1

            vis[x][y] = 1
            for dx, dy in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                nx = x + dx
                ny = y + dy
                if 0<=nx<m and 0<=ny<n and vis[nx][ny]==0:
                    s = sum(map(int, list(str(nx)))) + sum(map(int, list(str(ny))))
                    if s <= k:
                        dfs(nx, ny)
            # vis[x][y] = 0

        dfs(0, 0)
        return self.count
                        

if __name__ == '__main__':
    m, n, k = 3, 1, 0
    su = Solution()
    print su.movingCount(m, n, k)