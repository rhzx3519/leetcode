class Solution(object):
    def maxLengthBetweenEqualCharacters(self, s):
        """
        :type s: str
        :rtype: int
        """
        ans = -1
        last = {}
        for i, ch in enumerate(s):
            if ch in last:
                ans = max(ans, i - last[ch] - 1)
                continue
            last[ch] = i
        return ans

if __name__ == '__main__':
    s = "mgntdygtxrvxjnwksqhxuxtrv"
    su = Solution()
    print su.maxLengthBetweenEqualCharacters(s)