class Solution(object):
    def maxScore(self, cardPoints, k):
        """
        :type cardPoints: List[int]
        :type k: int
        :rtype: int
        """
        n = len(cardPoints)
        s = 0
        min_val = float('inf')
        for r, p in enumerate(cardPoints):
            s += p
            if r >= n-k:
                s -= cardPoints[r-n+k]
            if r >= n-k-1:
                min_val = min(min_val, s)
        return sum(cardPoints) - min_val

# 维持大小等于n-k的窗口，从左往右遍历，获取最小的窗口值
# time: O(N), space: O(N-K)
