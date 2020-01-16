class Solution(object):
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        s = haystack
        p = needle
        i = 0
        ls, lp = len(haystack), len(needle)
        while i <= ls - lp:
            j = 0
            while j < lp:
                if s[i+j] != p[j]:
                    break
                j += 1
            if j == lp:
                return i
            i += 1
        return -1
        