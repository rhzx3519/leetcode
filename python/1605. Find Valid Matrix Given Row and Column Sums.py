class Solution(object):
    def restoreMatrix(self, rowSum, colSum):
        """
        :type rowSum: List[int]
        :type colSum: List[int]
        :rtype: List[List[int]]
        """
        m, n = len(rowSum), len(colSum)
        grid = [[0]*n for _ in range(m)]
        for i in range(m):
            for j in range(n):
                now = min(rowSum[i], colSum[j])
                grid[i][j] = now
                rowSum[i] -= now
                colSum[j] -= now
        return grid
        