class Solution(object):
    def equalSubstring(self, s, t, maxCost):
        """
        :type s: str
        :type t: str
        :type maxCost: int
        :rtype: int
        """
        costs = []
        n = len(s)
        for i in range(n):
            dis = abs(ord(s[i]) - ord(t[i]))
            costs.append(dis)

        # print costs
        ans = max_len = 0
        l = 0
        for r in range(n):
            maxCost -= costs[r]
            while maxCost < 0:
                maxCost += costs[l]
                l += 1
            ans = max(ans, r - l + 1)
        return ans

# 思路：滑动窗口。
# time: O(N), space: O(N)
