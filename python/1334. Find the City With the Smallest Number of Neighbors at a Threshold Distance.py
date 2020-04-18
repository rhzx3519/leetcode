class Solution(object):
    def findTheCity(self, n, edges, distanceThreshold):
        """
        :type n: int
        :type edges: List[List[int]]
        :type distanceThreshold: int
        :rtype: int
        """
        INF = 1<<31
        grid = [[INF]*n for _ in range(n)]
        for i in range(n):
            grid[i][i] = 0

        for a, b, w in edges:
            grid[a][b] = w
            grid[b][a] = w

        def floyd():
            for k in range(n):
                for i in range(n):
                    for j in range(n):
                        if grid[i][k] + grid[k][j] < grid[i][j]:
                            grid[i][j] = grid[i][k] + grid[k][j]

        floyd()
        print grid
        ans = 0
        min_val = INF
        for i in range(n):
            count = 0
            for j in range(n):
                if i==j: continue
                if grid[i][j] <= distanceThreshold:
                    count += 1
            if count <= min_val:
                min_val = count
                ans = i
        return ans
        

if __name__ == '__main__':
    n = 4
    edges = [[0,1,3],[1,2,1],[1,3,4],[2,3,1]]
    distanceThreshold = 4
    su = Solution()
    print su.findTheCity(n, edges, distanceThreshold)