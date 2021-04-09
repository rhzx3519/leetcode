class Solution(object):
    def minCost(self, s, cost):
        """
        :type s: str
        :type cost: List[int]
        :rtype: int
        """
        ans = 0
        n = len(s)
        l = 0
        for i in range(1, n):
            if s[i-1] != s[i]:
                if i-1>l:
                    # print l, i
                    ans += sum(cost[l:i]) - max(cost[l:i])
                l = i
        if l != n-1:
            ans += sum(cost[l:]) - max(cost[l:])
        return ans