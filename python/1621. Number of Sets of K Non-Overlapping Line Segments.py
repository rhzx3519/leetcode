class Solution(object):
    def numberOfSets(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: int
        """
        Mod = 10**9 + 7
        mem = {}
        def dp(n, k):
            if k > n-1: return 0
            if k==n-1: return 1
            if k==1: return n*(n-1)//2
            if (n, k) in mem: return mem[(n, k)]
            mem[(n, k)] = (2*dp(n-1, k) + dp(n-1, k-1) - dp(n-2, k))%Mod
            return mem[(n, k)]

        return dp(n, k)

        
