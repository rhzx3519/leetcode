class Solution(object):
    def maxScoreSightseeingPair(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        n = len(A)
        if n==0:
            return
        res = -float('inf')
        left = A[0]
        for i in range(1, n):
            res = max(res, left + A[i] - i)
            left = max(left, A[i] + i)
        return res

# 思路: A[i] + A[j] + i - j, 可以分解为(A[i] + i) + (A[j] - j), 一次遍历分别求取两部分最大值
# time: O(N), space: O(1)
        