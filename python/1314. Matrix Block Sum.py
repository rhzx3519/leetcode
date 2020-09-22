class Solution(object):
    def matrixBlockSum(self, mat, K):
        """
        :type mat: List[List[int]]
        :type K: int
        :rtype: List[List[int]]
        """
        m, n = len(mat), len(mat[0])
        ans = [[0]*n for _ in range(m)]
        dp = [[0]*(n+1) for _ in range(m+1)]
        for i in range(1, m+1):
            for j in range(1, n+1):
                dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
        print dp
        for i in range(m):
            for j in range(n):
                # 左上角坐标
                r1 = max(0, i-K)
                c1 = max(0, j-K)
                # 右下角坐标
                r2 = min(m-1, i+K)
                c2 = min(n-1, j+K)
                ans[i][j] = dp[r2+1][c2+1] - dp[r1][c2+1] - dp[r2+1][c1] + dp[r1][c1]

        return ans

# time: O(m*n), space: O(m*n)