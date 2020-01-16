class Solution(object):
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        if n == 1:
            return '1'
        s = self.countAndSay(n-1)
        ls = len(s)
        i = 0
        t = ''
        while i < ls:
            j = i + 1
            while j < ls and s[i] == s[j]:
                j += 1
            k = j-i
            t += str(k) + s[i]    
            i += k
        
        return t
        
        