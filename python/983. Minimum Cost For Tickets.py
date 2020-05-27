class Solution(object):
    def mincostTickets(self, days, costs):
        """
        :type days: List[int]
        :type costs: List[int]
        :rtype: int
        """
        INF = float('inf')
        N = days[-1]
        dp = [0]*(N+1)
        for day in days:
            dp[day] = -1

        for i in range(1, N+1):
            if dp[i]==0:
                dp[i] = dp[i-1]
                continue
            a, b, c = costs
            a += dp[i-1]
            if i>=30:
                c += dp[i-30]
            if i>=7:
                b += dp[i-7]
            dp[i] = min(a, b, c)
        return dp[N]

# 思路 dp[i]表示第i天旅行的最低费用，dp[i] = min(dp[i-1] + costs[0], dp[i-7] + costs[1], dp[i-30] + costs[2])
# time: O(N), space: O(N)
            
                