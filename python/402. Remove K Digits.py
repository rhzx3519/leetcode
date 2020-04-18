class Solution(object):
    def removeKdigits(self, num, k):
        """
        :type num: str
        :type k: int
        :rtype: str
        """
        s = []
        for ch in num:
            while k!=0 and s and s[-1] > ch:
                s.pop()
                k -= 1
            if not s and ch=='0':
                continue
            s.append(ch)
        
        while s and k!=0:
            s.pop()
            k -= 1
        return ''.join(s)   

if __name__ == '__main__':
    num = "1432219"
    k = 3
    su = Solution()
    print su.removeKdigits(num, k)        