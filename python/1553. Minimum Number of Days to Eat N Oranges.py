class Solution(object):
    def __init__(self):
        self.mem = {}

    def minDays(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n<=1:
            return n
        if n in self.mem:
            return self.mem[n]
        ans = min(self.minDays(n/3) + n%3 + 1, self.minDays(n/2) + n%2 + 1)
        self.mem[n] = ans
        return ans

# 思路：记忆化搜索