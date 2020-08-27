class Solution(object):
    def checkIfPrerequisite(self, n, prerequisites, queries):
        """
        :type n: int
        :type prerequisites: List[List[int]]
        :type queries: List[List[int]]
        :rtype: List[bool]
        """
        grid = [[0]*n for _ in range(n)]
        for a, b in prerequisites:
            grid[a][b] = 1

        for k in range(n):
            for i in range(n):
                for j in range(n):
                    if not grid[i][j]:
                        grid[i][j] = grid[i][k] & grid[k][j]

        res = []
        for a, b in queries:
            res.append(bool(grid[a][b]))
        return res

# floyd，求有向图两点之间是否联通
# time: O(N^3), space: O(N^2)