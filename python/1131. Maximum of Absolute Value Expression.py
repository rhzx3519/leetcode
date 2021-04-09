class Solution(object):
    def maxAbsValExpr(self, arr1, arr2):
        """
        :type arr1: List[int]
        :type arr2: List[int]
        :rtype: int
        """
        n = len(arr1)
        dp = ((1,1,1), (1,1,-1), (1,-1,1), (1,-1,-1), (-1,1,1), (-1,-1,1), (-1,1,-1), (-1,-1,-1))
        k = 8
        
        ans = 0
        maxK = [float('-inf')]*k
        for i in range(k):
            for j in range(n):
                maxK[i] = max(maxK[i], j*dp[i][0] + arr1[j]*dp[i][1] + arr2[j]*dp[i][2])
        
        for i in range(k):
            for j in range(n):
                ans = max(ans, maxK[i] - j*dp[i][0] - arr1[j]*dp[i][1] - arr2[j]*dp[i][2])
        
        return ans