class Solution(object):
    def minimumOperations(self, leaves):
        """
        :type leaves: str
        :rtype: int
        """
        INFINITY = float('inf')
        n = len(leaves)
        dp = [[0]*3 for _ in range(n)]
        dp[0][0] = int(leaves[0]=='y')
        dp[0][1] = dp[0][2] = dp[1][2] = INFINITY
        for i in range(1, n):
            isYellow = int(leaves[i]=='y')
            isRed = int(leaves[i]=='r')
            dp[i][0] = dp[i-1][0] + isYellow
            dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + isRed
            if i >= 2:
                dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + isYellow
        return dp[n-1][2]

# 思路：动态规划
# dp[i][j] 表示leves[i]是状态j时的最小调整次数，j = {0|1|2}, 分别表示
# 数组的red/yellow/red三部分,
# 分成三种情况：
# 1. j=0时，dp[i][0] = dp[i-1][0] + isYellow(leves[i])
# 2. j=1时，dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + isRed(leves[i])
# 3. j=2时，dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + isYellow(leves[i])