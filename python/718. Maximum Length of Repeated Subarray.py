class Solution(object):
    def findLength(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: int
        """
        n1, n2 = len(A), len(B)
        max_val = 0
        dp = [[0]*(n2+1) for _ in range(n1+1)]
        for i in range(1, n1+1):
            for j in range(1, n2+1):
                if A[i-1] == B[j-1]:
                    dp[i][j] = dp[i-1][j-1] + 1
                    max_val = max(dp[i][j], max_val)
                else:
                    dp[i][j] = 0
        return max_val

# 思路：dp[i][j]表示以以i, j结尾的 A, B的公共子数组长度
# dp[i][j] = dp[i-1][j-1] + 1 if A[i-1]==B[j-1] else dp[i][j] = 0
# time: O(m*n), space: O(m*n)
                    