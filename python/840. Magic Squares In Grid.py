class Solution(object):
    def numMagicSquaresInside(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m, n = len(grid), len(grid[0])
        if m < 3 or n < 3:
            return 0
        row = [[0]*(n+1) for _ in range(m)]
        for i in range(m):
            for j in range(1, n+1):
                row[i][j] = row[i][j-1] + grid[i][j-1]
        # print row

        col = [[0]*(n) for _ in range(m+1)]
        for j in range(n):
            for i in range(1, m+1):
                col[i][j] = col[i-1][j] + grid[i-1][j]
        # print col

        def check(x, y):
            # print x, y
            # digit
            d = [0]*10
            for i in range(x, x+3):
                for j in range(y, y+3):
                    if grid[i][j]<1 or grid[i][j]>9:
                        return False
                    d[grid[i][j]] += 1
            for i in range(1, 10):
                if d[i] != 1:
                    return False
            # row
            r = row[x][y+3] - row[x][y]
            # print r
            for i in range(x, x+3):
                if r != (row[i][y+3] - row[i][y]):
                    return False
            # col
            for j in range(y, y+3):
                if r != (col[x+3][j] - col[x][j]):
                    return False
            # diagonal
            s1 = s2 = 0
            for k in range(3):
                s1 += grid[x+k][y+k]
                s2 += grid[x+2-k][y+k]

            if s1 != r or s2 != r:
                return False
            return True

        ans = 0
        for i in range(m-2):
            for j in range(n-2):
                if check(i, j):
                    ans += 1
        return ans
                    