class Solution(object):
    def numWays(self, s):
        """
        :type s: str
        :rtype: int
        """
        mod = 10**9 + 7
        def sn(n):
            return n*(n+1)/2

        n = len(s)
        if n < 3:
            return 0
        count = collections.Counter(s)
        cnt = count['1']
        if cnt%3!=0:
            return 0
        if cnt==0: # no 1
            return sn(n-2)%mod
        each = cnt/3
        cnt = 0
        i = 0
        while i < n and cnt < each:
            if s[i]=='1':
                cnt += 1
            i += 1
        l = r = 0
        while i < n and s[i]=='0':
            l += 1
            i += 1
        cnt = 0
        while i < n and cnt < each:
            if s[i]=='1':
                cnt += 1
            i += 1
        while i < n and s[i]=='0':
            r += 1
            i += 1
        # print each, l, r
        return (l+1)*(r+1)%mod
        
