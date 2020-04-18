class Solution(object):
    def numDecodings(self, s):
        """
        :type s: str
        :rtype: int
        """
        n = len(s)
        f1 = 0
        f2 = 1

        for i in range(n):
            f3 = 0
            d = ord(s[i]) - ord('0')
            if d>=1 and d<=26:
                f3 += f2
            if i>0 and s[i-1]!='0':
                d = ( ord(s[i-1]) - ord('0') ) * 10 + d
                if d>=1 and d<=26:            
                    f3 += f1
            f1 = f2
            f2 = f3
        return f3