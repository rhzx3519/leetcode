class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """
        if not matrix:
            return -1
        m, n = len(matrix), len(matrix[0])
        l = matrix[0][0]
        r = matrix[m-1][n-1]
        while l < r:
            mid = (l + r)//2
            count = 0
            j = n-1
            for i in range(m):
                while j>=0 and matrix[i][j] > mid:
                    j -= 1
                count += (j+1)
            if count < k:
                l = mid + 1
            else:
                r = mid
        return l

if __name__ == '__main__':
    matrix = [[1,5,9],[10,11,13],[12,13,15]]
    k = 8
    su = Solution()
    print su.kthSmallest(matrix, k)                    