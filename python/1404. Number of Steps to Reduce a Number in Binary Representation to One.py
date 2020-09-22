class Solution(object):
    def numSteps(self, s):
        """
        :type s: str
        :rtype: int
        """
        def foo(s):
            num = 0
            n = len(s)
            for i in range(n):
                num += 2**(n-i-1) * (ord(s[i]) - ord('0'))
            return num

        s = foo(s)
        ans = 0
        # print s
        
        while s != 1:
            ans += 1
            if s&1 == 0:
                s /= 2
            else:
                s += 1
        print ans
        return ans

if __name__ == '__main__':
    s = '1101'
    su = Solution()
    su.numSteps(s)        