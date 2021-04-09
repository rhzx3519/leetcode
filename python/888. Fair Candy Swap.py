class Solution(object):
    def fairCandySwap(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: List[int]
        """
        A.sort()
        B.sort()

        s = (sum(A) - sum(B))/2
        i = j = 0
        while i < len(A) and j < len(B):
            if A[i] - B[j] == s:
                return (A[i], B[j])
            elif A[i] - B[j] < s:
                i += 1
            else:
                j += 1
        return -1

# 思路： sumB - B + A = sumA - A + B -> sumA - sumB = 2(A - B)
