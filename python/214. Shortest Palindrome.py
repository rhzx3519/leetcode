class Solution(object):
    def shortestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        if not s:
            return s
        n = len(s)
        t = s[::-1]
        i = 0
        while True:
            if t[i:] == s[:n-i]:
                break
            i += 1
        return t[:i] + s