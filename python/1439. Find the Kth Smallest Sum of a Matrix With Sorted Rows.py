class Solution(object):
    def kthSmallest(self, mat, k):
        """
        :type mat: List[List[int]]
        :type k: int
        :rtype: int
        """
        res = [0]
        for row in mat:
            res = sorted([x+r for r in row for x in res])[:k]
        return res[-1]