class Solution(object):
    def matrixScore(self, A):
        """
        :type A: List[List[int]]
        :rtype: int
        """
        m, n = len(A), len(A[0])
        for i in range(m):
            if A[i][0] == 0:
                A[i][0] = 1
                for j in range(1, n):
                    A[i][j] ^= 1

        for j in range(1, n):
            cnt = 0
            for i in range(m):
                if A[i][j]==1:
                    cnt += 1
            if cnt < m-cnt:
                for i in range(m):
                    A[i][j] ^= 1

        ans = 0
        for i in range(m):
            ans += int(''.join(map(str, A[i])), 2)
        return ans
