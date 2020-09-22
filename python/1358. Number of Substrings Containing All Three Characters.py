class Solution(object):
    def numberOfSubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        n = len(s)
        ans = 0
        l = 0
        count = [0]*3
        for r, ch in enumerate(s):
            idx = ord(ch) - ord('a')
            count[idx] += 1
            while all([i>0 for i in count]):
                count[ord(s[l]) - ord('a')] -= 1
                l += 1
            ans += l
        return ans