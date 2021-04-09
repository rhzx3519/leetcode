class Solution(object):
    def bestTeamScore(self, scores, ages):
        """
        :type scores: List[int]
        :type ages: List[int]
        :rtype: int
        """
        ans = 0
        n = len(scores)
        a = [(ages[i], scores[i]) for i in range(n)]
        a.sort()
        dp = [0]*n
        for i in range(n):
            dp[i] = a[i][1]
            for j in range(i):
                if a[i][1] >= a[j][1]:
                    dp[i] = max(dp[i], dp[j] + a[i][1])
            ans = max(ans, dp[i])
        return ans

# 思路：最长上升子序列问题