class Solution(object):
    def maxCoins(self, piles):
        """
        :type piles: List[int]
        :rtype: int
        """
        ans = 0
        n = len(piles)
        piles.sort(reverse=True)
        for i in range(1, 2*(n/3), 2):
            ans += piles[i]
        return ans

# 排序去掉最小的1/3的石堆，间隔2取
# time: O(n)