class Solution(object):
    def numSubarrayBoundedMax(self, A, L, R):
        """
        :type A: List[int]
        :type L: int
        :type R: int
        :rtype: int
        """

        def foo(x):
            res = 0
            numSubarry = 0
            n = len(A)
            for i in range(n):
                if A[i]<=x:
                    numSubarry += 1
                    res += numSubarry
                else:
                    numSubarry = 0
            return res

        return foo(R) - foo(L-1)