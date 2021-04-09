class Solution(object):
    def minimumDeletions(self, s):
        """
        :type s: str
        :rtype: int
        """
        n = len(s)
        dp = [[0]*2 for _ in range(n+1)]
        for i in range(1, n+1):
            if s[i-1]=='a':
                dp[i][0] = dp[i-1][0]
                dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
            else:
                dp[i][0] = dp[i-1][0] + 1
                dp[i][1] = min(dp[i-1][0], dp[i-1][1])
        return min(dp[-1])

# 思路：dp[i][0]表示前i个字符，以'a'结尾的平衡字符串的最少删除次数
#       dp[i][1]表示前i个字符，以'b'结尾的平衡字符串的最少删除次数
