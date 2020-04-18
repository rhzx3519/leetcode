class Solution(object):
    def coinChange(self, coins, amount):
        """
        :type coins: List[int]
        :type amount: int
        :rtype: int
        """
        INF = 1<<31
        dp = [INF] * (amount+1)
        dp[0] = 0 
        for coin in coins:
            for i in range(1, amount+1):
                if i>=coin:
                    if dp[i-coin] != INF:
                        dp[i] = min(dp[i], dp[i-coin] + 1)
        if dp[amount]==INF:
            return -1
        return dp[amount]

