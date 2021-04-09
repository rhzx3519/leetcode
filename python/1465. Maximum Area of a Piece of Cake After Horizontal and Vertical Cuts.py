class Solution(object):
    def maxArea(self, h, w, horizontalCuts, verticalCuts):
        """
        :type h: int
        :type w: int
        :type horizontalCuts: List[int]
        :type verticalCuts: List[int]
        :rtype: int
        """
        MOD = 10**9 + 7
        horizontalCuts.sort()
        verticalCuts.sort()
        max_h = max_w = 0
        prev_h = prev_w = 0
        for i, t in enumerate(horizontalCuts + [h]):
            max_h = max(max_h, t - prev_h)
            prev_h = t

        for j, t in enumerate(verticalCuts + [w]):
            max_w = max(max_w, t - prev_w)
            prev_w = t
        return (max_h*max_w) % MOD
            