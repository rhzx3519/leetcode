class Solution(object):
    def kthGrammar(self, N, K):
        """
        :type N: int
        :type K: int
        :rtype: int
        """
        if N<=1 and K<=1:
            return 0
        last = self.kthGrammar(N-1, (K+1)/2)
        if last==0:
            return 0 if K&1 else 1
        return 1 if K&1 else 0
        