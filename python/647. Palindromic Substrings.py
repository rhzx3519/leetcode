class Solution(object):
    def countSubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        self.res = 0

        def count(l, r):
            while l>=0 and r<len(s) and s[l]==s[r]:
                self.res += 1
                l -= 1
                r += 1

        for i in range(len(s)):
            count(i, i)
            count(i, i+1)
        return self.res
