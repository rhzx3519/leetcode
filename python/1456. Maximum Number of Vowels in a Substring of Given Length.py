class Solution(object):
    def maxVowels(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: int
        """
        ans = cnt = l = 0
        vowels = ('a', 'e', 'i', 'o', 'u')
        for r, ch in enumerate(s):
            if ch in vowels:
                cnt += 1
            if r>=k:
                if s[l] in vowels:
                    cnt -= 1
                l += 1
            ans = max(cnt, ans)
        return ans
