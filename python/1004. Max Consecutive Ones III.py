class Solution(object):
    def longestOnes(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: int
        """
        max_val = 0
        l = r = 0
        zero = 0
        while r < len(A):
            if A[r] == 0:
                zero += 1
            while zero > K:
                if A[l] == 0:
                    zero -= 1
                l += 1
            max_val = max(max_val, r-l+1)
            r += 1
        return max_val