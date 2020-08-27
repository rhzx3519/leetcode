class Solution(object):
    def lastStoneWeightII(self, stones):
        """
        :type stones: List[int]
        :rtype: int
        """
        s = sum(stones)
        dp = [0] * (s//2 + 1)
        n = len(stones)
        for i in range(n):
            for j in range(s//2, stones[i]-1, -1):
                dp[j] = max(dp[j], stones[i] + dp[j - stones[i]])
        return s - 2*dp[-1]

# 转换成01背包问题，求两堆石头的最小差值。由于石头总和为sum.则问题转换成了
# 背包最多装sum / 2的石头,stones数组里有一大堆石头。求如何装能装下最多重量石头。