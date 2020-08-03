class Solution(object):
    def spiralOrder(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        if not matrix:
            return []
        m, n = len(matrix), len(matrix[0])
        li, hi = 0, m-1
        lj, hj = 0, n-1

        i = j = 0
        res = []
        while li<=hi and lj<=hj:
            i = li
            j = lj
            while j <= hj:
                print i, j
                res.append(matrix[i][j])
                j += 1
            i += 1
            j -= 1
            while i <= hi:
                res.append(matrix[i][j])
                i += 1
            i -= 1
            j -= 1
            while li != hi and j > lj:
                res.append(matrix[i][j])
                j -= 1

            while lj != hj and i > li:
                res.append(matrix[i][j])
                i -= 1

            li += 1
            lj += 1
            hi -= 1
            hj -= 1

        return res

if __name__ == '__main__':
    matrix = [[6,9,7]]

    su = Solution()
    print su.spiralOrder(matrix)