class Solution(object):
    def maxUncrossedLines(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: int
        """
        la, lb = len(A), len(B)
        dp = [[0]*(lb+1) for _ in range(la+1)]
        for i in range(1, la+1):
            for j in range(1, lb+1):
                val = int(A[i-1]==B[j-1])
                dp[i][j] = max(dp[i-1][j-1]+val, dp[i-1][j], dp[i][j-1])
        return dp[la][lb]
        

if __name__ == '__main__':
    A = [2,5,1,2,5]
    B = [10,5,2,1,5,2]
    su = Solution()
    su.maxUncrossedLines(A, B)