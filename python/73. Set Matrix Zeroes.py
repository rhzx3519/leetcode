class Solution(object):
    def setZeroes(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: None Do not return anything, modify matrix in-place instead.
        """
        if not matrix: return
        m = len(matrix)
        n = len(matrix[0])
        row0 = col0 = False
        for i in range(m):
            for j in range(n):
                if matrix[i][j]==0:
                    if i==0: row0 = True
                    if j==0: col0 = True
                    matrix[i][0] = matrix[0][j] = 0
        
        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][0]==0 or matrix[0][j]==0:
                    matrix[i][j] = 0
        if row0:
            for i in range(n):
                matrix[0][i] = 0
        if col0:
            for i in range(m):
                matrix[i][0] = 0

# 思路：使用第0列标记1 到 m-1行是否存在0，使用第0行标记1 到 n-1列是否存在0
#       另外使用两个变量row0和col0标记第0行和第0列是否存在0
# 时间复杂度：O(m*n), 空间复杂度: O(1)


