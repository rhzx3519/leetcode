class Solution(object):
    def diagonalSort(self, mat):
        """
        :type mat: List[List[int]]
        :rtype: List[List[int]]
        """
        m, n = len(mat), len(mat[0])
        starts = [(i, 0) for i in range(m-1, -1, -1)] + [(0, j) for j in range(1,  n)]
        for x, y in starts:
            i, j = x, y
            a = []
            while i < m and j < n:
                a.append(mat[i][j])
                i += 1
                j += 1
            a.sort()
            i, j = x, y
            k = 0
            while i < m and j < n:
                mat[i][j] = a[k]
                k += 1
                i += 1
                j += 1
        return mat

# 思路：寻找斜对角遍历的所有起点

if __name__ == '__main__':
    mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
    su = Solution()
    su.diagonalSort(mat)
