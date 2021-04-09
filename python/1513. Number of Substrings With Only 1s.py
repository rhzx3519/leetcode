class Solution(object):
    def numSub(self, s):
        """
        :type s: str
        :rtype: int
        """
        def foo(n):
            if n==1: return 1
            return n + foo(n-1)

        def bar(i, j):
            return (j-i+1)*(j-i)/2

        mod = 10**9+7
        i = 0
        l = -1
        ans = 0
        n = len(s)
        while i < n:
            if s[i]=='1':
                l = i
                while i<n and s[i]=='1':
                    i += 1
                cnt = i - l
                # print cnt
                ans = (ans + bar(l, i)) % mod
            i += 1
        
        return ans