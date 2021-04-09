class Solution(object):
    def smallestRepunitDivByK(self, K):
        """
        :type K: int
        :rtype: int
        """
        if K%2==0 or K%5==0:
            return -1
        x = 1
        ans = 1
        while x%K != 0:
            x = x%K
            x = x*10 + 1
            ans += 1
        return ans