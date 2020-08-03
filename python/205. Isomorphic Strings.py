class Solution(object):
    def isIsomorphic(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        n = len(s)
        for i in range(n):
            if s.index(s[i]) != t.index(t[i]):
                return False
        return True