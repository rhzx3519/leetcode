class Solution(object):
    def isInterleave(self, s1, s2, s3):
        """
        :type s1: str
        :type s2: str
        :type s3: str
        :rtype: bool
        """
        n1, n2, n3 = len(s1), len(s2), len(s3)
        if n1+n2!=n3: return False
        dp = [[False]*(n2+1) for _ in range(n1+1)]
        dp[0][0] = True
        for i in range(1, n1+1):
            dp[i][0] = dp[i-1][0] and s1[i-1]==s3[i-1]
        for j in range(1, n2+1):
            dp[0][j] = dp[0][j-1] and s2[j-1]==s3[j-1]
        
        for i in range(1, n1+1):
            for j in range(1, n2+1):
                dp[i][j] = (dp[i-1][j] and s1[i-1]==s3[i-1+j] ) or (dp[i][j-1] and s2[j-1]==s3[j-1+i])
        
        return dp[n1][n2]

# 思路：dp[i][j]表示s1的前i-1个字符，s2的前j-1个字符能够构成s3的前i+j-1个字符,
# dp[i][j] = (dp[i-1][j] and s1[i-1]==s3[i-1+j] ) or (dp[i][j-1] and s2[j-1]==s3[j-1+i])
