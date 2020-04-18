class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        mem = collections.defaultdict(int)
        l = r = 0
        max_len = 0
        while r < len(s):
            ch = s[r]
            mem[ch] += 1
            r += 1

            while mem[ch]>1:
                mem[s[l]] -= 1
                l += 1
            if r-l > max_len:
                max_len = r-l

        return max_len