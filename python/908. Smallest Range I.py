class Solution(object):
    def smallestRangeI(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: int
        """
        A.sort()
        diff = A[-1] - A[0]
        return min(diff, max(0,diff - 2*K))