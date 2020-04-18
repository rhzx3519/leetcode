class Solution(object):
    def maxSumTwoNoOverlap(self, A, L, M):
        """
        :type A: List[int]
        :type L: int
        :type M: int
        :rtype: int
        """
        n = len(A)
        max_L = 0
        max_LM = 0
        
        for i in range(n-L-M+1):
            max_L = max(max_L, sum(A[i:i+L]))
            max_LM = max(max_LM, max_L + sum(A[i+L:i+L+M]))

        max_M = 0
        max_ML = 0
        for i in range(n-L-M+1):
            max_M = max(max_M, sum(A[i:i+M]))
            max_ML = max(max_ML, max_M + sum(A[i+M:i+L+M]))            
        
        return max(max_LM, max_ML)