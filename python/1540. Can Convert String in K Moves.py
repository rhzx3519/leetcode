class Solution(object):
    def canConvertString(self, s, t, k):
        """
        :type s: str
        :type t: str
        :type k: int
        :rtype: bool
        """
        ls, lt = len(s), len(t)
        if ls != lt: return False

        dp = [0]*27
        num = k/26
        for i in range(1, 26):
            if 26*num + i <= k:
                dp[i] = 1 + num
            else:
                dp[i] = num
        # print dp

        for i in range(ls):
            s1 = ord(s[i]) - ord('a')
            t1 = ord(t[i]) - ord('a')
            x = t1 - s1 if s1 <= t1 else 26+t1-s1
            if x == 0: continue
            if x > k: return False
            if dp[x%26] == 0: return False
            dp[x%26] -= 1
        return True

if __name__ == '__main__':
    s = "input"
    t = "ouput"
    k = 9

    s = "aab"
    t = "bbb"
    k = 27

    s = "abc"
    t = "bcd"
    k = 10

    s = "aaaaaaaaaaaaaaaaaaaaaaaaaa"
    t = "zyxwvuysrqponmlkjihgfedcba"
    k = 100000000
    su = Solution()
    print su.canConvertString(s, t, k)