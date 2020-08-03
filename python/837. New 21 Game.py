class Solution(object):
    def new21Game(self, N, K, W):
        """
        :type N: int
        :type K: int
        :type W: int
        :rtype: float
        """
        dp = [0.0] * (K+W)                      # (N + W)
        
        for k in range(K, min(N + 1, K + W)):    # (K, N + 1)
            dp[k] = 1.0

        S = min(N - K + 1,W)
        for k in range(K - 1, -1, -1):
            dp[k] = S / float(W)
            S += dp[k] - dp[k + W]

        return dp[0]