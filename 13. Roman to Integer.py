class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        mp = {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
        i = 0
        res = 0
        while i < len(s) - 1:
            if mp[s[i]] < mp[s[i+1]]:
                res -= mp[s[i]]
            else:
                res += mp[s[i]]
            i += 1
        return res + mp[s[-1]]
        