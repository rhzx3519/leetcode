class Solution(object):
    def findKthBit(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """
        def dfs(n):
            if n==1: return [0]
            t = dfs(n-1)
            return t + [1] + map(lambda x: x^1, t)[::-1]
        
        s = dfs(n)
        # print s
        return str(s[k-1])
        