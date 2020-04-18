class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        n = len(prices)
        if n < 2: return 0
        dp1 = [0]*n # 表示i天之前的股票交易能取得的最大利润
        dp2 = [0]*n # 表示i天之后的股票交易能取得的最大利润

        min_val = prices[0]
        for i in range(1, n):
            dp1[i] = max(dp1[i-1], prices[i]-min_val)
            min_val = min(min_val, prices[i])

        max_val = prices[-1]
        for i in range(n-2, -1, -1):
            dp2[i] = max(dp2[i+1], max_val - prices[i])
            max_val = max(max_val, prices[i])
        
        return max([dp1[i]+dp2[i] for i in range(n)])
        