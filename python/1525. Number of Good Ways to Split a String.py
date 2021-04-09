class Solution(object):
    def numSplits(self, s):
        """
        :type s: str
        :rtype: int
        """
        right = collections.Counter(s)
        left = collections.defaultdict(int)
        n = len(s)
        ans = 0
        for i in range(n-1):
            right[s[i]] -= 1
            if right[s[i]] == 0:
                del right[s[i]]
            left[s[i]] += 1
            if len(left) == len(right):
                # print i
                ans += 1
        return ans