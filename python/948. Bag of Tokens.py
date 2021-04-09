class Solution(object):
    def bagOfTokensScore(self, tokens, P):
        """
        :type tokens: List[int]
        :type P: int
        :rtype: int
        """
        ans = 0
        score = 0
        tokens.sort()
        l, r = 0, len(tokens)-1
        while l <= r and (P >= tokens[l] or score):
            while l <= r and P >= tokens[l]:
                score += 1
                P -= tokens[l]
                l += 1

            ans = max(ans, score)
            if l <= r and score:
                P += tokens[r]
                score -= 1
                r -= 1
        return ans


# 思路：贪心, 类似滑动窗口