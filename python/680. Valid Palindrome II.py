class Solution(object):
    def validPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        def valid(start, end):
            l, r = start, end
            while l<r:
                if s[l]!=s[r]:
                    return False
                l += 1
                r -= 1
            return True
        l, r = 0, len(s)-1
        while l<r:
            if s[l]!=s[r]:
                return valid(l+1, r) or valid(l, r-1)
            l += 1
            r -= 1
        return True