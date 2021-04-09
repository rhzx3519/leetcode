class Solution(object):
    def twoCitySchedCost(self, costs):
        """
        :type costs: List[List[int]]
        :rtype: int
        """
        sumA = sum([c[0] for c in costs])
        t = [c[0] - c[1] for c in costs]
        t.sort(reverse=True)
        for i in range(len(t)/2):
            sumA -= t[i]
        return sumA

# 思路：先求出所有人都去A时的总花费，然后A与B之间的差值按照你序排列