class Solution(object):
    def minCostClimbingStairs(self, cost):
        """
        :type cost: List[int]
        :rtype: int
        """
        n = len(cost)
        f1 = f2 = 0
        for i in range(n-1, -1, -1):
            f1, f2 = cost[i] + min(f1, f2), f1
        return min(f1, f2)

            