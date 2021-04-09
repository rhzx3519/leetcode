class Solution(object):
    def countSubstrings(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: int
        """
        ans = 0
        ls, lt = len(s), len(t)
        for i in range(ls):
            for j in range(lt):
                k = 0
                diff = 0
                while i+k<ls and j+k<lt:
                    diff += int(s[i+k] != t[j+k])
                    if diff == 1:
                        ans += 1
                    elif diff > 1:
                        break
                    k += 1
        return ans