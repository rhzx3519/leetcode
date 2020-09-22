class Solution(object):
    def maxSatisfaction(self, satisfaction):
        """
        :type satisfaction: List[int]
        :rtype: int
        """
        satisfaction.sort(reverse=True)
        s = t = 0
        for p in satisfaction:
            t += p
            if t > 0:
                s += t
            else:
                break
        return s

# time: O(NlogN), space: O(1)
# 思路：贪心