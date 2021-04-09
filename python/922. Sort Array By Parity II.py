class Solution(object):
    def sortArrayByParityII(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """
        j = 1
        for k in range(0, len(A)-1, 2):
            if A[k]%2 != 0:
                while A[j]%2 != 0:
                    j += 2
                A[k], A[j] = A[j], A[k]
        return A