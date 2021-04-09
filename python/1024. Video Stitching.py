class Solution(object):
    def videoStitching(self, clips, T):
        """
        :type clips: List[List[int]]
        :type T: int
        :rtype: int
        """
        dp = [float('inf')] * (T+1)
        dp[0] = 0
        for i in range(1, T+1):
            for a, b in clips:
                if a<i<=b:
                    dp[i] = min(dp[i], dp[a] + 1)
        return dp[T] if dp[T]!=float('inf') else -1

# dp[i] 表示覆盖0~i所需要的最小区间树木
# clip=[a, b], 如果a < i <= b, 表示区间[a,b]能够覆盖i,