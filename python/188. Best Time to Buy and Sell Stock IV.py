class Solution(object):
    def maxProfit(self, k, prices):
        """
        :type k: int
        :type prices: List[int]
        :rtype: int
        """
        if k==0: return 0
        if k>=len(prices)/2: return self.greedy(prices)
        INT_MIN = -1<<31
        dp = [[INT_MIN, 0] for _ in range(k)]
        for p in prices:
            dp[0][0] = max(dp[0][0], -p)
            dp[0][1] = max(dp[0][1], dp[0][0] + p)
            for i in range(1, k):
                dp[i][0] = max(dp[i][0], dp[i-1][1] - p)
                dp[i][1] = max(dp[i][1], dp[i][0] + p)

        return dp[k-1][1]

    def greedy(self, prices):
        res = 0
        for i in range(1, len(prices)):
            if prices[i] > prices[i-1]:
                res += prices[i] - prices[i-1]
        return res

# dp[i][0]和dp[i][1]分别表示第i笔交易买入和卖出时的股票价格        